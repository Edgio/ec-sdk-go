// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.

package jsonutil

import (
	"bytes"
	"encoding/json"
	"log"
)

// LogRequestBody logs an JSON API request body, method, url in json pretty format.
func LogRequestBody(method string, url string, body interface{}) {
	fb, _ := json.MarshalIndent(body, "", "    ")
	// Read the Body content
	log.Print("=====================================================================")
	log.Printf("[[[REQUEST-URI]]]:[%s] %s", method, url)
	log.Printf("[[[REQUEST-BODY]]]:%s", fb)
	log.Print("=====================================================================")
}

// LogPrettyJson logs a json string with pretty format with a message
func LogPrettyJson(message string, jsonString string) {

	log.Print("=====================================================================")
	log.Printf("[[[%s]]]:", message)
	log.Printf("[[[Json]]]:%s", jsonPrettyPrint(jsonString))
	log.Print("=====================================================================")
}

// IsJSONString determines whether a string is JSON
func IsJSONString(s string) bool {
	var js string
	return json.Unmarshal([]byte(s), &js) != nil
}

// Make json string formatted in terraform.log
func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}
