package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/eclog"
	"github.com/EdgeCast/ec-sdk-go/edgecast/routedns"
)

func main() {
	// Setup
	apiToken := "ltd5n2U145YVRIopwgOEQx5D88ShPJuO"
	accountNumber := "CEA78"

	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.Logger = eclog.NewFileLogger("routedns")
	sdkConfig.APIToken = apiToken
	routeDNSService, err := routedns.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating Route DNS Service: %+v\n", err)
		return
	}

	//
	// Master Server Group
	//

	// Add Master Server Group
	masterServerGroup := buildMasterServerGroup()

	addParams := routedns.NewAddMasterServerGroupParams()
	addParams.AccountNumber = accountNumber
	addParams.MasterServerGroup = masterServerGroup

	masterServerGroupObj, err := routeDNSService.AddMasterServerGroup(
		*addParams,
	)

	if err != nil {
		fmt.Printf("error creating Master Server Group: %+v\n", err)
		return
	}

	fmt.Printf("Created MSG: %+v\n", masterServerGroupObj)

	// Get Master Server Group
	getParams := routedns.NewGetMasterServerGroupParams()
	getParams.AccountNumber = accountNumber
	getParams.MasterServerGroupID = masterServerGroupObj.MasterGroupID

	masterServerGroupObj, err = routeDNSService.GetMasterServerGroup(*getParams)

	if err != nil {
		fmt.Printf("error retrieving Master Server Group: %+v\n", err)
		return
	}

	fmt.Printf("Retrievd MSG: %+v\n", masterServerGroupObj)

	// Update Master Server Group
	masterServerGroupObj.Name = "SDK Test Group 1 Updated"
	updateParams := routedns.NewUpdateMasterServerGroupParams()
	updateParams.AccountNumber = accountNumber
	updateParams.MasterServerGroup.MasterGroupID =
		masterServerGroupObj.MasterGroupID
	updateParams.MasterServerGroup.MasterServerGroup =
		masterServerGroupObj.MasterServerGroup

	err = routeDNSService.UpdateMasterServerGroup(*updateParams)

	if err != nil {
		fmt.Printf("error updating Master Server Group: %+v\n", err)
		return
	}

	fmt.Println("Updated MSG")

	// Delete Master Server Group
	deleteParams := routedns.NewDeleteMasterServerGroupParams()
	deleteParams.AccountNumber = accountNumber
	deleteParams.MasterServerGroup = *masterServerGroupObj

	err = routeDNSService.DeleteMasterServerGroup(*deleteParams)

	if err != nil {
		fmt.Printf("error deleting Master Server Group: %+v\n", err)
		return
	}

	fmt.Println("Deleted MSG")

	//
	// Zone
	//

	// Add Zone
	zone := buildZone()

	addZoneParams := routedns.NewAddZoneParams()
	addZoneParams.AccountNumber = accountNumber
	addZoneParams.Zone = zone

	zoneID, err := routeDNSService.AddZone(*addZoneParams)

	if err != nil || zoneID == nil {
		fmt.Printf("error creating zone: %+v\n", err)
		return
	}

	// Get Zone
	getZoneParams := routedns.NewGetZoneParams()
	getZoneParams.AccountNumber = accountNumber
	getZoneParams.ZoneID = *zoneID

	zoneObj, err := routeDNSService.GetZone(*getZoneParams)

	if err != nil {
		fmt.Printf("error retrieving zone: %+v\n", err)
		return
	}

	// Update Zone
	zoneObj.Comment = "Test updated comment"

	updateZoneParams := routedns.NewUpdateZoneParams()
	updateZoneParams.AccountNumber = accountNumber
	updateZoneParams.Zone = *zoneObj

	err = routeDNSService.UpdateZone(*updateZoneParams)

	if err != nil {
		fmt.Printf("error updating zone: %+v\n", err)
		return
	}

	// Delete Zone
	deleteZoneParams := routedns.NewDeleteZoneParams()
	deleteZoneParams.AccountNumber = accountNumber
	deleteZoneParams.Zone = *zoneObj

	err = routeDNSService.DeleteZone(*deleteZoneParams)

	if err != nil {
		fmt.Printf("error deleting zone: %+v\n", err)
		return
	}

	//
	// Group
	//

	// Add Group
	group := buildLoadbalancedGroup(routedns.CName)

	addGroupParams := routedns.NewAddGroupParams()
	addGroupParams.AccountNumber = accountNumber
	addGroupParams.Group = group

	groupID, err := routeDNSService.AddGroup(*addGroupParams)

	if err != nil {
		fmt.Printf("error adding group: %+v\n", err)
		return
	}

	// Get Group
	getGroupParams := routedns.NewGetGroupParams()
	getGroupParams.AccountNumber = accountNumber
	getGroupParams.GroupID = *groupID
	getGroupParams.GroupProductType = routedns.LoadBalancing

	groupObj, err := routeDNSService.GetGroup(*getGroupParams)

	if err != nil {
		fmt.Printf("error retrieving group: %+v\n", err)
		return
	}

	// Update Group
	groupObj.Name = "UpdatedSDKName"

	updateGroupParams := routedns.NewUpdateGroupParams()
	updateGroupParams.AccountNumber = accountNumber
	updateGroupParams.Group = groupObj

	err = routeDNSService.UpdateGroup(updateGroupParams)

	if err != nil {
		fmt.Printf("error updating group: %+v\n", err)
		return
	}

	// Delete Group
	deleteGroupParams := routedns.NewDeleteGroupParams()
	deleteGroupParams.AccountNumber = accountNumber
	deleteGroupParams.Group = *groupObj

	err = routeDNSService.DeleteGroup(*deleteGroupParams)

	if err != nil {
		fmt.Printf("error deleting group: %+v\n", err)
		return
	}

	//
	// TSIG
	//

	// Add TSIG
	tsig := routedns.TSIG{
		Alias:       "Test SDK TSIG",
		KeyName:     "TSIG1",
		KeyValue:    "SAFJ34SJLFDSFL",
		AlgorithmID: routedns.HMAC_SHA512,
	}

	addTSIGParams := routedns.NewAddTSIGParams()
	addTSIGParams.AccountNumber = accountNumber
	addTSIGParams.TSIG = tsig

	tsigID, err := routeDNSService.AddTSIG(*addTSIGParams)

	if err != nil {
		fmt.Printf("error creating TSIG: %+v\n", err)
		return
	}

	// Get TSIG
	getTSIGParams := routedns.NewGetTSIGParams()
	getTSIGParams.AccountNumber = accountNumber
	getTSIGParams.TSIGID = *tsigID

	tsigObj, err := routeDNSService.GetTSIG(*getTSIGParams)

	if err != nil {
		fmt.Printf("error retreiving TSIG: %+v\n", err)
		return
	}

	// Update TSIG
	tsigObj.Alias = "Test SDK TSIG Updated"

	updateTSIGParams := routedns.NewUpdateTSIGParams()
	updateTSIGParams.AccountNumber = accountNumber
	updateTSIGParams.TSIG = *tsigObj

	err = routeDNSService.UpdateTSIG(*updateTSIGParams)

	if err != nil {
		fmt.Printf("error updating TSIG: %+v\n", err)
		return
	}

	// Delete TSIG
	deleteTSIGParams := routedns.NewDeleteTSIGParams()
	deleteTSIGParams.AccountNumber = accountNumber
	deleteTSIGParams.TSIG = *tsigObj

	err = routeDNSService.DeleteTSIG(*deleteTSIGParams)

	if err != nil {
		fmt.Printf("error deleting TSIG: %+v\n", err)
		return
	}

	//
	// Secondary Zone Group
	//

	// Add Secondary Zone Group
	addSecondaryParams := routedns.NewAddSecondaryZoneGroupParams()
	addSecondaryParams.AccountNumber = accountNumber
	addSecondaryParams.SecondaryZoneGroup = buildSecondaryZoneGroup()

	secondaryZoneResponse, err := routeDNSService.AddSecondaryZoneGroup(
		*addSecondaryParams,
	)

	if err != nil {
		fmt.Printf("error adding secondary zone group: %+v\n", err)
		return
	}

	// Get Secondary Zone group
	getSecondaryParams := routedns.NewGetSecondaryZoneGroupParams()
	getSecondaryParams.AccountNumber = accountNumber
	getSecondaryParams.ID = secondaryZoneResponse.ID

	secondaryZoneObj, err := routeDNSService.GetSecondaryZoneGroup(
		*getSecondaryParams,
	)

	if err != nil {
		fmt.Printf("error retrieving secondary zone group: %+v\n", err)
		return
	}

	// Update Secondary Zone Group
	secondaryZoneObj.Name = "TestSDKSecondaryZoneGroupUpdated"

	updateSecondaryParams := routedns.NewUpdateSecondaryZoneGroupParams()
	updateSecondaryParams.AccountNumber = accountNumber
	updateSecondaryParams.SecondaryZoneGroup = *secondaryZoneObj

	err = routeDNSService.UpdateSecondaryZoneGroup(*updateSecondaryParams)

	if err != nil {
		fmt.Printf("error updating secondary zone group: %+v\n", err)
		return
	}

	// Delete Secondary Zone Group
	deleteSecondaryParams := routedns.NewDeleteSecondaryZoneGroupParams()
	deleteSecondaryParams.AccountNumber = accountNumber
	deleteSecondaryParams.SecondaryZoneGroup = *secondaryZoneObj

	err = routeDNSService.DeleteSecondaryZoneGroup(*deleteSecondaryParams)

	if err != nil {
		fmt.Printf("error deleting secondary zone group: %+v\n", err)
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
	// Simple Zone Records
	aRecord1 := routedns.DNSRecord{
		Name:         "testarecord1",
		TTL:          300,
		Rdata:        "54.11.33.27",
		RecordTypeID: routedns.A,
	}
	aRecord2 := routedns.DNSRecord{
		Name:         "testarecord2",
		TTL:          300,
		Rdata:        "54.11.33.29",
		RecordTypeID: routedns.A,
	}

	dnsRecords := routedns.DNSRecords{}
	dnsRecords.A = append(dnsRecords.A, aRecord1, aRecord2)

	lbGroup := buildLoadbalancedGroup(routedns.PrimaryZone)
	failoverGroup := buildFailoverGroup(routedns.PrimaryZone)
	groups := make([]routedns.DnsRouteGroup, 0)
	groups = append(groups, lbGroup, failoverGroup)

	zone := routedns.Zone{
		DomainName: "sdkzone.com",
		Status:     1,
		ZoneType:   1,
		Comment:    "SDK test zone 1",
		Records:    dnsRecords,
		Groups:     groups,
	}

	return zone
}

