// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package ecretryablehttp

/*
	This file contains backoff strategies for calling APIs
*/

import (
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
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

	// calculate the initial sleep period before jitter
	// attemptNum starts at 0 so we add 1
	sleep := math.Pow(2, float64(attemptNum+1)) * float64(min)

	// The final sleep time will be a random number between sleep/2 and sleep
	sleepWithJitter := sleep/2 + randBetween(0, sleep/2)

	if sleepWithJitter > float64(max) {
		return max
	}

	return time.Duration(sleepWithJitter)
}

func randBetween(min float64, max float64) float64 {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + rand.Float64()*(max-min)
}
