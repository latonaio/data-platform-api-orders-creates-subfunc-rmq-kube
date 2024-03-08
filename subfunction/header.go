package subfunction

import (
	api_input_reader "data-platform-api-orders-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"data-platform-api-orders-creates-subfunc-rmq-kube/database/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (f *SubFunction) ExtractBusinessPartnerCustomerRecord(businessPartner *int, buyer *int) (*models.DataPlatformBusinessPartnerCustomerDatum, error) {
	res, err := models.DataPlatformBusinessPartnerCustomerData(
		qm.And("BusinessPartner=?", *businessPartner),
		qm.And("Customer=?", *buyer),
	).One(f.ctx, f.db)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (f *SubFunction) ExtractBusinessPartnerSupplierRecord(businessPartner *int, seller *int) (*models.DataPlatformBusinessPartnerSupplierDatum, error) {
	res, err := models.DataPlatformBusinessPartnerSupplierData(
		qm.And("BusinessPartner=?", *businessPartner),
		qm.And("Supplier=?", *seller),
	).One(f.ctx, f.db)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (f *SubFunction) ExtractNumberRangeLatestNumberRecord(serviceLabel string, property string) (*models.DataPlatformNumberRangeLatestNumberDatum, error) {
	res, err := models.DataPlatformNumberRangeLatestNumberData(
		qm.And("ServiceLabel=?", serviceLabel),
		qm.And("FieldNameWithNumberRange=?", property),
	).One(f.ctx, f.db)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (f *SubFunction) HoldNumberRangeLatestNumberRecord(
	processingData *api_processing_data_formatter.HeaderRelatedData,
	nRLatestNumberRecord *models.DataPlatformNumberRangeLatestNumberDatum,
) *api_processing_data_formatter.HeaderRelatedData {
	processingData.LatestNumber = &nRLatestNumberRecord.LatestNumber.Int
	return processingData
}

func (f *SubFunction) SetBusinessPartnerCustomerRecord(
	order *api_input_reader.Order,
	bPCustomerRecord *models.DataPlatformBusinessPartnerCustomerDatum,
) *api_input_reader.Order {
	order.Incoterms = bPCustomerRecord.Incoterms.String
	order.PaymentTerms = bPCustomerRecord.PaymentTerms.String
	order.PaymentMethod = bPCustomerRecord.PaymentMethod.String
	order.BPAccountAssignmentGroup = bPCustomerRecord.BPAccountAssignmentGroup.String
	return order
}

func (f *SubFunction) SetBusinessPartnerSupplierRecord(
	order *api_input_reader.Order,
	bPSupplierRecord *models.DataPlatformBusinessPartnerSupplierDatum,
) *api_input_reader.Order {
	order.Incoterms = bPSupplierRecord.Incoterms.String
	order.PaymentTerms = bPSupplierRecord.PaymentTerms.String
	order.PaymentMethod = bPSupplierRecord.PaymentMethod.String
	order.BPAccountAssignmentGroup = bPSupplierRecord.BPAccountAssignmentGroup.String
	return order
}

func (f *SubFunction) SetNumberRangeLatestNumberRecord(
	order *api_input_reader.Order,
	latestNumber *int,
) *api_input_reader.Order {
	order.OrderID = CalculateOrderId(*latestNumber)
	return order
}

func CalculateOrderId(latestNumber int) *int {
	orderId := latestNumber + 1
	return &orderId
}
