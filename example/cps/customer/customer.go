package main

import (
	"fmt"
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/customer"
)

func main() {

	// Setup - fill in the below variables before running this code
	idsCredentials := edgecast.IDSCredentials{
		ClientID:     "",
		ClientSecret: "",
		Scope:        "sec.cps.certificates",
	}

	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.IDSCredentials = idsCredentials

	cpsService, err := cps.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating service: %v\n", err)
		return
	}

	getCustomerNotificationsParams := customer.NewCustomerGetCustomerNotificationsParams()
	response, err := cpsService.Customer.CustomerGetCustomerNotifications(getCustomerNotificationsParams)
	if err != nil {
		fmt.Printf("error fetching customer notifications: %v\n", err)
		return
	}

	fmt.Printf("customer notifications:\n%+v", response)
}
