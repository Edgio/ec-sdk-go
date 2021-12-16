package customer

// CustomerUser -
type CustomerUser struct {
	Address1      string
	Address2      string
	City          string
	State         string
	Zip           string
	Country       string
	Mobile        string
	Phone         string
	Fax           string
	Email         string
	Title         string
	FirstName     string
	LastName      string
	Id            int    `json:",omitempty"` // read-only
	CustomID      string `json:"CustomId"`   // read-only
	IsAdmin       int8   // 1 true, 0 false
	LastLoginDate string // read-only
}

type AddCustomerUserParams struct {
	Customer     GetCustomer
	CustomerUser CustomerUser
}

func NewAddCustomerUserParams() *AddCustomerUserParams {
	return &AddCustomerUserParams{}
}

type GetCustomerUserParams struct {
	Customer       GetCustomer
	CustomerUserID int
}

func NewGetCustomerUserParams() *GetCustomerUserParams {
	return &GetCustomerUserParams{}
}

type UpdateCustomerUserParams struct {
	Customer     GetCustomer
	CustomerUser CustomerUser
}

func NewUpdateCustomerUserParams() *UpdateCustomerUserParams {
	return &UpdateCustomerUserParams{}
}

type DeleteCustomerUserParams struct {
	Customer     GetCustomer
	CustomerUser CustomerUser
}

func NewDeleteCustomerUserParams() *DeleteCustomerUserParams {
	return &DeleteCustomerUserParams{}
}
