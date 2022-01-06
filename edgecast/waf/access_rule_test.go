// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package waf

import (
	"errors"
	"reflect"
	"testing"

	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
)

func TestAccessRuleActions(t *testing.T) {
	// The actions all follow the same basic pattern, so we'll use one test func
	cases := []struct {
		name          string
		doFn          func(params client.DoParams) (interface{}, error)
		input         AddAccessRuleParams
		expected      *AddAccessRuleResponse
		expectedError bool
	}{
		{
			name: "Happy Path",
			doFn: func(params client.DoParams) (interface{}, error) {
				return &AddAccessRuleResponse{AddRuleResponse: AddRuleResponse{ID: "123", WAFResponse: WAFResponse{Success: true, Status: "Success", Errors: make([]WAFError, 0)}}}, nil
			},

			input:         AddAccessRuleParams{},
			expected:      &AddAccessRuleResponse{AddRuleResponse: AddRuleResponse{ID: "123", WAFResponse: WAFResponse{Success: true, Status: "Success", Errors: make([]WAFError, 0)}}},
			expectedError: false,
		},
		{
			name: "Error Path - Unexpected return object",
			doFn: func(params client.DoParams) (interface{}, error) {
				return "this is not the correct type!", nil
			},
			expectedError: true,
		},
		{
			name: "Error Path - error returned",
			doFn: func(params client.DoParams) (interface{}, error) {
				return nil, errors.New("some error occurred!")
			},
			expectedError: true,
		},
	}

	for _, c := range cases {
		client := mockClient{
			doFn: c.doFn,
		}
		svc := buildTestWAFService(client)

		actual, err := svc.AddAccessRule(c.input)
		if c.expectedError && err == nil {
			t.Fatalf("Case '%s': expected an error, but got none", c.name)
			return
		}
		if !reflect.DeepEqual(c.expected, actual) {
			t.Fatalf("%s: Expected %+v but got %+v", c.name, c.expected, actual)
		}
	}
}
