// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package testhelper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/EdgeCast/ec-sdk-go/edgecast/eclog"
)

// GetTestLoger creates a log file for a test run
func GetTestLogger(prefix string) eclog.Logger {
	timeNowUTC := time.Now().UTC()
	timeStamp := timeNowUTC.Format("20060102150405")
	filePath := fmt.Sprintf("%s_%s.log", prefix, timeStamp)
	return eclog.NewFileLogger(filePath)
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

func EmptyPointerFloat64() *float64 {
	var v float64
	return &v
}

func EmptyPointerString() *string {
	var v string
	return &v
}

func WrapStringInPointer(v string) *string {
	return &v
}

func WrapDurationInPointer(v time.Duration) *time.Duration {
	return &v
}

func WrapIntInPointer(v int) *int {
	return &v
}

func ToIOReadCloser(v interface{}) io.ReadCloser {
	str := ToString(v)
	return io.NopCloser(strings.NewReader(str))
}

func ToString(v interface{}) string {
	switch d := v.(type) {
	case string:
		return d
	case int:
		return strconv.Itoa(d)
	default:
		return fmt.Sprintf("%+v", d)
	}
}

func ToJSONBytes(v interface{}) []byte {
	bytes, _ := json.Marshal(v)
	return bytes
}

func TypeEqual(actual interface{}, expected interface{}) bool {
	actualType := reflect.TypeOf(actual)
	expectedType := reflect.TypeOf(expected)
	return actualType == expectedType
}

func JSONEqual(a string, b string) bool {
	var j, j2 interface{}
	if err := json.Unmarshal([]byte(a), &j); err != nil {
		return false
	}
	if err := json.Unmarshal([]byte(b), &j2); err != nil {
		return false
	}
	return reflect.DeepEqual(j, j2)
}
