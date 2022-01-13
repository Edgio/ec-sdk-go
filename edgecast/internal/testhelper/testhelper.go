// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package testhelper

import (
	"fmt"
	"net/url"
	"time"

	"github.com/EdgeCast/ec-sdk-go/edgecast/logging"
)

// GetTestLoger creates a log file for a test run
func GetTestLogger(prefix string) logging.Logger {
	timeNowUTC := time.Now().UTC()
	timeStamp := timeNowUTC.Format("20060102150405")
	filePath := fmt.Sprintf("%s_%s.log", prefix, timeStamp)
	return logging.NewFileLogger(filePath)
}

// URLParse provides a quick pass-through to url.Parse that will discard errors.
// This allows us to use in-line url.URLs in new struct declarations
func URLParse(urlRaw string) *url.URL {
	result, err := url.Parse(urlRaw)
	if err != nil {
		panic(err)
	}
	return result
}
