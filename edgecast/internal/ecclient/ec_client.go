// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package ecclient

/*
	This file contains the concrete client implementation for the EC SDK
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/collectionhelper"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecauth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/jsonhelper"
)

const (
	defaultHeaderAccept      string = "application/json"
	defaultHeaderContentType string = "application/json"
)

// SubmitRequest invokes an HTTP request with the given parameters
func (c ECClient) SubmitRequest(params SubmitRequestParams) (*Response, error) {
	req, err := c.reqBuilder.buildRequest(buildRequestParams{
		method:      params.Method,
		path:        params.Path,
		rawBody:     params.RawBody,
		queryParams: params.QueryParams,
		pathParams:  params.PathParams,
		userAgent:   c.Config.UserAgent,
		headers:     params.Headers,
	})

	if err != nil {
		return nil, fmt.Errorf("SubmitRequest: %w", err)
	}

	// Provides an object to be filled in when unmarshaling the API response
	req.parsedResponse = params.ParsedResponse

	c.Config.Logger.Debug(
		"[REQUEST-URI]:[%s] %s\n",
		req.method,
		req.url.String())
	c.Config.Logger.Debug("[REQUEST-BODY]:%v\n", req.rawBody)
	c.Config.Logger.Debug("[REQUEST-HEADERS]:%+v\n", req.headers)

	resp, err := c.reqSender.sendRequest(*req)
	if err != nil {
		return nil, fmt.Errorf("SubmitRequest: %w", err)
	}
	bodyAsString, _ := jsonhelper.ConvertToJSONString(resp.Data, true)
	c.Config.Logger.Debug("[RESPONSE-BODY]:%s\n", bodyAsString)
	return resp, nil
}

// buildRequest creates a new Request for the Edgecast API with query params,
// adding appropriate headers
func (eb ecRequestBuilder) buildRequest(
	params buildRequestParams,
) (*request, error) {
	eb.logger.Debug("Building Request: %+v", params)
	relativeURL, err := url.Parse(params.path)
	if err != nil {
		return nil,
			fmt.Errorf("ecRequestBuilder.buildRequest: url.Parse: %w", err)
	}

	if !params.method.IsValid() {
		return nil,
			fmt.Errorf(
				"ecRequestBuilder.buildRequest: invalid HTTP method: %d",
				params.method)
	}

	req := request{
		method:  params.method.String(),
		url:     eb.baseAPIURL.ResolveReference(relativeURL),
		headers: make(map[string]string),
	}

	err = req.setPathParams(params.pathParams)
	if err != nil {
		return nil,
			fmt.Errorf("ecRequestBuilder.buildRequest: %w", err)
	}
	req.setQueryParams(params.queryParams)

	req.headers["User-Agent"] = params.userAgent
	req.headers["Accept"] = defaultHeaderAccept

	if params.rawBody != nil {
		err := req.setBody(params.rawBody)
		if err != nil {
			return nil, fmt.Errorf(
				"ecRequestBuilder.buildRequest: %w", err)
		}
	}

	if eb.authProvider != nil {
		err := req.setAuthorization(eb.authProvider)
		if err != nil {
			return nil, fmt.Errorf(
				"ecRequestBuilder.buildRequest: %w", err)
		}
	}

	if len(params.headers) > 0 {
		for k, v := range params.headers {
			req.headers[k] = v
		}
	}

	return &req, nil
}

func (req *request) setPathParams(params map[string]string) error {
	// Apply path parameters
	// e.g.
	// path = "/customers/{customer_id}/policies/{policy_id}""
	// params = { "customer_id": "1", "policy_id": "99" }
	// result = "customers/1/policies/99"
	for k, v := range params {
		searchKey := fmt.Sprintf("{%s}", k)

		if !strings.Contains(req.url.Path, searchKey) {
			return fmt.Errorf(
				"Request.setPathParams: param not found in path: %s",
				k)
		}

		req.url.Path = strings.Replace(
			req.url.Path,
			searchKey,
			fmt.Sprintf("%v", v),
			-1)
	}
	return nil
}

func (req *request) setQueryParams(queryParams map[string]string) {
	// Adding Query Params
	query := req.url.Query()
	for k, v := range queryParams {
		query.Add(k, v)
	}
	// Encode the parameters and set the URL
	req.url.RawQuery = query.Encode()
}

func (req *request) setBody(rawBody interface{}) error {
	if req.headers == nil {
		req.headers = make(map[string]string)
	}
	switch b := rawBody.(type) {
	case string:
		req.rawBody = []byte(b)
		req.headers["Content-Type"] = "text/plain; charset=utf-8"
		req.headers["Accept"] = "application/json, text/html"
	default:
		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(rawBody)
		if err != nil {
			return err
		}
		req.rawBody = buf
		req.headers["Accept"] = defaultHeaderAccept
		req.headers["Content-Type"] = defaultHeaderContentType
	}
	return nil
}

func (req *request) setAuthorization(
	auth ecauth.AuthorizationProvider,
) error {
	authHeader, err := auth.GetAuthorizationHeader()
	if err != nil {
		return fmt.Errorf(
			"request.setAuthorization: failed to get authorization: %v",
			err)
	}
	if req.headers == nil {
		req.headers = make(map[string]string)
	}
	req.headers["Authorization"] = authHeader
	return nil
}

func (bp jsonBodyParser) parseBody(
	body []byte,
	parsedResponse interface{},
) error {
	var temp interface{}
	if err := json.Unmarshal(body, &temp); err != nil {
		return fmt.Errorf(
			"json.Unmarshal: malformed JSON response: %v",
			err)
	}

	if collectionhelper.IsInterfaceArray(temp) {
		if err := json.Unmarshal([]byte(body), &parsedResponse); err != nil {
			return fmt.Errorf(
				"json.Unmarshal: malformed JSON Array:%v",
				err)
		}
	} else {
		bodyAsString := string(body)
		if jsonhelper.IsJSONString(bodyAsString) {
			err := json.Unmarshal([]byte(bodyAsString), &parsedResponse)
			if err != nil {
				return fmt.Errorf(
					"json.Unmarshal: Decode error: %v",
					err)
			}
		} else {
			// if response is not json string
			switch v := parsedResponse.(type) {
			case literalResponse:
				rs, ok := parsedResponse.(literalResponse)
				if ok {
					rs.value = bodyAsString
					parsedResponse = rs
				}
			case float64:
				fmt.Println("float64:", v)
			default:
				fmt.Printf("%T:%+v\n", v, v)
			}
		}
	}
	return nil
}

// sendRequest sends a Request and returns the Response.
// If Request.ParsedResponse is non-nil, then the response body will be
// unmarshaled to it.
// Response.Data will always have the unmarshaled response body as a string.
func (es ecRequestSender) sendRequest(req request) (*Response, error) {
	httpResp, err := es.clientAdapter.Do(
		req.method,
		req.url,
		req.headers,
		req.rawBody)
	if err != nil {
		return nil, fmt.Errorf("sendRequest: %w", err)
	}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	bodyAsString := string(body)

	logMsg := jsonhelper.CreateJSONLogMessage("Reading Response", bodyAsString)
	es.logger.Debug(logMsg)

	if err != nil {
		return nil, fmt.Errorf(
			"sendRequest: ioutil.ReadAll: %v",
			err)
	}

	if httpResp.StatusCode >= 400 && httpResp.StatusCode <= 599 {
		return nil, fmt.Errorf(
			"sendRequest failed (HTTP StatusCode:%d): %s",
			httpResp.StatusCode,
			bodyAsString)
	}

	// If a schema was provided, use the parser.
	if req.parsedResponse != nil && httpResp.StatusCode != 204 {
		err = es.parser.parseBody(body, &req.parsedResponse)
		if err != nil {
			return nil, fmt.Errorf("sendRequest: parseBody: %w", err)
		}
	}

	return &Response{
		Data:         bodyAsString,
		HTTPResponse: httpResp,
	}, nil
}
