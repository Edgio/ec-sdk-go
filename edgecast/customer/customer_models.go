package customer

type Customer struct {
	AccountID                 string `json:"AccountId,omitempty"` // TODO: This might be completely unused. Verify
	Address1                  string
	Address2                  string
	City                      string
	State                     string
	Zip                       string
	Country                   string
	BandwidthUsageLimit       int64
	BillingAccountTag         string
	BillingAddress1           string
	BillingAddress2           string
	BillingCity               string
	BillingContactEmail       string
	BillingContactFax         string
	BillingContactFirstName   string
	BillingContactLastName    string
	BillingContactMobile      string
	BillingContactPhone       string
	BillingContactTitle       string
	BillingCountry            string
	BillingRateInfo           string
	BillingState              string
	BillingZip                string
	ContactEmail              string
	ContactFax                string
	ContactFirstName          string
	ContactLastName           string
	ContactMobile             string
	ContactPhone              string
	ContactTitle              string
	CompanyName               string
	DataTransferredUsageLimit int64
	Notes                     string
	PartnerUserID             int // Required when providing a PCC token
	ServiceLevelCode          string
	Website                   string
	Status                    int
}

type AddCustomerParams struct {
	Customer Customer
}

func NewAddCustomerParams() *AddCustomerParams {
	return &AddCustomerParams{}
}

// GetCustomer -
type GetCustomer struct {
	Customer
	ID                   int32  `json:"Id,omitempty"`
	CustomID             string `json:"CustomId,omitempty"`
	HexID                string
	UsageLimitUpdateDate string
	PartnerID            int `json:"PartnerId,omitempty"`
	PartnerName          string
	WholesaleID          int `json:"WholesaleId,omitempty"`
	WholesaleName        string
}

type GetCustomerParams struct {
	AccountNumber string
}

func NewGetCustomerParams() *GetCustomerParams {
	return &GetCustomerParams{}
}

type UpdateCustomerParams struct {
	Customer GetCustomer
}

func NewUpdateCustomerParams() *UpdateCustomerParams {
	return &UpdateCustomerParams{}
}

type DeleteCustomerParams struct {
	Customer GetCustomer
}

func NewDeleteCustomerParams() *DeleteCustomerParams {
	return &DeleteCustomerParams{}
}

// Service -
type Service struct {
	ID       int `json:"Id,omitempty"`
	Name     string
	ParentID int `json:"parentId,omitempty"`
	Status   int8
}

type GetCustomerServicesParams struct {
	Customer GetCustomer
}

func NewGetCustomerServicesParams() *GetCustomerServicesParams {
	return &GetCustomerServicesParams{}
}

type UpdateCustomerServicesParams struct {
	Customer   GetCustomer
	ServiceIDs []int
	Status     int
}

func NewUpdateCustomerServicesParams() *UpdateCustomerServicesParams {
	return &UpdateCustomerServicesParams{}
}

type GetCustomerDeliveryRegionParams struct {
	Customer GetCustomer
}

func NewGetCustomerDeliveryRegionParams() *GetCustomerDeliveryRegionParams {
	return &GetCustomerDeliveryRegionParams{}
}

type UpdateCustomerDeliveryRegionParams struct {
	Customer         GetCustomer
	DeliveryRegionID int
}

func NewUpdateCustomerDeliveryRegionParams() *UpdateCustomerDeliveryRegionParams {
	return &UpdateCustomerDeliveryRegionParams{}
}

type DomainType struct {
	Id   int
	Name string
}

type UpdateCustomerDomainURLParams struct {
	Customer   GetCustomer
	DomainType int
	Url        string
}

func NewUpdateCustomerDomainURLParams() *UpdateCustomerDomainURLParams {
	return &UpdateCustomerDomainURLParams{}
}

// AccessModule represents a feature that a customer has access to
type AccessModule struct {
	ID       int
	Name     string
	ParentID *int
}

func NewGetCustomerAccessModulesParams() *GetCustomerAccessModulesParams {
	return &GetCustomerAccessModulesParams{}
}

type GetCustomerAccessModulesParams struct {
	Customer GetCustomer
}

type UpdateCustomerAccessModuleParams struct {
	Customer        GetCustomer
	AccessModuleIDs []int
	Status          int
}

func NewUpdateCustomerAccessModuleParams() *UpdateCustomerAccessModuleParams {
	return &UpdateCustomerAccessModuleParams{}
}
