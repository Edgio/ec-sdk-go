package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/appendix"
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

	cpsService, err := cps.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating service: %v\n", err)
		return
	}

	fmt.Println("")
	fmt.Println("**** GET COUNTRY CODES ****")
	fmt.Println("")

	appendixParams := appendix.NewAppendixGetParams()
	countryCodeParam := "Bermuda"
	appendixParams.Name = &countryCodeParam
	resp, err := cpsService.Appendix.AppendixGet(appendixParams)

	if err != nil {
		fmt.Printf("failed to get country codes: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved country codes")
	fmt.Printf("%# v", pretty.Formatter(resp))

	fmt.Println("")
	fmt.Println("**** GET VALIDATION TYPES ****")
	fmt.Println("")

	validationTypesParams := appendix.NewAppendixGetValidationTypesParams()
	validationTypesResp, err :=
		cpsService.Appendix.AppendixGetValidationTypes(validationTypesParams)

	if err != nil {
		fmt.Printf("failed to get validation Types: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved validation types")
	fmt.Printf("%# v", pretty.Formatter(validationTypesResp))

	fmt.Println("")
	fmt.Println("**** GET CERTIFICATE AUTHORITIES ****")
	fmt.Println("")

	certAuthParams := appendix.NewAppendixGetCertificateAuthoritiesParams()
	certAuthResp, err :=
		cpsService.Appendix.AppendixGetCertificateAuthorities(certAuthParams)

	if err != nil {
		fmt.Printf("failed to get certificate authorities: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved certificate authorities")
	fmt.Printf("%# v", pretty.Formatter(certAuthResp))

	fmt.Println("")
	fmt.Println("**** GET PRODUCT TYPES ****")
	fmt.Println("")

	prodTypesParams := appendix.NewAppendixGetProductTypesParams()
	prodTypesResp, err :=
		cpsService.Appendix.AppendixGetProductTypes(prodTypesParams)

	if err != nil {
		fmt.Printf("failed to get product types: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved product types")
	fmt.Printf("%# v", pretty.Formatter(prodTypesResp))

	fmt.Println("")
	fmt.Println("**** GET DOMAIN STATUSES ****")
	fmt.Println("")

	domainStatusesParams := appendix.NewAppendixGetDomainStatusesParams()
	domainStatusResp, err :=
		cpsService.Appendix.AppendixGetDomainStatuses(domainStatusesParams)

	if err != nil {
		fmt.Printf("failed to get domain statuses: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved domain statuses")
	fmt.Printf("%# v", pretty.Formatter(domainStatusResp))

	fmt.Println("")
	fmt.Println("**** GET ORDER STATUSES ****")
	fmt.Println("")

	orderStatusesParams := appendix.NewAppendixGetOrderStatusesParams()
	orderStatusesResp, err :=
		cpsService.Appendix.AppendixGetOrderStatuses(orderStatusesParams)

	if err != nil {
		fmt.Printf("failed to get order statuses: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved order statuses")
	fmt.Printf("%# v", pretty.Formatter(orderStatusesResp))

	fmt.Println("")
	fmt.Println("**** GET REQUEST TYPES ****")
	fmt.Println("")

	reqTypesParams := appendix.NewAppendixGetRequestTypeParams()
	reqTypesResp, err :=
		cpsService.Appendix.AppendixGetRequestType(reqTypesParams)

	if err != nil {
		fmt.Printf("failed to get request types: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved request types")
	fmt.Printf("%# v", pretty.Formatter(reqTypesResp))

	fmt.Println("")
	fmt.Println("**** GET DCV TYPES ****")
	fmt.Println("")

	dcvTypesParams := appendix.NewAppendixGetDcvTypesParams()
	dcvTypesResp, err :=
		cpsService.Appendix.AppendixGetDcvTypes(dcvTypesParams)

	if err != nil {
		fmt.Printf("failed to get dcv types: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved dcv types")
	fmt.Printf("%# v", pretty.Formatter(dcvTypesResp))

	fmt.Println("")
	fmt.Println("**** GET CANCEL ACTIONS ****")
	fmt.Println("")

	cancelActionsParams := appendix.NewAppendixGetCancelActionsParams()
	cancelActionsResp, err :=
		cpsService.Appendix.AppendixGetCancelActions(cancelActionsParams)

	if err != nil {
		fmt.Printf("failed to get cancel actions: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved cancel actions")
	fmt.Printf("%# v", pretty.Formatter(cancelActionsResp))

	fmt.Println("")
	fmt.Println("**** GET CERTIFICATE STATUSES ****")
	fmt.Println("")

	certStatusesParams := appendix.NewAppendixGetCertificateStatusesParams()
	certStatusesResp, err :=
		cpsService.Appendix.AppendixGetCertificateStatuses(certStatusesParams)

	if err != nil {
		fmt.Printf("failed to get certificate statuses: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved certificate statuses")
	fmt.Printf("%# v", pretty.Formatter(certStatusesResp))

	fmt.Println("")
	fmt.Println("**** GET VALIDATION STATUSES ****")
	fmt.Println("")

	validationStatusesParams :=
		appendix.NewAppendixGetValidationStatusesParams()
	validationStatusesResp, err :=
		cpsService.Appendix.AppendixGetValidationStatuses(validationStatusesParams)

	if err != nil {
		fmt.Printf("failed to get validation statuses: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved validation statuses")
	fmt.Printf("%# v", pretty.Formatter(validationStatusesResp))

}
