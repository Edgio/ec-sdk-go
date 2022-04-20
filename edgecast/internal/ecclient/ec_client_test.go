// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package ecclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecauth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/jsonhelper"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/testhelper"
)

var (
	testLog = testhelper.GetTestLogger("ec_client_test")
)

func TestSetBody(t *testing.T) {
	cases := []struct {
		name     string
		body     interface{}
		expected string
	}{
		{
			name:     "Happy Path - string body",
			body:     "some string payload",
			expected: "some string payload",
		},
		{
			name: "Happy Path - struct body",
			body: sampleData{
				StringData: "some string",
				IntData:    100,
				Nested: nestedData{
					ID: "abcd",
				},
			},
			expected: `{"StringData":"some string","IntData":100,"Nested":{"ID":"abcd"}}`,
		},
	}
	for _, c := range cases {
		req := request{}
		req.setBody(c.body)

		var actual string
		switch c.body.(type) {
		case string:
			actual = string(req.rawBody.([]byte))
		default:
			buf := req.rawBody.(*bytes.Buffer)
			actual = buf.String()
			// JSON encoding adds a newline character to the end, so trim it
			actual = strings.TrimSuffix(actual, "\n")
		}

		if strings.Compare(c.expected, actual) != 0 {
			t.Fatalf("%s: Expected %+v but got %+v", c.name, c.expected, actual)
		}
	}
}

func TestSetAuthorization(t *testing.T) {
	cases := []struct {
		name       string
		token      string
		throwError bool
		expected   string
	}{
		{
			name:       "Happy Path",
			token:      "AUTH_TOKEN",
			throwError: false,
			expected:   "AUTH_TOKEN",
		},
		{
			name:       "Error Path",
			throwError: true,
		},
	}
	for _, c := range cases {
		req := request{}
		authProvider := testAuthProvider{
			Auth:  c.token,
			Error: c.throwError,
		}
		err := req.setAuthorization(authProvider)
		if c.throwError {
			if err == nil {
				t.Fatalf("Case '%s': expected an error, but got none", c.name)
			}
		} else {
			// Because of the way testAuthProvider is designed, err should never
			// be non-nil here. However, we'll include the check anyway.
			if err != nil {
				t.Fatalf("Case '%s': unexpected error: %v", c.name, err)
			}
			actual := req.headers["Authorization"]
			if strings.Compare(c.expected, actual) != 0 {
				t.Fatalf("%s: Expected %+v but got %+v", c.name, c.expected, actual)
			}
		}
	}
}

func TestSetQueryParams(t *testing.T) {
	cases := []struct {
		name     string
		baseURL  string
		params   map[string]string
		expected string
	}{
		{
			name:    "Happy Path",
			baseURL: "https://edgecast.com/test",
			params: map[string]string{
				"key1": "value1",
				"key2": "value 2 with spaces",
				"key3": `specialchars!@#$%^&*()+[]\;',/<>:"{}|`,
			},
			expected: `https://edgecast.com/test?key1=value1&key2=value+2+with+spaces&key3=specialchars%21%40%23%24%25%5E%26%2A%28%29%2B%5B%5D%5C%3B%27%2C%2F%3C%3E%3A%22%7B%7D%7C`,
		},
		{
			name:     "Empty Params",
			baseURL:  "https://edgecast.com/test",
			params:   nil,
			expected: `https://edgecast.com/test`,
		},
	}
	for _, c := range cases {
		req := request{
			url: testhelper.URLParse(c.baseURL),
		}
		req.setQueryParams(c.params)
		actual := req.url.String()
		if !reflect.DeepEqual(c.expected, actual) {
			t.Fatalf("%s: Expected %+v but got %+v", c.name, c.expected, actual)
		}
	}
}

