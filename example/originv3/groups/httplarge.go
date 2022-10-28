package main

import (
	"fmt"
	"strconv"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/ecutils"
	"github.com/EdgeCast/ec-sdk-go/edgecast/originv3"
	"github.com/EdgeCast/ec-sdk-go/edgecast/shared/enums"
	"github.com/kr/pretty"
)

func main() {

	// Setup
	idsCredentials := edgecast.IDSCredentials{
		ClientID:     "",
		ClientSecret: "",
		Scope:        "cdn.origins",
	}

	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.IDSCredentials = idsCredentials

	originV3Service, err := originv3.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating service: %v\n", err)
		return
	}

	// Create Group
	originGroupRequest := createOriginGroupRequest()
	createGroupParams := originv3.NewPostHttpLargeGroupsParams()
	createGroupParams.CustomerOriginGroupHTTPRequest = originGroupRequest
	originGroup, err := originV3Service.HttpLargeOnly.PostHttpLargeGroups(
		createGroupParams,
	)

	if err != nil {
		fmt.Printf("error creating origin group: %v\n", err)
		return
	}

	fmt.Println("successfully created origin group")
	fmt.Printf("%# v", pretty.Formatter(originGroup))

	// Get Group by ID
	getGroupParams := originv3.NewGetHttpLargeGroupsGroupIdParams()
	getGroupParams.GroupId = strconv.Itoa(int(*originGroup.Id))
	originGroup, err = originV3Service.HttpLargeOnly.GetHttpLargeGroupsGroupId(
		getGroupParams,
	)

	if err != nil {
		fmt.Printf("error retrieving origin group by ID: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved origin group by ID")
	fmt.Printf("%# v", pretty.Formatter(originGroup))

	// Get Shield POPs
	getShieldPOPsParams := originv3.NewGetHttpLargeShieldPopsParams()
	edgeNodes, err := originV3Service.HttpLargeOnly.GetHttpLargeShieldPops(
		getShieldPOPsParams,
	)

	if err != nil {
		fmt.Printf("error retrieving shield POPs: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved origin shield POPs")
	fmt.Printf("%# v", pretty.Formatter(edgeNodes))

	// Convert group retreived from API to proper update model
	updateGroup := originv3.CustomerOriginGroupHTTPRequest{}
	err = ecutils.Convert(originGroup, &updateGroup)

	if err != nil {
		fmt.Printf("error preparing group update respose: %v\n", err)
		return
	}

	// Update Group with shield POP
	shieldPOPs := []*string{}
	shieldPOPs = append(shieldPOPs,
		edgeNodes[0].Pops[0].Code,
		edgeNodes[1].Pops[1].Code,
	)

	updateGroup.ShieldPops = shieldPOPs

	updateGroupParams := originv3.NewPutHttplargeGroupsGroupIdParams()
	updateGroupParams.GroupId = strconv.Itoa(int(*originGroup.Id))
	updateGroupParams.CustomerOriginGroupHTTPRequest = updateGroup

	originGroup, err = originV3Service.HttpLargeOnly.PutHttplargeGroupsGroupId(
		updateGroupParams,
	)

	if err != nil {
		fmt.Printf("error updating origin group: %v\n", err)
		return
	}

	fmt.Println("successfully updated origin group")
	fmt.Printf("%# v", pretty.Formatter(originGroup))

	// Get all Groups
	originGroups, err := originV3Service.HttpLargeOnly.GetHttpLargeGroups()

	if err != nil {
		fmt.Printf("error retrieving all origin groups: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved all origin groups")
	fmt.Printf("%# v", pretty.Formatter(originGroups))

	// Delete Group
	deleteOriginGroupParams := originv3.NewDeleteMediaTypeGroupsGroupIdParams()
	deleteOriginGroupParams.GroupId = strconv.Itoa(int(*originGroup.Id))
	deleteOriginGroupParams.MediaType = enums.HttpLarge.String()

	err = originV3Service.Common.DeleteMediaTypeGroupsGroupId(
		deleteOriginGroupParams,
	)

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

	origin := originv3.CustomerOriginGroupHTTPRequest{
		Name:        "TestSDKOriginGroup",
		TlsSettings: &tlsSettings,
	}
	origin.SetHostHeader("override-hostheader.example.com")
	origin.SetNetworkTypeId(2)          // Prefer IPv6 over IPv4
	origin.SetStrictPciCertified(false) // Allow non-PCI regions

	return origin
}
