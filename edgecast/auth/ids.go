// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.

package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const idsAddressDefault string = "https://id.vdms.io"

// IDSToken holds the OAuth 2.0 token for calling EdgeCast APIs
type IDSToken struct {
	AccessToken    string
	ExpirationTime time.Time
}

type IDSAuthProvider struct {
	IDSURL       *url.URL
	ClientID     string
	ClientSecret string
	Scope        string
	idsToken     *IDSToken
}

func DefaultIDSAuthProvider(clientID string, clientSecret string, scope string) *IDSAuthProvider {
	idsURL, _ := url.Parse(idsAddressDefault)

	return &IDSAuthProvider{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        scope,
		IDSURL:       idsURL,
	}
}

// GetIdsToken returns the cached token, refreshing it if it has expired
func (ip *IDSAuthProvider) GetAuthorization() (string, error) {

	if len(ip.ClientID) == 0 || len(ip.ClientSecret) == 0 || len(ip.Scope) == 0 {
		return "", errors.New("GetIdsToken: Client ID, Secret, and Scope required")
	}

	if ip.idsToken == nil || ip.idsToken.ExpirationTime.Before(time.Now()) {
		data := url.Values{}
		data.Set("grant_type", "client_credentials")
		data.Add("scope", ip.Scope)
		data.Add("client_id", ip.ClientID)
		data.Add("client_secret", ip.ClientSecret)

		idsTokenEndpoint := fmt.Sprintf("%s/connect/token", ip.IDSURL.String())
		newTokenRequest, err := http.NewRequest("POST", idsTokenEndpoint, bytes.NewBufferString(data.Encode()))

		if err != nil {
			return "", fmt.Errorf("IDSAuthProvider.GetAuthorization: NewRequest: %v", err)
		}

		newTokenRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		newTokenRequest.Header.Add("Cache-Control", "no-cache")
		newTokenRequest.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

		httpClient := &http.Client{}
		newTokenResponse, err := httpClient.Do(newTokenRequest)

		if err != nil {
			return "", fmt.Errorf("GetIdsToken: Do: %v", err)
		}

		var tokenMap map[string]interface{}
		err = json.NewDecoder(newTokenResponse.Body).Decode(&tokenMap)
		if err != nil {
			return "", fmt.Errorf("GetIdsToken: Decode: %v", err)
		}

		expiresIn := time.Second * time.Duration((tokenMap["expires_in"].(float64)))

		ip.idsToken = &IDSToken{
			AccessToken:    tokenMap["access_token"].(string),
			ExpirationTime: time.Now().Add(expiresIn),
		}
	}

	return "Bearer " + ip.idsToken.AccessToken, nil
}
