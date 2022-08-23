// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package ecclient

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/EdgeCast/ec-sdk-go/edgecast/eclog"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecauth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient/ecretryablehttp"
)

// APIClient describes structs that can send HTTP requests to an API given the
// request's method, path relative to the API base path.
type APIClient interface {
	SubmitRequest(params SubmitRequestParams) (*Response, error)
}

// SubmitRequestParams contains the parameters for SubmitRequest.
type SubmitRequestParams struct {
	Method      HTTPMethod
	Path        string
	RawBody     interface{}
	QueryParams map[string]string
	PathParams  map[string]string
	Headers     map[string]string
	// ParsedResponse will be filled in using the API response
	ParsedResponse interface{}
}

// NewSubmitRequestParams creates a new instance of SubmitRequestParams.
func NewSubmitRequestParams() SubmitRequestParams {
	return SubmitRequestParams{
		QueryParams: make(map[string]string),
		PathParams:  make(map[string]string),
		Headers:     make(map[string]string),
	}
}

// response contains the response body as a string from a request along with the
// raw http.Response itself if present
type Response struct {
	Data         string
	HTTPResponse *http.Response
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

// buildRequestParams contains the parameters necessary to construct a new
// request
type buildRequestParams struct {
	method      HTTPMethod
	path        string
	rawBody     interface{}
	queryParams map[string]string
	pathParams  map[string]string
	headers     map[string]string
	userAgent   string
}

// requestBuilder builds a new request using the given parameters
type requestBuilder interface {
	buildRequest(params buildRequestParams) (*request, error)
}

// requestSender sends a request to an API
type requestSender interface {
	sendRequest(req request) (*Response, error)
}

// Describes structs that can pass requests to a 3rd party http library, and
// return the http.Response from the library
type clientAdapter interface {
	Do(
		method string,
		url *url.URL,
		headers map[string]string,
		rawBody interface{},
	) (*http.Response, error)
}

// ECClient -
type ECClient struct {
	reqBuilder requestBuilder
	reqSender  requestSender
	Config     ClientConfig
}

// New creates a default instance of ECClient using the provided
// configuration
func New(config ClientConfig) ECClient {
	clientAdapter := ecretryablehttp.NewRetryableHTTPClientAdapter(
		ecretryablehttp.RetryConfig{
			Logger:       config.Logger,
			RetryWaitMin: config.RetryWaitMin,
			RetryWaitMax: config.RetryWaitMax,
			RetryMax:     config.RetryMax,
			CheckRetry:   ecretryablehttp.CheckRetry(config.CheckRetry),
		})
	return ECClient{
		reqBuilder: newECRequestBuilder(config),
		reqSender:  newECRequestSender(config, clientAdapter),
		Config:     config,
	}
}

// ecRequestBuilder builds requests to be sent to the Edgecast API
type ecRequestBuilder struct {
	baseAPIURL   url.URL
	authProvider ecauth.AuthorizationProvider
	userAgent    string
	logger       eclog.Logger
}

// newECRequestBuilder creates a default instance of ecRequestBuilder using the
// provided configuration
func newECRequestBuilder(config ClientConfig) ecRequestBuilder {
	return ecRequestBuilder{
		baseAPIURL:   config.BaseAPIURL,
		authProvider: config.AuthProvider,
		userAgent:    config.UserAgent,
		logger:       config.Logger,
	}
}

type bodyParser interface {
	parseBody(body []byte, parsedResponse interface{}) error
}

// ecRequestSender sends requests to the Edgecast API
type ecRequestSender struct {
	clientAdapter clientAdapter
	logger        eclog.Logger
	parser        bodyParser
}

type jsonBodyParser struct{}

func newJSONBodyParser() jsonBodyParser {
	return jsonBodyParser{}
}

// newECRequestSender creates a new instance of ecRequestSender
func newECRequestSender(config ClientConfig, ca clientAdapter) ecRequestSender {
	return ecRequestSender{
		clientAdapter: ca,
		logger:        config.Logger,
		parser:        newJSONBodyParser(),
	}
}

// literalResponse is used for unmarshaling response data
// that is in an unrecognized format
type literalResponse struct {
	value interface{}
}

type HTTPMethod int

const (
	Connect HTTPMethod = iota
	Delete
	Get
	Head
	Options
	Patch
	Post
	Put
	Trace
	Unknown
)

func (m HTTPMethod) String() string {
	switch m {
	case Connect:
		return "CONNECT"
	case Delete:
		return "DELETE"
	case Get:
		return "GET"
	case Head:
		return "HEAD"
	case Options:
		return "OPTIONS"
	case Patch:
		return "PATCH"
	case Post:
		return "POST"
	case Put:
		return "PUT"
	case Trace:
		return "TRACE"
	}
	return "UNKNOWN"
}

func ToHTTPMethod(m string) (HTTPMethod, error) {
	switch m {
	case "CONNECT":
		return Connect, nil
	case "DELETE":
		return Delete, nil
	case "GET":
		return Get, nil
	case "HEAD":
		return Head, nil
	case "OPTIONS":
		return Options, nil
	case "PATCH":
		return Patch, nil
	case "POST":
		return Post, nil
	case "PUT":
		return Put, nil
	case "TRACE":
		return Trace, nil
	}
	return Unknown, fmt.Errorf("unknown string type")
}

func (m HTTPMethod) IsValid() bool {
	switch m {
	case Connect, Delete, Get, Head, Options, Patch, Post, Put, Trace:
		return true
	}
	return false
}
