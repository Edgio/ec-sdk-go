// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

// Package jsonutil provides helper methods for working with JSON
package jsonutil

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// RequestBodyToLogMessage creates a log message from an JSON API request body, method, url in json pretty format.
func CreateRequestBodyLogMessage(method string, url string, body interface{}) string {
	fb, _ := json.MarshalIndent(body, "", "    ")
	s := fmt.Sprintf("[[[REQUEST-URI]]]:[%s] %s\n", method, url)
	s += fmt.Sprintf("[[[REQUEST-BODY]]]:%s\n", fb)
	return s
}

// ShowAsJson shows obj in json pretty format.
func ShowAsJson(objName string, body interface{}) string {
	fb, _ := json.MarshalIndent(body, "", "    ")
	s := fmt.Sprintf("Object: %s\n", objName)
	s += fmt.Sprintf("Marshall as JSON:%s\n", fb)
	return s
}

// JSONToLogMessage logs a json string with pretty format with a message
func CreateJSONLogMessage(message string, jsonString string) string {
	s := fmt.Sprintf("[[[%s]]]:", message)
	s += fmt.Sprintf("[[[Json]]]:%s", printPrettyJson(jsonString))
	return s
}

// IsJSONString determines whether a string is JSON
func IsJSONString(s string) bool {
	var js string
	return json.Unmarshal([]byte(s), &js) != nil
}

// Make json string formatted in terraform.log
func printPrettyJson(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}
