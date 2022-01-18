// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package client

import (
	"fmt"
	"net/http"
	"net/url"
)

// APIClient describes structs that can send HTTP requests to an API given the
// request's method, path relative to the API base path.
type APIClient interface {
	Do(params DoParams) (*Response, error)
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
type Request struct {
	method  string
	url     *url.URL
	headers map[string]string
	rawBody interface{}
	// parsedResponse will be filled in using the API response
	parsedResponse interface{}
}

func (r Request) String() string {
	s := fmt.Sprintf("%s %s", r.method, r.url)
	if len(r.headers) > 0 {
		s = s + fmt.Sprintf("\nHeaders:%v", r.headers)
	}
	if r.rawBody != nil {
		s = s + fmt.Sprintf("\nBody:%v", r.rawBody)
	}
	return s
}

// response contains the parsed response from a request along with the raw
// http.Response itself if present
type Response struct {
	Data         interface{}
	HTTPResponse *http.Response
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
	buildRequest(params buildRequestParams) (*Request, error)
}

// requestSender sends a request to an API
type requestSender interface {
	sendRequest(req Request) (*Response, error)
	sendRequestWithStringResponse(req Request) (*string, error)
}

// Describes structs can take client.Response objects, adapt them to a 3rd
// party http library, and return the http.Response from the library
type clientAdapter interface {
	do(req Request) (*http.Response, error)
}
