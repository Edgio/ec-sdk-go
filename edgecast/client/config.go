// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package client

import (
	"net/url"
	"time"

	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/logging"
)

// ClientConfig provides configuration for the base client
type ClientConfig struct {
	// Generates Authorization Header values for HTTP requests
	AuthProvider auth.AuthorizationProvider

	// APIURL contains the base URL for the target API
	BaseAPIURL url.URL

	// The User Agent for outgoing HTTP requests
	UserAgent string

	Logger logging.Logger

	RetryWaitMin *time.Duration

	RetryWaitMax *time.Duration

	RetryMax *int
}
