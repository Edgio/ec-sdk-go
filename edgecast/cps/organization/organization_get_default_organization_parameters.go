// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package organization

// This file was generated by codegen-sdk-go.
// Any changes made to this file may be overwritten.

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
)

// NewOrganizationGetDefaultOrganizationParams creates a new OrganizationGetDefaultOrganizationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrganizationGetDefaultOrganizationParams() OrganizationGetDefaultOrganizationParams {
	return OrganizationGetDefaultOrganizationParams{}
}

// OrganizationGetDefaultOrganizationParams contains all the parameters to send to the API
// endpoint for the organization get default organization operation. Typically these are written
// to a http.Request.
type OrganizationGetDefaultOrganizationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the organization get default organization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationGetDefaultOrganizationParams) WithDefaults() *OrganizationGetDefaultOrganizationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the organization get default organization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationGetDefaultOrganizationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WriteToRequest extracts parameters and sets for the request to be consumed
func WriteToRequestOrganizationGetDefaultOrganizationParams(o OrganizationGetDefaultOrganizationParams) (RequestParameters, error) {

	var res []error

	params := NewRequestParameters()

	if len(res) > 0 {
		return *params, errors.CompositeValidationError(res...)
	}
	return *params, nil
}