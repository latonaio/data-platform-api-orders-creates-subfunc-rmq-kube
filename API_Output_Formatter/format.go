package api_output_formatter

import (
	api_input_reader "data-platform-api-orders-creates-subfunc-rmq-kube/API_Input_Reader"
	"data-platform-api-orders-creates-subfunc-rmq-kube/database/models"
)

func ConvertToHeader(
	order *api_input_reader.Order,
	bPCustomerRecord *models.DataPlatformBusinessPartnerCustomerDatum,
	nRLatestNumberRecord *models.DataPlatformNumberRangeLatestNumberDatum,
	bPCustomerPartnerFunctionArray models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatumSlice,
	bPGeneralArray models.DataPlatformBusinessPartnerGeneralDatumSlice,
	bPCustomerPartnerPlantArray models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice,
) *api_input_reader.Order {
	// order.OrderID = CalculateOrderId(nRLatestNumberRecord.LatestNumber.Int)
	// order.Incoterms = bPCustomerRecord.Incoterms.String
	// order.PaymentTerms = bPCustomerRecord.PaymentTerms.String
	// order.PaymentMethod = bPCustomerRecord.PaymentMethod.String
	// order.BPAccountAssignmentGroup = bPCustomerRecord.BPAccountAssignmentGroup.String
	order.HeaderPartner = ConvertToHeaderPartner(order.HeaderPartner, bPCustomerPartnerFunctionArray, bPGeneralArray, bPCustomerPartnerPlantArray)

	return order
}

func ConvertToHeaderPartner(
	inoutHeaderPartner []api_input_reader.HeaderPartner,
	bPCustomerPartnerFunctionArray models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatumSlice,
	bPGeneralArray models.DataPlatformBusinessPartnerGeneralDatumSlice,
	bPCustomerPartnerPlantArray models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice,
) []api_input_reader.HeaderPartner {
	headerPartners := make(map[int]api_input_reader.HeaderPartner, len(bPCustomerPartnerFunctionArray))
	inoutHeaderPartnerMap := make(map[int]api_input_reader.HeaderPartner, len(inoutHeaderPartner))
	bPCustomerPartnerFunctionArrayMap := make(map[int]models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatum, len(bPCustomerPartnerFunctionArray))
	bPGeneralArrayMap := make(map[int]models.DataPlatformBusinessPartnerGeneralDatum, len(bPGeneralArray))
	bPCustomerPartnerPlantArrayMap := make(map[int]models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice, len(bPCustomerPartnerPlantArray))

	for i, v := range inoutHeaderPartner {
		inoutHeaderPartnerMap[v.BusinessPartner] = inoutHeaderPartner[i]
	}

	for i, v := range bPCustomerPartnerFunctionArray {
		bPCustomerPartnerFunctionArrayMap[v.PartnerFunctionBusinessPartner.Int] = *bPCustomerPartnerFunctionArray[i]
	}

	for i, v := range bPGeneralArray {
		bPGeneralArrayMap[v.BusinessPartner] = *bPGeneralArray[i]
	}

	for i, v := range bPCustomerPartnerPlantArray {
		bPCustomerPartnerPlantArrayMap[v.PartnerFunctionBusinessPartner] = append(bPCustomerPartnerPlantArrayMap[v.PartnerFunctionBusinessPartner], bPCustomerPartnerPlantArray[i])
	}

	for businessPartnerID, _ := range bPCustomerPartnerFunctionArrayMap {
		// bPGeneralRecord := bPGeneralArrayMap[bPCustomerPartnerFunctionRecord.PartnerFunctionBusinessPartner.Int]

		if _, ok := inoutHeaderPartnerMap[businessPartnerID]; !ok {
			inoutHeaderPartnerMap[businessPartnerID] = api_input_reader.HeaderPartner{}
		}

		newHeaderPartner := inoutHeaderPartnerMap[businessPartnerID]

		// newHeaderPartner.PartnerFunction = bPCustomerPartnerFunctionRecord.PartnerFunction.String
		// newHeaderPartner.BusinessPartner = bPCustomerPartnerFunctionRecord.PartnerFunctionBusinessPartner.Int
		// newHeaderPartner.BusinessPartnerFullName = bPGeneralRecord.BusinessPartnerFullName.String
		// newHeaderPartner.BusinessPartnerName = bPGeneralRecord.BusinessPartnerName
		// newHeaderPartner.Country = bPGeneralRecord.Country
		// newHeaderPartner.Language = bPGeneralRecord.Language
		// newHeaderPartner.Currency = bPGeneralRecord.Currency
		// newHeaderPartner.AddressID = bPGeneralRecord.AddressID.Ptr()

		// bPCustomerPartnerPlantArray, ok := bPCustomerPartnerPlantArrayMap[businessPartnerID]
		// if ok {
		// 	for i, v := range newHeaderPartner.HeaderPartnerPlant {
		// 		if v.Plant != "" {
		// 			break
		// 		}
		// 		if i == len(newHeaderPartner.HeaderPartnerPlant)-1 {
		// 			newHeaderPartner.HeaderPartnerPlant = nil
		// 		}
		// 	}
		// 	for _, bPCustomerPartnerPlantRecord := range bPCustomerPartnerPlantArray {
		// 		newHeaderPartner.HeaderPartnerPlant = append(newHeaderPartner.HeaderPartnerPlant, api_input_reader.HeaderPartnerPlant{
		// 			Plant: bPCustomerPartnerPlantRecord.Plant.String,
		// 		})
		// 	}
		// }

		headerPartners[businessPartnerID] = newHeaderPartner
	}

	res := make([]api_input_reader.HeaderPartner, 0, len(headerPartners))
	for i := range headerPartners {
		res = append(res, headerPartners[i])
	}

	return res
}

