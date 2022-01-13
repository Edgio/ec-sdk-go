package client

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/testhelper"
)

type TestOKResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func TestSendRequest(t *testing.T) {
	cases := []struct {
		name          string
		rawResponse   string
		request       Request
		expected      TestOKResponse
		expectedError bool
	}{
		{
			name:        "Happy Path",
			rawResponse: `{"id":"1","name":"test 1"}`,
			request: Request{
				method:         "GET",
				url:            testhelper.URLParse("https://edgecast.com/tests/1"),
				parsedResponse: &TestOKResponse{},
			},
			expected:      TestOKResponse{ID: "1", Name: "test 1"},
			expectedError: false,
		},
	}
	for _, c := range cases {
		sender := ecRequestSender{
			clientAdapter: testClientAdapter{json: c.rawResponse},
		}
		resp, err := sender.sendRequest(c.request)
		if c.expectedError && err == nil {
			t.Fatalf("Case '%s': expected an error, but got none", c.name)
			return
		}
		actual := *resp.Data.(*TestOKResponse)
		if !reflect.DeepEqual(c.expected, actual) {
			t.Fatalf("%s: Expected %+v but got %+v", c.name, c.expected, actual)
		}
	}
}

type testClientAdapter struct {
	doFn func(req Request) (*http.Response, error)
	json string
}

func (c testClientAdapter) do(req Request) (*http.Response, error) {
	// Check for custom function
	if c.doFn != nil {
		return c.doFn(req)
	}

	data := c.json

	// Default implementation
	resp := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(data)),
	}
	return resp, nil
}

func TestECClientDo(t *testing.T) {
	cases := []struct {
		name          string
		input         DoParams
		reqBuilder    requestBuilder
		reqSender     requestSender
		expected      interface{}
		expectedError bool
	}{
		{
			name:  "Happy Path",
			input: DoParams{},
			reqBuilder: testReqBuilderAlwaysSafe{
				method: "GET",
				url:    "https://edgecast.com/test/1",
			},
			reqSender: testReqSenderAlwaysSafe{
				returnData: sampleData{
					StringData: "some string",
					IntData:    1,
					Nested:     nestedData{ID: "abcd"},
				},
			},
			expected: sampleData{
				StringData: "some string",
				IntData:    1,
				Nested:     nestedData{ID: "abcd"},
			},
			expectedError: false,
		},
	}

	for _, c := range cases {
		client := ECClient{
			reqBuilder: c.reqBuilder,
			reqSender:  c.reqSender,
			config: ClientConfig{
				Logger: testhelper.GetTestLogger("ec_client_test"),
			},
		}

		resp, err := client.Do(c.input)
		if c.expectedError && err == nil {
			t.Fatalf("Case '%s': expected an error, but got none", c.name)
			return
		}
		actual := resp.Data.(sampleData)
		if !reflect.DeepEqual(c.expected, actual) {
			t.Fatalf("%s: Expected %+v but got %+v", c.name, c.expected, actual)
		}
	}
}

type testReqBuilderAlwaysSafe struct {
	method string
	url    string
}

func (rb testReqBuilderAlwaysSafe) buildRequest(
	params buildRequestParams,
) (*Request, error) {
	var u *url.URL
	if len(rb.url) > 0 {
		u, _ = url.Parse(rb.url)
	} else {
		u, _ = url.Parse("https://edgecast.com")
	}
	return &Request{
		method: rb.method,
		url:    u,
	}, nil
}

type testReqSenderAlwaysSafe struct {
	returnData interface{}
}

func (rs testReqSenderAlwaysSafe) sendRequest(
	req Request,
) (*Response, error) {
	var data interface{}
	data = "some data"
	if rs.returnData != nil {
		data = rs.returnData
	}
	return &Response{Data: data}, nil
}

func (rb testReqSenderAlwaysSafe) sendRequestWithStringResponse(
	req Request,
) (*string, error) {
	data := "some data"
	return &data, nil
}