func TestSetPathParams(t *testing.T) {
	cases := []struct {
		name          string
		baseURL       string
		params        map[string]string
		expected      string
		expectedError bool
	}{
		{
			name:    "Happy Path",
			baseURL: "https://edgecast.com/{key1}/{key2}/{key3}",
			params: map[string]string{
				"key1": "value1",
				"key2": "value 2 with spaces",
				"key3": `specialchars!@#$%^&*()+[]\;',/<>:"{}|`,
			},
			expected:      `https://edgecast.com/value1/value%202%20with%20spaces/specialchars%21@%23$%25%5E&%2A%28%29+%5B%5D%5C;%27,/%3C%3E:%22%7B%7D%7C`,
			expectedError: false,
		},
		{
			name:    "Empty Params",
			baseURL: "https://edgecast.com/{key1}/{key2}/{key3}",
			params:  nil,
			// the curly braces will be encoded...
			expected:      "https://edgecast.com/%7Bkey1%7D/%7Bkey2%7D/%7Bkey3%7D",
			expectedError: false,
		},
		{
			name:    "Error Path: Param Key not in Path",
			baseURL: "https://edgecast.com/{key1}/{key3}",
			params: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
			expectedError: true,
		},
	}
	for _, c := range cases {
		req := request{
			url: testhelper.URLParse(c.baseURL),
		}
		err := req.setPathParams(c.params)
		if c.expectedError {
			if err == nil {
				t.Fatalf("Case '%s': expected an error, but got none", c.name)
			}
		} else {
			actual := req.url.String()
			if !reflect.DeepEqual(c.expected, actual) {
				t.Fatalf("%s: Expected %+v but got %+v", c.name, c.expected, actual)
			}
		}
	}

}

type testBodyParser struct {
	bodyValue     interface{}
	errorToReturn error
}

func (bp testBodyParser) parseBody(
	body []byte,
	parsedResponse interface{},
) error {
	if bp.errorToReturn != nil {
		return bp.errorToReturn
	}
	pv := reflect.ValueOf(parsedResponse)
	if pv.Kind() != reflect.Ptr {
		panic("parsedResponse not a pointer")
	}
	pv.Elem().Set(reflect.ValueOf(bp.bodyValue))
	return nil
}

func TestSendRequest(t *testing.T) {
	// Used for Happy Path below - has to be the same struct for the
	// equality check to work
	jsonData := `{ID:"1"}`
	literalStringData := "hello world!"
	jsonBody := testhelper.ToIOReadCloser(jsonData)
	stringBody := testhelper.ToIOReadCloser(literalStringData)

	cases := []struct {
		name          string
		request       request
		clientAdapter clientAdapter
		parser        bodyParser
		expected      *Response
		expectedError bool
	}{
		{
			name: "Happy Path - json response",
			clientAdapter: testClientAdapter{
				response: http.Response{
					StatusCode: 200,
					Body:       jsonBody,
				},
			},
			request: request{
				parsedResponse: testOKResponse{},
			},
			parser: testBodyParser{
				bodyValue: &testOKResponse{ID: "1"},
			},
			expected: &Response{
				Data: jsonData,
				HTTPResponse: &http.Response{
					StatusCode: 200,
					Body:       jsonBody,
				},
			},
		},
		{
			name: "Happy Path - non-JSON string response",
			clientAdapter: testClientAdapter{
				response: http.Response{
					StatusCode: 200,
					Body:       stringBody,
				},
			},
			expected: &Response{
				Data: literalStringData,
				HTTPResponse: &http.Response{
					StatusCode: 200,
					Body:       stringBody,
				},
			},
		},
		{
			name: "Error Path: client adapter error",
			clientAdapter: testClientAdapter{
				errorToReturn: errors.New("error from client adapter!"),
			},
			expectedError: true,
		},
		{
			name: "Error Path: HTTP Error Code",
			clientAdapter: testClientAdapter{
				response: http.Response{
					StatusCode: 400,
					Body:       testhelper.ToIOReadCloser("bad request!"),
				},
			},
			expectedError: true,
		},
		{
			name: "Error Path: unable to read body",
			clientAdapter: testClientAdapter{
				response: http.Response{
					StatusCode: 200,
					Body:       badReadCloser{},
				},
			},
			expectedError: true,
		},
		{
			name: "Error Path: body parser error",
			clientAdapter: testClientAdapter{
				response: http.Response{
					StatusCode: 200,
					Body:       testhelper.ToIOReadCloser(`{"id":"1"}`),
				},
			},
			request: request{
				parsedResponse: &sampleData{},
			},
			parser: testBodyParser{
				errorToReturn: errors.New("error from body parser!"),
			},
			expectedError: true,
		},
	}
	for _, c := range cases {
		sender := ecRequestSender{
			clientAdapter: c.clientAdapter,
			parser:        c.parser,
			logger:        testLog,
		}
		actual, err := sender.sendRequest(c.request)
		if c.expectedError {
			if err == nil {
				t.Fatalf("Case '%s': expected an error, but got none", c.name)
			}
		} else {
			if err != nil {
				t.Fatalf("Case '%s': unexpected error: %v", c.name, err)
			}
			if !reflect.DeepEqual(c.expected, actual) {
				t.Fatalf("%s: Expected %v but got %v", c.name, c.expected, actual)
			}
		}
	}
}

