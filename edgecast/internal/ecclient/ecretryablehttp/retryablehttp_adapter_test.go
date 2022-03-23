// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package ecretryablehttp

import (
	"context"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/testhelper"
	"github.com/hashicorp/go-retryablehttp"
)

func TestNewRetryableHTTPClientAdapterDefaults(t *testing.T) {
	actual := NewRetryableHTTPClientAdapter(RetryConfig{})

	if !reflect.DeepEqual(
		DefaultRetryWaitMinSeconds*time.Second,
		actual.RetryableHttpClient.RetryWaitMin) {
		t.Fatalf(
			"Expected %+v but got %+v",
			DefaultRetryWaitMinSeconds*time.Second,
			actual.RetryableHttpClient.RetryWaitMin)
	}
	if !reflect.DeepEqual(
		DefaultRetryWaitMaxSeconds*time.Second,
		actual.RetryableHttpClient.RetryWaitMax) {
		t.Fatalf(
			"Expected %+v but got %+v",
			DefaultRetryWaitMaxSeconds*time.Second,
			actual.RetryableHttpClient.RetryWaitMax)
	}
	if !reflect.DeepEqual(
		DefaultRetryMax,
		actual.RetryableHttpClient.RetryMax) {
		t.Fatalf(
			"Expected %+v but got %+v",
			DefaultRetryMax,
			actual.RetryableHttpClient.RetryMax)
	}
	if actual.HasCustomRetry {
		t.Fatalf(
			"Expected %+v but got %+v",
			false,
			actual.HasCustomRetry)
	}
}

func TestNewRetryableHTTPClientAdapter(t *testing.T) {

	config := RetryConfig{
		RetryWaitMin: testhelper.WrapDurationInPointer(99 * time.Second),
		RetryWaitMax: testhelper.WrapDurationInPointer(899 * time.Second),
		RetryMax:     testhelper.WrapIntInPointer(20),
		CheckRetry: func(
			ctx context.Context,
			resp *http.Response,
			err error,
		) (bool, error) {
			return true, nil
		},
	}

	actual := NewRetryableHTTPClientAdapter(config)
	if !reflect.DeepEqual(
		*config.RetryWaitMin,
		actual.RetryableHttpClient.RetryWaitMin) {
		t.Fatalf(
			"Expected %+v but got %+v",
			*config.RetryWaitMin,
			actual.RetryableHttpClient.RetryWaitMin)
	}
	if !reflect.DeepEqual(
		*config.RetryWaitMax,
		actual.RetryableHttpClient.RetryWaitMax) {
		t.Fatalf(
			"Expected %+v but got %+v",
			*config.RetryWaitMax,
			actual.RetryableHttpClient.RetryWaitMax)
	}
	if !reflect.DeepEqual(
		*config.RetryMax,
		actual.RetryableHttpClient.RetryMax) {
		t.Fatalf(
			"Expected %+v but got %+v",
			*config.RetryMax,
			actual.RetryableHttpClient.RetryMax)
	}
	if !actual.HasCustomRetry {
		t.Fatalf(
			"Expected %+v but got %+v",
			true,
			actual.HasCustomRetry)
	}
}

func TestSetHeaders(t *testing.T) {
	cases := []struct {
		name    string
		headers map[string]string
	}{
		{
			name: "Happy Path",
			headers: map[string]string{
				"Authorization": "auth-key",
				"Header2":       "v2",
				"header3":       "v3", // should get title-cased to "Header3"
			},
		},
		{
			name:    "Empty headers",
			headers: map[string]string{},
		},
	}

	for _, c := range cases {
		req, _ := retryablehttp.NewRequest("GET", "http://edgecast.com", nil)
		setHeaders(req, c.headers)

		for k, v := range c.headers {
			// Headers get title-cased
			e := req.Header[strings.Title(k)]
			if len(e) == 0 || e[0] != v {
				t.Fatalf("%s: {%s,%s} not found in headers", c.name, k, v)
			}
		}
	}
}
