// Copyright Edgecast, Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package edgecname

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
	"github.com/EdgeCast/ec-sdk-go/edgecast/logging"
)

// Cname service interacts with the EdgeCast API for managing Cnames
type CnameService struct {
	client.Client
	Logger logging.Logger
}

// New creates a new Cname service
func New(config edgecast.SDKConfig) (*CnameService, error) {

	authProvider, err := auth.NewTokenAuthorizationProvider(config.APIToken)

	if err != nil {
		return nil, fmt.Errorf("cname.New(): %v", err)
	}

	c := client.NewClient(client.ClientConfig{
		AuthProvider: authProvider,
		BaseAPIURL:   config.BaseAPIURLLegacy,
		UserAgent:    config.UserAgent,
		Logger:       config.Logger,
	})

	return &CnameService{
		Client: c,
		Logger: config.Logger,
	}, nil
}
