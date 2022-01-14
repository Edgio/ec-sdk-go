package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/testhelper"
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
		req := Request{}
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
		req := Request{}
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
		req := Request{
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
		req := Request{
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

func TestSendRequest(t *testing.T) {
	cases := []struct {
		name          string
		rawResponse   string
		request       Request
		expected      testOKResponse
		expectedError bool
	}{
		{
			name:        "Happy Path",
			rawResponse: `{"id":"1","name":"test 1"}`,
			request: Request{
				method:         "GET",
				url:            testhelper.URLParse("https://edgecast.com/tests/1"),
				parsedResponse: &testOKResponse{},
			},
			expected:      testOKResponse{ID: "1", Name: "test 1"},
			expectedError: false,
		},
	}
	for _, c := range cases {
		sender := ecRequestSender{
			clientAdapter: testClientAdapter{json: c.rawResponse},
		}
		resp, err := sender.sendRequest(c.request)
		if c.expectedError {
			if err == nil {
				t.Fatalf("Case '%s': expected an error, but got none", c.name)
			}
		} else {
			actual := *resp.Data.(*testOKResponse)
			if !reflect.DeepEqual(c.expected, actual) {
				t.Fatalf("%s: Expected %+v but got %+v", c.name, c.expected, actual)
			}
		}
	}
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
		if c.expectedError {
			if err == nil {
				t.Fatalf("Case '%s': expected an error, but got none", c.name)
			}
		} else {
			actual := resp.Data.(sampleData)
			if !reflect.DeepEqual(c.expected, actual) {
				t.Fatalf("%s: Expected %+v but got %+v", c.name, c.expected, actual)
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
		authProvider  auth.AuthorizationProvider
		input         buildRequestParams
		expected      *Request
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
				method:  "POST",
				path:    "/customers/{customer_id}/policies/{policy_id}",
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
			expected: &Request{
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
			name: "Error path - malformed url path",
			input: buildRequestParams{
				path: "h ttp://edgecast.com",
			},
			expectedError: true,
		},
	}

	for _, c := range cases {
		baseAPIURL, _ := url.Parse(c.baseAPIURL)
		builder := ecRequestBuilder{
			baseAPIURL:   *baseAPIURL,
			authProvider: &c.authProvider,
			userAgent:    "test",
		}

		actual, err := builder.buildRequest(c.input)
		if c.expectedError {
			if err == nil {
				t.Fatalf("Case '%s': expected an error, but got none", c.name)
			}
			return
		}

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
