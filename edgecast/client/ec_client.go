package client

/*
	This file contains the concrete client implementation for the EC SDK
*/

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/collectionhelper"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/jsonhelper"
)

const (
	defaultHeaderAccept      string = "application/json"
	defaultHeaderContentType string = "application/json"
)

// Do invokes an HTTP request with the given parameters
func (c ECClient) SubmitRequest(params SubmitRequestParams) (*Response, error) {
	req, err := c.reqBuilder.buildRequest(buildRequestParams{
		method:      params.Method,
		path:        params.Path,
		rawBody:     params.Body,
		queryParams: params.QueryParams,
		pathParams:  params.PathParams,
		userAgent:   c.config.UserAgent,
	})

	if err != nil {
		return nil, fmt.Errorf("SubmitRequest: %v", err)
	}

	// Provides an object to be filled in when unmarshaling the API response
	req.parsedResponse = params.ParsedResponse

	c.config.Logger.Debug(
		"[REQUEST-URI]:[%s] %s\n",
		req.method,
		req.url.String())
	c.config.Logger.Debug("[REQUEST-BODY]:%v\n", req.rawBody)
	c.config.Logger.Debug("[REQUEST-HEADERS]:%+v\n", req.headers)

	resp, err := c.reqSender.sendRequest(*req)
	if err != nil {
		return nil, fmt.Errorf("SubmitRequest: %v", err)
	}
	bodyAsString, _ := jsonhelper.ConvertToJSONString(resp.Data, true)
	c.config.Logger.Debug("[RESPONSE-BODY]:%s\n", bodyAsString)
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
			fmt.Errorf("ecRequestBuilder.buildRequest: url.Parse: %v", err)
	}

	if !params.method.IsValid() {
		return nil,
			errors.New("ecRequestBuilder.buildRequest: invalid HTTP method")
	}

	req := request{
		method:  params.method.String(),
		url:     eb.baseAPIURL.ResolveReference(relativeURL),
		headers: make(map[string]string),
	}

	err = req.setPathParams(params.pathParams)
	if err != nil {
		return nil,
			fmt.Errorf("ecRequestBuilder.buildRequest: %v", err)
	}
	req.setQueryParams(params.queryParams)

	req.headers["User-Agent"] = params.userAgent
	req.headers["Accept"] = defaultHeaderAccept

	if params.rawBody != nil {
		err := req.setBody(params.rawBody)
		if err != nil {
			return nil, fmt.Errorf(
				"ecRequestBuilder.buildRequest: %v", err)
		}
	}

	if eb.authProvider != nil {
		err := req.setAuthorization(*eb.authProvider)
		if err != nil {
			return nil, fmt.Errorf(
				"ecRequestBuilder.buildRequest: %v", err)
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
	auth auth.AuthorizationProvider,
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
				fmt.Println("unknown")
			}
		}
	}
	return nil
}

// sendRequest sends a Request and returns the Response.
// If Request.ParsedResponse is non-nil, then the response body will be
// unmarshaled to it.
// Response.Data will always have the unmarshaled response body. Note that if
// Request.ParsedResponse is nil, then Response.Data will be a map[string]string
// as a result of unmarshaling JSON.
func (es ecRequestSender) sendRequest(req request) (*Response, error) {
	httpResp, err := es.clientAdapter.Do(
		req.method,
		req.url,
		req.headers,
		req.rawBody)
	if err != nil {
		return nil, fmt.Errorf("sendRequest: %v", err)
	}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf(
			"sendRequest: ioutil.ReadAll: %v",
			err)
	}

	if httpResp.StatusCode >= 400 && httpResp.StatusCode <= 599 {
		bodyAsString := string(body)
		return nil, fmt.Errorf(
			"sendRequest failed (HTTP StatusCode:%d): %s",
			httpResp.StatusCode,
			bodyAsString)
	}

	// If a string response is expected, do not use the parser
	if _, ok := req.parsedResponse.(*string); ok {
		bodyAsString := string(body)
		req.parsedResponse = &bodyAsString
	} else {
		err = es.parser.parseBody(body, &req.parsedResponse)
		if err != nil {
			return nil, fmt.Errorf("sendRequest: parseBody: %v", err)
		}
	}

	return &Response{
		Data:         req.parsedResponse,
		HTTPResponse: httpResp,
	}, nil
}
