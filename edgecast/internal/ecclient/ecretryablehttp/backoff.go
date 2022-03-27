// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package ecretryablehttp

/*
	This file contains backoff strategies for calling APIs
*/

import (
	"net/http"
	"strconv"
	"time"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/mathhelper"
)

// ExponentialJitterBackoff calculates exponential backoff, with jitter,
// based on the attempt number and limited by the provided
// minimum and maximum durations.
//
// It also tries to parse Retry-After response header when a
// http.StatusTooManyRequests (HTTP Code 429) is found in the resp parameter.
// Hence it will return the number of seconds the server states it may be
// ready to process more requests from this client.
func exponentialJitterBackoff(
	min, max time.Duration,
	attemptNum int,
	resp *http.Response,
) time.Duration {

	if resp != nil {
		if resp.StatusCode == http.StatusTooManyRequests ||
			resp.StatusCode == http.StatusServiceUnavailable {
			if s, ok := resp.Header["Retry-After"]; ok {
				if sleep, err := strconv.ParseInt(s[0], 10, 64); err == nil {
					return time.Second * time.Duration(sleep)
				}
			}
		}
	}

	return mathhelper.CalculateSleepWithJitter(min, max, attemptNum)
}
