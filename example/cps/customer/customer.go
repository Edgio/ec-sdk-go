package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/customer"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/models"
	"github.com/kr/pretty"
)

func main() {

	// !!!!
	// NOTE: update the test emails in this script to an email address that
	// is registered to a user for the account.
	// !!!!

	// Setup
	idsCredentials := edgecast.IDSCredentials{
		ClientID:     "CLIENT_ID",
		ClientSecret: "CLIENT_SECRET",
		Scope:        "SCOPE",
	}

	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.IDSCredentials = idsCredentials

	cpsService, err := cps.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating service: %v\n", err)
		return
	}

	fmt.Println("")
	fmt.Println("**** GET CUSTOMER COMMITS ****")
	fmt.Println("")

	getCommitsParams := customer.NewCustomerGetCustomerCommitsParams()
	getCommitsResp, err :=
		cpsService.Customer.CustomerGetCustomerCommits(getCommitsParams)
	if err != nil {
		fmt.Printf("error fetching customer commits: %v\n", err)
		return
	}

	fmt.Printf("%# v", pretty.Formatter(getCommitsResp))

	fmt.Println("")
	fmt.Println("**** UPDATE CUSTOMER NOTIFICATIONS ****")
	fmt.Println("")

	updateNotifParams := customer.NewCustomerUpdateCustomerNotificationsParams()
	updateNotifParams.Notifications = []*models.EmailNotification{
		{
			NotificationType: "CertificateRenewal",
			Emails:           []string{"customer@account.com"}, //customer or partner user for the account
			Enabled:          true,
		},
		{
			NotificationType: "CertificateExpiring",
			Emails:           []string{"customer@account.com"}, //customer or partner user for the account
			Enabled:          true,
		},
	}
	updateNotifResp, err :=
		cpsService.Customer.CustomerUpdateCustomerNotifications(updateNotifParams)
	if err != nil {
		fmt.Printf("error fetching customer notifications: %v\n", err)
		return
	}

	fmt.Printf("%# v", pretty.Formatter(updateNotifResp))

	fmt.Println("")
	fmt.Println("**** GET CUSTOMER NOTIFICATIONS ****")
	fmt.Println("")

	getNotifParams := customer.NewCustomerGetCustomerNotificationsParams()
	getNotifResp, err :=
		cpsService.Customer.CustomerGetCustomerNotifications(getNotifParams)
	if err != nil {
		fmt.Printf("error fetching customer notifications: %v\n", err)
		return
	}

	fmt.Printf("%# v", pretty.Formatter(getNotifResp))
}
