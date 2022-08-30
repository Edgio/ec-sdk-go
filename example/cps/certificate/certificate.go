package main

import (
	"fmt"
	"log"
	"time"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/certificate"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/models"
	"github.com/kr/pretty"
)

const timestampFormat = "2006-01-02T15_04_05Z07_00"

func main() {

	// !!!!
	// NOTE: update the test emails in this script to an email address that
	// is registered to a user for the account.
	// !!!!

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
	fmt.Println("**** CREATE A NEW CERTIFICATE ****")
	fmt.Println("")

	certParams := certificate.NewCertificatePostParams()
	certParams.Certificate = &models.CertificateCreate{
		CertificateLabel:     "C_" + time.Now().Format(timestampFormat),
		Description:          "test cert",
		AutoRenew:            false,
		CertificateAuthority: "DigiCert",
		ValidationType:       "DV",
		Organization: &models.OrganizationDetail{
			AdditionalContacts: nil,
			City:               "L.A.",
			CompanyAddress:     "111 fantastic way",
			CompanyName:        "Test Co.",
			ContactEmail:       "user3@test.com",
			ContactFirstName:   "test3",
			ContactLastName:    "user",
			ContactPhone:       "111-111-1111",
			ContactTitle:       "N/A",
			Country:            "US",
			OrganizationalUnit: "Dept1",
			State:              "CA",
			ZipCode:            "90001",
		},
		DcvMethod: "Email",

		Domains: []*models.DomainCreateUpdate{
			{
				IsCommonName: true,
				Name:         "testssdomain.com",
			},
		},
	}

	certResp, err := cpsService.Certificate.CertificatePost(certParams)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully created certificate")
	fmt.Printf("%# v", pretty.Formatter(certResp))

	certID := int64(certResp.ID)

	fmt.Println("")
	fmt.Println("**** GET CERTIFICATE ****")
	fmt.Println("")

	getByIDParams := certificate.NewCertificateGetParams()
	getByIDParams.ID = certID
	getByIDResp, err := cpsService.Certificate.CertificateGet(getByIDParams)
	if err != nil {
		fmt.Printf("failed to get certificate: %v\n", err)
	}

	fmt.Println("successfully retrieved certificate")
	fmt.Printf("%# v", pretty.Formatter(getByIDResp))

	fmt.Println("")
	fmt.Println("**** GET ALL CERTIFICATES ****")
	fmt.Println("")

	getParams := certificate.NewCertificateFindParams()
	page := int32(1)
	pagesize := int32(10)
	getParams.Page = &page
	getParams.PageSize = &pagesize

	resp, err := cpsService.Certificate.CertificateFind(getParams)
	if err != nil {
		fmt.Printf("failed to get certificates: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved certificates")
	fmt.Printf("%# v", pretty.Formatter(resp))

	fmt.Println("")
	fmt.Println("**** UPDATE CERTIFICATE ****")
	fmt.Println("")

	patchParams := certificate.NewCertificatePatchParams()
	patchParams.ID = certID
	patchParams.CertificateRequest = &models.CertificateUpdate{
		CertificateLabel: "C_" + time.Now().Format(timestampFormat),
		AutoRenew:        true,
	}

	patchResp, err := cpsService.Certificate.CertificatePatch(patchParams)
	if err != nil {
		fmt.Printf("failed to update certificate: %v\n", err)
	}

	fmt.Println("successfully updated certificate")
	fmt.Printf("%# v", pretty.Formatter(patchResp))

	fmt.Println("")
	fmt.Println("**** UPDATE CERTIFICATE NOTIFICATIONS ****")
	fmt.Println("")

	patchNotifParams := certificate.NewCertificateUpdateRequestNotificationsParams()
	patchNotifParams.ID = certID
	patchNotifParams.Notifications = []*models.EmailNotification{
		{
			NotificationType: "CertificateRenewal",
			Emails:           []string{"testaccount@test.com"}, //customer or partner user for the account
			Enabled:          true,
		},
	}
	patchNotifResp, err := cpsService.Certificate.CertificateUpdateRequestNotifications(patchNotifParams)
	if err != nil {
		fmt.Printf("failed to update certificate notifications: %v\n", err)
	}

	fmt.Println("successfully updated certificate notifications")
	fmt.Printf("%# v", pretty.Formatter(patchNotifResp))

	fmt.Println("")
	fmt.Println("**** GET CERTIFICATE NOTIFICATIONS ****")
	fmt.Println("")

	getNotifParams := certificate.NewCertificateGetRequestNotificationsParams()
	getNotifParams.ID = certID
	getNotifResp, err := cpsService.Certificate.CertificateGetRequestNotifications(getNotifParams)
	if err != nil {
		fmt.Printf("failed to retrieve certificate notifications: %v\n", err)
	}

	fmt.Println("successfully retrieved certificates notifications")
	fmt.Printf("%# v", pretty.Formatter(getNotifResp))

	fmt.Println("")
	fmt.Println("**** GET CERTIFICATE STATUS ****")
	fmt.Println("")

	getStatusParams := certificate.NewCertificateGetCertificateStatusParams()
	getStatusParams.ID = certID
	getStatusResp, err := cpsService.Certificate.CertificateGetCertificateStatus(getStatusParams)
	if err != nil {
		fmt.Printf("failed to get certificate status: %v\n", err)
	}
	fmt.Println("successfully retrieved certificate status")
	fmt.Printf("%# v", pretty.Formatter(getStatusResp))

	fmt.Println("")
	fmt.Println("**** CANCEL CERTIFICATE ****")
	fmt.Println("")

	cancelParams := certificate.NewCertificateCancelParams()
	cancelParams.Apply = true
	cancelParams.ID = certID
	cancelResp, err := cpsService.Certificate.CertificateCancel(cancelParams)
	if err != nil {
		fmt.Printf("failed to cancel certificate: %v\n", err)
		return
	}

	fmt.Println("successfully canceled certificate")
	fmt.Printf("%# v", pretty.Formatter(cancelResp))

}