func buildLoadbalancedGroup(
	groupTypeID routedns.GroupType,
) routedns.DnsRouteGroup {
	// Load Balanced Group with Records
	cnameRecord1 := routedns.DNSRecord{
		Name:         "testcnamerecord1",
		TTL:          300,
		Rdata:        "lb1.sdkzone.com",
		RecordTypeID: routedns.CNAME,
		Weight:       100,
	}
	cnameRecord2 := routedns.DNSRecord{
		Name:         "testcnamerecord2",
		TTL:          300,
		Rdata:        "lb2.sdkzone.com",
		RecordTypeID: routedns.CNAME,
		Weight:       50,
	}

	lbGroupRecord1 := routedns.DNSGroupRecord{
		Record: cnameRecord1,
	}

	lbGroupRecord2 := routedns.DNSGroupRecord{
		Record: cnameRecord2,
	}

	lbGroupRecords := routedns.DNSGroupRecords{}
	lbGroupRecords.CNAME = append(
		lbGroupRecords.CNAME,
		lbGroupRecord1,
		lbGroupRecord2,
	)

	lbGroup := routedns.DnsRouteGroup{
		Name:             "sdklbgroup01",
		GroupTypeID:      groupTypeID,
		GroupProductType: routedns.LoadBalancing,
		GroupComposition: lbGroupRecords,
	}

	return lbGroup
}

