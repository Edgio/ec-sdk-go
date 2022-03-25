package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/customer"
)

func main() {

	// Setup - fill in the below variables before running this code
	apiToken := "MY_API_TOKEN"

	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = apiToken
	customerService, err := customer.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	// Create a new Customer
	newCustomer := customer.Customer{
		AccountID:                 "AcctId",
		CompanyName:               "EdgeCast SDK",
		Website:                   "https://www.edgecast.com",
		ContactFirstName:          "ContactFirst",
		ContactLastName:           "ContactLast",
		Address1:                  "123 sdk street",
		Address2:                  "suite sdk",
		City:                      "SDK City",
		State:                     "CA",
		ZIP:                       "90210",
		Country:                   "US",
		ContactEmail:              "SDKContactEmail01@sharedectest.com",
		ContactFax:                "888-555-1212",
		ContactMobile:             "888-444-1212",
		ContactPhone:              "888-777-1212",
		ContactTitle:              "Mr",
		BillingAccountTag:         "",
		BillingContactFirstName:   "BillFirst",
		BillingContactLastName:    "BillLast",
		BillingContactTitle:       "Ms",
		BillingAddress1:           "456 bill me",
		BillingAddress2:           "suite bills",
		BillingCity:               "Billing City",
		BillingState:              "CA",
		BillingZIP:                "90210",
		BillingCountry:            "US",
		BillingContactEmail:       "SDKBillingEmail01@sharedectest.com",
		BillingContactFax:         "800-555-1212",
		BillingContactMobile:      "800-444-1212",
		BillingContactPhone:       "800-777-1212",
		Notes:                     "User created via the Go SDK",
		BillingRateInfo:           "",
		BandwidthUsageLimit:       0,
		DataTransferredUsageLimit: 0,
		PartnerUserID:             0,
		ServiceLevelCode:          "STND",
		Status:                    1,
	}

	addParams := customer.NewAddCustomerParams()
	addParams.Customer = newCustomer
	accountNumber, err := customerService.AddCustomer(addParams)

	if err != nil {
		fmt.Printf("error creating customer: %v\n", err)
		return
	}

	fmt.Printf("Customer created. Account number: %s", accountNumber)

	// Get Customer
	getParams := customer.NewGetCustomerParams()
	getParams.AccountNumber = accountNumber
	customerResponse, err := customerService.GetCustomer(*getParams)

	if err != nil {
		fmt.Printf("error retrieving customer: %v\n", err)
		return
	}

	fmt.Printf(
		"Customer retrieved. Account number: %s, Contact Email: %s",
		accountNumber, customerResponse.ContactEmail)

	// Configure customer services and access
	err = ConfigureCustomerAccount(*customerService, *customerResponse)
	if err != nil {
		fmt.Printf("Error configuring customer account: %v", err)
		return
	}

	fmt.Printf("Customer configured. Account number: %s", accountNumber)

	// Update Customer
	updateParams := customer.NewUpdateCustomerParams()
	updateParams.Customer = *customerResponse

	updateParams.Customer.ContactFirstName = "ContactFirstUpdated"

	err = customerService.UpdateCustomer(*updateParams)

	if err != nil {
		fmt.Printf("error updating customer: %v\n", err)
		return
	}

	fmt.Printf("Customer updated. Account number: %s", customerResponse.HexID)

	// Create a Customer User under the above Customer account
	customerUser := customer.CustomerUser{
		Address1:  "123 sdk street",
		Address2:  "suite sdk",
		City:      "SDK City",
		State:     "CA",
		ZIP:       "90210",
		Country:   "US",
		Mobile:    "888-444-1212",
		Phone:     "888-777-1212",
		Fax:       "888-555-1212",
		Email:     "SDKUserEmail01@sharedectest.com",
		Title:     "Ms",
		FirstName: "CustomerUser",
		LastName:  "SDK",
		IsAdmin:   0, // false
	}

	// Add Customer User
	addCustomerUserParams := customer.NewAddCustomerUserParams()
	addCustomerUserParams.Customer = *customerResponse
	addCustomerUserParams.CustomerUser = customerUser
	customerUserID, err := customerService.AddCustomerUser(
		*addCustomerUserParams,
	)

	if err != nil {
		fmt.Printf("error adding customer user: %v\n", err)
		return
	}

	// Get Customer User
	getCustomerUserParams := customer.NewGetCustomerUserParams()
	getCustomerUserParams.Customer = *customerResponse
	getCustomerUserParams.CustomerUserID = customerUserID
	customerUserResponse, err := customerService.GetCustomerUser(
		*getCustomerUserParams,
	)

	if err != nil {
		fmt.Printf("error retrieving customer user: %v\n", err)
		return
	}

	// Update Customer User
	updateCustomerUserParams := customer.NewUpdateCustomerUserParams()
	updateCustomerUserParams.Customer = *customerResponse
	updateCustomerUserParams.CustomerUser = *customerUserResponse
	updateCustomerUserParams.CustomerUser.FirstName = "CustomerUserUpdated"
	err = customerService.UpdateCustomerUser(*updateCustomerUserParams)

	if err != nil {
		fmt.Printf("error updating customer user: %v\n", err)
		return
	}

	// Delete Customer User
	deleteCustomerUserParams := customer.NewDeleteCustomerUserParams()
	deleteCustomerUserParams.Customer = *customerResponse
	deleteCustomerUserParams.CustomerUser = *customerUserResponse
	err = customerService.DeleteCustomerUser(*deleteCustomerUserParams)

	if err != nil {
		fmt.Printf("error deleting customer user: %v\n", err)
		return
	}

	// Delete Customer
	deleteParams := customer.NewDeleteCustomerParams()
	deleteParams.Customer = *customerResponse
	err = customerService.DeleteCustomer(*deleteParams)

	if err != nil {
		fmt.Printf("error deleting customer: %v\n", err)
		return
	}

	fmt.Printf("Customer deleted. Account number: %s", customerResponse.HexID)
}

