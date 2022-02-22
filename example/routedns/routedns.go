package main

import (
	"flag"
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/routedns"
)

func main() {
	// Setup
	apiToken := flag.String("api-token", "", "API Token provided to you")
	accountNumber := flag.String(
		"account-number",
		"",
		"Account number you wish to manage Route DNS records for",
	)

	flag.Parse()

	// Route DNS management does not use IDS credentials
	idsCredentials := auth.OAuth2Credentials{}

	sdkConfig := edgecast.NewSDKConfig(*apiToken, idsCredentials)
	sdkConfig.BaseAPIURLLegacy.Host = "qa-api.edgecast.com"
	routeDNSService, err := routedns.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating Route DNS Service: %v\n", err)
		return
	}

	//
	// Master Server Group
	//

	// Add Master Server Group
	masterServerGroup := buildMasterServerGroup()

	addParams := routedns.NewAddMasterServerGroupParams()
	addParams.AccountNumber = *accountNumber
	addParams.MasterServerGroup = masterServerGroup

	masterServerGroupObj, err := routeDNSService.AddMasterServerGroup(
		*addParams,
	)

	if err != nil {
		fmt.Printf("error creating Master Server Group: %v\n", err)
		return
	}

	fmt.Printf("msg: %v", masterServerGroupObj)

	// Get Master Server Group
	getParams := routedns.NewGetMasterServerGroupParams()
	getParams.AccountNumber = *accountNumber
	getParams.MasterServerGroupID = masterServerGroupObj.MasterGroupID

	masterServerGroupObj, err = routeDNSService.GetMasterServerGroup(*getParams)

	if err != nil {
		fmt.Printf("error retrieving Master Server Group: %v\n", err)
		return
	}

	fmt.Printf("msg: %v", masterServerGroupObj)

	// Update Master Server Group
	masterServerGroupObj.Name = "SDK Test Group 1 Updated"
	updateParams := routedns.NewUpdateMasterServerGroupParams()
	updateParams.AccountNumber = *accountNumber
	updateParams.MasterServerGroup.MasterGroupID =
		masterServerGroupObj.MasterGroupID
	updateParams.MasterServerGroup.MasterServerGroup =
		masterServerGroupObj.MasterServerGroup

	err = routeDNSService.UpdateMasterServerGroup(*updateParams)

	if err != nil {
		fmt.Printf("error updating Master Server Group: %v\n", err)
		return
	}

	// Delete Master Server Group
	deleteParams := routedns.NewDeleteMasterServerGroupParams()
	deleteParams.AccountNumber = *accountNumber
	deleteParams.MasterServerGroup = *masterServerGroupObj

	err = routeDNSService.DeleteMasterServerGroup(*deleteParams)

	if err != nil {
		fmt.Printf("error deleting Master Server Group: %v\n", err)
		return
	}

	//
	// Zone
	//

	// Add Zone
	zone := buildZone()

	addZoneParams := routedns.NewAddZoneParams()
	addZoneParams.AccountNumber = *accountNumber
	addZoneParams.Zone = zone

	zoneID, err := routeDNSService.AddZone(*addZoneParams)

	if err != nil || zoneID == nil {
		fmt.Printf("error creating zone: %v\n", err)
		return
	}

	// Get Zone
	getZoneParams := routedns.NewGetZoneParams()
	getZoneParams.AccountNumber = *accountNumber
	getZoneParams.ZoneID = *zoneID

	zoneObj, err := routeDNSService.GetZone(*getZoneParams)

	if err != nil {
		fmt.Printf("error retrieving zone: %v\n", err)
		return
	}

	// Update Zone
	zoneObj.Comment = "Test updated comment"

	updateZoneParams := routedns.NewUpdateZoneParams()
	updateZoneParams.AccountNumber = *accountNumber
	updateZoneParams.Zone = *zoneObj

	err = routeDNSService.UpdateZone(*updateZoneParams)

	if err != nil {
		fmt.Printf("error updating zone: %v\n", err)
		return
	}

	// Delete Zone
	deleteZoneParams := routedns.NewDeleteZoneParams()
	deleteZoneParams.AccountNumber = *accountNumber
	deleteZoneParams.Zone = *zoneObj

	err = routeDNSService.DeleteZone(*deleteZoneParams)

	if err != nil {
		fmt.Printf("error deleting zone: %v\n", err)
		return
	}

}

func buildMasterServerGroup() routedns.MasterServerGroupAddRequest {
	masterServer01 := routedns.MasterServer{
		Name:      "Master 01",
		IPAddress: "35.11.100.10",
	}

	masterServer02 := routedns.MasterServer{
		Name:      "Master 02",
		IPAddress: "35.11.100.11",
	}

	masterServerGroup := routedns.MasterServerGroupAddRequest{
		Name: "SDK Test Group 1",
	}

	masterServerGroup.Masters = append(
		masterServerGroup.Masters,
		masterServer01,
		masterServer02,
	)

	return masterServerGroup
}

func buildZone() routedns.Zone {
	aRecord1 := routedns.DNSRecord{
		Name:         "testrecord1",
		TTL:          300,
		Rdata:        "54.11.33.27",
		RecordTypeID: routedns.A,
	}
	aRecord2 := routedns.DNSRecord{
		Name:         "testrecord2",
		TTL:          300,
		Rdata:        "54.11.33.29",
		RecordTypeID: routedns.A,
	}

	dnsRecord := routedns.DNSRecords{}
	dnsRecord.A = append(dnsRecord.A, aRecord1, aRecord2)

	zone := routedns.Zone{
		DomainName: "sdkzone.com",
		Status:     1,
		ZoneType:   1,
		Comment:    "SDK test zone 1",
	}

	return zone
}
