package api_processing_data_formatter

import (
	"data-platform-api-orders-creates-subfunc-rmq-kube/database/models"
)

func ConvertToHeaderRelatedData(
	headerRelatedData *HeaderRelatedData,
	bPCustomerRecord *models.DataPlatformBusinessPartnerCustomerDatum,
	nRLatestNumberRecord *models.DataPlatformNumberRangeLatestNumberDatum,
	bPCustomerPartnerFunctionArray models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatumSlice,
	bPCustomerPartnerPlantArray models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice,
) *HeaderRelatedData {
	// headerRelatedData.LatestNumber = &nRLatestNumberRecord.LatestNumber.Int
	headerRelatedData.HeaderPartnerRelatedData = ConvertToHeaderPartnerRelatedData(bPCustomerPartnerFunctionArray, bPCustomerPartnerPlantArray)
	return headerRelatedData
}

func ConvertToHeaderPartnerRelatedData(
	bPCustomerPartnerFunctionArray models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatumSlice,
	bPCustomerPartnerPlantArray models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice,
) []HeaderPartnerRelatedData {
	headerPartnerRelatedData := make([]HeaderPartnerRelatedData, 0, len(bPCustomerPartnerFunctionArray))
	// for i, bPCustomerPartnerFunctionRecord := range bPCustomerPartnerFunctionArray {
	// 	headerPartnerRelatedData = append(headerPartnerRelatedData, HeaderPartnerRelatedData{
	// 		PartnerFunction: ConvertToCustomerPartnerFunction(bPCustomerPartnerFunctionRecord),
	// 		PartnerPlant:    ConvertToCustomerPartnerPlant(bPCustomerPartnerPlantArray[i]),
	// 	})
	// }
	return headerPartnerRelatedData
}

func ConvertToCustomerPartnerFunction(
	bPCustomerPartnerFunctionRecord *models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatum,
) PartnerFunction {
	customerPartnerFunctionData := PartnerFunction{
		PartnerCounter: &bPCustomerPartnerFunctionRecord.PartnerCounter,
		DefaultPartner: bPCustomerPartnerFunctionRecord.DefaultPartner.Ptr(),
	}
	return customerPartnerFunctionData
}

func ConvertToCustomerPartnerPlant(
	bPCustomerPartnerPlantRecord *models.DataPlatformBusinessPartnerCustomerPartnerPlantDatum,
) PartnerPlant {
	customerPartnerPlantData := PartnerPlant{
		PlantCounter: &bPCustomerPartnerPlantRecord.PlantCounter,
		DefaultPlant: &bPCustomerPartnerPlantRecord.DefaultPlant.Bool,
	}
	return customerPartnerPlantData
}
