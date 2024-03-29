// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package customer

import "github.com/EdgeCast/ec-sdk-go/edgecast/cps/models"

// This file was generated by codegen-sdk-go.
// Any changes made to this file may be overwritten.

// NewCustomerUpdateCustomerNotificationsOK creates a CustomerUpdateCustomerNotificationsOK with default headers values
func NewCustomerUpdateCustomerNotificationsOK() *CustomerUpdateCustomerNotificationsOK {
	return &CustomerUpdateCustomerNotificationsOK{}
}

/*
	CustomerUpdateCustomerNotificationsOK describes a response with status code 200, with default header values.

Success
*/
type CustomerUpdateCustomerNotificationsOK struct {
	models.HyperionCollectionEmailNotification
}

// NewCustomerUpdateCustomerNotificationsBadRequest creates a CustomerUpdateCustomerNotificationsBadRequest with default headers values
func NewCustomerUpdateCustomerNotificationsBadRequest() *CustomerUpdateCustomerNotificationsBadRequest {
	return &CustomerUpdateCustomerNotificationsBadRequest{}
}

/*
	CustomerUpdateCustomerNotificationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CustomerUpdateCustomerNotificationsBadRequest struct {
	models.HyperionErrorReponse
}
