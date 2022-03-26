// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package routedns

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/eclog"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecauth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// RouteDNS service interacts with the EdgeCast API to manage Route DNS
// configurations
type RouteDNSService struct {
	client ecclient.APIClient
	logger eclog.Logger
}

// New creates a new Route DNS service
func New(config edgecast.SDKConfig) (*RouteDNSService, error) {

	authProvider, err := ecauth.NewTokenAuthorizationProvider(config.APIToken)

	if err != nil {
		return nil, fmt.Errorf("RouteDNS.New(): %w", err)
	}

	c := ecclient.New(ecclient.ClientConfig{
		AuthProvider: authProvider,
		BaseAPIURL:   config.BaseAPIURLLegacy,
		UserAgent:    config.UserAgent,
		Logger:       config.Logger,
	})

	return &RouteDNSService{
		client: c,
		logger: config.Logger,
	}, nil
}
