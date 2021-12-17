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
	Address1 string

	// A string providing additional location information for the customer.
	// Typically, this parameter is used to define a suite number or an
	// apartment number.
	Address2 string

	// A string identifying the city associated with the customer's address.
	City string

	// A string identifying the state associated with the customer's address.
	// If the country associated with this account is "United States," then
	// valid values for this parameter are restricted to the two character
	// state abbreviations defined in the State/Province section in Appendix D.
	// of the Web Services REST API (Partner) document
	State string

	// A string identifying the zip code associated with a customers's address.
	Zip string

	// A string identifying the country associated with the customers's
	// address. Valid values for this parameter must be an exact match to a
	// country defined in the Country section in Appendix D of the Web Services
	// REST API (Partner) document
	Country string

	// A 64-bit integer that is reserved for future use.
	BandwidthUsageLimit int64

	// A string that is reserved for future use.
	BillingAccountTag string

	// A string identifying the billing address for the customer.
	BillingAddress1 string

	// A string providing additional location information for the billing
	// address associated with the customer. Typically, this parameter is
	// used to define a suite number or an apartment number.
	BillingAddress2 string

	// A string identifying the city associated with the customers's
	// billing address.
	BillingCity string

	// A string identifying the billing e-mail address associated with the
	// customer.
	BillingContactEmail string

	// A string identifying the billing fax number associated with the
	// customer.
	BillingContactFax string

	// A string identifying the billing contact first name of the customer.
	BillingContactFirstName string

	// A string identifying the billing contact last name of the customer.
	BillingContactLastName string

	// A string identifying the billing contact cell phone number of the
	// customer.
	BillingContactMobile string

	// A string identifying the billing contact main phone number of the
	// customer.
	BillingContactPhone string

	// A string identifying the billing contact title of the customer.
	BillingContactTitle string

	// A string identifying the country associated with the customer's
	// billing address. Valid values for this parameter must be an exact match
	// to a country defined in the Country section in Appendix D of the Web
	// Services REST API (Partner) document
	BillingCountry string

	// A string that is reserved for future use.
	BillingRateInfo string

	// A string identifying the state associated with the customer's billing
	// address. If the country associated with this account is "United States,"
	// then valid values for this parameter are restricted to the two character
	// state abbreviations defined in the State/Province section in Appendix D
	// of the Web Services REST API (Partner) document
	BillingState string

	// A string identifying the zip code associated with a customer's address.
	BillingZip string

	// A string identifying the e-mail address associated with the customer.
	ContactEmail string

	// A string identifying the fax number associated with the customer.
	ContactFax string

	// A string identifying the first name of the customer.
	ContactFirstName string

	// A string identifying the last name of the customer.
	ContactLastName string

	// A string identifying the cell phone number of the customer.
	ContactMobile string

	// A string identifying the main phone number of the customer.
	ContactPhone string

	// A string identifying the title of the customer.
	ContactTitle string

	// A string identifying the company name of the customer
	CompanyName string

	// A 64-bit integer that is reserved for future use.
	DataTransferredUsageLimit int64

	// A string that may contain any notes associated to the customer.
	Notes string

	// Required when providing a PCC token. An integer that contains the Partner
	// User ID associated with the PCC token provided.
	PartnerUserID int

	// A string that should always be set to "STND"
	ServiceLevelCode string

	// A string identifying the website for the customer
	Website string

	// An integer that indicates whether the customer is active or inactive
	// 0: Indicates that the customer is inactive
	// 1: Indicates that the customer is active
	Status int
}

// GetCustomer -
type GetCustomer struct {
	Customer
	// An integer that indicates a customer's ID.
	ID int32 `json:"Id,omitempty"`

	// A string that indicates the unique custom identifier assigned to a
	// customer. This value was assigned to the customer by your organization.
	CustomID string `json:"CustomId,omitempty"`

	// A string, also referred to as Account Number, that represents the
	// hexadecimal value of the ID
	HexID string

	// A string indicating the date and time (UTC) for the last time that the
	// usage limit was updated. Syntax: YYYY-MM-DD hh:mm:ssZ
	UsageLimitUpdateDate string

	// An integer identifying the Partner associated with this Customer
	PartnerID int `json:"PartnerId,omitempty"`

	// A string representing the Partner name associated with this Customer
	PartnerName string

	// An integer identifying the Wholesaler associated with Partner that is
	// then associated with this Customer
	WholesaleID int `json:"WholesaleId,omitempty"`

	// A string representing the Wholsaler name associated with Partner that is
	// then associated with this Customer
	WholesaleName string
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
	Customer GetCustomer
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
	Customer GetCustomer
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
	Name string

	// An integer that represents a service's parent ID
	ParentID int `json:"parentId,omitempty"`

	// An integer that indicates whether the service is active or inactive
	// 0: Indicates that the service is inactive
	// 1: Indicates that the service is active
	Status int8
}

// GetCustomerServicesParams object contains the properties necessary to
// retrieve a Customer's available services via the API.
type GetCustomerServicesParams struct {
	Customer GetCustomer
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
	Customer GetCustomer
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

// GetCustomerDeliveryRegionParams object contains the properties necessary to
// retrieve a Customer's available delivery regions via the API.
type GetCustomerDeliveryRegionParams struct {
	Customer GetCustomer
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
	Customer         GetCustomer
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
	Customer GetCustomer

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
	ID int

	// An string that indicates an access module's name.
	Name string

	// An integer that indicates the parent ID of an access module.
	ParentID *int
}

// GetCustomerAccessModulesParams object contains the properties necessary
// to retrieve a Customer's access modules via the API.
type GetCustomerAccessModulesParams struct {
	Customer GetCustomer
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
	Customer GetCustomer

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
