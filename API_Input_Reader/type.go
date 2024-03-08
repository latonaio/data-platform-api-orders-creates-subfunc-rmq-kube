package api_input_reader

type SDC struct {
	ConnectionKey    string       `json:"connection_key"`
	Result           bool         `json:"result"`
	RedisKey         string       `json:"redis_key"`
	Filepath         string       `json:"filepath"`
	RuntimeSessionID string       `json:"runtime_session_id"`
	BusinessPartner  *int         `json:"business_partner"`
	ServiceLabel     string       `json:"service_label"`
	Orders           Order        `json:"Orders"`
	APISchema        string       `json:"api_schema"`
	Accepter         []string     `json:"accepter"`
	OrderID          *interface{} `json:"order_id"`
	Deleted          bool         `json:"deleted"`
}
type HeaderPartnerContact struct {
	ContactID         *interface{} `json:"ContactID"`
	ContactPersonName string       `json:"ContactPersonName"`
	EmailAddress      string       `json:"EmailAddress"`
	PhoneNumber       string       `json:"PhoneNumber"`
	MobilePhoneNumber string       `json:"MobilePhoneNumber"`
	FaxNumber         string       `json:"FaxNumber"`
	ContactTag1       string       `json:"ContactTag1"`
	ContactTag2       string       `json:"ContactTag2"`
	ContactTag3       string       `json:"ContactTag3"`
	ContactTag4       string       `json:"ContactTag4"`
}
type HeaderPartnerPlant struct {
	Plant string `json:"Plant"`
}

