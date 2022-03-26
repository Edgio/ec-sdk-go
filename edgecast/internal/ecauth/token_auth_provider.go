// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package ecauth

import (
	"errors"
	"fmt"
)

const (
	tokenAuthHeaderFormat string = "TOK:%s"
)

// TokenAuthorizationProvider creates authorization header values for legacy
// EdgeCast API calls
type TokenAuthorizationProvider struct {
	APIToken string
}

// GetAuthorizationHeader gets the authorization header value for EdgeCast APIs
// that require an API Token
func (lp *TokenAuthorizationProvider) GetAuthorizationHeader() (string, error) {
	if len(lp.APIToken) == 0 {
		return "", errors.New("api token is required")
	}

	return fmt.Sprintf(tokenAuthHeaderFormat, lp.APIToken), nil
}

// NewTokenAuthorizationProvider -
func NewTokenAuthorizationProvider(
	apiToken string,
) (*TokenAuthorizationProvider, error) {
	if len(apiToken) == 0 {
		return nil, errors.New("api token is required")
	}

	return &TokenAuthorizationProvider{APIToken: apiToken}, nil
}
