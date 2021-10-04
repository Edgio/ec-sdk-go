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
	//accountNumber := flag.String("account-number", "", "Account number you wish to retrieve")

	flag.Parse()

	idsCredentials := auth.OAuth2Credentials{} // Customer does not use these credentials
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

	fmt.Println(accountNumber)

	// Get Customer
	customer, err := customerService.GetCustomer(accountNumber)

	if err != nil {
		fmt.Printf("error retrieving customer: %v\n", err)
		return
	}

	fmt.Println(customer)
}
