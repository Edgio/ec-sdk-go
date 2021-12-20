// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

// Package jsonhelper provides helper methods for working with JSON
package jsonhelper

import (
	"bytes"
	"encoding/json"
)

// JSONHelper defines structs that can provide JSON helper functions
type JSONHelper interface {
	IsJSONString(s string) bool
	ConvertToJSONBuffer(b interface{}) (*bytes.Buffer, error)
	PrintPrettyJSON(in string) (string, error)
	ConvertToJSONString(in interface{}, prettyPrint bool) (string, error)
}

type ECJSONHelper struct {
}

func New() ECJSONHelper {
	return ECJSONHelper{}
}

// IsJSONString determines whether a string is JSON
func (h ECJSONHelper) IsJSONString(s string) bool {
	var js string
	return json.Unmarshal([]byte(s), &js) != nil
}

// Converts any object into a JSON representation, returned in a buffer
func (h ECJSONHelper) ConvertToJSONBuffer(b interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(b)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// Returns a nicely formatted copy of a JSON string
func (h ECJSONHelper) PrintPrettyJSON(in string) (string, error) {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in, err
	}
	return out.String(), nil
}

func (h ECJSONHelper) ConvertToJSONString(
	in interface{},
	prettyPrint bool,
) (string, error) {
	bytes, err := json.MarshalIndent(in, "", "    ")
	if err != nil {
		return "", err
	}
	s := string(bytes)
	if prettyPrint {
		return h.PrintPrettyJSON(s)
	}
	return s, nil
}
