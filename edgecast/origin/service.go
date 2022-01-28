// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package origin

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/eclog"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecauth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// Origin service interacts with the EdgeCast API for managing Origins
type OriginService struct {
	client ecclient.APIClient
	Logger eclog.Logger
}

// New creates a new Origin service
func New(config edgecast.SDKConfig) (*OriginService, error) {

	authProvider, err := ecauth.NewTokenAuthorizationProvider(config.APIToken)

	if err != nil {
		return nil, fmt.Errorf("origin.New(): %v", err)
	}

	c := ecclient.New(ecclient.ClientConfig{
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
