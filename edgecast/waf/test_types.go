package waf

import (
	"errors"

	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
	"github.com/EdgeCast/ec-sdk-go/edgecast/logging"
)

type mockClient struct {
	doFn func(params client.DoParams) (interface{}, error)
}

func (c mockClient) Do(params client.DoParams) (interface{}, error) {
	// Check for custom function
	if c.doFn != nil {
		return c.doFn(params)
	}

	// Default implementation - test case must provide a function
	return "error", errors.New("implementation for Do not provided")
}

func buildTestWAFService(apiClient client.APIClient) WAFService {
	return WAFService{
		client: apiClient,
		Logger: logging.NewStandardLogger(),
	}
}
