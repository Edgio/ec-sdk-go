// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/ecutils"
	"github.com/EdgeCast/ec-sdk-go/edgecast/originv3"
	"github.com/EdgeCast/ec-sdk-go/edgecast/shared/enums"
	"github.com/kr/pretty"
)

func main() {
	// Setup.
	idsCredentials := edgecast.IDSCredentials{
		ClientID:     "",
		ClientSecret: "",
		Scope:        "cdn.origins",
	}

	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.IDSCredentials = idsCredentials

	svc, err := originv3.New(sdkConfig)
	if err != nil {
		fmt.Printf("error creating service: %v\n", err)
		return
	}

	// Add Group.
	addParams := originv3.NewAddHttpLargeGroupParams()
	addParams.CustomerOriginGroupHTTPRequest = createOriginGroupRequest()

	addResp, err := svc.HttpLargeOnly.AddHttpLargeGroup(addParams)
	if err != nil {
		fmt.Printf("error creating origin group: %v\n", err)
		return
	}

	fmt.Println("successfully created origin group")
	fmt.Printf("%# v", pretty.Formatter(addResp))

	// The response model contains the newly generated group ID.
	groupID := *addResp.Id

	// Get Group by ID.
	getParams := originv3.NewGetHttpLargeGroupParams()
	getParams.GroupId = groupID

	getResp, err := svc.HttpLargeOnly.GetHttpLargeGroup(getParams)
	if err != nil {
		fmt.Printf("error retrieving origin group by ID: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved origin group by ID")
	fmt.Printf("%# v", pretty.Formatter(getResp))

	// Get Shield POPs
	getShieldPOPsParams := originv3.NewGetOriginShieldPopsParams()

	edgeNodes, err := svc.HttpLargeOnly.GetOriginShieldPops(getShieldPOPsParams)
	if err != nil {
		fmt.Printf("error retrieving shield POPs: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved origin shield POPs")
	fmt.Printf("%# v", pretty.Formatter(edgeNodes))

	// Convert Read model to Update Model.
	updateReq := originv3.CustomerOriginGroupHTTPRequest{}
	err = ecutils.Convert(getResp, &updateReq)

	if err != nil {
		fmt.Printf("error preparing group update model: %v\n", err)
		return
	}

	updateReq.SetShieldPOPsFromEdgeNodes(edgeNodes)
	updateParams := originv3.NewUpdateHttpLargeGroupParams()
	updateParams.GroupId = groupID
	updateParams.CustomerOriginGroupHTTPRequest = updateReq

	updateResp, err := svc.HttpLargeOnly.UpdateHttpLargeGroup(updateParams)
	if err != nil {
		fmt.Printf("error updating origin group: %v\n", err)
		return
	}

	fmt.Println("successfully updated origin group")
	fmt.Printf("%# v", pretty.Formatter(updateResp))

	// Get all Groups.
	originGroups, err := svc.HttpLargeOnly.GetAllHttpLargeGroups()
	if err != nil {
		fmt.Printf("error retrieving all origin groups: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved all origin groups")
	fmt.Printf("%# v", pretty.Formatter(originGroups))

	// Cleanup - Delete Group.
	deleteParams := originv3.NewDeleteGroupParams()
	deleteParams.GroupId = groupID
	deleteParams.MediaType = enums.HttpLarge.String()

	err = svc.Common.DeleteGroup(deleteParams)
	if err != nil {
		fmt.Printf("error deleting origin group: %v\n", err)
		return
	}
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
	grp := originv3.CustomerOriginGroupHTTPRequest{
		Name:        "TestSDKOriginGroup",
		TlsSettings: &tlsSettings,
	}

	grp.SetHostHeader("override-hostheader.example.com")
	grp.SetNetworkTypeId(2)          // Prefer IPv6 over IPv4
	grp.SetStrictPciCertified(false) // Allow non-PCI regions

	return grp
}
