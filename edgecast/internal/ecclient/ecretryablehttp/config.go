// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package ecretryablehttp

import (
	"context"
	"net/http"
	"time"

	"github.com/EdgeCast/ec-sdk-go/edgecast/eclog"
)

const (
	DefaultRetryWaitMinSeconds = 1
	DefaultRetryWaitMaxSeconds = 60
	DefaultRetryMax            = 5
)

type RetryConfig struct {
	// Logger -
	Logger eclog.Logger

	// The minimum wait time for retries on API errors
	RetryWaitMin *time.Duration

	// The maximum wait time for retries on API errors
	RetryWaitMax *time.Duration

	// The maximum number of retries on API errors
	RetryMax *int

	// CheckRetry is a handler that allows users to define custom logic
	// to determine whether the API Client should retry a failed API call
	CheckRetry CheckRetry
}

type CheckRetry *func(
	ctx context.Context,
	resp *http.Response,
	err error,
) (bool, error)
