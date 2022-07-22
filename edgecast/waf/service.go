// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package waf

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecauth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules/access"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules/bot"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules/custom"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules/managed"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules/rate"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/scopes"
	"github.com/hashicorp/go-retryablehttp"
)

// WafService interacts with the EdgeCast API for WAF
type WafService struct {
	Access  access.ClientService
	Bot     bot.ClientService
	Custom  custom.ClientService
	Managed managed.ClientService
	Rate    rate.ClientService
	Scopes  scopes.ClientService
}

// New creates a new instance of WafService using the provided configuration
func New(config edgecast.SDKConfig) (*WafService, error) {
	authProvider, err := ecauth.NewTokenAuthorizationProvider(config.APIToken)
	if err != nil {
		return nil, fmt.Errorf("error creating WafService: %w", err)
	}

	c := ecclient.New(ecclient.ClientConfig{
		AuthProvider: authProvider,
		BaseAPIURL:   config.BaseAPIURLLegacy,
		UserAgent:    config.UserAgent,
		Logger:       config.Logger,
		CheckRetry:   checkRetryForWAFScopes,
	})

	return &WafService{
		Access:  access.New(c, c.Config.BaseAPIURL.String()),
		Bot:     bot.New(c, c.Config.BaseAPIURL.String()),
		Custom:  custom.New(c, c.Config.BaseAPIURL.String()),
		Managed: managed.New(c, c.Config.BaseAPIURL.String()),
		Rate:    rate.New(c, c.Config.BaseAPIURL.String()),
		Scopes:  scopes.New(c, c.Config.BaseAPIURL.String()),
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
	if resp != nil &&
		resp.StatusCode == http.StatusBadRequest &&
		resp.Request.Method == "POST" &&
		strings.Contains(resp.Request.URL.String(), "waf/v1.0/scopes") {
		return true, nil
	}

	return retryablehttp.DefaultRetryPolicy(ctx, resp, err)
}
