package customer

// Customer represents data for a Customer that is a child object of
// the associated Partner. Customers contain Customer Users that are used to
// login to my.edgecast.com (the MCC). For more information on user management,
// please review the Web Services REST API (Partner) document available at the
// below url:
// https://partner.edgecast.com/support/default.aspx
type Customer struct {
	// A string that is reserved for future use.
	AccountID string `json:"AccountId,omitempty"`

	// A string identifying the street address for the customer.
	Address1 string `json:"Address1,omitempty"`

	// A string providing additional location information for the customer.
	// Typically, this parameter is used to define a suite number or an
	// apartment number.
	Address2 string `json:"Address2,omitempty"`

	// A string identifying the city associated with the customer's address.
	City string `json:"City,omitempty"`

	// A string identifying the state associated with the customer's address.
	// If the country associated with this account is "United States," then
	// valid values for this parameter are restricted to the two character
	// state abbreviations defined in the State/Province section in Appendix D.
	// of the Web Services REST API (Partner) document
	State string `json:"State,omitempty"`

	// A string identifying the ZIP code associated with a customers's address.
	ZIP string `json:"Zip,omitempty"`

	// A string identifying the country associated with the customers's
	// address. Valid values for this parameter must be an exact match to a
	// country defined in the Country section in Appendix D of the Web Services
	// REST API (Partner) document
	Country string `json:"Country,omitempty"`

	// A 64-bit integer that is reserved for future use.
	BandwidthUsageLimit int64 `json:"BandwidthUsageLimit,omitempty"`

	// A string that is reserved for future use.
	BillingAccountTag string `json:"BillingAccountTag,omitempty"`

	// A string identifying the billing address for the customer.
	BillingAddress1 string `json:"BillingAddress1,omitempty"`

	// A string providing additional location information for the billing
	// address associated with the customer. Typically, this parameter is
	// used to define a suite number or an apartment number.
	BillingAddress2 string `json:"BillingAddress2,omitempty"`

	// A string identifying the city associated with the customers's
	// billing address.
	BillingCity string `json:"BillingCity,omitempty"`

	// A string identifying the billing e-mail address associated with the
	// customer.
	BillingContactEmail string `json:"BillingContactEmail,omitempty"`

	// A string identifying the billing fax number associated with the
	// customer.
	BillingContactFax string `json:"BillingContactFax,omitempty"`

	// A string identifying the billing contact first name of the customer.
	BillingContactFirstName string `json:"BillingContactFirstName,omitempty"`

	// A string identifying the billing contact last name of the customer.
	BillingContactLastName string `json:"BillingContactLastName,omitempty"`

	// A string identifying the billing contact cell phone number of the
	// customer.
	BillingContactMobile string `json:"BillingContactMobile,omitempty"`

	// A string identifying the billing contact main phone number of the
	// customer.
	BillingContactPhone string `json:"BillingContactPhone,omitempty"`

	// A string identifying the billing contact title of the customer.
	BillingContactTitle string `json:"BillingContactTitle,omitempty"`

	// A string identifying the country associated with the customer's
	// billing address. Valid values for this parameter must be an exact match
	// to a country defined in the Country section in Appendix D of the Web
	// Services REST API (Partner) document
	BillingCountry string `json:"BillingCountry,omitempty"`

	// A string that is reserved for future use.
	BillingRateInfo string `json:"BillingRateInfo,omitempty"`

	// A string identifying the state associated with the customer's billing
	// address. If the country associated with this account is "United States,"
	// then valid values for this parameter are restricted to the two character
	// state abbreviations defined in the State/Province section in Appendix D
	// of the Web Services REST API (Partner) document
	BillingState string `json:"BillingState,omitempty"`

	// A string identifying the ZIP code associated with a customer's address.
	BillingZIP string `json:"BillingZip,omitempty"`

	// A string identifying the e-mail address associated with the customer.
	ContactEmail string `json:"ContactEmail,omitempty"`

	// A string identifying the fax number associated with the customer.
	ContactFax string `json:"ContactFax,omitempty"`

	// A string identifying the first name of the customer.
	ContactFirstName string `json:"ContactFirstName,omitempty"`

	// A string identifying the last name of the customer.
	ContactLastName string `json:"ContactLastName,omitempty"`

	// A string identifying the cell phone number of the customer.
	ContactMobile string `json:"ContactMobile,omitempty"`

	// A string identifying the main phone number of the customer.
	ContactPhone string `json:"ContactPhone,omitempty"`

	// A string identifying the title of the customer.
	ContactTitle string `json:"ContactTitle,omitempty"`

	// A string identifying the company name of the customer
	CompanyName string `json:"CompanyName,omitempty"`

	// A 64-bit integer that is reserved for future use.
	DataTransferredUsageLimit int64 `json:"DataTransferredUsageLimit,omitempty"`

	// A string that may contain any notes associated to the customer.
	Notes string `json:"Notes,omitempty"`

	// Required when providing a PCC token. An integer that contains the Partner
	// User ID associated with the PCC token provided.
	PartnerUserID int `json:"PartnerUserId,omitempty"`

	// A string that should always be set to "STND"
	ServiceLevelCode string `json:"ServiceLevelCode,omitempty"`

	// A string identifying the website for the customer
	Website string `json:"Website,omitempty"`

	// An integer that indicates whether the customer is active or inactive
	// 0: Indicates that the customer is inactive
	// 1: Indicates that the customer is active
	Status int `json:"Status,omitempty"`
}