func TestJSONParseBody(t *testing.T) {
	cases := []struct {
		name           string
		body           []byte
		parsedResponse interface{}
		expected       interface{}
		expectedError  bool
	}{
		{
			name: "Happy Path - parsed response schema",
			body: testhelper.ToJSONBytes(
				testOKResponse{
					ID:   "1",
					Name: "resource 1",
				}),
			parsedResponse: &testOKResponse{},
			expected: &testOKResponse{
				ID:   "1",
				Name: "resource 1",
			},
		},
		{
			name:     "Happy Path - literal/non-JSON response",
			body:     testhelper.ToJSONBytes(100.99),
			expected: 100.99,
		},
		{
			// json.Unmarshal does not support simple strings
			name:          "Error Path - non-JSON string response",
			body:          []byte("hello world!"),
			expectedError: true,
		},
		{
			name:           "Error Path: string is not json",
			body:           []byte("hello world"),
			parsedResponse: testhelper.EmptyPointerFloat64(),
			expected:       100.99,
			expectedError:  true,
		},
	}
	for _, c := range cases {
		parser := newJSONBodyParser()
		err := parser.parseBody(c.body, &c.parsedResponse)

		if c.expectedError {
			if err == nil {
				t.Fatalf("Case '%s': expected an error, but got none", c.name)
			}
		} else {
			if err != nil {
				t.Fatalf("Case '%s': unexpected error: %v", c.name, err)
			}
			if !reflect.DeepEqual(c.expected, c.parsedResponse) {
				t.Fatalf("%s: Expected %+v but got %+v", c.name, c.expected, c.parsedResponse)
			}
		}
	}
}

func TestECClientSubmitRequest(t *testing.T) {
	cases := []struct {
		name           string
		input          SubmitRequestParams
		reqBuilder     requestBuilder
		reqSender      requestSender
		expected       interface{}
		expectedString string
		expectedError  bool
	}{
		{
			name: "Happy Path",
			input: SubmitRequestParams{
				ParsedResponse: &sampleData{},
			},
			reqBuilder: testReqBuilder{
				method: "GET",
				url:    "https://edgecast.com/test/1",
			},
			reqSender: testReqSender{
				returnData: sampleData{
					StringData: "some string",
					IntData:    1,
					Nested:     nestedData{ID: "abcd"},
				},
			},
			expected: &sampleData{
				StringData: "some string",
				IntData:    1,
				Nested:     nestedData{ID: "abcd"},
			},
			expectedString: `{"StringData":"some string","IntData":1,"Nested":{"ID":"abcd"}}`,
			expectedError:  false,
		},
		{
			name: "Error Path - builder error",
			reqBuilder: testReqBuilder{
				errorToReturn: errors.New("something happened!"),
			},
			expectedError: true,
		},
		{
			name: "Error Path - sender error",
			reqBuilder: testReqBuilder{
				method: "GET",
				url:    "https://edgecast.com/test/1",
			},
			reqSender: testReqSender{
				errorToReturn: errors.New("something happened!"),
			},
			expectedError: true,
		},
	}

	for _, c := range cases {
		client := ECClient{
			reqBuilder: c.reqBuilder,
			reqSender:  c.reqSender,
			Config: ClientConfig{
				Logger: testLog,
			},
		}

		resp, err := client.SubmitRequest(c.input)
		if c.expectedError {
			if err == nil {
				t.Fatalf("Case '%s': expected an error, but got none", c.name)
			}
		} else {
			actual := c.input.ParsedResponse.(*sampleData)
			if !reflect.DeepEqual(c.expected, actual) {
				t.Fatalf("%s: Expected %+v but got %+v", c.name, c.expected, actual)
			}
			if len(c.expectedString) > 0 &&
				!testhelper.JSONEqual(c.expectedString, resp.Data) {
				t.Fatalf("%s: Expected '%s' but got '%s'", c.name, c.expectedString, resp.Data)
			}

		}
	}
}

