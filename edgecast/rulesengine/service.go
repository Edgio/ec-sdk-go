// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package rulesengine

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/eclog"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecauth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// Rules Engine service interacts with the EdgeCast API for managing Rules
type RulesEngineService struct {
	client ecclient.APIClient
	logger eclog.Logger
}

// New creates a new Rules Engine service
func New(config edgecast.SDKConfig) (*RulesEngineService, error) {

	authProvider, err := ecauth.NewIDSAuthorizationProvider(
		config.BaseIDSURL,
		ecauth.OAuth2Credentials{
			ClientID:     config.IDSCredentials.ClientID,
			ClientSecret: config.IDSCredentials.ClientSecret,
			Scope:        config.IDSCredentials.Scope,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("rulesengine.New(): %v", err)
	}

	c := ecclient.New(ecclient.ClientConfig{
		AuthProvider: authProvider,
		BaseAPIURL:   config.BaseAPIURL,
		UserAgent:    config.UserAgent,
		Logger:       config.Logger,
	})

	return &RulesEngineService{
		client: c,
		logger: config.Logger,
	}, nil
}
