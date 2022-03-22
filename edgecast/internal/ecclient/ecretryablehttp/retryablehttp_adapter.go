// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package ecretryablehttp

/*
	This file contains code that adapts thee retryablehttp client for use in
	this SDK's core client code.
*/

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// Adapts the client from the retryablehttp library
type RetryableHTTPClientAdapter struct {
	RetryableHttpClient *retryablehttp.Client
	HasCustomRetry      bool
}

func NewRetryableHTTPClientAdapter(
	config RetryConfig,
) *RetryableHTTPClientAdapter {
	httpClient := retryablehttp.NewClient()
	httpClient.ErrorHandler = retryablehttp.PassthroughErrorHandler
	httpClient.Logger = config.Logger
	httpClient.Backoff = exponentialJitterBackoff

	adapter := &RetryableHTTPClientAdapter{}

	if config.CheckRetry != nil {
		adapter.HasCustomRetry = true
		httpClient.CheckRetry = retryablehttp.CheckRetry(*config.CheckRetry)
	}

	if config.RetryWaitMin != nil {
		httpClient.RetryWaitMin = *config.RetryWaitMin
	} else {
		httpClient.RetryWaitMin = DefaultRetryWaitMinSeconds * time.Second
	}
	if config.RetryWaitMax != nil {
		httpClient.RetryWaitMax = *config.RetryWaitMax
	} else {
		httpClient.RetryWaitMax = DefaultRetryWaitMaxSeconds * time.Second
	}
	if config.RetryMax != nil {
		httpClient.RetryMax = *config.RetryMax
	} else {
		httpClient.RetryMax = DefaultRetryMax
	}

	adapter.RetryableHttpClient = httpClient
	return adapter
}

func (c *RetryableHTTPClientAdapter) Do(
	method string,
	url *url.URL,
	headers map[string]string,
	rawBody interface{},
) (*http.Response, error) {
	retryablehttpReq, err := retryablehttp.NewRequest(
		method,
		url.String(),
		rawBody,
	)
	if err != nil {
		return nil, fmt.Errorf("RetryableHTTPClientAdapter.Do:%v", err)
	}

	setHeaders(retryablehttpReq, headers)

	return c.RetryableHttpClient.Do(retryablehttpReq)
}

func setHeaders(req *retryablehttp.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}