type HeaderPartner struct {
	PartnerFunction         string                 `json:"PartnerFunction"`
	BusinessPartner         int                    `json:"BusinessPartner"`
	BusinessPartnerFullName string                 `json:"BusinessPartnerFullName"`
	BusinessPartnerName     string                 `json:"BusinessPartnerName"`
	Organization            string                 `json:"Organization"`
	Country                 string                 `json:"Country"`
	Language                string                 `json:"Language"`
	Currency                string                 `json:"Currency"`
	ExternalDocumentID      string                 `json:"ExternalDocumentID"`
	AddressID               *int                   `json:"AddressID"`
	HeaderPartnerContact    []HeaderPartnerContact `json:"HeaderPartnerContact"`
	HeaderPartnerPlant      []HeaderPartnerPlant   `json:"HeaderPartnerPlant"`
}
type Address struct {
	AddressID   *int         `json:"AddressID"`
	PostalCode  string       `json:"PostalCode"`
	LocalRegion string       `json:"LocalRegion"`
	Country     string       `json:"Country"`
	District    string       `json:"District"`
	StreetName  string       `json:"StreetName"`
	CityName    string       `json:"CityName"`
	Building    string       `json:"Building"`
	Floor       *interface{} `json:"Floor"`
	Room        *interface{} `json:"Room"`
}
type HeaderPDF struct {
	DocType                  string       `json:"DocType"`
	DocVersionID             *interface{} `json:"DocVersionID"`
	DocID                    string       `json:"DocID"`
	DocIssuerBusinessPartner *interface{} `json:"DocIssuerBusinessPartner"`
	FileName                 string       `json:"FileName"`
}
type ItemPartnerPlant struct {
	Plant string `json:"Plant"`
}
type ItemPartner struct {
	PartnerFunction  string           `json:"PartnerFunction"`
	BusinessPartner  *interface{}     `json:"BusinessPartner"`
	ItemPartnerPlant ItemPartnerPlant `json:"ItemPartnerPlant"`
}
type ItemPricingElement struct {
	PricingProcedureStep       *interface{} `json:"PricingProcedureStep"`
	PricingProcedureCounter    *interface{} `json:"PricingProcedureCounter"`
	ConditionType              string       `json:"ConditionType"`
	PricingDate                string       `json:"PricingDate"`
	ConditionRateValue         *interface{} `json:"ConditionRateValue"`
	ConditionCurrency          string       `json:"ConditionCurrency"`
	ConditionQuantity          *interface{} `json:"ConditionQuantity"`
	ConditionQuantityUnit      string       `json:"ConditionQuantityUnit"`
	ConditionRecord            *interface{} `json:"ConditionRecord"`
	ConditionSequentialNumber  *interface{} `json:"ConditionSequentialNumber"`
	TaxCode                    string       `json:"TaxCode"`
	ConditionAmount            *interface{} `json:"ConditionAmount"`
	TransactionCurrency        string       `json:"TransactionCurrency"`
	ConditionIsManuallyChanged *interface{} `json:"ConditionIsManuallyChanged"`
}
type ItemSchedulingLine struct {
	ScheduleLine                                 *interface{} `json:"ScheduleLine"`
	Product                                      string       `json:"Product"`
	StockConfirmationPartnerFunction             string       `json:"StockConfirmationPartnerFunction"`
	StockConfirmationBusinessPartner             *interface{} `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                       string       `json:"StockConfirmationPlant"`
	StockConfirmationPlantBatch                  string       `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate string       `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate   string       `json:"StockConfirmationPlantBatchValidityEndDate"`
	ConfirmedDeliveryDate                        string       `json:"ConfirmedDeliveryDate"`
	RequestedDeliveryDate                        string       `json:"RequestedDeliveryDate"`
	OrderQuantityInBaseUnit                      *interface{} `json:"OrderQuantityInBaseUnit"`
	ConfdOrderQtyByPDTAvailCheck                 *interface{} `json:"ConfdOrderQtyByPDTAvailCheck"`
	DeliveredQtyInOrderQtyUnit                   *interface{} `json:"DeliveredQtyInOrderQtyUnit"`
	OpenConfdDelivQtyInOrdQtyUnit                *interface{} `json:"OpenConfdDelivQtyInOrdQtyUnit"`
	DelivBlockReasonForSchedLine                 *interface{} `json:"DelivBlockReasonForSchedLine"`
	PlusMinusFlag                                string       `json:"PlusMinusFlag"`
}
type Item struct {
	OrderItem                                     *int                 `json:"OrderItem"`
	OrderItemCategory                             string               `json:"OrderItemCategory"`
	OrderItemText                                 string               `json:"OrderItemText"`
	Product                                       string               `json:"Product"`
	ProductStandardID                             string               `json:"ProductStandardID"`
	ProductGroup                                  string               `json:"ProductGroup"`
	BaseUnit                                      string               `json:"BaseUnit"`
	PricingDate                                   string               `json:"PricingDate"`
	PriceDetnExchangeRate                         *interface{}         `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate                         string               `json:"RequestedDeliveryDate"`
	StockConfirmationPartnerFunction              string               `json:"StockConfirmationPartnerFunction"`
	StockConfirmationBusinessPartner              *interface{}         `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                        string               `json:"StockConfirmationPlant"`
	StockConfirmationPlantBatch                   string               `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate  string               `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate    string               `json:"StockConfirmationPlantBatchValidityEndDate"`
	ProductIsBatchManagedInStockConfirmationPlant *interface{}         `json:"ProductIsBatchManagedInStockConfirmationPlant"`
	OrderQuantityInBaseUnit                       *interface{}         `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInIssuingUnit                    *interface{}         `json:"OrderQuantityInIssuingUnit"`
	OrderQuantityInReceivingUnit                  *interface{}         `json:"OrderQuantityInReceivingUnit"`
	OrderIssuingUnit                              string               `json:"OrderIssuingUnit"`
	OrderReceivingUnit                            string               `json:"OrderReceivingUnit"`
	StockConfirmationPolicy                       string               `json:"StockConfirmationPolicy"`
	StockConfirmationStatus                       string               `json:"StockConfirmationStatus"`
	ConfdDelivQtyInOrderQtyUnit                   *interface{}         `json:"ConfdDelivQtyInOrderQtyUnit"`
	ItemWeightUnit                                string               `json:"ItemWeightUnit"`
	ProductGrossWeight                            *float32             `json:"ProductGrossWeight"`
	ItemGrossWeight                               *interface{}         `json:"ItemGrossWeight"`
	ProductNetWeight                              *float32             `json:"ProductNetWeight"`
	ItemNetWeight                                 *interface{}         `json:"ItemNetWeight"`
	NetAmount                                     *interface{}         `json:"NetAmount"`
	TaxAmount                                     *interface{}         `json:"TaxAmount"`
	GrossAmount                                   *interface{}         `json:"GrossAmount"`
	BillingDocumentDate                           string               `json:"BillingDocumentDate"`
	ProductionPlantPartnerFunction                string               `json:"ProductionPlantPartnerFunction"`
	ProductionPlantBusinessPartner                *interface{}         `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               string               `json:"ProductionPlant"`
	ProductionPlantTimeZone                       string               `json:"ProductionPlantTimeZone"`
	ProductionPlantStorageLocation                string               `json:"ProductionPlantStorageLocation"`
	IssuingPlantPartnerFunction                   string               `json:"IssuingPlantPartnerFunction"`
	IssuingPlantBusinessPartner                   *interface{}         `json:"IssuingPlantBusinessPartner"`
	IssuingPlant                                  string               `json:"IssuingPlant"`
	IssuingPlantTimeZone                          string               `json:"IssuingPlantTimeZone"`
	IssuingPlantStorageLocation                   string               `json:"IssuingPlantStorageLocation"`
	ReceivingPlantPartnerFunction                 string               `json:"ReceivingPlantPartnerFunction"`
	ReceivingPlantBusinessPartner                 *interface{}         `json:"ReceivingPlantBusinessPartner"`
	ReceivingPlant                                string               `json:"ReceivingPlant"`
	ReceivingPlantTimeZone                        string               `json:"ReceivingPlantTimeZone"`
	ReceivingPlantStorageLocation                 string               `json:"ReceivingPlantStorageLocation"`
	ProductIsBatchManagedInProductionPlant        *interface{}         `json:"ProductIsBatchManagedInProductionPlant"`
	ProductIsBatchManagedInIssuingPlant           *interface{}         `json:"ProductIsBatchManagedInIssuingPlant"`
	ProductIsBatchManagedInReceivingPlant         *interface{}         `json:"ProductIsBatchManagedInReceivingPlant"`
	BatchMgmtPolicyInProductionPlant              string               `json:"BatchMgmtPolicyInProductionPlant"`
	BatchMgmtPolicyInIssuingPlant                 string               `json:"BatchMgmtPolicyInIssuingPlant"`
	BatchMgmtPolicyInReceivingPlant               string               `json:"BatchMgmtPolicyInReceivingPlant"`
	ProductionPlantBatch                          string               `json:"ProductionPlantBatch"`
	IssuingPlantBatch                             string               `json:"IssuingPlantBatch"`
	ReceivingPlantBatch                           string               `json:"ReceivingPlantBatch"`
	ProductionPlantBatchValidityStartDate         string               `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityEndDate           string               `json:"ProductionPlantBatchValidityEndDate"`
	IssuingPlantBatchValidityStartDate            string               `json:"IssuingPlantBatchValidityStartDate"`
	IssuingPlantBatchValidityEndDate              string               `json:"IssuingPlantBatchValidityEndDate"`
	ReceivingPlantBatchValidityStartDate          string               `json:"ReceivingPlantBatchValidityStartDate"`
	ReceivingPlantBatchValidityEndDate            string               `json:"ReceivingPlantBatchValidityEndDate"`
	Incoterms                                     string               `json:"Incoterms"`
	BPTaxClassification                           string               `json:"BPTaxClassification"`
	ProductTaxClassification                      string               `json:"ProductTaxClassification"`
	BPAccountAssignmentGroup                      string               `json:"BPAccountAssignmentGroup"`
	ProductAccountAssignmentGroup                 string               `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                                  string               `json:"PaymentTerms"`
	PaymentMethod                                 string               `json:"PaymentMethod"`
	DocumentRjcnReason                            *interface{}         `json:"DocumentRjcnReason"`
	ItemBillingBlockReason                        *interface{}         `json:"ItemBillingBlockReason"`
	Project                                       string               `json:"Project"`
	AccountingExchangeRate                        *interface{}         `json:"AccountingExchangeRate"`
	ReferenceDocument                             *interface{}         `json:"ReferenceDocument"`
	ReferenceDocumentItem                         *interface{}         `json:"ReferenceDocumentItem"`
	ItemCompleteDeliveryIsDefined                 *interface{}         `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus                            string               `json:"ItemDeliveryStatus"`
	IssuingStatus                                 string               `json:"IssuingStatus"`
	ReceivingStatus                               string               `json:"ReceivingStatus"`
	BillingStatus                                 string               `json:"BillingStatus"`
	TaxCode                                       string               `json:"TaxCode"`
	TaxRate                                       *interface{}         `json:"TaxRate"`
	CountryOfOrigin                               string               `json:"CountryOfOrigin"`
	ItemPartner                                   []ItemPartner        `json:"ItemPartner"`
	ItemPricingElement                            []ItemPricingElement `json:"ItemPricingElement"`
	ItemSchedulingLine                            []ItemSchedulingLine `json:"ItemSchedulingLine"`
}
type Order struct {
	OrderID                         *int            `json:"OrderID"`
	OrderDate                       string          `json:"OrderDate"`
	OrderType                       string          `json:"OrderType"`
	Buyer                           *int            `json:"Buyer"`
	Seller                          *int            `json:"Seller"`
	CreationDate                    string          `json:"CreationDate"`
	LastChangeDate                  string          `json:"LastChangeDate"`
	ContractType                    string          `json:"ContractType"`
	ValidityStartDate               string          `json:"ValidityStartDate"`
	ValidityEndDate                 string          `json:"ValidityEndDate"`
	InvoiceScheduleStartDate        string          `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate          string          `json:"InvoiceScheduleEndDate"`
	TotalNetAmount                  *interface{}    `json:"TotalNetAmount"`
	TotalTaxAmount                  *interface{}    `json:"TotalTaxAmount"`
	TotalGrossAmount                *interface{}    `json:"TotalGrossAmount"`
	OverallDeliveryStatus           string          `json:"OverallDeliveryStatus"`
	TotalBlockStatus                *interface{}    `json:"TotalBlockStatus"`
	OverallOrdReltdBillgStatus      string          `json:"OverallOrdReltdBillgStatus"`
	OverallDocReferenceStatus       string          `json:"OverallDocReferenceStatus"`
	TransactionCurrency             string          `json:"TransactionCurrency"`
	PricingDate                     string          `json:"PricingDate"`
	PriceDetnExchangeRate           *interface{}    `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate           string          `json:"RequestedDeliveryDate"`
	HeaderCompleteDeliveryIsDefined *interface{}    `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderBillingBlockReason        *interface{}    `json:"HeaderBillingBlockReason"`
	DeliveryBlockReason             *interface{}    `json:"DeliveryBlockReason"`
	Incoterms                       string          `json:"Incoterms"`
	PaymentTerms                    string          `json:"PaymentTerms"`
	PaymentMethod                   string          `json:"PaymentMethod"`
	ReferenceDocument               *interface{}    `json:"ReferenceDocument"`
	ReferenceDocumentItem           *interface{}    `json:"ReferenceDocumentItem"`
	BPAccountAssignmentGroup        string          `json:"BPAccountAssignmentGroup"`
	AccountingExchangeRate          *interface{}    `json:"AccountingExchangeRate"`
	BillingDocumentDate             string          `json:"BillingDocumentDate"`
	IsExportImportDelivery          *interface{}    `json:"IsExportImportDelivery"`
	HeaderText                      string          `json:"HeaderText"`
	HeaderPartner                   []HeaderPartner `json:"HeaderPartner"`
	Address                         []Address       `json:"Address"`
	HeaderPDF                       []HeaderPDF     `json:"HeaderPDF"`
	Item                            []Item          `json:"Item"`
}
