// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package certificate

// This file was generated by codegen-sdk-go.
// Any changes made to this file may be overwritten.

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NewCertificateFindParams creates a new CertificateFindParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCertificateFindParams() CertificateFindParams {
	return CertificateFindParams{}
}

// CertificateFindParams contains all the parameters to send to the API
// endpoint for the certificate find operation. Typically these are written
// to a http.Request.
type CertificateFindParams struct {

	// Page.
	//
	// Format: int32
	// Default: 1
	Page *int32

	// PageSize.
	//
	// Format: int32
	PageSize *int32

	// Query.
	Query *string

	// Sort.
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the certificate find params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CertificateFindParams) WithDefaults() *CertificateFindParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the certificate find params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CertificateFindParams) SetDefaults() {
	var (
		pageDefault = int32(1)
	)

	val := CertificateFindParams{
		Page: &pageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WriteToRequest extracts parameters and sets for the request to be consumed
func WriteToRequestCertificateFindParams(o CertificateFindParams) (RequestParameters, error) {

	var res []error

	params := NewRequestParameters()

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			params.QueryParams["page"] = qPage
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			params.QueryParams["page_size"] = qPageSize
		}
	}

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			params.QueryParams["query"] = qQuery
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			params.QueryParams["sort"] = qSort
		}
	}

	if len(res) > 0 {
		return *params, errors.CompositeValidationError(res...)
	}
	return *params, nil
}