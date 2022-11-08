package main

import (
	"fmt"
	"strconv"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	originv3 "github.com/EdgeCast/ec-sdk-go/edgecast/originv3"
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

	originV3Service, err := originv3.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating service: %v\n", err)
		return
	}

	// Create a Group
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

	groupid := *originGroup.Id

	// add origin entries
	// 1.
	addParams1 := originv3.NewAddAdnParams()
	addParams1.MediaType = enums.HttpLarge.String()
	origin1 := originv3.NewCustomerOriginRequest(
		"cdn-origin-example.com",
		true,
		int32(groupid),
	)
	protocoltypeid := int32(3)
	origin1.ProtocolTypeId.Set(&protocoltypeid)
	addParams1.CustomerOriginRequest = *origin1

	resp1, err := originV3Service.Common.AddAdn(addParams1)
	if err != nil {
		fmt.Printf("failed to add origin entry: %v\n", err)
		return
	}

	fmt.Println("successfully added origin entry")
	fmt.Printf("%# v", pretty.Formatter(resp1))

	// 2.
	addParams2 := originv3.NewAddAdnParams()
	addParams2.MediaType = enums.HttpLarge.String()
	origin2 := originv3.NewCustomerOriginRequest(
		"cdn-origin-example2.com",
		false,
		int32(groupid),
	)
	origin2.ProtocolTypeId.Set(&protocoltypeid)
	addParams2.CustomerOriginRequest = *origin2

	resp2, err := originV3Service.Common.AddAdn(addParams2)
	if err != nil {
		fmt.Printf("failed to add origin entry: %v\n", err)
		return
	}

	fmt.Println("successfully added origin entry")
	fmt.Printf("%# v", pretty.Formatter(resp2))

	//3.
	addParams3 := originv3.NewAddAdnParams()
	addParams3.MediaType = enums.HttpLarge.String()
	origin3 := originv3.NewCustomerOriginRequest(
		"cdn-origin-example3.com",
		false,
		int32(groupid),
	)
	addParams3.CustomerOriginRequest = *origin3

	resp3, err := originV3Service.Common.AddAdn(addParams3)
	if err != nil {
		fmt.Printf("failed to add origin entry: %v\n", err)
		return
	}

	fmt.Println("successfully added origin entry")
	fmt.Printf("%# v", pretty.Formatter(resp3))

	fmt.Println("")
	fmt.Println("**** SET LOAD BALANCING ****")
	fmt.Println("")

	loadbalancingParams := originv3.NewPatchMediaTypeGroupsGroupIdOriginsParams()
	loadbalancingParams.MediaType = enums.HttpLarge.String()
	loadbalancingParams.GroupId = strconv.Itoa(int(groupid))
	loadbalancingParams.FailoverOrder = []originv3.FailoverOrder{
		{
			Id:            *resp1.Id,
			Host:          "http://cdn-origin-example.com",
			FailoverOrder: 0,
		},
		{
			Id:            *resp2.Id,
			Host:          "http://cdn-origin-example2.com",
			FailoverOrder: 2,
		},
		{
			Id:            *resp3.Id,
			Host:          "http://cdn-origin-example3.com",
			FailoverOrder: 1,
		},
	}

	err = originV3Service.Common.PatchMediaTypeGroupsGroupIdOrigins(loadbalancingParams)
	if err != nil {
		fmt.Printf("failed to set load balancing: %v\n", err)
		return
	}

	fmt.Println("successfully set load balancing")

	//Delete Group
	deleteOriginGroupParams := originv3.NewDeleteMediaTypeGroupsGroupIdParams()
	deleteOriginGroupParams.GroupId = strconv.Itoa(int(groupid))
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
