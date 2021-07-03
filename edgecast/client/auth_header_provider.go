// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.
package client

import (
	"errors"
	"fmt"
	"time"
)

const (
	legacyAuthHeaderFormat string = "TOK:%s"
	idsAddressDefault      string = "https://id.vdms.io"
	idsAuthHeaderFormat    string = "Bearer %s"
)

// AuthorizationHeaderProvider defines structs that can provide Authorization headers
type AuthorizationHeaderProvider interface {
	GetAuthorizationHeader() (string, error)
}

// LegacyAuthorizationHeaderProvider creates authorization header values for legacy EdgeCast API calls
type LegacyAuthorizationHeaderProvider struct {
	APIToken string
}

// GetAuthorizationHeader gets the authorization header value for legacy EdgeCast API calls
func (lp *LegacyAuthorizationHeaderProvider) GetAuthorizationHeader() (string, error) {
	if len(lp.APIToken) == 0 {
		return "", errors.New("GetAuthorizationHeader: API Token is required")
	}

	return fmt.Sprintf(legacyAuthHeaderFormat, lp.APIToken), nil
}

// NewLegacyAuthorizationHeaderProvider -
func NewLegacyAuthorizationHeaderProvider(apiToken string) (*LegacyAuthorizationHeaderProvider, error) {
	if len(apiToken) == 0 {
		return nil, errors.New("NewLegacyAuthorizationHeaderProvider: API Token is required")
	}

	return &LegacyAuthorizationHeaderProvider{APIToken: apiToken}, nil
}

// IDSToken holds the OAuth 2.0 token for calling EdgeCast APIs
type IDSToken struct {
	AccessToken    string
	ExpirationTime time.Time
}

// Holds a customer's IDS Credentials
type IDSCredentials struct {
	ClientID     string
	ClientSecret string
	Scope        string
}

// Generates IDS Authoriation Header values
type IDSAuthorizationHeaderProvider struct {
	// The latest IDS token. May be expired.
	CurrentToken *IDSToken

	// Calls the IDS token endpoint for new tokens
	IDSClient IDSClient

	Credentials IDSCredentials
}

// Creates a new IDSAuthorizationHeaderProvider with the given credentials
func NewIDSAuthorizationHeaderProvider(credentials IDSCredentials) (*IDSAuthorizationHeaderProvider, error) {
	if len(credentials.ClientID) == 0 || len(credentials.ClientSecret) == 0 || len(credentials.Scope) == 0 {
		return nil, errors.New("NewIDSAuthorizationHeaderProvider: Client ID, Secret, and Scope required")
	}

	return &IDSAuthorizationHeaderProvider{
		Credentials: credentials,
		IDSClient:   NewDefaultIDSClient(),
	}, nil
}

// GetAuthorizationHeader creates an authorization value for the current token, refreshing it if it has expired.
// Used for EdgeCast APIs that use IDS OAuth 2.0 tokens.
func (ip *IDSAuthorizationHeaderProvider) GetAuthorizationHeader() (string, error) {
	if ip.CurrentToken == nil || ip.CurrentToken.ExpirationTime.Before(time.Now()) {

		model, err := ip.IDSClient.GetIDSToken(ip.Credentials)

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

	return fmt.Sprintf(idsAuthHeaderFormat, ip.CurrentToken.AccessToken), nil
}
