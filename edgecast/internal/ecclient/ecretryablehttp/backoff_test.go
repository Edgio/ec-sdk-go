// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package ecretryablehttp

import (
	"net/http"
	"testing"
	"time"
)

func TestExponentialJitterBackoff(t *testing.T) {
	cases := []struct {
		name       string
		min        time.Duration
		max        time.Duration
		attemptNum int
		resp       *http.Response
		expected   time.Duration
	}{
		{
			name: "Too Many Requests Server Response",
			resp: &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Header: map[string][]string{
					"Retry-After": {"10"},
				},
			},
			expected: 10 * time.Second, // 10 seconds
		},
		{
			name: "Service Unavailable Server Response",
			resp: &http.Response{
				StatusCode: http.StatusServiceUnavailable,
				Header: map[string][]string{
					"Retry-After": {"10"},
				},
			},
			expected: 10 * time.Second, // 10 seconds
		},
		{
			name:       "Testing with Random/Jitter Path",
			resp:       nil,
			min:        20 * time.Second, // Min more than max to force max
			max:        1 * time.Second,
			attemptNum: 3,
			expected:   1000000000,
		},
	}

	for _, c := range cases {
		actual := exponentialJitterBackoff(c.min, c.max, c.attemptNum, c.resp)
		if c.expected != actual {
			t.Fatalf("'%s': Expected %d but got %d", c.name, c.expected, actual)
		}
	}
}
