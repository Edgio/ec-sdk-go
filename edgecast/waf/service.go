// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package waf

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/eclog"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecauth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
	"github.com/hashicorp/go-retryablehttp"
)

// WAFService interacts with the EdgeCast API for WAF
type WAFService struct {
	client ecclient.APIClient
	Logger eclog.Logger
}

// New creates a new instance of WAFservice using the provided configuration
func New(config edgecast.SDKConfig) (*WAFService, error) {
	authProvider, err := ecauth.NewTokenAuthorizationProvider(config.APIToken)

	if err != nil {
		return nil, fmt.Errorf("waf.New(): %v", err)
	}

	retryPolicy := checkRetryForWAFScopes

	c := ecclient.New(ecclient.ClientConfig{
		AuthProvider: authProvider,
		BaseAPIURL:   config.BaseAPIURLLegacy,
		UserAgent:    config.UserAgent,
		Logger:       config.Logger,
		CheckRetry:   &retryPolicy,
	})

	return &WAFService{
		client: c,
		Logger: config.Logger,
	}, nil
}

// checkRetryForWAFScopes provides a callback to check if we
// will retry on connection errors and server errors.
func checkRetryForWAFScopes(
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
