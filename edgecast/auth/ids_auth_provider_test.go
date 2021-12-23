// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package auth

import (
	"testing"
	"time"
)

func TestGetAuthorizationHeader_IDS(t *testing.T) {
	cases := []struct {
		Name          string
		CurrentToken  *IDSToken
		NewToken      *OAuth2TokenResponse
		Expected      string
		ExpectedError bool
	}{
		{
			Name:         "No cached token - new token retrieved",
			CurrentToken: nil,
			NewToken: &OAuth2TokenResponse{
				AccessToken: "abcd",
				ExpiresIn:   100000,
			},
			Expected:      "Bearer abcd",
			ExpectedError: false,
		},
		{
			Name: "Expired token - new token retrieved",
			CurrentToken: &IDSToken{
				AccessToken:    "abcd",
				ExpirationTime: time.Now().Add(-1),
			},
			NewToken: &OAuth2TokenResponse{
				AccessToken: "efgh",
				ExpiresIn:   100000,
			},
			Expected:      "Bearer efgh",
			ExpectedError: false,
		},
		{
			Name: "Non-Expired token - same token retrieved",
			CurrentToken: &IDSToken{
				AccessToken:    "abcd",
				ExpirationTime: time.Now().Add(time.Hour * 24),
			},
			NewToken:      nil,
			Expected:      "Bearer abcd",
			ExpectedError: false,
		},
		{
			Name:          "New token retrieval failure",
			CurrentToken:  nil,
			NewToken:      nil,
			Expected:      "",
			ExpectedError: true,
		},
	}

	for _, v := range cases {
		provider := IDSAuthorizationProvider{
			CurrentToken: v.CurrentToken,
			TokenClient: TestIDSClient{
				StaticToken: v.NewToken,
			},
		}

		actual, err := provider.GetAuthorizationHeader()

		if v.ExpectedError {
			if err == nil {
				t.Fatalf("Failed for case: '%+v'. Expected an error, but did not get one", v.Name)
			}
		} else {
			if v.Expected != actual {
				t.Fatalf("Failed for case: '%+v'. Expected '%s', but got '%s'", v.Name, v.Expected, actual)
			}
		}
	}
}

// A test client that implements OAuth2Client
type TestIDSClient struct {
	StaticToken *OAuth2TokenResponse
}

// GetToken is a test implementation that returns the same token every time
func (c TestIDSClient) GetToken(
	credentials OAuth2Credentials,
) (*OAuth2TokenResponse, error) {
	return c.StaticToken, nil
}
