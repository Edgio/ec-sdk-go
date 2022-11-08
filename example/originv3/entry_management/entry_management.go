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

	// Create a Group, or provide a groupID if one exists.
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

	fmt.Println("")
	fmt.Println("**** ADD ORIGIN ENTRY ****")
	fmt.Println("")

	addParams := originv3.NewAddAdnParams()
	addParams.MediaType = enums.HttpLarge.String()

	origin := originv3.NewCustomerOriginRequest(
		"cdn-origin-example.com",
		false,
		int32(groupid),
	)
	addParams.CustomerOriginRequest = *origin

	resp, err := originV3Service.Common.AddAdn(addParams)

	if err != nil {
		fmt.Printf("failed to add origin entry: %v\n", err)
		return
	}

	fmt.Println("successfully added origin entry")
	fmt.Printf("%# v", pretty.Formatter(resp))

	originid := resp.Id

	fmt.Println("")
	fmt.Println("**** GET ORIGIN ENTRY BY ID ****")
	fmt.Println("")

	getParams := originv3.NewGetAdnIdParams()
	getParams.Id = int32(*originid)
	getParams.MediaType = enums.HttpLarge.String()

	getResp, err := originV3Service.Common.GetAdnId(getParams)
	if err != nil {
		fmt.Printf("failed to retrieve origin entry: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved origin entry")
	fmt.Printf("%# v", pretty.Formatter(getResp))

	fmt.Println("")
	fmt.Println("**** GET ORIGIN ENTRIES BY GROUP ****")
	fmt.Println("")

	getByGroupParams := originv3.NewGetMediaTypeGroupsIdOriginsParams()
	getByGroupParams.MediaType = enums.HttpLarge.String()
	getByGroupParams.GroupId = strconv.Itoa(int(groupid))

	getByGroupResp, err := originV3Service.Common.GetMediaTypeGroupsIdOrigins(getByGroupParams)

	if err != nil {
		fmt.Printf("failed to retrieve origin entries by group: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved origin entries by group")
	fmt.Printf("%# v", pretty.Formatter(getByGroupResp))

	fmt.Println("")
	fmt.Println("**** UPDATE ORIGIN ENTRY ****")
	fmt.Println("")

	updateParams := originv3.NewPatchAdnIdParams()
	updateParams.MediaType = enums.HttpLarge.String()
	updateParams.Id = int32(*originid)
	origin.IsPrimary = true
	updateParams.CustomerOriginRequest = *origin

	updateResp, err := originV3Service.Common.PatchAdnId(updateParams)

	if err != nil {
		fmt.Printf("failed to update origin entry: %v\n", err)
		return
	}

	fmt.Println("successfully updated origin entry")
	fmt.Printf("%# v", pretty.Formatter(updateResp))

	fmt.Println("")
	fmt.Println("**** DELETE ORIGIN ENTRY ****")
	fmt.Println("")

	deleteParams := originv3.NewDeleteMediaTypeIdParams()
	deleteParams.MediaType = enums.HttpLarge.String()
	deleteParams.Id = int32(*originid)

	err = originV3Service.Common.DeleteMediaTypeId(deleteParams)

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
