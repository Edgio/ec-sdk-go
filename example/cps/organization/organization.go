package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/organization"
	"github.com/kr/pretty"
)

func main() {

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
	fmt.Println("**** GET DEFAULT ORGANIZATION ****")
	fmt.Println("")

	getdefaultOrgParams :=
		organization.NewOrganizationGetDefaultOrganizationParams()
	getdefaultOrgResp, err :=
		cpsService.Organization.OrganizationGetDefaultOrganization(getdefaultOrgParams)
	if err != nil {
		fmt.Printf("error fetching default organization: %v\n", err)
		return
	}
	fmt.Println("successfully retrieved default organization")
	fmt.Printf("%# v", pretty.Formatter(getdefaultOrgResp))

	defaultOrgID := int64(getdefaultOrgResp.ID)
	defaultOrgName := string(getdefaultOrgResp.CompanyName)

	fmt.Println("")
	fmt.Println("**** GET ORGANIZATION BY ID ****")
	fmt.Println("")

	getOrgParams := organization.NewOrganizationGetParams()
	getOrgParams.ID = defaultOrgID
	getOrgResp, err := cpsService.Organization.OrganizationGet(getOrgParams)
	if err != nil {
		fmt.Printf("error fetching organization by id: %v\n", err)
		return
	}
	fmt.Println("successfully retrieved organization by id")
	fmt.Printf("%# v", pretty.Formatter(getOrgResp))

	fmt.Println("")
	fmt.Println("**** GET ORGANIZATION BY NAME ****")
	fmt.Println("")

	findOrgParams := organization.NewOrganizationFindParams()
	findOrgParams.Name = defaultOrgName
	findOrgResp, err := cpsService.Organization.OrganizationFind(findOrgParams)
	if err != nil {
		fmt.Printf("error fetching organization by name: %v\n", err)
		return
	}
	fmt.Println("successfully retrieved organization by name")
	fmt.Printf("%# v", pretty.Formatter(findOrgResp))

}
