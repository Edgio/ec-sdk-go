// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package auth

// Holds a customer's OAuth 2.0 Credentials
type OAuth2Credentials struct {
	ClientID     string
	ClientSecret string
	Scope        string
}

// OAuth2TokenResponse represents the response from an identity server when retrieving a new token
type OAuth2TokenResponse struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
	TokenType   string  `json:"token_type"`
	Scope       string  `json:"scope"`
}

// Defines structs that can retrieve OAuth 2.0 Tokens
type OAuth2Client interface {
	GetToken(credentials OAuth2Credentials) (*OAuth2TokenResponse, error)
}
