package main

import (
	"flag"
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/edgecname"
)

func main() {
	// Setup
	apiToken := flag.String("api-token", "", "API Token provided to you")
	accountNumber := flag.String("account-number", "", "Account number, aka hex")

	flag.Parse()

	// Customer management does not use IDS credentials
	idsCredentials := auth.OAuth2Credentials{}

	sdkConfig := edgecast.NewSDKConfig(*apiToken, idsCredentials)
	edgecnameService, err := edgecname.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating Edgecname Service: %v\n", err)
		return
	}

	// Create Edge CNAME
	cname := edgecname.EdgeCname{
		Name:                "test001.sharedectest.com",
		DirPath:             "/my/origin/path",
		EnableCustomReports: 1,
		OriginID:            -1,
		MediaTypeID:         3,
	}

	addParams := edgecname.NewAddEdgeCnameParams()
	addParams.AccountNumber = *accountNumber
	addParams.EdgeCname = cname

	edgeCnameID, err := edgecnameService.AddEdgeCname(*addParams)

	if err != nil {
		fmt.Printf("error creating Edge CNAME: %v\n", err)
		return
	}

	// Get Edge CNAME
	getParams := edgecname.NewGetEdgeCnameParams()
	getParams.AccountNumber = *accountNumber
	getParams.EdgeCnameID = *edgeCnameID

	edgeCnameObj, err := edgecnameService.GetEdgeCname(*getParams)

	if err != nil {
		fmt.Printf("error retrieving Edge CNAME: %v\n", err)
		return
	}

	// Update Edge CNAME
	edgeCnameObj.EnableCustomReports = 0
	updateParams := edgecname.NewUpdateEdgeCnameParams()
	updateParams.AccountNumber = *accountNumber
	updateParams.EdgeCname = *edgeCnameObj

	_, err = edgecnameService.UpdateEdgeCname(*updateParams)

	if err != nil {
		fmt.Printf("error updating Edge CNAME: %v\n", err)
		return
	}

	// Delete Edge CNAME
	deleteParams := edgecname.NewDeleteEdgeCnameParams()
	deleteParams.AccountNumber = *accountNumber
	deleteParams.EdgeCname = *edgeCnameObj

	err = edgecnameService.DeleteEdgeCname(*deleteParams)

	if err != nil {
		fmt.Printf("error deleting Edge CNAME: %v\n", err)
		return
	}
}
