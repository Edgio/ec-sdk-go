package client

/*
	This file contains code that adapts thee retryablehttp client for use in
	this SDK's core client code.
*/

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// Adapts the client from the retryablehttp library
type retryableHTTPClientAdapter struct {
	retryablehttpClient retryablehttp.Client
}

func newRetryableHTTPClientAdapter(
	config ClientConfig,
) retryableHTTPClientAdapter {
	httpClient := retryablehttp.NewClient()
	httpClient.ErrorHandler = retryablehttp.PassthroughErrorHandler
	httpClient.Logger = config.Logger
	httpClient.Backoff = exponentialJitterBackoff

	if config.CheckRetry != nil {
		httpClient.CheckRetry = retryablehttp.CheckRetry(*config.CheckRetry)
	}

	if config.RetryWaitMin != nil {
		httpClient.RetryWaitMin = *config.RetryWaitMin
	} else {
		httpClient.RetryWaitMin = defaultRetryWaitMinSeconds * time.Second
	}
	if config.RetryWaitMax != nil {
		httpClient.RetryWaitMax = *config.RetryWaitMax
	} else {
		httpClient.RetryWaitMax = defaultRetryWaitMaxSeconds * time.Second
	}
	if config.RetryMax != nil {
		httpClient.RetryMax = *config.RetryMax
	} else {
		httpClient.RetryMax = defaultRetryMax
	}

	return retryableHTTPClientAdapter{
		retryablehttpClient: *httpClient,
	}
}

func (c retryableHTTPClientAdapter) do(req Request) (*http.Response, error) {
	retryablehttpReq, err := retryablehttp.NewRequest(
		req.method,
		req.url.String(),
		req.rawBody,
	)
	if err != nil {
		return nil, fmt.Errorf("retryableHTTPClientAdapter.Do:%v", err)
	}
	for k, v := range req.headers {
		retryablehttpReq.Header.Set(k, v)
	}

	return c.retryablehttpClient.Do(retryablehttpReq)
}
