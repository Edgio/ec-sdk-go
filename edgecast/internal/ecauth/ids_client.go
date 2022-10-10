// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package ecauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const errorPrefix string = "authentication error:"

// Calls the IDS token endpoint
type IDSClient struct {
	IDSBaseUrl *url.URL
}

// NewIDSClientWithURL -
func NewIDSClient(baseURL url.URL) IDSClient {
	return IDSClient{IDSBaseUrl: &baseURL}
}

// Gets a new token from the IDS Token Endpoint
func (c IDSClient) GetToken(
	credentials OAuth2Credentials,
) (*OAuth2TokenResponse, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Add("scope", credentials.Scope)
	data.Add("client_id", credentials.ClientID)
	data.Add("client_secret", credentials.ClientSecret)
	idsTokenEndpoint := fmt.Sprintf("%s/connect/token", c.IDSBaseUrl.String())
	dataString := data.Encode()
	newTokenRequest, err := http.NewRequest(
		"POST",
		idsTokenEndpoint,
		bytes.NewBufferString(dataString))

	if err != nil {
		return nil, fmt.Errorf("%s failed creating HTTP request: %w",
			errorPrefix, err)
	}

	newTokenRequest.Header.Add(
		"Content-Type",
		"application/x-www-form-urlencoded")
	newTokenRequest.Header.Add("Cache-Control", "no-cache")
	newTokenRequest.Header.Add("Content-Length", strconv.Itoa(len(dataString)))
	httpClient := &http.Client{}
	resp, err := httpClient.Do(newTokenRequest)

	if err != nil {
		return nil, fmt.Errorf("%s HTTP request failed: %w", errorPrefix, err)
	}

	if resp.StatusCode == http.StatusBadRequest {
		oAuth2Error := &OAuth2ErrorResponse{}
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return nil, fmt.Errorf("%s error reading HTTP response: %w",
				errorPrefix, err)
		}

		err = json.Unmarshal(bodyBytes, oAuth2Error)
		if err != nil {
			// Cannot decode to oAuth2Error so return complete response body
			return nil, fmt.Errorf("%s error parsing oAuth2Error response: %s",
				errorPrefix, bodyBytes)
		}

		return nil, fmt.Errorf("%s bad request: %s",
			errorPrefix, oAuth2Error.Error)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"%s expected 200 OK, received status code %d", errorPrefix,
			resp.StatusCode)
	}

	tokenResponse := &OAuth2TokenResponse{}
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)

	if err != nil {
		return nil, fmt.Errorf("%s error decoding token response: %w",
			errorPrefix, err)
	}

	return tokenResponse, nil
}
