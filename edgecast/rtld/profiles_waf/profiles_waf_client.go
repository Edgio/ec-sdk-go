// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package profiles_waf

// This file was generated by codegen-sdk-go.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// New creates a new profiles waf API client.
func New(c ecclient.APIClient, cc ecclient.ClientConfig) ClientService {
	return &Client{c, cc}
}

/*
Client for profiles waf API
*/
type Client struct {
	client ecclient.APIClient
	config ecclient.ClientConfig
}

// ClientService is the interface for Client methods
type ClientService interface {
	ProfilesWafAddCustomerSetting(params *ProfilesWafAddCustomerSettingParams) (*ProfilesWafAddCustomerSettingOK, error)

	ProfilesWafDeleteCustomerSettingsByID(params *ProfilesWafDeleteCustomerSettingsByIDParams) (*ProfilesWafDeleteCustomerSettingsByIDNoContent, error)

	ProfilesWafGetCustomerSettings(params *ProfilesWafGetCustomerSettingsParams) (*ProfilesWafGetCustomerSettingsOK, error)

	ProfilesWafGetCustomerSettingsByID(params *ProfilesWafGetCustomerSettingsByIDParams) (*ProfilesWafGetCustomerSettingsByIDOK, error)

	ProfilesWafUpdateCustomerSetting(params *ProfilesWafUpdateCustomerSettingParams) (*ProfilesWafUpdateCustomerSettingOK, error)
}

/*
  ProfilesWafAddCustomerSetting profiles waf add customer setting API
*/
func (a *Client) ProfilesWafAddCustomerSetting(params *ProfilesWafAddCustomerSettingParams) (*ProfilesWafAddCustomerSettingOK, error) {
	// Validate the params before sending
	if params == nil {
		params = NewProfilesWafAddCustomerSettingParams()
	}

	//Set parameters
	results, err := WriteToRequestProfilesWafAddCustomerSettingParams(params)
	if err != nil {
		return nil, err
	}

	method, err := ecclient.ToHTTPMethod("POST")
	if err != nil {
		return nil, fmt.Errorf("ProfilesWafAddCustomerSetting: %v", err)
	}

	parsedResponse := &ProfilesWafAddCustomerSettingOK{}

	_, respErr := a.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:         method,
		Path:           a.config.BaseAPIURL.String() + "/v1.0/waf/profiles",
		RawBody:        results.Body,
		PathParams:     results.PathParams,
		QueryParams:    results.QueryParams,
		ParsedResponse: parsedResponse,
	})

	if respErr != nil {
		return nil, fmt.Errorf("ProfilesWafAddCustomerSetting: %v", respErr)
	}

	return parsedResponse, nil
}

/*
  ProfilesWafDeleteCustomerSettingsByID profiles waf delete customer settings by Id API
*/
func (a *Client) ProfilesWafDeleteCustomerSettingsByID(params *ProfilesWafDeleteCustomerSettingsByIDParams) (*ProfilesWafDeleteCustomerSettingsByIDNoContent, error) {
	// Validate the params before sending
	if params == nil {
		params = NewProfilesWafDeleteCustomerSettingsByIDParams()
	}

	//Set parameters
	results, err := WriteToRequestProfilesWafDeleteCustomerSettingsByIDParams(params)
	if err != nil {
		return nil, err
	}

	method, err := ecclient.ToHTTPMethod("DELETE")
	if err != nil {
		return nil, fmt.Errorf("ProfilesWafDeleteCustomerSettingsByID: %v", err)
	}

	parsedResponse := &ProfilesWafDeleteCustomerSettingsByIDNoContent{}

	_, respErr := a.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:         method,
		Path:           a.config.BaseAPIURL.String() + "/v1.0/waf/profiles/{id}",
		RawBody:        results.Body,
		PathParams:     results.PathParams,
		QueryParams:    results.QueryParams,
		ParsedResponse: parsedResponse,
	})

	if respErr != nil {
		return nil, fmt.Errorf("ProfilesWafDeleteCustomerSettingsByID: %v", respErr)
	}

	return parsedResponse, nil
}

