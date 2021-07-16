// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package client

import (
	"testing"
	"time"
)

func TestGetAuthorizationHeader_IDS(t *testing.T) {
	cases := []struct {
		Name          string
		CurrentToken  *IDSToken
		NewToken      *GetIDSTokenResponse
		Expected      string
		ExpectedError bool
	}{
		{
			Name:         "No cached token - new token retrieved",
			CurrentToken: nil,
			NewToken: &GetIDSTokenResponse{
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
			NewToken: &GetIDSTokenResponse{
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
		provider := IDSAuthorizationHeaderProvider{
			CurrentToken: v.CurrentToken,
			IDSClient: TestIDSClient{
				NewToken: v.NewToken,
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

// A test client that always returns the same token
type TestIDSClient struct {
	NewToken *GetIDSTokenResponse
}

// Returns a static token
func (c TestIDSClient) GetIDSToken(credentials IDSCredentials) (*GetIDSTokenResponse, error) {
	return c.NewToken, nil
}

func TestGetAuthorizationHeader_Legacy(t *testing.T) {
	cases := []struct {
		APIToken      string
		Expected      string
		ExpectedError bool
	}{
		{
			APIToken:      "abcd",
			Expected:      "TOK:abcd",
			ExpectedError: false,
		},
		{
			APIToken:      "",
			Expected:      "",
			ExpectedError: true,
		},
	}

	for _, v := range cases {
		provider := LegacyAuthorizationHeaderProvider{
			APIToken: v.APIToken,
		}

		actual, err := provider.GetAuthorizationHeader()

		if v.ExpectedError {
			if err == nil {
				t.Fatalf("Expected an error, but did not get one")
			}
		} else {
			if v.Expected != actual {
				t.Fatalf("Expected '%s', but got '%s'", v.Expected, actual)
			}
		}
	}
}
