package customer

// CustomerUser represents data for a Customer User that is a child object of
// the associated Customer. Customer Users are used to login to my.edgecast.com
// (the MCC) and are associated with API tokens for legacy APIs and oath2
// clients. For more information on user management, please review the
// Web Services REST API (Partner) document available at the below url:
// https://partner.edgecast.com/support/default.aspx
type CustomerUser struct {
	// A string identifying the street address for the user.
	Address1 string

	// A string providing additional location information for the user.
	// Typically, this parameter is used to define a suite number or an
	// apartment number.
	Address2 string

	// A string identifying the city associated with the user's address.
	City string

	// A string identifying the state associated with the user's address.
	// If the country associated with this account is "United States," then
	// valid values for this parameter are restricted to the two character
	// state abbreviations defined in the State/Province section in Appendix D.
	// of the Web Services REST API (Partner) document
	State string

	// A string identifying the ZIP code associated with a user's address.
	ZIP string `json:"Zip"`

	// A string identifying the country associated with the user's address.
	// Valid values for this parameter must be an exact match to a country
	// defined in the Country section in Appendix D of the Web Services REST
	// API (Partner) document
	Country string

	// A string identifying the cell phone number of the user.
	Mobile string

	// A string identifying the main phone number of the user.
	Phone string

	// A string identifying the fax number associated with the user.
	Fax string

	// Required. A string identifying the e-mail address associated with the
	// user. The user will use this e-mail address to log in to the MCC.
	// The specified e-mail address must be unique.
	Email string

	// A string identifying the title of the user.
	Title string

	// Required. A string identifying the first name of the user.
	FirstName string

	// Required. A string identifying the last name of the user.
	LastName string

	// Determines whether the user will be the administrator for that
	// customer's MCC account. An MCC administrator has complete access to all
	// MCC configuration pages that are available for that customer. Keep in
	// mind that there can only be a single administrator per customer account
	// and that user cannot be deleted. Valid values for this parameter are:
	// 0: Indicates that the user will not be an administrator.
	// This is the default value.
	// 1: Indicates that the user will be assigned the administrator role
	// for the specified customer account. It is important to note that the
	// first user to be assigned the administrator will become the permanent
	// administrator for that customer account.
	IsAdmin int8
}

// GetCustomerUser is used specifically when retrieving a Customer User and
// contains additional read-only properties.
type GetCustomerUser struct {
	CustomerUser

	// A string indicating the date and time (UTC) for the last time that the
	// user logged into the MCC. Syntax: YYYY-MM-DD hh:mm:ssZ
	LastLoginDate string // read-only

	// An integer that indicates a user's ID.
	Id int `json:",omitempty"` // read-only

	// A string that indicates the unique custom identifier assigned to a user.
	// This value was assigned to the user by your organization.
	CustomID string `json:"CustomId"` // read-only
}

// AddCustomerUserParams object contains the properties necessary to create a
// new Customer User via the API.
type AddCustomerUserParams struct {
	Customer     GetCustomer
	CustomerUser CustomerUser
}

// NewAddCustomerUserParams creates an object with the parameters necessary to
// provide to the AddCustomerUser function in order to create a Customer User
// via the API.
func NewAddCustomerUserParams() *AddCustomerUserParams {
	return &AddCustomerUserParams{}
}

// GetCustomerUserParams object contains the properties necessary to retrieve a
// Customer User via the API.
type GetCustomerUserParams struct {
	Customer       GetCustomer
	CustomerUserID int
}

// NewGetCustomerUserParams creates an object with the parameters necessary to
// provide to the GetCustomerUser function in order to retrieve a Customer User
// via the API.
func NewGetCustomerUserParams() *GetCustomerUserParams {
	return &GetCustomerUserParams{}
}

// UpdateCustomerUserParams object contains the properties necessary to update a
// Customer User via the API.
type UpdateCustomerUserParams struct {
	Customer     GetCustomer
	CustomerUser GetCustomerUser
}

// NewUpdateCustomerUserParams creates an object with the parameters necessary
// to provide to the UpdateCustomerUser function in order to update a Customer
// User via the API.
func NewUpdateCustomerUserParams() *UpdateCustomerUserParams {
	return &UpdateCustomerUserParams{}
}

// DeleteCustomerUserParams object contains the properties necessary to delete a
// Customer User via the API.
type DeleteCustomerUserParams struct {
	Customer     GetCustomer
	CustomerUser GetCustomerUser
}

// NewDeleteCustomerUserParams creates an object with the parameters necessary
// to provide to the DeleteCustomerUser function in order to delete a Customer
// User via the API.
func NewDeleteCustomerUserParams() *DeleteCustomerUserParams {
	return &DeleteCustomerUserParams{}
}
