package subfunction

import (
	api_input_reader "data-platform-api-orders-creates-subfunc-rmq-kube/API_Input_Reader"
	"data-platform-api-orders-creates-subfunc-rmq-kube/database/models"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (f *SubFunction) ExtractBusinessPartnerCustomerTaxRecord(
	businessPartner *int,
	buyer *int,
	country string,
) (*models.DataPlatformBusinessPartnerCustomerTaxDatum, error) {
	res, err := models.DataPlatformBusinessPartnerCustomerTaxData(
		qm.And("BusinessPartner=?", *businessPartner),
		qm.And("Customer=?", *buyer),
		qm.And("DepartureCountry=?", country),
	).One(f.ctx, f.db)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (f *SubFunction) ExtractBusinessPartnerSupplierTaxRecord(
	businessPartner *int,
	seller *int,
	country string,
) (*models.DataPlatformBusinessPartnerSupplierTaxDatum, error) {
	res, err := models.DataPlatformBusinessPartnerSupplierTaxData(
		qm.And("BusinessPartner=?", *businessPartner),
		qm.And("Supplier=?", *seller),
		qm.And("DepartureCountry=?", country),
	).One(f.ctx, f.db)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (f *SubFunction) ExtractProductMasterGeneralArray(item []api_input_reader.Item) (models.DataPlatformProductMasterGeneralDatumSlice, error) {
	var res models.DataPlatformProductMasterGeneralDatumSlice
	for _, v := range item {
		tmp, err := models.DataPlatformProductMasterGeneralData(
			qm.Where("Product=?", v.Product),
		).One(f.ctx, f.db)
		if err != nil {
			return nil, err
		}
		res = append(res, tmp)
	}
	return res, nil
}

func (f *SubFunction) ExtractProductMasterProductDescriptionArray(
	item []api_input_reader.Item,
	language string,
) (models.DataPlatformProductMasterProductDescriptionDatumSlice, error) {
	var res models.DataPlatformProductMasterProductDescriptionDatumSlice

	where := make([]qm.QueryMod, 0, len(item))
	for _, v := range item {
		where = append(where,
			qm.Or("(`Product`,`Language`)=(?, ?)", v.Product, language),
		)
	}

	res, err := models.DataPlatformProductMasterProductDescriptionData(
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

func (f *SubFunction) SetOrderItem(
	order *api_input_reader.Order,
) *api_input_reader.Order {
	inoutItem := order.Item
	items := make(map[string]api_input_reader.Item, len(inoutItem))
	inoutItemMap := make(map[string]api_input_reader.Item, len(inoutItem))
	var orderItems []int
	for i, v := range inoutItem {
		inoutItemMap[v.Product] = inoutItem[i]
		orderItems = append(orderItems, i+1)
	}

	i := 0
	for productKey := range inoutItemMap {
		newItem := inoutItemMap[productKey]

		newItem.OrderItem = &orderItems[i]

		items[productKey] = newItem
		i++
	}

	res := make([]api_input_reader.Item, 0, len(items))
	for i := range items {
		res = append(res, items[i])
	}

	order.Item = res

	return order
}

func (f *SubFunction) SetBusinessPartnerCustomerTaxRecord(
	order *api_input_reader.Order,
	bPCustomerTaxRecord *models.DataPlatformBusinessPartnerCustomerTaxDatum,
) *api_input_reader.Order {
	inoutItem := order.Item
	items := make(map[string]api_input_reader.Item, len(inoutItem))
	inoutItemMap := make(map[string]api_input_reader.Item, len(inoutItem))

	for i, v := range inoutItem {
		inoutItemMap[v.Product] = inoutItem[i]
	}

	for productKey := range inoutItemMap {

		newItem := inoutItemMap[productKey]

		newItem.BPTaxClassification = bPCustomerTaxRecord.BPTaxClassification.String

		items[productKey] = newItem
	}

	res := make([]api_input_reader.Item, 0, len(items))
	for i := range items {
		res = append(res, items[i])
	}

	order.Item = res

	return order
}

func (f *SubFunction) SetBusinessPartnerSupplierTaxRecord(
	order *api_input_reader.Order,
	bPSupplierTaxRecord *models.DataPlatformBusinessPartnerSupplierTaxDatum,
) *api_input_reader.Order {
	inoutItem := order.Item
	items := make(map[string]api_input_reader.Item, len(inoutItem))
	inoutItemMap := make(map[string]api_input_reader.Item, len(inoutItem))

	for i, v := range inoutItem {
		inoutItemMap[v.Product] = inoutItem[i]
	}

	for productKey := range inoutItemMap {

		newItem := inoutItemMap[productKey]

		newItem.BPTaxClassification = bPSupplierTaxRecord.BPTaxClassification.String

		items[productKey] = newItem
	}

	res := make([]api_input_reader.Item, 0, len(items))
	for i := range items {
		res = append(res, items[i])
	}

	order.Item = res

	return order
}

func (f *SubFunction) SetProductMasterGeneralArray(
	order *api_input_reader.Order,
	pMGeneralArray models.DataPlatformProductMasterGeneralDatumSlice,
) *api_input_reader.Order {
	inoutItem := order.Item
	items := make(map[string]api_input_reader.Item, len(inoutItem))
	inoutItemMap := make(map[string]api_input_reader.Item, len(inoutItem))
	pMGeneralArrayMap := make(map[string]models.DataPlatformProductMasterGeneralDatum, len(pMGeneralArray))

	for i, v := range inoutItem {
		inoutItemMap[v.Product] = inoutItem[i]
	}

	for i, v := range pMGeneralArray {
		pMGeneralArrayMap[v.Product] = *pMGeneralArray[i]
	}

	for productKey, pMGeneralRecord := range pMGeneralArrayMap {
		if _, ok := inoutItemMap[productKey]; !ok {
			inoutItemMap[productKey] = api_input_reader.Item{}
		}

		newItem := inoutItemMap[productKey]

		newItem.ProductStandardID = pMGeneralRecord.ProductStandardID.String
		newItem.ProductGroup = pMGeneralRecord.ProductGroup.String
		newItem.BaseUnit = pMGeneralRecord.BaseUnit.String
		newItem.ItemWeightUnit = pMGeneralRecord.WeightUnit.String
		newItem.ProductGrossWeight = pMGeneralRecord.GrossWeight.Ptr()
		newItem.ProductNetWeight = pMGeneralRecord.NetWeight.Ptr()
		newItem.ProductAccountAssignmentGroup = pMGeneralRecord.ProductAccountAssignmentGroup.String
		newItem.CountryOfOrigin = pMGeneralRecord.CountryOfOrigin.String

		items[productKey] = newItem
	}

	res := make([]api_input_reader.Item, 0, len(items))
	for i := range items {
		res = append(res, items[i])
	}

	order.Item = res

	return order
}

func (f *SubFunction) SetProductMasterProductDescriptionArray(
	order *api_input_reader.Order,
	pMProductDescriptionArray models.DataPlatformProductMasterProductDescriptionDatumSlice,
) *api_input_reader.Order {
	inoutItem := order.Item
	items := make(map[string]api_input_reader.Item, len(inoutItem))
	inoutItemMap := make(map[string]api_input_reader.Item, len(inoutItem))
	pMProductDescriptionArrayMap := make(map[string]models.DataPlatformProductMasterProductDescriptionDatum, len(pMProductDescriptionArray))

	for i, v := range inoutItem {
		inoutItemMap[v.Product] = inoutItem[i]
	}

	for i, v := range pMProductDescriptionArray {
		pMProductDescriptionArrayMap[v.Product] = *pMProductDescriptionArray[i]
	}

	for productKey := range inoutItemMap {

		newItem := inoutItemMap[productKey]

		pMProductDescriptionRecord, ok := pMProductDescriptionArrayMap[productKey]
		if ok {
			newItem.OrderItemText = pMProductDescriptionRecord.ProductDescription.String
		}

		items[productKey] = newItem
	}

	res := make([]api_input_reader.Item, 0, len(items))
	for i := range items {
		res = append(res, items[i])
	}

	order.Item = res

	return order
}
