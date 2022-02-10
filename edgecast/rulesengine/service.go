// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package rulesengine

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
	"github.com/EdgeCast/ec-sdk-go/edgecast/logging"
)

// Rules Engine service interacts with the EdgeCast API for managing Rules
type RulesEngineService struct {
	client.Client
	Logger logging.Logger
}

// New creates a new Rules Engine service
func New(config edgecast.SDKConfig) (*RulesEngineService, error) {

	authProvider, err := auth.NewIDSAuthorizationProvider(
		config.BaseIDSURL,
		config.IDSCredentials,
	)

	if err != nil {
		return nil, fmt.Errorf("rulesengine.New(): %v", err)
	}

	c := client.NewClient(client.ClientConfig{
		AuthProvider: authProvider,
		BaseAPIURL:   config.BaseAPIURLLegacy,
		UserAgent:    config.UserAgent,
		Logger:       config.Logger,
	})

	return &RulesEngineService{
		Client: c,
		Logger: config.Logger,
	}, nil
}