// CustomerGetOK is used specifically when retrieving a Customer and contains
// additional read-only properties.
type CustomerGetOK struct {
	Customer
	// An integer that indicates a customer's ID.
	ID int32 `json:"Id,omitempty"` // read-only

	// A string that indicates the unique custom identifier assigned to a
	// customer. This value was assigned to the customer by your organization.
	CustomID string `json:"CustomId,omitempty"` // read-only

	// A string, also referred to as Account Number, that represents the
	// hexadecimal value of the ID
	HexID string `json:"HexId,omitempty"` // read-only

	// A string indicating the date and time (UTC) for the last time that the
	// usage limit was updated. Syntax: YYYY-MM-DD hh:mm:ssZ
	UsageLimitUpdateDate string `json:"UsageLimitUpdateDate,omitempty"` // read-only

	// An integer identifying the Partner associated with this Customer
	PartnerID int `json:"PartnerId,omitempty"` // read-only

	// A string representing the Partner name associated with this Customer
	PartnerName string `json:"PartnerName,omitempty"` // read-only

	// An integer identifying the Wholesaler associated with Partner that is
	// then associated with this Customer
	WholesaleID int `json:"WholesaleId,omitempty"` // read-only

	// A string representing the Wholsaler name associated with Partner that is
	// then associated with this Customer
	WholesaleName string `json:"WholesaleName,omitempty"` // read-only
}

// AddCustomerParams object contains the properties necessary to create a
// new Customer via the API.
type AddCustomerParams struct {
	Customer Customer
}

// NewAddCustomerParams creates an object with the parameters necessary to
// provide to the AddCustomer function in order to create a Customer via the API.
func NewAddCustomerParams() *AddCustomerParams {
	return &AddCustomerParams{}
}

// GetCustomerParams object contains the properties necessary to retrieve a
// Customer via the API.
type GetCustomerParams struct {
	AccountNumber string
}

// NewGetCustomerParams creates an object with the parameters necessary to
// provide to the GetCustomer function in order to retrieve a Customer via
// the API.
func NewGetCustomerParams() *GetCustomerParams {
	return &GetCustomerParams{}
}

// UpdateCustomerParams object contains the properties necessary to update a
// Customer via the API.
type UpdateCustomerParams struct {
	Customer CustomerGetOK
}

// NewUpdateCustomerParams creates an object with the parameters necessary
// to provide to the UpdateCustomer function in order to update a Customer
// via the API.
func NewUpdateCustomerParams() *UpdateCustomerParams {
	return &UpdateCustomerParams{}
}

// DeleteCustomerParams object contains the properties necessary to delete a
// Customer via the API.
type DeleteCustomerParams struct {
	Customer CustomerGetOK
}

// NewDeleteCustomerParams creates an object with the parameters necessary
// to provide to the DeleteCustomer function in order to delete a Customer
// via the API.
func NewDeleteCustomerParams() *DeleteCustomerParams {
	return &DeleteCustomerParams{}
}

// Service represents data for services (features) available to be enabled or
// disabled for your Customer
type Service struct {
	// An integer that indicates a services's ID.
	ID int `json:"Id,omitempty"`

	// A string identifying the name of the service.
	Name string `json:"Name,omitempty"`

	// An integer that represents a service's parent ID
	ParentID int `json:"ParentId,omitempty"`

	// An integer that indicates whether the service is active or inactive
	// 0: Indicates that the service is inactive
	// 1: Indicates that the service is active
	Status int8 `json:"Status,omitempty"`
}

// GetCustomerServicesParams object contains the properties necessary to
// retrieve a Customer's available services via the API.
type GetCustomerServicesParams struct {
	Customer CustomerGetOK
}

// NewGetCustomerServicesParams creates an object with the parameters necessary
// to provide to the GetCustomerServices function in order to retrieve available
// services for a Customer via the API.
func NewGetCustomerServicesParams() *GetCustomerServicesParams {
	return &GetCustomerServicesParams{}
}

