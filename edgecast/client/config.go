package client

import (
	"fmt"
	"net/url"

	"github.com/VerizonDigital/ec-sdk-go/edgecast"
)

const (
	defaultBaseAPIURL       string = "https://api.vdms.io"
	defaultBaseAPIURLLegacy string = "https://api.edgecast.com"
	defaultUserAgentFormat  string = "edgecast/%s:%s"
)

// ClientConfig provides configuration for the base client
type ClientConfig struct {
	// Generates Authorization Header values for HTTP requests
	AuthHeaderProvider AuthorizationHeaderProvider

	// APIURL contains the base URL for the target API
	BaseAPIURL *url.URL

	// The User Agent specifed for HTTP requests
	UserAgent string

	// Logger -
	Logger edgecast.Logger
}

// NewLegacyAPIClientConfig creates a new ClientConfig for targeting the legacy EdgeCast API
// The legacy APIs use an API Token
func NewLegacyAPIClientConfig(apiToken string) (*ClientConfig, error) {
	authProvider, err := NewLegacyAuthorizationHeaderProvider(apiToken)

	if err != nil {
		return nil, fmt.Errorf("Failed to create Legacy API config: %v", err)
	}

	baseAPIURL, _ := url.Parse(defaultBaseAPIURLLegacy)

	return &ClientConfig{
		AuthHeaderProvider: authProvider,
		BaseAPIURL:         baseAPIURL,
		UserAgent:          getDefaultUserAgent(),
		Logger:             edgecast.NewNullLogger(),
	}, nil
}

// NewAPIClientConfig creates a new ClientConfig for targeting the EdgeCast API
// The EdgeCast API uses OAuth 2.0 credentials
func NewAPIClientConfig(credentials IDSCredentials) (*ClientConfig, error) {
	authProvider, err := NewIDSAuthorizationHeaderProvider(credentials)

	if err != nil {
		return nil, fmt.Errorf("Failed to create API config: %v", err)
	}

	baseAPIURL, _ := url.Parse(defaultBaseAPIURL)

	return &ClientConfig{
		AuthHeaderProvider: authProvider,
		BaseAPIURL:         baseAPIURL,
		UserAgent:          getDefaultUserAgent(),
		Logger:             edgecast.NewNullLogger(),
	}, nil
}

func getDefaultUserAgent() string {
	return fmt.Sprintf(defaultUserAgentFormat, edgecast.SDKName, edgecast.SDKVersion)
}
