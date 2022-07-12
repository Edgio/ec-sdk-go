package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/dcv"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/models"
	"github.com/kr/pretty"
)

func main() {

	// Setup
	idsCredentials := edgecast.IDSCredentials{
		ClientID:     "CLIENT_ID",
		ClientSecret: "CLIENT_SECRET",
		Scope:        "SCOPE",
	}

	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.IDSCredentials = idsCredentials

	cpsService, err := cps.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating service: %v\n", err)
		return
	}

	certID := int64(0)                //certificate ID of a provisioned certificate
	domainIDs := string("DOMAIN_IDs") //comma separated list of domain ids

	fmt.Println("")
	fmt.Println("**** GET DCV ****")
	fmt.Println("")

	getDCVParams := dcv.NewDcvGetCertificateDomainDetailsParams()
	getDCVParams.ID = certID
	getDCVParams.DomainIds = &domainIDs
	getDCVResp, err :=
		cpsService.Dcv.DcvGetCertificateDomainDetails(getDCVParams)
	if err != nil {
		fmt.Printf("failed to retrieve certificate dcv details: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved certificate dcv details")
	fmt.Printf("%# v", pretty.Formatter(getDCVResp))

	fmt.Println("")
	fmt.Println("**** RESEND DCV EMAILS ****")
	fmt.Println("")

	postDCVEmailParams := dcv.NewDcvPostEmailResendParams()
	postDCVEmailParams.ID = certID
	postDCVEmailParams.DomainIds = &domainIDs
	postDCVEmailResp, err :=
		cpsService.Dcv.DcvPostEmailResend(postDCVEmailParams)
	if err != nil {
		fmt.Printf("failed to resend dcv emails: %v\n", err)
		return
	}

	fmt.Println("successfully resent dcv emails")
	fmt.Printf("%# v", pretty.Formatter(postDCVEmailResp))

	fmt.Println("")
	fmt.Println("**** UPDATE DCV METHOD ****")
	fmt.Println("")

	setDCVmethodParams := dcv.NewDcvSetCertificateDcvMethodParams()
	setDCVmethodParams.ID = certID
	setDCVmethodParams.DcvMethod = &models.DcvMethodRequest{
		DcvMethod: "DnsCnameToken",
	}
	setDCVmethodResp, err :=
		cpsService.Dcv.DcvSetCertificateDcvMethod(setDCVmethodParams)
	if err != nil {
		fmt.Printf("failed to set dcv validation method: %v\n", err)
		return
	}

	fmt.Println("successfully set dcv validation method")
	fmt.Printf("%# v", pretty.Formatter(setDCVmethodResp))

	fmt.Println("")
	fmt.Println("**** REGENERATE DCV TOKEN ****")
	fmt.Println("")

	regenerateDcvTokenParams := dcv.NewDcvRegenerateDcvTokensParams()
	regenerateDcvTokenParams.ID = certID
	regenerateDcvTokenParams.DomainIds = &domainIDs
	regenerateDcvTokenResp, err :=
		cpsService.Dcv.DcvRegenerateDcvTokens(regenerateDcvTokenParams)
	if err != nil {
		fmt.Printf("failed to regenerate dcv token: %v\n", err)
		return
	}

	fmt.Println("successfully regenerated dcv token")
	fmt.Printf("%# v", pretty.Formatter(regenerateDcvTokenResp))
}
