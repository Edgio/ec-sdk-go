// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package auth

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

const (
	bearerAuthHeaderFormat string = "Bearer %s"
)

// IDSToken holds an authorization token and its expiration time
type IDSToken struct {
	AccessToken    string
	ExpirationTime time.Time
}

// Generates IDS Authorization Header values
type IDSAuthorizationProvider struct {
	// The latest token. May be expired.
	CurrentToken *IDSToken

	// TokenClient retrievees new tokens
	TokenClient OAuth2Client

	// The Credentials used to retrieve new tokens
	Credentials OAuth2Credentials
}

// Creates a new IDSAuthorizationProvider with the given credentials
// that retrieves tokens from the specified URL
func NewIDSAuthorizationProvider(baseIDSURL url.URL, credentials OAuth2Credentials) (*IDSAuthorizationProvider, error) {
	if len(credentials.ClientID) == 0 || len(credentials.ClientSecret) == 0 || len(credentials.Scope) == 0 {
		return nil, errors.New("NewIDSAuthorizationProvider: Client ID, Secret, and Scope required")
	}

	return &IDSAuthorizationProvider{
		Credentials: credentials,
		TokenClient: NewIDSClient(baseIDSURL),
	}, nil
}

// GetAuthorizationHeader creates an authorization header value for the current token, refreshing it if it has expired.
// Used for EdgeCast APIs that use IDS OAuth 2.0 tokens.
func (ip *IDSAuthorizationProvider) GetAuthorizationHeader() (string, error) {

	// If there is no cached token or it's expired, get a new one
	if ip.CurrentToken == nil || ip.CurrentToken.ExpirationTime.Before(time.Now()) {

		model, err := ip.TokenClient.GetToken(ip.Credentials)

		if err != nil {
			return "", err
		}

		if model == nil {
			return "", fmt.Errorf("no access token retrieved, please check your IDS credentials")
		}

		expiresIn := time.Second * time.Duration(model.ExpiresIn)

		ip.CurrentToken = &IDSToken{
			AccessToken:    model.AccessToken,
			ExpirationTime: time.Now().Add(expiresIn),
		}
	}

	return fmt.Sprintf(bearerAuthHeaderFormat, ip.CurrentToken.AccessToken), nil
}