func buildFailoverGroup(
	groupTypeID routedns.GroupType,
) routedns.DnsRouteGroup {
	// Failover Group with Records
	aaaaRecord1 := routedns.DNSRecord{
		Name:         "testaaaarecord1",
		TTL:          300,
		Rdata:        "2001:2011:c002:0000:0000:0000:0000:0000",
		RecordTypeID: routedns.AAAA,
		Weight:       50,
	}

	aaaaRecord2 := routedns.DNSRecord{
		Name:         "testaaaarecord2",
		TTL:          300,
		Rdata:        "2001:2011:c002:0000:0000:0000:0000:0001",
		RecordTypeID: routedns.AAAA,
		Weight:       100,
	}

	failoverGroupRecord1 := routedns.DNSGroupRecord{
		Record: aaaaRecord1,
	}

	failoverGroupRecord2 := routedns.DNSGroupRecord{
		Record: aaaaRecord2,
	}

	failoverGroupRecords := routedns.DNSGroupRecords{}
	failoverGroupRecords.AAAA = append(
		failoverGroupRecords.AAAA,
		failoverGroupRecord1,
		failoverGroupRecord2,
	)

	failoverGroup := routedns.DnsRouteGroup{
		Name:             "sdkfogroup01",
		GroupTypeID:      groupTypeID,
		GroupProductType: routedns.Failover,
		GroupComposition: failoverGroupRecords,
	}

	return failoverGroup
}

func buildSecondaryZoneGroup() routedns.SecondaryZoneGroup {
	// These are static values associated with an account. You may use values
	// obtained by Master Server Group or TSIG API calls above or hard code your
	// own values.
	masterGroupID := 872
	masterServer := routedns.MasterServerID{ID: 2256}
	tsigIDs := routedns.TSIGID{ID: 400}

	masterServerTSIG := routedns.MasterServerTSIGIDs{
		MasterServer: masterServer,
		TSIG:         tsigIDs,
	}

	secondaryZoneRequest := routedns.SecondaryZone{
		Comment:    "Test SDK Secondary Zone",
		DomainName: "sdkseczone1.com",
		Status:     1, // Enabled
	}

	zoneCompositionRequest := routedns.ZoneComposition{}
	zoneCompositionRequest.MasterGroupID = masterGroupID
	zoneCompositionRequest.MasterServerTSIGs = append(
		zoneCompositionRequest.MasterServerTSIGs,
		masterServerTSIG,
	)
	zoneCompositionRequest.Zones = append(
		zoneCompositionRequest.Zones,
		secondaryZoneRequest,
	)

	secondaryZoneGroup := routedns.SecondaryZoneGroup{
		Name:            "TestSDKSecondaryZoneGroup",
		ZoneComposition: zoneCompositionRequest,
	}

	return secondaryZoneGroup
}