/*
  ProfilesWafGetCustomerSettings profiles waf get customer settings API
*/
func (a *Client) ProfilesWafGetCustomerSettings(params *ProfilesWafGetCustomerSettingsParams) (*ProfilesWafGetCustomerSettingsOK, error) {
	// Validate the params before sending
	if params == nil {
		params = NewProfilesWafGetCustomerSettingsParams()
	}

	//Set parameters
	results, err := WriteToRequestProfilesWafGetCustomerSettingsParams(params)
	if err != nil {
		return nil, err
	}

	method, err := ecclient.ToHTTPMethod("GET")
	if err != nil {
		return nil, fmt.Errorf("ProfilesWafGetCustomerSettings: %v", err)
	}

	parsedResponse := &ProfilesWafGetCustomerSettingsOK{}

	_, respErr := a.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:         method,
		Path:           a.config.BaseAPIURL.String() + "/v1.0/waf/profiles",
		RawBody:        results.Body,
		PathParams:     results.PathParams,
		QueryParams:    results.QueryParams,
		ParsedResponse: parsedResponse,
	})

	if respErr != nil {
		return nil, fmt.Errorf("ProfilesWafGetCustomerSettings: %v", respErr)
	}

	return parsedResponse, nil
}

/*
  ProfilesWafGetCustomerSettingsByID profiles waf get customer settings by Id API
*/
func (a *Client) ProfilesWafGetCustomerSettingsByID(params *ProfilesWafGetCustomerSettingsByIDParams) (*ProfilesWafGetCustomerSettingsByIDOK, error) {
	// Validate the params before sending
	if params == nil {
		params = NewProfilesWafGetCustomerSettingsByIDParams()
	}

	//Set parameters
	results, err := WriteToRequestProfilesWafGetCustomerSettingsByIDParams(params)
	if err != nil {
		return nil, err
	}

	method, err := ecclient.ToHTTPMethod("GET")
	if err != nil {
		return nil, fmt.Errorf("ProfilesWafGetCustomerSettingsByID: %v", err)
	}

	parsedResponse := &ProfilesWafGetCustomerSettingsByIDOK{}

	_, respErr := a.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:         method,
		Path:           a.config.BaseAPIURL.String() + "/v1.0/waf/profiles/{id}",
		RawBody:        results.Body,
		PathParams:     results.PathParams,
		QueryParams:    results.QueryParams,
		ParsedResponse: parsedResponse,
	})

	if respErr != nil {
		return nil, fmt.Errorf("ProfilesWafGetCustomerSettingsByID: %v", respErr)
	}

	return parsedResponse, nil
}

/*
  ProfilesWafUpdateCustomerSetting profiles waf update customer setting API
*/
func (a *Client) ProfilesWafUpdateCustomerSetting(params *ProfilesWafUpdateCustomerSettingParams) (*ProfilesWafUpdateCustomerSettingOK, error) {
	// Validate the params before sending
	if params == nil {
		params = NewProfilesWafUpdateCustomerSettingParams()
	}

	//Set parameters
	results, err := WriteToRequestProfilesWafUpdateCustomerSettingParams(params)
	if err != nil {
		return nil, err
	}

	method, err := ecclient.ToHTTPMethod("PUT")
	if err != nil {
		return nil, fmt.Errorf("ProfilesWafUpdateCustomerSetting: %v", err)
	}

	parsedResponse := &ProfilesWafUpdateCustomerSettingOK{}

	_, respErr := a.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:         method,
		Path:           a.config.BaseAPIURL.String() + "/v1.0/waf/profiles/{id}",
		RawBody:        results.Body,
		PathParams:     results.PathParams,
		QueryParams:    results.QueryParams,
		ParsedResponse: parsedResponse,
	})

	if respErr != nil {
		return nil, fmt.Errorf("ProfilesWafUpdateCustomerSetting: %v", respErr)
	}

	return parsedResponse, nil
}

type RequestParameters struct {
	QueryParams map[string]string
	PathParams  map[string]string
	Body        interface{}
}

func NewRequestParameters() *RequestParameters {
	return &RequestParameters{
		QueryParams: make(map[string]string),
		PathParams:  make(map[string]string),
	}
}
