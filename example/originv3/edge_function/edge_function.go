package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	originv3 "github.com/EdgeCast/ec-sdk-go/edgecast/originv3"
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

	params := originv3.NewGetMediaTypeEdgeFunctionsParams()
	params.MediaType = "http-large"

	resp, err := originV3Service.Common.GetMediaTypeEdgeFunctions(params)
	if err != nil {
		fmt.Printf("failed to get functions at edge: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved functions at edge")
	fmt.Printf("%# v", pretty.Formatter(resp))
}
