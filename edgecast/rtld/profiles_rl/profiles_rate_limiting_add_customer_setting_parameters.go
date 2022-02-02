// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package profiles_rl

// This file was generated by codegen-sdk-go.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"

	"github.com/EdgeCast/ec-sdk-go/edgecast/rtldmodels"
)

// NewProfilesRateLimitingAddCustomerSettingParams creates a new ProfilesRateLimitingAddCustomerSettingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProfilesRateLimitingAddCustomerSettingParams() *ProfilesRateLimitingAddCustomerSettingParams {
	return &ProfilesRateLimitingAddCustomerSettingParams{}
}

/* ProfilesRateLimitingAddCustomerSettingParams contains all the parameters to send to the API endpoint
   for the profiles rate limiting add customer setting operation.

   Typically these are written to a http.Request.
*/
type ProfilesRateLimitingAddCustomerSettingParams struct {

	// SettingDto.
	SettingDto *rtldmodels.RateLimitingProfileDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the profiles rate limiting add customer setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProfilesRateLimitingAddCustomerSettingParams) WithDefaults() *ProfilesRateLimitingAddCustomerSettingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the profiles rate limiting add customer setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProfilesRateLimitingAddCustomerSettingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WriteToRequest extracts parameters and sets for the request to be consumed
func WriteToRequestProfilesRateLimitingAddCustomerSettingParams(o *ProfilesRateLimitingAddCustomerSettingParams) (RequestParameters, error) {

	var res []error

	params := NewRequestParameters()
	if o.SettingDto != nil {
		params.Body = o.SettingDto
	}

	if len(res) > 0 {
		return *params, errors.CompositeValidationError(res...)
	}
	return *params, nil
}