// UpdateCustomerServicesParams object contains the properties necessary to
// update a Customer's available services via the API.
type UpdateCustomerServicesParams struct {
	Customer CustomerGetOK
	// Array of integers representing Service.ID values to be enabled or disabled
	ServiceIDs []int

	// An integer that indicates whether the service will be activated or
	// inactivated
	// 0: Indicates that the service is inactive
	// 1: Indicates that the service is active
	Status int
}

// NewUpdateCustomerServicesParams creates an object with the parameters
// necessary to provide to the UpdateCustomerServices function in order to
// update available services for a Customer via the API.
func NewUpdateCustomerServicesParams() *UpdateCustomerServicesParams {
	return &UpdateCustomerServicesParams{}
}

type DeliveryRegion struct {
	AccountNumber    string `json:"AccountNumber,omitempty"`
	CustomID         string `json:"CustomId,omitempty"`
	DeliveryRegionID int    `json:"DeliveryRegionId,omitempty"`
}

// GetCustomerDeliveryRegionParams object contains the properties necessary to
// retrieve a Customer's available delivery regions via the API.
type GetCustomerDeliveryRegionParams struct {
	Customer CustomerGetOK
}

// NewGetCustomerDeliveryRegionParams creates an object with the parameters
// necessary to provide to the GetCustomerDeliveryRegion function in order to
// retrieve available delivery regions for a Customer via the API.
func NewGetCustomerDeliveryRegionParams() *GetCustomerDeliveryRegionParams {
	return &GetCustomerDeliveryRegionParams{}
}

// UpdateCustomerDeliveryRegionParams object contains the properties necessary
// to update a Customer's available delivery region via the API.
type UpdateCustomerDeliveryRegionParams struct {
	Customer         CustomerGetOK
	DeliveryRegionID int
}

// NewUpdateCustomerDeliveryRegionParams creates an object with the parameters
// necessary to provide to the UpdateCustomerDeliveryRegion function in order to
// update a delivery region for a Customer via the API.
func NewUpdateCustomerDeliveryRegionParams() *UpdateCustomerDeliveryRegionParams {
	return &UpdateCustomerDeliveryRegionParams{}
}

// DomainType represents data for domain types available when updating a
// customer's URLs
type DomainType struct {
	// An integer that indicates the domain type's ID.
	Id int

	// A string identifying the domain type name.
	Name string
}

// UpdateCustomerDomainURLParams object contains the properties necessary
// to update a Customer's URLs via the API.
type UpdateCustomerDomainURLParams struct {
	Customer CustomerGetOK

	// An integer that indicates the DomainType.ID value.
	DomainType int

	// A string representing the Domain URL to update
	Url string
}

// NewUpdateCustomerDomainURLParams creates an object with the parameters
// necessary to provide to the UpdateCustomerDomainURL function in order to
// update a customer's domain URL via the API.
func NewUpdateCustomerDomainURLParams() *UpdateCustomerDomainURLParams {
	return &UpdateCustomerDomainURLParams{}
}

// AccessModule represents data for privileges that may be enabled or disabled
// for a customer
type AccessModule struct {
	// An integer that indicates an access module's ID.
	ID int `json:"Id,omitempty"`

	// An string that indicates an access module's name.
	Name string `json:"Name,omitempty"`

	// An integer that indicates the parent ID of an access module.
	ParentID *int `json:"ParentId,omitempty"`
}

// GetCustomerAccessModulesParams object contains the properties necessary
// to retrieve a Customer's access modules via the API.
type GetCustomerAccessModulesParams struct {
	Customer CustomerGetOK
}

// NewGetCustomerAccessModulesParams creates an object with the parameters
// necessary to provide to the GetCustomerAccessModules function in order to
// retrieve a customer's access modules via the API.
func NewGetCustomerAccessModulesParams() *GetCustomerAccessModulesParams {
	return &GetCustomerAccessModulesParams{}
}

// UpdateCustomerAccessModuleParams object contains the properties necessary
// to update a Customer's access modules enablement status via the API.
type UpdateCustomerAccessModuleParams struct {
	Customer CustomerGetOK

	// A list of integers that indicate the AccessModule.ID values to enable or
	// disable
	AccessModuleIDs []int

	// An integer that indicates whether the service will be activated or
	// inactivated
	// 0: Indicates that the service is inactive
	// 1: Indicates that the service is active
	Status int
}

// NewUpdateCustomerAccessModuleParams creates an object with the parameters
// necessary to provide to the UpdateCustomerAccessModule function in order to
// update a customer's access modules enablement status via the API.
func NewUpdateCustomerAccessModuleParams() *UpdateCustomerAccessModuleParams {
	return &UpdateCustomerAccessModuleParams{}
}