func TestBuildRequest(t *testing.T) {
	goodSampleData := sampleData{
		StringData: "some string",
		IntData:    1,
		Nested:     nestedData{ID: "abcd"},
	}

	cases := []struct {
		name          string
		baseAPIURL    string
		authProvider  ecauth.AuthorizationProvider
		input         buildRequestParams
		expected      *request
		expectedError bool
	}{
		{
			name:       "Happy path - Request With Body",
			baseAPIURL: "https://edgecast.com",
			authProvider: testAuthProvider{
				Auth:  "valid-token",
				Error: false,
			},
			input: buildRequestParams{
				method:  Post,
				path:    "customers/{customer_id}/policies/{policy_id}",
				rawBody: goodSampleData,
				queryParams: map[string]string{
					"q1": "val1",
					"q2": "val2",
				},
				pathParams: map[string]string{
					"customer_id": "HEX",
					"policy_id":   "100",
				},
				userAgent: "test-app",
			},
			expected: &request{
				method: "POST",
				url:    testhelper.URLParse("https://edgecast.com/customers/HEX/policies/100?q1=val1&q2=val2"),
				headers: map[string]string{
					"Authorization": "valid-token",
					"Accept":        "application/json",
					"Content-Type":  "application/json",
					"User-Agent":    "test-app",
				},
				rawBody: goodSampleData,
			},
		},
		{
			name: "malformed url path",
			input: buildRequestParams{
				path: "h ttp://edgecast.com",
			},
			expectedError: true,
		},
		{
			name: "Error path - invalid method",
			input: buildRequestParams{
				method: HTTPMethod(999),
			},
			expectedError: true,
		},
		{
			name: "Error path - set path params error",
			input: buildRequestParams{
				method:     Get,
				path:       "test/{key}",
				pathParams: map[string]string{"otherkey": "val"},
			},
			expectedError: true,
		},
		{
			name: "Error path - set body error",
			input: buildRequestParams{
				method:  Get,
				path:    "test/1",
				rawBody: func() {}, // funcs cannot be marshaled to json
			},
			expectedError: true,
		},
		{
			name:       "Error path - auth provider error",
			baseAPIURL: "https://edgecast.com",
			authProvider: testAuthProvider{
				Error: true,
			},
			input: buildRequestParams{
				method: Get,
				path:   "customers/1",
			},
			expectedError: true,
		},
	}

	for _, c := range cases {
		baseAPIURL, _ := url.Parse(c.baseAPIURL)
		builder := ecRequestBuilder{
			baseAPIURL:   *baseAPIURL,
			authProvider: c.authProvider,
			userAgent:    "test",
			logger:       testLog,
		}

		actual, err := builder.buildRequest(c.input)
		if c.expectedError {
			if err == nil {
				t.Fatalf("Case '%s': expected an error, but got none", c.name)
			}
		} else {
			// Need to check each property because of rawBody
			if !reflect.DeepEqual(c.expected.method, actual.method) {
				t.Fatalf("%s: Expected %+v but got %+v", c.name, c.expected.method, actual.method)
			}
			if !reflect.DeepEqual(c.expected.url, actual.url) {
				t.Fatalf("%s: Expected %+v but got %+v", c.name, c.expected.url, actual.url)
			}
			if !reflect.DeepEqual(c.expected.headers, actual.headers) {
				t.Fatalf("%s: Expected %+v but got %+v", c.name, c.expected.headers, actual.headers)
			}
			if bodyBuffer, ok := actual.rawBody.(*bytes.Buffer); ok {
				var resultParsed sampleData
				bufBytes := bodyBuffer.Bytes()
				err := json.Unmarshal(bufBytes, &resultParsed)
				if err == nil {
					if !reflect.DeepEqual(c.expected.rawBody, resultParsed) {
						t.Fatalf("%s: Expected %+v but got %+v", c.name, c.expected.rawBody, resultParsed)
					}
				} else {
					t.Fatalf("%s: failed to parse rawBody result as json:%+v", c.name, actual.rawBody)
				}
			} else {
				t.Fatalf("%s: rawBody expected to be bytes.Buffer, Actual:%T", c.name, actual.rawBody)
			}
		}
	}
}

