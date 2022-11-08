// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/originv3"
	"github.com/EdgeCast/ec-sdk-go/edgecast/shared/enums"
	"github.com/kr/pretty"
)

func main() {
	// Setup
	apiToken := "MY_API_TOKEN"

	idsCredentials := edgecast.IDSCredentials{
		ClientID:     "CLIENT_ID",
		ClientSecret: "CLIENT_SECRET",
		Scope:        "SCOPE",
	}

	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = apiToken
	sdkConfig.IDSCredentials = idsCredentials

	svc, err := originv3.New(sdkConfig)
	if err != nil {
		fmt.Printf("error creating service: %v\n", err)
		return
	}

	// Create Customer Origin Group.
	addGroupParams := originv3.NewAddHttpLargeGroupParams()
	addGroupParams.CustomerOriginGroupHTTPRequest = createOriginGroupRequest()

	grp, err := svc.HttpLargeOnly.AddHttpLargeGroup(addGroupParams)
	if err != nil {
		fmt.Printf("error creating origin group: %v\n", err)
		return
	}

	fmt.Println("successfully created origin group")
	fmt.Printf("%# v", pretty.Formatter(grp))

	groupID := *grp.Id

	fmt.Println("")
	fmt.Println("**** ADD ORIGIN ENTRY ****")
	fmt.Println("")
	addOriginParams := originv3.NewAddOriginParams()
	addOriginParams.MediaType = enums.HttpLarge.String()
	originRequest := originv3.NewCustomerOriginRequest(
		"cdn-origin-example.com",
		false,
		groupID,
	)
	addOriginParams.CustomerOriginRequest = *originRequest

	origin, err := svc.Common.AddOrigin(addOriginParams)
	if err != nil {
		fmt.Printf("failed to add origin entry: %v\n", err)
		return
	}

	fmt.Println("successfully added origin entry")
	fmt.Printf("%# v", pretty.Formatter(origin))
	fmt.Println("")

	fmt.Println("**** GET ORIGIN ENTRY BY ID ****")
	fmt.Println("")

	getParams := originv3.NewGetOriginParams()
	getParams.Id = *origin.Id
	getParams.MediaType = enums.HttpLarge.String()

	getResp, err := svc.Common.GetOrigin(getParams)
	if err != nil {
		fmt.Printf("failed to retrieve origin entry: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved origin entry")
	fmt.Printf("%# v", pretty.Formatter(getResp))
	fmt.Println("")

	fmt.Println("**** GET ORIGIN ENTRIES BY GROUP ****")
	fmt.Println("")

	getByGrpParams := originv3.NewGetOriginsByGroupParams()
	getByGrpParams.MediaType = enums.HttpLarge.String()
	getByGrpParams.GroupId = groupID

	getByGroupResp, err := svc.Common.GetOriginsByGroup(getByGrpParams)
	if err != nil {
		fmt.Printf("failed to retrieve origin entries by group: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved origin entries by group")
	fmt.Printf("%# v", pretty.Formatter(getByGroupResp))
	fmt.Println("")

	fmt.Println("**** UPDATE ORIGIN ENTRY ****")
	fmt.Println("")

	updateParams := originv3.NewUpdateOriginParams()
	updateParams.MediaType = enums.HttpLarge.String()
	updateParams.Id = *origin.Id
	originRequest.IsPrimary = true // reuse request obj from earlier
	updateParams.CustomerOriginRequest = *originRequest

	updateResp, err := svc.Common.UpdateOrigin(updateParams)
	if err != nil {
		fmt.Printf("failed to update origin entry: %v\n", err)
		return
	}

	fmt.Println("successfully updated origin entry")
	fmt.Printf("%# v", pretty.Formatter(updateResp))

	fmt.Println("")
	fmt.Println("**** DELETE ORIGIN ENTRY ****")
	fmt.Println("")

	deleteParams := originv3.NewDeleteOriginParams()
	deleteParams.MediaType = enums.HttpLarge.String()
	deleteParams.Id = *origin.Id

	err = svc.Common.DeleteOrigin(deleteParams)
	if err != nil {
		fmt.Printf("failed to delete origin entry: %v\n", err)
		return
	}

	fmt.Println("successfully deleted origin entry")
}

func createOriginGroupRequest() originv3.CustomerOriginGroupHTTPRequest {
	tlsSettings := originv3.TlsSettings{
		PublicKeysToVerify: []string{
			"ff8b4a82b08ea5f7be124e6b4363c00d7462655f",
			"c571398b01fce46a8a177abdd6174dfee6137358",
		},
	}

	tlsSettings.SetAllowSelfSigned(false)
	tlsSettings.SetSniHostname("origin.example.com")
	origin := originv3.CustomerOriginGroupHTTPRequest{
		Name:        "TestSDKOriginGroup",
		TlsSettings: &tlsSettings,
	}

	origin.SetHostHeader("override-hostheader.example.com")
	origin.SetNetworkTypeId(2)          // Prefer IPv6 over IPv4
	origin.SetStrictPciCertified(false) // Allow non-PCI regions

	return origin
}
