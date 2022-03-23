// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package edgecname

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/eclog"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecauth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// Edge Cname service interacts with the EdgeCast API for managing Edge Cnames
type EdgeCnameService struct {
	client ecclient.APIClient
	logger eclog.Logger
}

// New creates a new Edge Cname service
func New(config edgecast.SDKConfig) (*EdgeCnameService, error) {

	authProvider, err := ecauth.NewTokenAuthorizationProvider(config.APIToken)

	if err != nil {
		return nil, fmt.Errorf("edgecname.New(): %v", err)
	}

	c := ecclient.New(ecclient.ClientConfig{
		AuthProvider: authProvider,
		BaseAPIURL:   config.BaseAPIURLLegacy,
		UserAgent:    config.UserAgent,
		Logger:       config.Logger,
	})

	return &EdgeCnameService{
		client: c,
		logger: config.Logger,
	}, nil
}
