// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package ecclient

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/EdgeCast/ec-sdk-go/edgecast/eclog"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecauth"
)

// ClientConfig provides configuration for the core SDK client code
type ClientConfig struct {
	// Generates Authorization Header values for HTTP requests
	AuthProvider ecauth.AuthorizationProvider

	// APIURL contains the base URL for the target API
	BaseAPIURL url.URL

	// The User Agent for outgoing HTTP requests
	UserAgent string

	// The Logger that APIClients will use to write messages
	Logger eclog.Logger

	// The minimum wait time for retries on API errors
	RetryWaitMin *time.Duration

	// The maximum wait time for retries on API errors
	RetryWaitMax *time.Duration

	// The maximum number of retries on API errors
	RetryMax *int

	// CheckRetry is a handler that allows users to define custom logic
	// to determine whether the API Client should retry a failed API call
	CheckRetry CheckRetry
}

type CheckRetry func(
	ctx context.Context,
	resp *http.Response,
	err error,
) (bool, error)
