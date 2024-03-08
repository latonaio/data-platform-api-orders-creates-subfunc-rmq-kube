package subfunction

import (
	api_input_reader "data-platform-api-orders-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"data-platform-api-orders-creates-subfunc-rmq-kube/database/models"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (f *SubFunction) ExtractBusinessPartnerCustomerPartnerFunctionArray(
	businessPartner *int,
	buyer *int,
) (models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatumSlice, error) {
	res, err := models.DataPlatformBusinessPartnerCustomerPartnerFunctionData(
		qm.And("BusinessPartner=?", *businessPartner),
		qm.And("Customer=?", *buyer),
	).All(f.ctx, f.db)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("sql: no rows in result set")
	}

	return res, nil
}

func (f *SubFunction) ExtractBusinessPartnerSupplierPartnerFunctionArray(
	businessPartner *int,
	seller *int,
) (models.DataPlatformBusinessPartnerSupplierPartnerFunctionDatumSlice, error) {
	res, err := models.DataPlatformBusinessPartnerSupplierPartnerFunctionData(
		qm.And("BusinessPartner=?", *businessPartner),
		qm.And("Supplier=?", *seller),
	).All(f.ctx, f.db)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("sql: no rows in result set")
	}

	return res, nil
}

func (f *SubFunction) ExtractBusinessPartnerGeneralArray(
	headerPartner []api_input_reader.HeaderPartner,
) (models.DataPlatformBusinessPartnerGeneralDatumSlice, error) {
	var res models.DataPlatformBusinessPartnerGeneralDatumSlice
	where := make([]qm.QueryMod, 0, len(headerPartner))
	for _, v := range headerPartner {
		where = append(where,
			qm.Or("BusinessPartner=?", v.BusinessPartner),
		)
	}

	res, err := models.DataPlatformBusinessPartnerGeneralData(
		where...,
	).All(f.ctx, f.db)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("sql: no rows in result set")
	}

	return res, nil
}

func (f *SubFunction) HoldBusinessPartnerCustomerPartnerFunctionArray(
	processingData *api_processing_data_formatter.HeaderRelatedData,
	bPCustomerPartnerFunctionArray models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatumSlice,
) *api_processing_data_formatter.HeaderRelatedData {
	headerPartnerRelatedData := make([]api_processing_data_formatter.HeaderPartnerRelatedData, 0, len(bPCustomerPartnerFunctionArray))
	for _, bPCustomerPartnerFunctionRecord := range bPCustomerPartnerFunctionArray {
		headerPartnerRelatedData = append(headerPartnerRelatedData, api_processing_data_formatter.HeaderPartnerRelatedData{
			PartnerFunction: api_processing_data_formatter.PartnerFunction{
				BusinessPartner: &bPCustomerPartnerFunctionRecord.PartnerFunctionBusinessPartner.Int,
				PartnerCounter:  &bPCustomerPartnerFunctionRecord.PartnerCounter,
				DefaultPartner:  bPCustomerPartnerFunctionRecord.DefaultPartner.Ptr(),
			},
		})
	}

	processingData.HeaderPartnerRelatedData = headerPartnerRelatedData
	return processingData
}

func (f *SubFunction) HoldBusinessPartnerSupplierPartnerFunctionArray(
	processingData *api_processing_data_formatter.HeaderRelatedData,
	bPSupplierPartnerFunctionArray models.DataPlatformBusinessPartnerSupplierPartnerFunctionDatumSlice,
) *api_processing_data_formatter.HeaderRelatedData {
	headerPartnerRelatedData := make([]api_processing_data_formatter.HeaderPartnerRelatedData, 0, len(bPSupplierPartnerFunctionArray))
	for _, bPSupplierPartnerFunctionRecord := range bPSupplierPartnerFunctionArray {
		headerPartnerRelatedData = append(headerPartnerRelatedData, api_processing_data_formatter.HeaderPartnerRelatedData{
			PartnerFunction: api_processing_data_formatter.PartnerFunction{
				BusinessPartner: &bPSupplierPartnerFunctionRecord.PartnerFunctionBusinessPartner.Int,
				PartnerCounter:  &bPSupplierPartnerFunctionRecord.PartnerCounter,
				DefaultPartner:  bPSupplierPartnerFunctionRecord.DefaultPartner.Ptr(),
			},
		})
	}

	processingData.HeaderPartnerRelatedData = headerPartnerRelatedData
	return processingData
}

