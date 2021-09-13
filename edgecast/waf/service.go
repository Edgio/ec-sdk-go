// Copyright Edgecast, Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package waf

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
	"github.com/EdgeCast/ec-sdk-go/edgecast/logging"
	"github.com/hashicorp/go-retryablehttp"
)

// WAF service interacts with the EdgeCast API for WAF
type WAFService struct {
	client.Client
	Logger logging.Logger
}

// New creates a new WAF service
func New(config edgecast.SDKConfig) (*WAFService, error) {

	authProvider, err := auth.NewTokenAuthorizationProvider(config.APIToken)

	if err != nil {
		return nil, fmt.Errorf("waf.New(): %v", err)
	}

	c := client.NewClient(client.ClientConfig{
		AuthProvider: authProvider,
		BaseAPIURL:   config.BaseAPIURLLegacy,
		UserAgent:    config.UserAgent,
		Logger:       config.Logger,
	})

	// Special retry policy for WAF Scopes
	c.HTTPClient.CheckRetry = CheckRetryScopes

	return &WAFService{
		Client: c,
		Logger: config.Logger,
	}, nil
}

// RetryPolicy provides a callback to check if we
// will retry on connection errors and server errors.
func CheckRetryScopes(
	ctx context.Context,
	resp *http.Response,
	err error,
) (bool, error) {

	// The WAF API throws a 400 Bad Request when the rules
	// being used for a scope have not been fully processed
	// We will retry in that situation until a more specific error is provided
	if resp.StatusCode == http.StatusBadRequest &&
		resp.Request.Method == "POST" &&
		strings.Contains(resp.Request.URL.String(), "waf/v1.0/scopes") {
		return true, nil
	}
	return retryablehttp.DefaultRetryPolicy(ctx, resp, err)
}
