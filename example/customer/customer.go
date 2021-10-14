package main

import (
	"flag"
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/customer"
)

func main() {
	// Setup
	apiToken := flag.String("api-token", "", "API Token provided to you")

	flag.Parse()

	// Customer management does not use IDS credentials
	idsCredentials := auth.OAuth2Credentials{}

	sdkConfig := edgecast.NewSDKConfig(*apiToken, idsCredentials)
	customerService, err := customer.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating Customer Service: %v\n", err)
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
		Zip:                       "90210",
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
		BillingZip:                "90210",
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

	accountNumber, err := customerService.AddCustomer(&newCustomer)

	if err != nil {
		fmt.Printf("error creating customer: %v\n", err)
		return
	}

	fmt.Printf("Customer created. Account number: %s", accountNumber)

	// Get Customer
	customerResponse, err := customerService.GetCustomer(accountNumber)

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
		fmt.Println("error configuring customer account")
		return
	}

	fmt.Printf("Customer configured. Account number: %s", accountNumber)

	// Update Customer
	customerResponse.ContactFirstName = "ContactFirstUpdated"
	err = customerService.UpdateCustomer(customerResponse)

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
		Zip:       "90210",
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
	customerUserID, err := customerService.AddCustomerUser(
		customerResponse,
		&customerUser,
	)

	if err != nil {
		fmt.Printf("error adding customer user: %v\n", err)
		return
	}

	// Get Customer User
	customerUserResponse, err := customerService.GetCustomerUser(
		customerResponse,
		customerUserID,
	)

	if err != nil {
		fmt.Printf("error retrieving customer user: %v\n", err)
		return
	}

	// Update Customer User
	customerUserResponse.FirstName = "CustomerUserUpdated"
	err = customerService.UpdateCustomerUser(
		customerResponse,
		customerUserResponse,
	)

	if err != nil {
		fmt.Printf("error updating customer user: %v\n", err)
		return
	}

	// Delete Customer User
	err = customerService.DeleteCustomerUser(
		customerResponse,
		*customerUserResponse,
	)

	if err != nil {
		fmt.Printf("error deleting customer user: %v\n", err)
		return
	}

	// Delete Customer
	err = customerService.DeleteCustomer(customerResponse)

	if err != nil {
		fmt.Printf("error deleting customer: %v\n", err)
		return
	}

	fmt.Printf("Customer deleted. Account number: %s", customerResponse.HexID)
}

func ConfigureCustomerAccount(
	customerService customer.CustomerService,
	customer customer.GetCustomerResponse,
) error {

	// Get customers available services
	services, err := customerService.GetCustomerServices(customer.HexID)

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
	status := 1 // 1 Enable, 0 Disable
	err = customerService.UpdateCustomerServices(customer.HexID, serviceIDs, status)

	if err != nil {
		fmt.Printf("error updating customer services: %v\n", err)
		return err
	}

	fmt.Printf("Enabled services with IDs: %v", serviceIDs)

	// Get customer delivery region
	deliveryRegionID, err := customerService.GetCustomerDeliveryRegion(customer.HexID)

	if err != nil {
		fmt.Printf("error retriving customer delivery region: %v\n", err)
		return err
	}

	fmt.Printf("Currently enabled regionID: %v", deliveryRegionID)

	// Update delivery regions
	deliveryRegionID = 2 // North America and Europe
	err = customerService.UpdateCustomerDeliveryRegion(customer, deliveryRegionID)

	if err != nil {
		fmt.Printf("error updating customer delivery region: %v\n", err)
		return err
	}

	fmt.Printf("Updated customer regionID: %v", deliveryRegionID)

	// Get customer access modules
	accessModules, err := customerService.GetCustomerAccessModules(customer)

	if err != nil {
		fmt.Printf("error retrieving customer access modules: %v\n", err)
		return err
	}

	// DEBUG delete
	fmt.Println(accessModules)

	// Update customer access modules
	var accessModuleIDs []int
	var disableAttribute = 0
	for _, accessModule := range *accessModules {
		if accessModule.Name == "Storage" {
			accessModuleIDs = append(accessModuleIDs, accessModule.ID)
		}
	}
	err = customerService.UpdateCustomerAccessModule(
		customer,
		accessModuleIDs,
		disableAttribute)

	if err != nil {
		fmt.Printf("error updating customer access modules: %v\n", err)
		return err
	}

	return nil
}
