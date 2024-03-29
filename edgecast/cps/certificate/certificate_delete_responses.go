// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package certificate

import "github.com/EdgeCast/ec-sdk-go/edgecast/cps/models"

// This file was generated by codegen-sdk-go.
// Any changes made to this file may be overwritten.

// NewCertificateDeleteNoContent creates a CertificateDeleteNoContent with default headers values
func NewCertificateDeleteNoContent() *CertificateDeleteNoContent {
	return &CertificateDeleteNoContent{}
}

/*
	CertificateDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type CertificateDeleteNoContent struct {
}

// NewCertificateDeleteBadRequest creates a CertificateDeleteBadRequest with default headers values
func NewCertificateDeleteBadRequest() *CertificateDeleteBadRequest {
	return &CertificateDeleteBadRequest{}
}

/*
	CertificateDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CertificateDeleteBadRequest struct {
	models.HyperionErrorReponse
}

// NewCertificateDeleteNotFound creates a CertificateDeleteNotFound with default headers values
func NewCertificateDeleteNotFound() *CertificateDeleteNotFound {
	return &CertificateDeleteNotFound{}
}

/*
	CertificateDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CertificateDeleteNotFound struct {
	models.HyperionErrorReponse
}