func ConvertToItem(
	order *api_input_reader.Order,
	bPCustomerTaxRecord *models.DataPlatformBusinessPartnerCustomerTaxDatum,
	pMGeneralArray models.DataPlatformProductMasterGeneralDatumSlice,
	pMProductDescriptionArray models.DataPlatformProductMasterProductDescriptionDatumSlice,
) *api_input_reader.Order {
	inoutItem := order.Item
	items := make(map[string]api_input_reader.Item, len(pMGeneralArray))
	inoutItemMap := make(map[string]api_input_reader.Item, len(inoutItem))
	pMGeneralArrayMap := make(map[string]models.DataPlatformProductMasterGeneralDatum, len(pMGeneralArray))
	pMProductDescriptionArrayMap := make(map[string]models.DataPlatformProductMasterProductDescriptionDatum, len(pMProductDescriptionArray))
	var orderItems []int

	for i, v := range inoutItem {
		inoutItemMap[v.Product] = inoutItem[i]
		orderItems = append(orderItems, i+1)
	}

	for i, v := range pMGeneralArray {
		pMGeneralArrayMap[v.Product] = *pMGeneralArray[i]
	}

	for i, v := range pMProductDescriptionArray {
		pMProductDescriptionArrayMap[v.Product] = *pMProductDescriptionArray[i]
	}

	i := 0
	for product, _ := range pMGeneralArrayMap {
		if _, ok := inoutItemMap[product]; !ok {
			inoutItemMap[product] = api_input_reader.Item{}
		}

		newItem := inoutItemMap[product]

		// newItem.OrderItem = &orderItems[i]

		// newItem.BPTaxClassification = bPCustomerTaxRecord.BPTaxClassification.String

		// newItem.ProductStandardID = pMGeneralRecord.ProductStandardID.String
		// newItem.ProductGroup = pMGeneralRecord.ProductGroup.String
		// newItem.BaseUnit = pMGeneralRecord.BaseUnit.String
		// newItem.ItemWeightUnit = pMGeneralRecord.WeightUnit.String
		// newItem.ProductGrossWeight = pMGeneralRecord.GrossWeight.Ptr()
		// newItem.ProductNetWeight = pMGeneralRecord.NetWeight.Ptr()
		// newItem.ProductAccountAssignmentGroup = pMGeneralRecord.ProductAccountAssignmentGroup.String
		// newItem.CountryOfOrigin = pMGeneralRecord.CountryOfOrigin.String

		// pMProductDescriptionRecord, ok := pMProductDescriptionArrayMap[product]
		// if ok {
		// 	newItem.OrderItemText = pMProductDescriptionRecord.ProductDescription.String
		// }

		items[product] = newItem
		i++
	}

	res := make([]api_input_reader.Item, 0, len(items))
	for i := range items {
		res = append(res, items[i])
	}

	order.Item = res

	return order
}

func CalculateOrderId(latestNumber int) *int {
	orderId := latestNumber + 1
	return &orderId
}
