package subfunction

import (
	api_input_reader "data-platform-api-orders-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"data-platform-api-orders-creates-subfunc-rmq-kube/database/models"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (f *SubFunction) ExtractBusinessPartnerCustomerPartnerPlantArray(
	businessPartner *int,
	buyer *int,
	bPCustomerPartnerFunctionArray models.DataPlatformBusinessPartnerCustomerPartnerFunctionDatumSlice,
) (models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice, error) {
	where := make([]qm.QueryMod, 0, len(bPCustomerPartnerFunctionArray))
	for i := range bPCustomerPartnerFunctionArray {
		where = append(where,
			qm.Or(
				fmt.Sprintf("(`BusinessPartner`, `Customer`, `PartnerCounter`, `PartnerFunction`, `PartnerFunctionBusinessPartner`) = (%d, %d, %d,'%s',%d)", *businessPartner, *buyer, bPCustomerPartnerFunctionArray[i].PartnerCounter, bPCustomerPartnerFunctionArray[i].PartnerFunction.String, bPCustomerPartnerFunctionArray[i].PartnerFunctionBusinessPartner.Int),
			),
		)
	}

	res, err := models.DataPlatformBusinessPartnerCustomerPartnerPlantData(
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

func (f *SubFunction) ExtractBusinessPartnerSupplierPartnerPlantArray(
	businessPartner *int,
	seller *int,
	bPSupplierPartnerFunctionArray models.DataPlatformBusinessPartnerSupplierPartnerFunctionDatumSlice,
) (models.DataPlatformBusinessPartnerSupplierPartnerPlantDatumSlice, error) {
	where := make([]qm.QueryMod, 0, len(bPSupplierPartnerFunctionArray))
	for i := range bPSupplierPartnerFunctionArray {
		where = append(where,
			qm.Or(
				fmt.Sprintf("(`BusinessPartner`, `Supplier`, `PartnerCounter`, `PartnerFunction`, `PartnerFunctionBusinessPartner`) = (%d, %d, %d,'%s',%d)", *businessPartner, *seller, bPSupplierPartnerFunctionArray[i].PartnerCounter, bPSupplierPartnerFunctionArray[i].PartnerFunction.String, bPSupplierPartnerFunctionArray[i].PartnerFunctionBusinessPartner.Int),
			),
		)
	}

	res, err := models.DataPlatformBusinessPartnerSupplierPartnerPlantData(
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

func (f *SubFunction) HoldBusinessPartnerCustomerPartnerPlantArray(
	processingData *api_processing_data_formatter.HeaderRelatedData,
	bPCustomerPartnerPlantArray models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice,
) *api_processing_data_formatter.HeaderRelatedData {
	headerPartnerRelatedDataMap := make(map[int]api_processing_data_formatter.HeaderPartnerRelatedData, len(processingData.HeaderPartnerRelatedData))
	bPCustomerPartnerPlantArrayMap := make(map[int]models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice, len(bPCustomerPartnerPlantArray))

	for i, v := range processingData.HeaderPartnerRelatedData {
		headerPartnerRelatedDataMap[*v.PartnerFunction.BusinessPartner] = processingData.HeaderPartnerRelatedData[i]
	}

	for i, v := range bPCustomerPartnerPlantArray {
		bPCustomerPartnerPlantArrayMap[v.PartnerFunctionBusinessPartner] = append(bPCustomerPartnerPlantArrayMap[v.PartnerFunctionBusinessPartner], bPCustomerPartnerPlantArray[i])
	}

	for businessPartnerKey := range headerPartnerRelatedDataMap {

		newHeaderPartnerRelatedData := headerPartnerRelatedDataMap[businessPartnerKey]

		bPCustomerPartnerPlantArray, ok := bPCustomerPartnerPlantArrayMap[businessPartnerKey]
		if ok {
			for _, bPCustomerPartnerPlantRecord := range bPCustomerPartnerPlantArray {
				newHeaderPartnerRelatedData.PartnerPlant = append(newHeaderPartnerRelatedData.PartnerPlant, api_processing_data_formatter.PartnerPlant{
					PlantCounter: &bPCustomerPartnerPlantRecord.PlantCounter,
					DefaultPlant: &bPCustomerPartnerPlantRecord.DefaultPlant.Bool,
				})
			}
		}
		headerPartnerRelatedDataMap[businessPartnerKey] = newHeaderPartnerRelatedData
	}

	res := make([]api_processing_data_formatter.HeaderPartnerRelatedData, 0, len(headerPartnerRelatedDataMap))
	for i := range headerPartnerRelatedDataMap {
		res = append(res, headerPartnerRelatedDataMap[i])
	}

	processingData.HeaderPartnerRelatedData = res

	return processingData
}

func (f *SubFunction) HoldBusinessPartnerSupplierPartnerPlantArray(
	processingData *api_processing_data_formatter.HeaderRelatedData,
	bPSupplierPartnerPlantArray models.DataPlatformBusinessPartnerSupplierPartnerPlantDatumSlice,
) *api_processing_data_formatter.HeaderRelatedData {
	headerPartnerRelatedDataMap := make(map[int]api_processing_data_formatter.HeaderPartnerRelatedData, len(processingData.HeaderPartnerRelatedData))
	bPSupplierPartnerPlantArrayMap := make(map[int]models.DataPlatformBusinessPartnerSupplierPartnerPlantDatumSlice, len(bPSupplierPartnerPlantArray))

	for i, v := range processingData.HeaderPartnerRelatedData {
		headerPartnerRelatedDataMap[*v.PartnerFunction.BusinessPartner] = processingData.HeaderPartnerRelatedData[i]
	}

	for i, v := range bPSupplierPartnerPlantArray {
		bPSupplierPartnerPlantArrayMap[v.PartnerFunctionBusinessPartner] = append(bPSupplierPartnerPlantArrayMap[v.PartnerFunctionBusinessPartner], bPSupplierPartnerPlantArray[i])
	}

	for businessPartnerKey := range headerPartnerRelatedDataMap {

		newHeaderPartnerRelatedData := headerPartnerRelatedDataMap[businessPartnerKey]

		bPSupplierPartnerPlantArray, ok := bPSupplierPartnerPlantArrayMap[businessPartnerKey]
		if ok {
			for _, bPSupplierPartnerPlantRecord := range bPSupplierPartnerPlantArray {
				newHeaderPartnerRelatedData.PartnerPlant = append(newHeaderPartnerRelatedData.PartnerPlant, api_processing_data_formatter.PartnerPlant{
					PlantCounter: &bPSupplierPartnerPlantRecord.PlantCounter,
					DefaultPlant: &bPSupplierPartnerPlantRecord.DefaultPlant.Bool,
				})
			}
		}
		headerPartnerRelatedDataMap[businessPartnerKey] = newHeaderPartnerRelatedData
	}

	res := make([]api_processing_data_formatter.HeaderPartnerRelatedData, 0, len(headerPartnerRelatedDataMap))
	for i := range headerPartnerRelatedDataMap {
		res = append(res, headerPartnerRelatedDataMap[i])
	}

	processingData.HeaderPartnerRelatedData = res

	return processingData
}

func (f *SubFunction) SetBusinessPartnerCustomerPartnerPlantArray(
	order *api_input_reader.Order,
	bPCustomerPartnerPlantArray models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice,
) *api_input_reader.Order {
	headerPartners := make(map[int]api_input_reader.HeaderPartner, len(order.HeaderPartner))
	inoutHeaderPartnerMap := make(map[int]api_input_reader.HeaderPartner, len(order.HeaderPartner))
	bPCustomerPartnerPlantArrayMap := make(map[int]models.DataPlatformBusinessPartnerCustomerPartnerPlantDatumSlice, len(bPCustomerPartnerPlantArray))

	for i, v := range order.HeaderPartner {
		inoutHeaderPartnerMap[v.BusinessPartner] = order.HeaderPartner[i]
	}

	for i, v := range bPCustomerPartnerPlantArray {
		bPCustomerPartnerPlantArrayMap[v.PartnerFunctionBusinessPartner] = append(bPCustomerPartnerPlantArrayMap[v.PartnerFunctionBusinessPartner], bPCustomerPartnerPlantArray[i])
	}

	for businessPartnerKey := range inoutHeaderPartnerMap {

		newHeaderPartner := inoutHeaderPartnerMap[businessPartnerKey]

		bPCustomerPartnerPlantArray, ok := bPCustomerPartnerPlantArrayMap[businessPartnerKey]
		if ok {
			for i, v := range newHeaderPartner.HeaderPartnerPlant {
				if v.Plant != "" {
					break
				}
				if i == len(newHeaderPartner.HeaderPartnerPlant)-1 {
					newHeaderPartner.HeaderPartnerPlant = nil
				}
			}
			for _, bPCustomerPartnerPlantRecord := range bPCustomerPartnerPlantArray {
				newHeaderPartner.HeaderPartnerPlant = append(newHeaderPartner.HeaderPartnerPlant, api_input_reader.HeaderPartnerPlant{
					Plant: bPCustomerPartnerPlantRecord.Plant.String,
				})
			}
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

func (f *SubFunction) SetBusinessPartnerSupplierPartnerPlantArray(
	order *api_input_reader.Order,
	bPSupplierPartnerPlantArray models.DataPlatformBusinessPartnerSupplierPartnerPlantDatumSlice,
) *api_input_reader.Order {
	headerPartners := make(map[int]api_input_reader.HeaderPartner, len(order.HeaderPartner))
	inoutHeaderPartnerMap := make(map[int]api_input_reader.HeaderPartner, len(order.HeaderPartner))
	bPSupplierPartnerPlantArrayMap := make(map[int]models.DataPlatformBusinessPartnerSupplierPartnerPlantDatumSlice, len(bPSupplierPartnerPlantArray))

	for i, v := range order.HeaderPartner {
		inoutHeaderPartnerMap[v.BusinessPartner] = order.HeaderPartner[i]
	}

	for i, v := range bPSupplierPartnerPlantArray {
		bPSupplierPartnerPlantArrayMap[v.PartnerFunctionBusinessPartner] = append(bPSupplierPartnerPlantArrayMap[v.PartnerFunctionBusinessPartner], bPSupplierPartnerPlantArray[i])
	}

	for businessPartnerKey := range inoutHeaderPartnerMap {

		newHeaderPartner := inoutHeaderPartnerMap[businessPartnerKey]

		bPSupplierPartnerPlantArray, ok := bPSupplierPartnerPlantArrayMap[businessPartnerKey]
		if ok {
			for i, v := range newHeaderPartner.HeaderPartnerPlant {
				if v.Plant != "" {
					break
				}
				if i == len(newHeaderPartner.HeaderPartnerPlant)-1 {
					newHeaderPartner.HeaderPartnerPlant = []api_input_reader.HeaderPartnerPlant{}

				}
			}
			for _, bPSupplierPartnerPlantRecord := range bPSupplierPartnerPlantArray {
				newHeaderPartner.HeaderPartnerPlant = append(newHeaderPartner.HeaderPartnerPlant, api_input_reader.HeaderPartnerPlant{
					Plant: bPSupplierPartnerPlantRecord.Plant.String,
				})
			}
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