type sampleData struct {
	StringData string
	IntData    int
	Nested     nestedData
}

type nestedData struct {
	ID string
}

type testAuthProvider struct {
	Auth  string
	Error bool
}

func (a testAuthProvider) GetAuthorizationHeader() (string, error) {
	if a.Error {
		return "", errors.New("auth error")
	}
	return a.Auth, nil
}

type testOKResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type testClientAdapter struct {
	response      http.Response
	errorToReturn error
}

func (c testClientAdapter) Do(
	method string,
	url *url.URL,
	headers map[string]string,
	rawBody interface{},
) (*http.Response, error) {
	if c.errorToReturn != nil {
		return nil, c.errorToReturn
	}
	return &c.response, nil
}

type testReqBuilder struct {
	method        string
	url           string
	errorToReturn error
}

func (rb testReqBuilder) buildRequest(
	params buildRequestParams,
) (*request, error) {
	if rb.errorToReturn != nil {
		return nil, rb.errorToReturn
	}
	var u *url.URL
	if len(rb.url) > 0 {
		u, _ = url.Parse(rb.url)
	} else {
		u, _ = url.Parse("https://edgecast.com")
	}
	return &request{
		method: rb.method,
		url:    u,
	}, nil
}

type testReqSender struct {
	returnData    interface{}
	errorToReturn error
}

func (rs testReqSender) sendRequest(
	req request,
) (*Response, error) {
	if rs.errorToReturn != nil {
		return nil, rs.errorToReturn
	}
	data := "some data"
	if rs.returnData != nil {
		// mock parser functionality - set pointer value with reflection
		targetPtrVal := reflect.ValueOf(req.parsedResponse)
		if targetPtrVal.Kind() != reflect.Ptr {
			panic("targetPtrVal not a pointer")
		}
		targetPtrVal.Elem().Set(reflect.ValueOf(rs.returnData))

		// return the json representation
		json, err := jsonhelper.ConvertToJSONString(rs.returnData, false)
		if err != nil {
			panic(fmt.Errorf("error converting test data to json: %v", err))
		}
		data = json
	}
	return &Response{Data: data}, nil
}

type badReadCloser struct {
	data          string
	errorToReturn error
}

func (rc badReadCloser) Read(p []byte) (n int, err error) {
	if rc.errorToReturn != nil {
		return 0, rc.errorToReturn
	}
	if len(rc.data) == 0 {
		return 0, errors.New("no data")
	}
	b := []byte(rc.data)
	copy(p, b)
	return len(b), nil
}
func (rc badReadCloser) Close() error {
	return nil
}

// MockAPIClient is a generic mock struct for unit testing feature services
type MockAPIClient struct {
	// MockAPIClient will always return this data on SubmitRequest
	ResponseData interface{}

	// MockAPIClient will always return this err on SubmitRequest
	Err error
}

func (m MockAPIClient) SubmitRequest(
	params SubmitRequestParams,
) (*Response, error) {
	// Use reflection to set the parsed response data
	pv := reflect.ValueOf(params.ParsedResponse)
	if pv.Kind() != reflect.Ptr {
		panic("ParsedResponse not a pointer")
	}
	pv.Elem().Set(reflect.ValueOf(m.ResponseData))

	dataAsString, err := jsonhelper.ConvertToJSONString(m.ResponseData, true)

	if err != nil {
		panic("could not convert mock response data as json")
	}

	return &Response{
		Data:         dataAsString,
		HTTPResponse: &http.Response{StatusCode: 200},
	}, m.Err
}
