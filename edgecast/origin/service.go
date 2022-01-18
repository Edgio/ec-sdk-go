// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package origin

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
	"github.com/EdgeCast/ec-sdk-go/edgecast/eclog"
)

// Origin service interacts with the EdgeCast API for managing Origins
type OriginService struct {
	client client.APIClient
	Logger eclog.Logger
}

// New creates a new Origin service
func New(config edgecast.SDKConfig) (*OriginService, error) {

	authProvider, err := auth.NewTokenAuthorizationProvider(config.APIToken)

	if err != nil {
		return nil, fmt.Errorf("origin.New(): %v", err)
	}

	c := client.NewECClient(client.ClientConfig{
		AuthProvider: authProvider,
		BaseAPIURL:   config.BaseAPIURLLegacy,
		UserAgent:    config.UserAgent,
		Logger:       config.Logger,
	})

	return &OriginService{
		client: c,
		Logger: config.Logger,
	}, nil
}
