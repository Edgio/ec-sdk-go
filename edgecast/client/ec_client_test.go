package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/url"
	"reflect"
	"testing"

	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ectest"
)

func TestECClientDo(t *testing.T) {
	// The actions all follow the same basic pattern, so we'll use one test func
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
				Logger: ectest.GetTestLogger("ec_client_test"),
			},
		}

		actual, err := client.Do(c.input)
		if c.expectedError && err == nil {
			t.Fatalf("Case '%s': expected an error, but got none", c.name)
			return
		}
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
) (*request, error) {
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

type testReqSenderAlwaysSafe struct {
	returnData interface{}
}

func (rs testReqSenderAlwaysSafe) sendRequest(
	req request,
) (*response, error) {
	var data interface{}
	data = "some data"
	if rs.returnData != nil {
		data = rs.returnData
	}
	return &response{data: data}, nil
}

func (rb testReqSenderAlwaysSafe) sendRequestWithStringResponse(
	req request,
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
			expected: &request{
				method: "POST",
				url:    urlParse("https://edgecast.com/customers/HEX/policies/100?q1=val1&q2=val2"),
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
		if c.expectedError && err == nil {
			t.Fatalf("Case '%s': expected an error, but got none", c.name)
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

func urlParse(urlRaw string) *url.URL {
	result, _ := url.Parse(urlRaw)
	return result
}
