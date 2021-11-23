// Copyright 2021 Edgecast Inc. Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package edgecast

import (
	"fmt"
	"net/url"

	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/logging"
)

const (
	defaultBaseAPIURL       string = "https://api.vdms.io"
	defaultBaseAPIURLLegacy string = "https://api.edgecast.com"
	defaultBaseIDSURL       string = "https://id.vdms.io"
	defaultUserAgentFormat  string = "edgecast/%s:%s"
)

// Config holds the configuration for SDK services
type SDKConfig struct {

	// APIURL contains the base URL for the EdgeCast APIs
	BaseAPIURL url.URL

	// BaseLegacyAPIURL contains the base URL for the legacy EdgeCast APIs
	BaseAPIURLLegacy url.URL

	// BaseLegacyAPIURL contains the base URL for retrieving IDS tokens
	BaseIDSURL url.URL

	Logger logging.Logger

	APIToken string

	IDSCredentials auth.OAuth2Credentials

	// The User Agent for outgoing HTTP requests
	UserAgent string
}

func NewSDKConfig(
	apiToken string,
	idsCredentials auth.OAuth2Credentials,
) SDKConfig {
	baseAPIURL, _ := url.Parse(defaultBaseAPIURL)
	baseAPIURLLegacy, _ := url.Parse(defaultBaseAPIURLLegacy)
	baseIDSURL, _ := url.Parse(defaultBaseIDSURL)

	return SDKConfig{
		BaseAPIURL:       *baseAPIURL,
		BaseAPIURLLegacy: *baseAPIURLLegacy,
		BaseIDSURL:       *baseIDSURL,
		Logger:           logging.NewStandardLogger(),
		APIToken:         apiToken,
		IDSCredentials:   idsCredentials,
		UserAgent:        getDefaultUserAgent(),
	}
}

func getDefaultUserAgent() string {
	return fmt.Sprintf(defaultUserAgentFormat, SDKName, SDKVersion)
}
