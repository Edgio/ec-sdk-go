// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// Defines a way of getting an IDS Token
type IDSClient interface {
	GetIDSToken(credentials IDSCredentials) (*GetIDSTokenResponse, error)
}

// Calls the IDS token endpoint
type DefaultIDSClient struct {
	IDSBaseUrl *url.URL
}

// NewDefaultIDSClient -
func NewDefaultIDSClient() DefaultIDSClient {
	idsURL, _ := url.Parse(idsAddressDefault)
	return DefaultIDSClient{IDSBaseUrl: idsURL}
}

// Gets a new token from the IDS Token Endpoint
func (c DefaultIDSClient) GetIDSToken(credentials IDSCredentials) (*GetIDSTokenResponse, error) {
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

	tokenResponse := &GetIDSTokenResponse{}
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)

	if err != nil {
		return nil, err
	}

	return tokenResponse, nil
}

// GetIDSTokenResponse -
type GetIDSTokenResponse struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
	TokenType   string  `json:"token_type"`
	Scope       string  `json:"scope"`
}