type sampleData struct {
	StringData string
	IntData    int
	Nested     nestedData
}

type nestedData struct {
	ID string
}

// func TestBuildRequest(t *testing.T) {
// 	goodSampleData := sampleData{
// 		StringData: "some string",
// 		IntData:    1,
// 		Nested:     nestedData{ID: "abcd"},
// 	}

// 	cases := []struct {
// 		name          string
// 		baseAPIURL    string
// 		authProvider  auth.AuthorizationProvider
// 		input         buildRequestParams
// 		expected      *Request
// 		expectedError bool
// 	}{
// 		{
// 			name:       "Happy path - Request With Body",
// 			baseAPIURL: "https://edgecast.com",
// 			authProvider: testAuthProvider{
// 				Auth:  "valid-token",
// 				Error: false,
// 			},
// 			input: buildRequestParams{
// 				method:  "POST",
// 				path:    "/customers/{customer_id}/policies/{policy_id}",
// 				rawBody: goodSampleData,
// 				queryParams: map[string]string{
// 					"q1": "val1",
// 					"q2": "val2",
// 				},
// 				pathParams: map[string]string{
// 					"customer_id": "HEX",
// 					"policy_id":   "100",
// 				},
// 				userAgent: "test-app",
// 			},
// 			expected: &Request{
// 				method: "POST",
// 				url:    testhelper.URLParse("https://edgecast.com/customers/HEX/policies/100?q1=val1&q2=val2"),
// 				headers: map[string]string{
// 					"Authorization": "valid-token",
// 					"Accept":        "application/json",
// 					"Content-Type":  "application/json",
// 					"User-Agent":    "test-app",
// 				},
// 				rawBody: goodSampleData,
// 			},
// 		},
// 		{
// 			name: "Error path - malformed url path",
// 			input: buildRequestParams{
// 				path: "h ttp://edgecast.com",
// 			},
// 			expectedError: true,
// 		},
// 	}

// 	for _, c := range cases {
// 		baseAPIURL, _ := url.Parse(c.baseAPIURL)
// 		builder := ecRequestBuilder{
// 			baseAPIURL:   *baseAPIURL,
// 			authProvider: &c.authProvider,
// 			userAgent:    "test",
// 		}

// 		actual, err := builder.buildRequest(c.input)
// 		if c.expectedError && err == nil {
// 			t.Fatalf("Case '%s': expected an error, but got none", c.name)
// 			return
// 		}
// 		// Need to check each property because of rawBody
// 		if !reflect.DeepEqual(c.expected.method, actual.method) {
// 			t.Fatalf("%s: Expected %+v but got %+v", c.name, c.expected.method, actual.method)
// 		}
// 		if !reflect.DeepEqual(c.expected.url, actual.url) {
// 			t.Fatalf("%s: Expected %+v but got %+v", c.name, c.expected.url, actual.url)
// 		}
// 		if !reflect.DeepEqual(c.expected.headers, actual.headers) {
// 			t.Fatalf("%s: Expected %+v but got %+v", c.name, c.expected.headers, actual.headers)
// 		}
// 		if bodyBuffer, ok := actual.rawBody.(*bytes.Buffer); ok {
// 			var resultParsed sampleData
// 			bufBytes := bodyBuffer.Bytes()
// 			err := json.Unmarshal(bufBytes, &resultParsed)
// 			if err == nil {
// 				if !reflect.DeepEqual(c.expected.rawBody, resultParsed) {
// 					t.Fatalf("%s: Expected %+v but got %+v", c.name, c.expected.rawBody, resultParsed)
// 				}
// 			} else {
// 				t.Fatalf("%s: failed to parse rawBody result as json:%+v", c.name, actual.rawBody)
// 			}
// 		} else {
// 			t.Fatalf("%s: rawBody expected to be bytes.Buffer, Actual:%T", c.name, actual.rawBody)
// 		}
// 	}

// }

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
