// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package ecclient

import (
	"reflect"
	"strings"
	"testing"

	"github.com/EdgeCast/ec-sdk-go/edgecast/eclog"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient/ecretryablehttp"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/testhelper"
)

func TestHTTPMethodString(t *testing.T) {
	cases := []struct {
		input    HTTPMethod
		expected string
	}{
		{
			input:    Connect,
			expected: "CONNECT",
		},
		{
			input:    Delete,
			expected: "DELETE",
		},
		{
			input:    Get,
			expected: "GET",
		},
		{
			input:    Head,
			expected: "HEAD",
		},
		{
			input:    Options,
			expected: "OPTIONS",
		},
		{
			input:    Patch,
			expected: "PATCH",
		},
		{
			input:    Post,
			expected: "POST",
		},
		{
			input:    Put,
			expected: "PUT",
		},
		{
			input:    Trace,
			expected: "TRACE",
		},
		{
			input:    HTTPMethod(99),
			expected: "UNKNOWN",
		},
	}

	for _, c := range cases {
		actual := c.input.String()
		if strings.Compare(c.expected, actual) != 0 {
			t.Fatalf("Expected %+v but got %+v", c.expected, actual)
		}
	}
}

func TestNewECClient(t *testing.T) {
	config := ClientConfig{
		Logger:       eclog.NewNullLogger(),
		AuthProvider: testAuthProvider{},
		BaseAPIURL:   *testhelper.URLParse("http://example.com"),
		UserAgent:    "TestNewECClient",
	}

	actual := New(config)
	if !reflect.DeepEqual(config, actual.config) {
		t.Fatalf("Expected %+v but got %+v", config, actual.config)
	}
	if !testhelper.TypeEqual(ecRequestBuilder{}, actual.reqBuilder) {
		t.Fatalf("Expected Type: ecRequestBuilder but got %T", actual.reqBuilder)
	}
	if !testhelper.TypeEqual(ecRequestSender{}, actual.reqSender) {
		t.Fatalf("Expected Type: ecRequestSender but got %T", actual.reqSender)
	}

	b := actual.reqBuilder.(ecRequestBuilder)
	if !reflect.DeepEqual(config.AuthProvider, *b.authProvider) {
		t.Fatalf("Expected %+v but got %+v", config.AuthProvider, *b.authProvider)
	}
	if !reflect.DeepEqual(config.BaseAPIURL, b.baseAPIURL) {
		t.Fatalf("Expected %+v but got %+v", config.BaseAPIURL, b.baseAPIURL)
	}
	if !reflect.DeepEqual(config.Logger, b.logger) {
		t.Fatalf("Expected %+v but got %+v", config.Logger, b.logger)
	}
	if !reflect.DeepEqual(config.UserAgent, b.userAgent) {
		t.Fatalf("Expected %+v but got %+v", config.UserAgent, b.userAgent)
	}

	s := actual.reqSender.(ecRequestSender)
	if !reflect.DeepEqual(config.Logger, s.logger) {
		t.Fatalf("Expected %+v but got %+v", config.Logger, s.logger)
	}
	if !testhelper.TypeEqual(jsonBodyParser{}, s.parser) {
		t.Fatalf("Expected Type: jsonBodyParser but got %T", s.parser)
	}
	if !testhelper.TypeEqual(ecretryablehttp.RetryableHTTPClientAdapter{}, s.clientAdapter) {
		t.Fatalf("Expected Type: ecretryablehttp.RetryableHTTPClientAdapter but got %T", s.clientAdapter)
	}
}