func (f *SubFunction) SetBusinessPartnerCustomerPartnerFunctionArray(
	order *api_input_reader.Order,
	bPCustomerPartnerFunctionArray models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatumSlice,
) *api_input_reader.Order {
	headerPartners := make(map[int]api_input_reader.HeaderPartner, len(bPCustomerPartnerFunctionArray))
	inoutHeaderPartnerMap := make(map[int]api_input_reader.HeaderPartner, len(order.HeaderPartner))
	bPCustomerPartnerFunctionArrayMap := make(map[int]models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatum, len(bPCustomerPartnerFunctionArray))

	for i, v := range order.HeaderPartner {
		inoutHeaderPartnerMap[v.BusinessPartner] = order.HeaderPartner[i]
	}

	for i, v := range bPCustomerPartnerFunctionArray {
		bPCustomerPartnerFunctionArrayMap[v.PartnerFunctionBusinessPartner.Int] = *bPCustomerPartnerFunctionArray[i]
	}

	for businessPartnerKey, bPCustomerPartnerFunctionRecord := range bPCustomerPartnerFunctionArrayMap {
		if _, ok := inoutHeaderPartnerMap[businessPartnerKey]; !ok {
			inoutHeaderPartnerMap[businessPartnerKey] = api_input_reader.HeaderPartner{}
		}

		newHeaderPartner := inoutHeaderPartnerMap[businessPartnerKey]

		newHeaderPartner.PartnerFunction = bPCustomerPartnerFunctionRecord.PartnerFunction.String
		newHeaderPartner.BusinessPartner = bPCustomerPartnerFunctionRecord.PartnerFunctionBusinessPartner.Int

		headerPartners[businessPartnerKey] = newHeaderPartner
	}

	res := make([]api_input_reader.HeaderPartner, 0, len(headerPartners))
	for i := range headerPartners {
		res = append(res, headerPartners[i])
	}

	order.HeaderPartner = res

	return order
}

func (f *SubFunction) SetBusinessPartnerSupplierPartnerFunctionArray(
	order *api_input_reader.Order,
	bPSupplierPartnerFunctionArray models.DataPlatformBusinessPartnerSupplierPartnerFunctionDatumSlice,
) *api_input_reader.Order {
	headerPartners := make(map[int]api_input_reader.HeaderPartner, len(bPSupplierPartnerFunctionArray))
	inoutHeaderPartnerMap := make(map[int]api_input_reader.HeaderPartner, len(order.HeaderPartner))
	bPSupplierPartnerFunctionArrayMap := make(map[int]models.DataPlatformBusinessPartnerSupplierPartnerFunctionDatum, len(bPSupplierPartnerFunctionArray))

	for i, v := range order.HeaderPartner {
		inoutHeaderPartnerMap[v.BusinessPartner] = order.HeaderPartner[i]
	}

	for i, v := range bPSupplierPartnerFunctionArray {
		bPSupplierPartnerFunctionArrayMap[v.PartnerFunctionBusinessPartner.Int] = *bPSupplierPartnerFunctionArray[i]
	}

	for businessPartnerKey, bPSupplierPartnerFunctionRecord := range bPSupplierPartnerFunctionArrayMap {
		if _, ok := inoutHeaderPartnerMap[businessPartnerKey]; !ok {
			inoutHeaderPartnerMap[businessPartnerKey] = api_input_reader.HeaderPartner{}
		}

		newHeaderPartner := inoutHeaderPartnerMap[businessPartnerKey]

		newHeaderPartner.PartnerFunction = bPSupplierPartnerFunctionRecord.PartnerFunction.String
		newHeaderPartner.BusinessPartner = bPSupplierPartnerFunctionRecord.PartnerFunctionBusinessPartner.Int

		headerPartners[businessPartnerKey] = newHeaderPartner
	}

	res := make([]api_input_reader.HeaderPartner, 0, len(headerPartners))
	for i := range headerPartners {
		res = append(res, headerPartners[i])
	}

	order.HeaderPartner = res

	return order
}

func (f *SubFunction) SetBusinessPartnerGeneralArray(
	order *api_input_reader.Order,
	bPGeneralArray models.DataPlatformBusinessPartnerGeneralDatumSlice,
) *api_input_reader.Order {
	headerPartners := make(map[int]api_input_reader.HeaderPartner, len(order.HeaderPartner))
	inoutHeaderPartnerMap := make(map[int]api_input_reader.HeaderPartner, len(order.HeaderPartner))
	bPGeneralArrayMap := make(map[int]models.DataPlatformBusinessPartnerGeneralDatum, len(bPGeneralArray))

	for i, v := range order.HeaderPartner {
		inoutHeaderPartnerMap[v.BusinessPartner] = order.HeaderPartner[i]
	}
	for i, v := range bPGeneralArray {
		bPGeneralArrayMap[v.BusinessPartner] = *bPGeneralArray[i]
	}

	for businessPartnerKey := range inoutHeaderPartnerMap {

		newHeaderPartner := inoutHeaderPartnerMap[businessPartnerKey]

		bPGeneralRecord, ok := bPGeneralArrayMap[businessPartnerKey]
		if ok {
			newHeaderPartner.BusinessPartnerFullName = bPGeneralRecord.BusinessPartnerFullName.String
			newHeaderPartner.BusinessPartnerName = bPGeneralRecord.BusinessPartnerName
			newHeaderPartner.Country = bPGeneralRecord.Country
			newHeaderPartner.Language = bPGeneralRecord.Language
			newHeaderPartner.Currency = bPGeneralRecord.Currency
			newHeaderPartner.AddressID = bPGeneralRecord.AddressID.Ptr()
		}

		headerPartners[businessPartnerKey] = newHeaderPartner
	}

	res := make([]api_input_reader.HeaderPartner, 0, len(headerPartners))
	for i := range headerPartners {
		res = append(res, headerPartners[i])
	}

	order.HeaderPartner = res

	return order
}

