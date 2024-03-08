package api_processing_data_formatter

type HeaderRelatedData struct {
	BuyerOrSeller            string                     `json:"BuyerOrSeller"`
	LatestNumber             *int                       `json:"LatestNumber"`
	HeaderPartnerRelatedData []HeaderPartnerRelatedData `json:"HeaderPartnerRelatedData"`
}

type HeaderPartnerRelatedData struct {
	PartnerFunction PartnerFunction `json:"PartnerFunction"`
	PartnerPlant    []PartnerPlant  `json:"PartnerPlant"`
}

type PartnerFunction struct {
	BusinessPartner *int  `json:"BusinessPartner"`
	PartnerCounter  *int  `json:"PartnerCounter"`
	DefaultPartner  *bool `json:"DefaultPartner"`
}

type PartnerPlant struct {
	PlantCounter *int  `json:"PlantCounter"`
	DefaultPlant *bool `json:"DefaultPlant"`
}