func ConfigureCustomerAccount(
	customerService customer.CustomerService,
	customerResponse customer.CustomerGetOK,
) error {

	// Get customers available services
	getCustomerServicesParams := customer.NewGetCustomerServicesParams()
	getCustomerServicesParams.Customer = customerResponse
	services, err := customerService.GetCustomerServices(
		*getCustomerServicesParams,
	)

	if err != nil {
		fmt.Printf("error retrieving available customer services: %v\n", err)
		return err
	}

	// DEBUG delete
	fmt.Println(services)

	var serviceIDs []int
	for _, service := range *services {
		if service.Name == "HTTP Large Object" || service.Name == "HTTPS Large Object" {
			serviceIDs = append(serviceIDs, service.ID)
		}
	}

	fmt.Printf("HTTP large service IDs identified: %v", serviceIDs)

	// Update customer services
	updateCustomerServicesParams := customer.NewUpdateCustomerServicesParams()
	updateCustomerServicesParams.Customer = customerResponse
	updateCustomerServicesParams.ServiceIDs = serviceIDs
	updateCustomerServicesParams.Status = 1 // 1 Enable, 0 Disable
	err = customerService.UpdateCustomerServices(*updateCustomerServicesParams)

	if err != nil {
		fmt.Printf("error updating customer services: %v\n", err)
		return err
	}

	fmt.Printf("Enabled services with IDs: %v", serviceIDs)

	// Get customer delivery region
	getCustomerDeliveryRegionParams := customer.NewGetCustomerDeliveryRegionParams()
	getCustomerDeliveryRegionParams.Customer = customerResponse
	deliveryRegion, err := customerService.GetCustomerDeliveryRegion(
		*getCustomerDeliveryRegionParams,
	)

	if err != nil {
		fmt.Printf("error retriving customer delivery region: %v\n", err)
		return err
	}

	fmt.Printf("Currently selected delivery region: %v", deliveryRegion)

	// Update delivery regions
	updateCustomerDeliveryRegionParams := customer.NewUpdateCustomerDeliveryRegionParams()
	updateCustomerDeliveryRegionParams.Customer = customerResponse
	updateCustomerDeliveryRegionParams.DeliveryRegionID = 2 // North America and Europe
	err = customerService.UpdateCustomerDeliveryRegion(
		*updateCustomerDeliveryRegionParams,
	)

	if err != nil {
		fmt.Printf("error updating customer delivery region: %v\n", err)
		return err
	}

	fmt.Printf(
		"Updated customer delivery region id: %v",
		updateCustomerDeliveryRegionParams.DeliveryRegionID,
	)

	// Get customer access modules
	getCustomerAccessModulesParams := customer.NewGetCustomerAccessModulesParams()
	getCustomerAccessModulesParams.Customer = customerResponse
	accessModules, err := customerService.GetCustomerAccessModules(
		*getCustomerAccessModulesParams,
	)

	if err != nil {
		fmt.Printf("error retrieving customer access modules: %v\n", err)
		return err
	}

	// DEBUG delete
	fmt.Println(accessModules)

	// Update customer access modules
	var accessModuleIDs []int
	for _, accessModule := range *accessModules {
		if accessModule.Name == "Storage" {
			accessModuleIDs = append(accessModuleIDs, accessModule.ID)
		}
	}

	updateCustomerAccessModuleParams := customer.NewUpdateCustomerAccessModuleParams()
	updateCustomerAccessModuleParams.Customer = customerResponse
	updateCustomerAccessModuleParams.AccessModuleIDs = accessModuleIDs
	updateCustomerAccessModuleParams.Status = 0 // 0 Disable, 1 Enable

	err = customerService.UpdateCustomerAccessModule(
		*updateCustomerAccessModuleParams,
	)

	if err != nil {
		fmt.Printf("error updating customer access modules: %v\n", err)
		return err
	}

	return nil
}
