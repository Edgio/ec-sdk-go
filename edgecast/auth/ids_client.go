// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// Calls the IDS token endpoint
type IDSClient struct {
	IDSBaseUrl *url.URL
}

// NewIDSClientWithURL -
func NewIDSClient(baseURL url.URL) IDSClient {
	return IDSClient{IDSBaseUrl: &baseURL}
}

// Gets a new token from the IDS Token Endpoint
func (c IDSClient) GetToken(credentials OAuth2Credentials) (*OAuth2TokenResponse, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Add("scope", credentials.Scope)
	data.Add("client_id", credentials.ClientID)
	data.Add("client_secret", credentials.ClientSecret)

	idsTokenEndpoint := fmt.Sprintf("%s/connect/token", c.IDSBaseUrl.String())

	dataString := data.Encode()
	newTokenRequest, err := http.NewRequest("POST", idsTokenEndpoint, bytes.NewBufferString(dataString))

	if err != nil {
		return nil, err
	}

	newTokenRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	newTokenRequest.Header.Add("Cache-Control", "no-cache")
	newTokenRequest.Header.Add("Content-Length", strconv.Itoa(len(dataString)))

	httpClient := &http.Client{}
	resp, err := httpClient.Do(newTokenRequest)

	if err != nil {
		return nil, err
	}

	tokenResponse := &OAuth2TokenResponse{}
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)

	if err != nil {
		return nil, err
	}

	return tokenResponse, nil
}
