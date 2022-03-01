// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package routedns

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
	"github.com/EdgeCast/ec-sdk-go/edgecast/logging"
)

// RouteDNS service interacts with the EdgeCast API to manage Route DNS
// configurations
type RouteDNSService struct {
	client.Client
	Logger logging.Logger
}

// New creates a new Route DNS service
func New(config edgecast.SDKConfig) (*RouteDNSService, error) {

	authProvider, err := auth.NewTokenAuthorizationProvider(config.APIToken)

	if err != nil {
		return nil, fmt.Errorf("RouteDNS.New(): %v", err)
	}

	c := client.NewClient(client.ClientConfig{
		AuthProvider: authProvider,
		BaseAPIURL:   config.BaseAPIURLLegacy,
		UserAgent:    config.UserAgent,
		Logger:       config.Logger,
	})

	return &RouteDNSService{
		Client: c,
		Logger: config.Logger,
	}, nil
}
