// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package customer

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/eclog"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecauth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// Customer service interacts with the EdgeCast API for Customer
type CustomerService struct {
	client ecclient.APIClient
	Logger eclog.Logger
}

// New creates a new Customer service
func New(config edgecast.SDKConfig) (*CustomerService, error) {
	authProvider, err := ecauth.NewTokenAuthorizationProvider(config.APIToken)

	if err != nil {
		return nil, fmt.Errorf("customer.New(): %v", err)
	}

	c := ecclient.New(ecclient.ClientConfig{
		AuthProvider: authProvider,
		BaseAPIURL:   config.BaseAPIURLLegacy,
		UserAgent:    config.UserAgent,
		Logger:       config.Logger,
	})

	return &CustomerService{
		client: c,
		Logger: config.Logger,
	}, nil
}
