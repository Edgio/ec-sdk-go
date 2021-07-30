// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package waf

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
	"github.com/EdgeCast/ec-sdk-go/edgecast/logging"
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

	return &WAFService{
		Client: client.NewClient(client.ClientConfig{
			AuthProvider: authProvider,
			BaseAPIURL:   config.BaseAPIURLLegacy,
			UserAgent:    config.UserAgent,
			Logger:       config.Logger,
		}),
		Logger: config.Logger,
	}, nil
}
