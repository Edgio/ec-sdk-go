// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package client

import (
	"net/http"
	"net/url"
)

// APIClient describes structs that can send HTTP requests to an API given the
// request's method, path relative to the API base path.
type APIClient interface {
	Do(params DoParams) (interface{}, error)
}

type DoParams struct {
	Method      string
	Path        string
	Body        interface{}
	QueryParams map[string]string
	PathParams  map[string]string
	// ParsedResponse will be filled in using the API response
	ParsedResponse interface{}
}

// request contains the properties of an HTTP request
type request struct {
	method  string
	url     *url.URL
	headers map[string]string
	rawBody interface{}
	// parsedResponse will be filled in using the API response
	parsedResponse interface{}
}

// response contains the parsed response from a request along with the raw
// http.Response itself if present
type response struct {
	data         interface{}
	httpResponse *http.Response
}

// buildRequestParams contains the parameters necessary to construct a new
// request
type buildRequestParams struct {
	method      string
	path        string
	rawBody     interface{}
	queryParams map[string]string
	pathParams  map[string]string
	userAgent   string
}

// requestBuilder builds a new request using the given parameters
type requestBuilder interface {
	buildRequest(params buildRequestParams) (*request, error)
}

// requestSender sends a request to an API
type requestSender interface {
	sendRequest(req request) (*response, error)
	sendRequestWithStringResponse(req request) (*string, error)
}

// Describes structs can take client.Response objects, adapt them to a 3rd
// party http library, and return the http.Response from the library
type clientAdapter interface {
	do(req request) (*http.Response, error)
}
