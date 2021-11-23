// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package auth

import (
	"testing"
)

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
		provider := TokenAuthorizationProvider{
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
