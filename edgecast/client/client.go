// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/VerizonDigital/ec-sdk-go/edgecast/internal/collections"
	"github.com/VerizonDigital/ec-sdk-go/edgecast/internal/jsonutil"
	"github.com/hashicorp/go-retryablehttp"
)

// LiteralResponse is used for unmarshaling response data that is in an unrecognized format
type LiteralResponse struct {
	Value interface{}
}

// Client is the primary means for services to interact with the EdgeCast API
type Client struct {
	// Config holds the configuration values for this client
	Config ClientConfig

	// Internal HTTP client
	HTTPClient *retryablehttp.Client
}

// Creates a new client pointing to EdgeCast APIs
func NewClient(config ClientConfig) *Client {
	// Use PassthroughErrorHandler so that retryablehttp.Client does not obscure API errors
	httpClient := retryablehttp.NewClient()
	httpClient.ErrorHandler = retryablehttp.PassthroughErrorHandler

	return &Client{
		Config:     config,
		HTTPClient: httpClient,
	}
}

// BuildRequest creates a new Request for the Edgecast API, adding appropriate headers
func (c *Client) BuildRequest(method, path string, body interface{}) (*retryablehttp.Request, error) {
	relativeURL, err := url.Parse(path)

	if err != nil {
		return nil, fmt.Errorf("BuildRequest: url.Parse: %v", err)
	}

	absoluteURL := c.Config.BaseAPIURL.ResolveReference(relativeURL)

	var payload interface{}

	if body != nil {
		switch b := body.(type) {
		case string:
			payload = []byte(b)
		default:
			buf := new(bytes.Buffer)
			err := json.NewEncoder(buf).Encode(body)
			if err != nil {
				return nil, err
			}
			logMsg := jsonutil.CreateRequestBodyLogMessage(method, absoluteURL.String(), body)
			c.Config.Logger.LogDebug(logMsg)
			payload = buf

		}
	}

	req, err := retryablehttp.NewRequest(method, absoluteURL.String(), payload)

	if err != nil {
		return nil, fmt.Errorf("BuildRequest: NewRequest: %v", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")

	authHeader, err := c.Config.AuthHeaderProvider.GetAuthorizationHeader()

	if err != nil {
		return nil, fmt.Errorf("BuildRequest: Failed to get authorization: %v", err)
	}

	req.Header.Set("Authorization", authHeader)
	req.Header.Set("User-Agent", c.Config.UserAgent)

	return req, nil
}

// SendRequest sends an HTTP request and, if applicable, sets the response to parsedResponse
func (c *Client) SendRequest(req *retryablehttp.Request, parsedResponse interface{}) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("SendRequest: Do: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	bodyAsString := string(body)
	jsonutil.CreateJSONLogMessage("Response", bodyAsString)
	if resp.StatusCode >= 400 && resp.StatusCode <= 599 {
		if err != nil {
			return nil, fmt.Errorf("SendRequest: ioutil.ReadAll: %v", err)
		}

		return nil, fmt.Errorf("SendRequest failed: %s", bodyAsString)
	}
	if parsedResponse == nil {
		return nil, nil
	}

	var f interface{}
	if jsonUnmarshalErr := json.Unmarshal(body, &f); err != nil {
		return nil, fmt.Errorf("malformed Json response:%v", jsonUnmarshalErr)
	}

	if collections.IsInterfaceArray(f) {
		if jsonArryErr := json.Unmarshal([]byte(body), parsedResponse); jsonArryErr != nil {
			return nil, fmt.Errorf("malformed Json Array response:%v", jsonArryErr)
		}
	} else {
		if jsonutil.IsJSONString(bodyAsString) {
			err = json.Unmarshal([]byte(bodyAsString), parsedResponse)

			if err != nil {
				return nil, fmt.Errorf("SendRequest: Decode error: %v", err)
			}
		} else {

			// if response is not json string
			switch v := parsedResponse.(type) {
			case LiteralResponse:
				rs, ok := parsedResponse.(LiteralResponse)
				if ok {
					rs.Value = bodyAsString
					parsedResponse = rs
				}
			case float64:
				fmt.Println("float64:", v)
			default:
				fmt.Println("unknown")
			}

		}
	}
	return resp, nil
}

// SendRequest sends an HTTP request and, if applicable, sets the response to parsedResponse
func (c *Client) SendRequestWithStringResponse(req *retryablehttp.Request) (*string, error) {
	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("SendRequest: Do: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	bodyAsString := string(body)

	if resp.StatusCode >= 400 && resp.StatusCode <= 599 {
		if err != nil {
			return nil, fmt.Errorf("SendRequest: ioutil.ReadAll: %v", err)
		}

		return nil, fmt.Errorf("SendRequest failed: %s", bodyAsString)
	}

	c.Config.Logger.LogDebug("SendRequest: Response Body: %s", body)

	return &bodyAsString, nil
}
