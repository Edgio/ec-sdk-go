// Code generated by the Code Generation Kit (CGK) using OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
//
// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

/*
WAF API

The WAF API is a RESTful server application for managing customer configuration settings.

API version: 1.0
*/

package waf_bot_manager

import (
	"fmt"
	"net/url"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/eclog"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecauth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

const (
	DefaultBasePath string = "/v2/mcc/customers"
)

// Service manages
// communication with the WAF API API v1.0
type Service struct {
	client ecclient.APIClient

	clientConfig ecclient.ClientConfig

	Logger eclog.Logger

	BotManagers BotManagersClientService
}

// New creates a new Service
func New(config edgecast.SDKConfig) (*Service, error) {
	apiURL, err := url.Parse(config.BaseAPIURLLegacy.String() + DefaultBasePath)
	if err != nil {
		return nil, fmt.Errorf("waf_bot_manager.New(): %w", err)
	}

	var auth ecauth.AuthorizationProvider

	auth, err = ecauth.NewIDSAuthorizationProvider(
		config.BaseIDSURL,
		ecauth.OAuth2Credentials(config.IDSCredentials))
	if err != nil {
		// Fall back to token authentication
		auth, err = ecauth.NewTokenAuthorizationProvider(config.APIToken)
		if err != nil {
			return nil,
				fmt.Errorf("error initializing waf_bot_manager Service: %w", err)
		}
	}

	c := ecclient.New(ecclient.ClientConfig{
		BaseAPIURL:   *apiURL,
		UserAgent:    config.UserAgent,
		Logger:       config.Logger,
		AuthProvider: auth,
	})

	return &Service{
		client:      c,
		Logger:      config.Logger,
		BotManagers: NewBotManagersClient(c, c.Config.BaseAPIURL.String()),
	}, nil
}