// TODO: Buyer, Sellerの判断を中で行なっている。
func (f *SubFunction) SetBusinessPartnerGeneralArray1(
	order *api_input_reader.Order,
	bPCustomerPartnerFunctionArray models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatumSlice,
	bPSupplierPartnerFunctionArray models.DataPlatformBusinessPartnerSupplierPartnerFunctionDatumSlice,
	bPGeneralArray models.DataPlatformBusinessPartnerGeneralDatumSlice,
	buyerOrSeller string,
) *api_input_reader.Order {
	headerPartners := make(map[int]api_input_reader.HeaderPartner, len(bPSupplierPartnerFunctionArray))
	inoutHeaderPartnerMap := make(map[int]api_input_reader.HeaderPartner, len(order.HeaderPartner))
	bPCustomerPartnerFunctionArrayMap := make(map[int]models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatum, len(bPCustomerPartnerFunctionArray))
	bPSupplierPartnerFunctionArrayMap := make(map[int]models.DataPlatformBusinessPartnerSupplierPartnerFunctionDatum, len(bPSupplierPartnerFunctionArray))
	bPGeneralArrayMap := make(map[int]models.DataPlatformBusinessPartnerGeneralDatum, len(bPGeneralArray))

	for i, v := range order.HeaderPartner {
		inoutHeaderPartnerMap[v.BusinessPartner] = order.HeaderPartner[i]
	}

	for i, v := range bPGeneralArray {
		bPGeneralArrayMap[v.BusinessPartner] = *bPGeneralArray[i]
	}

	if buyerOrSeller == "Seller" {
		for i, v := range bPCustomerPartnerFunctionArray {
			bPCustomerPartnerFunctionArrayMap[v.PartnerFunctionBusinessPartner.Int] = *bPCustomerPartnerFunctionArray[i]
		}

		for businessPartnerKey, bPCustomerPartnerFunctionRecord := range bPCustomerPartnerFunctionArrayMap {
			bPGeneralRecord := bPGeneralArrayMap[bPCustomerPartnerFunctionRecord.PartnerFunctionBusinessPartner.Int]
			if _, ok := inoutHeaderPartnerMap[businessPartnerKey]; !ok {
				inoutHeaderPartnerMap[businessPartnerKey] = api_input_reader.HeaderPartner{}
			}

			newHeaderPartner := inoutHeaderPartnerMap[businessPartnerKey]

			newHeaderPartner.BusinessPartnerFullName = bPGeneralRecord.BusinessPartnerFullName.String
			newHeaderPartner.BusinessPartnerName = bPGeneralRecord.BusinessPartnerName
			newHeaderPartner.Country = bPGeneralRecord.Country
			newHeaderPartner.Language = bPGeneralRecord.Language
			newHeaderPartner.Currency = bPGeneralRecord.Currency
			newHeaderPartner.AddressID = bPGeneralRecord.AddressID.Ptr()

			headerPartners[businessPartnerKey] = newHeaderPartner
		}

	} else if buyerOrSeller == "Buyer" {
		for i, v := range bPSupplierPartnerFunctionArray {
			bPSupplierPartnerFunctionArrayMap[v.PartnerFunctionBusinessPartner.Int] = *bPSupplierPartnerFunctionArray[i]
		}

		for businessPartnerKey, bPSupplierPartnerFunctionRecord := range bPSupplierPartnerFunctionArrayMap {
			bPGeneralRecord := bPGeneralArrayMap[bPSupplierPartnerFunctionRecord.PartnerFunctionBusinessPartner.Int]
			if _, ok := inoutHeaderPartnerMap[businessPartnerKey]; !ok {
				inoutHeaderPartnerMap[businessPartnerKey] = api_input_reader.HeaderPartner{}
			}

			newHeaderPartner := inoutHeaderPartnerMap[businessPartnerKey]

			newHeaderPartner.BusinessPartnerFullName = bPGeneralRecord.BusinessPartnerFullName.String
			newHeaderPartner.BusinessPartnerName = bPGeneralRecord.BusinessPartnerName
			newHeaderPartner.Country = bPGeneralRecord.Country
			newHeaderPartner.Language = bPGeneralRecord.Language
			newHeaderPartner.Currency = bPGeneralRecord.Currency
			newHeaderPartner.AddressID = bPGeneralRecord.AddressID.Ptr()

			headerPartners[businessPartnerKey] = newHeaderPartner
		}
	}

	res := make([]api_input_reader.HeaderPartner, 0, len(headerPartners))
	for i := range headerPartners {
		res = append(res, headerPartners[i])
	}

	order.HeaderPartner = res

	return order
}
