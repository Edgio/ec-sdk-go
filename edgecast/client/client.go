// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/collections"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/jsonutil"
	"github.com/hashicorp/go-retryablehttp"
)

var (
	defaultRetryWaitMin = 1 * time.Second
	defaultRetryWaitMax = 60 * time.Second
	defaultRetryMax     = 5
)

// LiteralResponse is used for unmarshaling response data
// that is in an unrecognized format
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
func NewClient(config ClientConfig) Client {
	httpClient := retryablehttp.NewClient()
	httpClient.ErrorHandler = retryablehttp.PassthroughErrorHandler
	httpClient.Logger = config.Logger
	httpClient.Backoff = ExponentialJitterBackoff

	if config.RetryWaitMin != nil {
		httpClient.RetryWaitMin = *config.RetryWaitMin
	} else {
		httpClient.RetryWaitMin = defaultRetryWaitMin
	}
	if config.RetryWaitMax != nil {
		httpClient.RetryWaitMax = *config.RetryWaitMax
	} else {
		httpClient.RetryWaitMax = defaultRetryWaitMax
	}
	if config.RetryMax != nil {
		httpClient.RetryMax = *config.RetryMax
	} else {
		httpClient.RetryMax = defaultRetryMax
	}

	return Client{
		HTTPClient: httpClient,
		Config:     config,
	}
}

// BuildRequest creates a new Request for the Edgecast API,
// adding appropriate headers
func (c Client) BuildRequest(
	method, path string,
	body interface{},
) (*retryablehttp.Request, error) {

	relativeURL, err := url.Parse(path)

	if err != nil {
		return nil, fmt.Errorf("Client.BuildRequest: url.Parse: %v", err)
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
			logMsg := jsonutil.CreateRequestBodyLogMessage(
				method,
				absoluteURL.String(),
				body)
			c.Config.Logger.Debug(logMsg)
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

	authHeader, err := c.Config.AuthProvider.GetAuthorizationHeader()

	if err != nil {
		return nil,
			fmt.Errorf("BuildRequest: Failed to get authorization: %v", err)
	}

	req.Header.Set("Authorization", authHeader)
	req.Header.Set("User-Agent", c.Config.UserAgent)

	return req, nil
}

// BuildRequest creates a new Request for the Edgecast API,
// adding appropriate headers
func (c Client) PrepareRequest(
	method, path string,
	body interface{},
	queryParams map[string]string,
	pathParams map[string]string,
) (*retryablehttp.Request, error) {

	// Adding Path Params
	for k, v := range pathParams {
		path = strings.Replace(path, "{"+k+"}", fmt.Sprintf("%v", v), -1)
	}

	relativeURL, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("Client.BuildRequest: url.Parse: %v", err)
	}

	absoluteURL := c.Config.BaseAPIURL.ResolveReference(relativeURL)

	// Adding Query Param
	query := absoluteURL.Query()
	for k, v := range queryParams {
		query.Add(k, v)
	}
	// Encode the parameters.
	absoluteURL.RawQuery = query.Encode()

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
			logMsg := jsonutil.CreateRequestBodyLogMessage(
				method,
				absoluteURL.String(),
				body)
			c.Config.Logger.Debug(logMsg)
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

	authHeader, err := c.Config.AuthProvider.GetAuthorizationHeader()

	if err != nil {
		return nil,
			fmt.Errorf("BuildRequest: Failed to get authorization: %v", err)
	}

	req.Header.Set("Authorization", authHeader)
	req.Header.Set("User-Agent", c.Config.UserAgent)

	return req, nil
}

// SendRequest sends an HTTP request and,
// if applicable, sets the response to parsedResponse
func (c *Client) SendRequest(
	req *retryablehttp.Request,
	parsedResponse interface{},
) (*http.Response, error) {

	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("SendRequest: Do: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	bodyAsString := string(body)

	logMsg := jsonutil.CreateJSONLogMessage("Response", bodyAsString)
	c.Config.Logger.Debug(logMsg)

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
		if err := json.Unmarshal([]byte(body), parsedResponse); err != nil {
			return nil, fmt.Errorf("malformed Json Array response:%v", err)
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

// SendRequest sends an HTTP request and,
// if applicable, sets the response to parsedResponse
func (c *Client) SendRequestWithStringResponse(
	req *retryablehttp.Request,
) (*string, error) {

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

	c.Config.Logger.Debug("SendRequest: Response Body: %s", body)

	return &bodyAsString, nil
}

// ExponentialJitterBackoff calculates exponential backoff, with jitter,
// based on the attempt number and limited by the provided
// minimum and maximum durations.
//
// It also tries to parse Retry-After response header when a
// http.StatusTooManyRequests (HTTP Code 429) is found in the resp parameter.
// Hence it will return the number of seconds the server states it may be
// ready to process more requests from this client.
func ExponentialJitterBackoff(
	min, max time.Duration,
	attemptNum int,
	resp *http.Response,
) time.Duration {

	if resp != nil {
		if resp.StatusCode == http.StatusTooManyRequests ||
			resp.StatusCode == http.StatusServiceUnavailable {
			if s, ok := resp.Header["Retry-After"]; ok {
				if sleep, err := strconv.ParseInt(s[0], 10, 64); err == nil {
					return time.Second * time.Duration(sleep)
				}
			}
		}
	}

	// calculate the initial sleep period before jitter
	// attemptNum starts at 0 so we add 1
	sleep := math.Pow(2, float64(attemptNum+1)) * float64(min)

	// The final sleep time will be a random number between sleep/2 and sleep
	sleepWithJitter := sleep/2 + randBetween(0, sleep/2)

	if sleepWithJitter > float64(max) {
		return max
	}

	return time.Duration(sleepWithJitter)
}

func randBetween(min float64, max float64) float64 {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + rand.Float64()*(max-min)
}
