package subfunction

import (
	"context"
	api_input_reader "data-platform-api-orders-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"data-platform-api-orders-creates-subfunc-rmq-kube/database"
	"data-platform-api-orders-creates-subfunc-rmq-kube/database/models"
	"fmt"
	"sync"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type SubFunction struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewSubFunction(ctx context.Context, db *database.Mysql, l *logger.Logger) *SubFunction {
	return &SubFunction{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (f *SubFunction) Controller(sdc *api_input_reader.SDC) error {
	var bPCustomerRecord *models.DataPlatformBusinessPartnerCustomerDatum
	var bPSupplierRecord *models.DataPlatformBusinessPartnerSupplierDatum
	var nRLatestNumberRecord *models.DataPlatformNumberRangeLatestNumberDatum
	var bPCustomerPartnerFunctionArray models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatumSlice
	var bPSupplierPartnerFunctionArray models.DataPlatformBusinessPartnerSupplierPartnerFunctionDatumSlice
	var bPGeneralArray models.DataPlatformBusinessPartnerGeneralDatumSlice
	var bPCustomerPartnerPlantArray models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice
	var bPSupplierPartnerPlantArray models.DataPlatformBusinessPartnerSupplierPartnerPlantDatumSlice
	var bPCustomerTaxRecord *models.DataPlatformBusinessPartnerCustomerTaxDatum
	var bPSupplierTaxRecord *models.DataPlatformBusinessPartnerSupplierTaxDatum
	var pMGeneralArray models.DataPlatformProductMasterGeneralDatumSlice
	var pMProductDescriptionArray models.DataPlatformProductMasterProductDescriptionDatumSlice
	var err error
	var e error

	businessPartner := sdc.BusinessPartner
	serviceLabel := sdc.ServiceLabel
	property := "OrderID"
	buyer := sdc.Orders.Buyer
	seller := sdc.Orders.Seller
	processingData := &api_processing_data_formatter.HeaderRelatedData{}

	// 1-0. 入力ファイルのbusiness_partnerがBuyerであるかSellerであるかの判断
	if *businessPartner == *buyer && *businessPartner != *seller {
		processingData.BuyerOrSeller = "Buyer"
	} else if *businessPartner != *buyer && *businessPartner == *seller {
		processingData.BuyerOrSeller = "Seller"
	} else {
		return fmt.Errorf("business_partnerがBuyerまたはSellerと一致しません")
	}
	f.l.Info(processingData.BuyerOrSeller)

	wg := sync.WaitGroup{}
	wg.Add(4)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-1. ビジネスパートナ 得意先データ/仕入先データ の取得
		if processingData.BuyerOrSeller == "Seller" {
			bPCustomerRecord, e = f.ExtractBusinessPartnerCustomerRecord(businessPartner, buyer)
			if e != nil {
				err = e
				return
			}
			sdc.Orders = *f.SetBusinessPartnerCustomerRecord(&sdc.Orders, bPCustomerRecord)
		} else if processingData.BuyerOrSeller == "Buyer" {
			bPSupplierRecord, e = f.ExtractBusinessPartnerSupplierRecord(businessPartner, seller)
			if e != nil {
				err = e
				return
			}
			sdc.Orders = *f.SetBusinessPartnerSupplierRecord(&sdc.Orders, bPSupplierRecord)
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-2. OrderID
		nRLatestNumberRecord, e = f.ExtractNumberRangeLatestNumberRecord(serviceLabel, property)
		if e != nil {
			err = e
			return
		}
		processingData = f.HoldNumberRangeLatestNumberRecord(processingData, nRLatestNumberRecord)
		sdc.Orders = *f.SetNumberRangeLatestNumberRecord(&sdc.Orders, processingData.LatestNumber)
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		start := time.Now()
		// 2-1. ビジネスパートナマスタの取引先機能データの取得
		if processingData.BuyerOrSeller == "Seller" {
			bPCustomerPartnerFunctionArray, e = f.ExtractBusinessPartnerCustomerPartnerFunctionArray(businessPartner, buyer)
			if e != nil {
				err = e
				return
			}
			processingData = f.HoldBusinessPartnerCustomerPartnerFunctionArray(processingData, bPCustomerPartnerFunctionArray)
			sdc.Orders = *f.SetBusinessPartnerCustomerPartnerFunctionArray(&sdc.Orders, bPCustomerPartnerFunctionArray)
		} else if processingData.BuyerOrSeller == "Buyer" {
			bPSupplierPartnerFunctionArray, e = f.ExtractBusinessPartnerSupplierPartnerFunctionArray(businessPartner, seller)
			if e != nil {
				err = e
				return
			}
			processingData = f.HoldBusinessPartnerSupplierPartnerFunctionArray(processingData, bPSupplierPartnerFunctionArray)
			sdc.Orders = *f.SetBusinessPartnerSupplierPartnerFunctionArray(&sdc.Orders, bPSupplierPartnerFunctionArray)
		}
		fmt.Printf("duration: %d [ms]\n", time.Since(start).Milliseconds())

		// 2-2. ビジネスパートナの一般データの取得
		bPGeneralArray, e = f.ExtractBusinessPartnerGeneralArray(sdc.Orders.HeaderPartner)
		if e != nil {
			err = e
			return
		}
		sdc.Orders = *f.SetBusinessPartnerGeneralArray(&sdc.Orders, bPGeneralArray)
		fmt.Printf("duration: %d [ms]\n", time.Since(start).Milliseconds())

		// 4-1. ビジネスパートナマスタの取引先プラントデータの取得
		if processingData.BuyerOrSeller == "Seller" {
			bPCustomerPartnerPlantArray, e = f.ExtractBusinessPartnerCustomerPartnerPlantArray(businessPartner, buyer, bPCustomerPartnerFunctionArray)
			if e != nil {
				err = e
				return
			}
			processingData = f.HoldBusinessPartnerCustomerPartnerPlantArray(processingData, bPCustomerPartnerPlantArray)
			sdc.Orders = *f.SetBusinessPartnerCustomerPartnerPlantArray(&sdc.Orders, bPCustomerPartnerPlantArray)
		} else if processingData.BuyerOrSeller == "Buyer" {
			bPSupplierPartnerPlantArray, e = f.ExtractBusinessPartnerSupplierPartnerPlantArray(businessPartner, seller, bPSupplierPartnerFunctionArray)
			if e != nil {
				err = e
				return
			}
			processingData = f.HoldBusinessPartnerSupplierPartnerPlantArray(processingData, bPSupplierPartnerPlantArray)
			sdc.Orders = *f.SetBusinessPartnerSupplierPartnerPlantArray(&sdc.Orders, bPSupplierPartnerPlantArray)
		}
		fmt.Printf("duration: %d [ms]\n", time.Since(start).Milliseconds())

		// PartnerFunction=”BUYER”(1-0の処理結果が”Seller”の場合)、または、PartnerFunction=”SELLER”(1-0の処理結果が”Buyer”の場合)のHeaderPartnerの取得
		var buyerOrSellerData *api_input_reader.HeaderPartner
		for _, v := range sdc.Orders.HeaderPartner {
			if processingData.BuyerOrSeller == "Seller" {
				if v.PartnerFunction == "BUYER" {
					buyerOrSellerData = &v
					break
				}
			} else if processingData.BuyerOrSeller == "Buyer" {
				if v.PartnerFunction == "SELLER" {
					buyerOrSellerData = &v
					break
				}
			}
		}

		// 5-1. BPTaxClassification
		if processingData.BuyerOrSeller == "Seller" {
			bPCustomerTaxRecord, e = f.ExtractBusinessPartnerCustomerTaxRecord(businessPartner, buyer, buyerOrSellerData.Country)
			if e != nil {
				err = e
				return
			}
			sdc.Orders = *f.SetBusinessPartnerCustomerTaxRecord(&sdc.Orders, bPCustomerTaxRecord)
		} else if processingData.BuyerOrSeller == "Buyer" {
			bPSupplierTaxRecord, e = f.ExtractBusinessPartnerSupplierTaxRecord(businessPartner, seller, buyerOrSellerData.Country)
			if e != nil {
				err = e
				return
			}
			sdc.Orders = *f.SetBusinessPartnerSupplierTaxRecord(&sdc.Orders, bPSupplierTaxRecord)
		}
		fmt.Printf("duration: %d [ms]\n", time.Since(start).Milliseconds())

		// 5-3. OrderItemText
		pMProductDescriptionArray, e = f.ExtractProductMasterProductDescriptionArray(sdc.Orders.Item, buyerOrSellerData.Language)
		if e != nil {
			err = e
			return
		}
		sdc.Orders = *f.SetProductMasterProductDescriptionArray(&sdc.Orders, pMProductDescriptionArray)
		fmt.Printf("duration: %d [ms]\n", time.Since(start).Milliseconds())
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 5-0. OrderItem
		sdc.Orders = *f.SetOrderItem(&sdc.Orders)

		// 5-2. 品目マスタ一般データの取得
		pMGeneralArray, e = f.ExtractProductMasterGeneralArray(sdc.Orders.Item)
		if e != nil {
			err = e
			return
		}
		sdc.Orders = *f.SetProductMasterGeneralArray(&sdc.Orders, pMGeneralArray)
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	f.l.Info(processingData)

	return nil
}
