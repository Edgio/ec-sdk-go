// Code generated by the Code Generation Kit (CGK) using OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
//
// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

/*
Customer Origins API v3

List of API of config Customer Origin.

API version: 0.5.0
Contact: portals-streaming@edgecast.com
*/

package originv3

import (
	"fmt"
	"path"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
	"github.com/go-openapi/errors"
)

// HttpLargeOnlyClient is the concrete client implementation for HttpLargeOnly
type HttpLargeOnlyClient struct {
	apiClient  ecclient.APIClient
	baseAPIURL string
}

// NewHttpLargeOnlyClient creates a new instance of HttpLargeOnlyClient
func NewHttpLargeOnlyClient(
	c ecclient.APIClient,
	baseAPIURL string,
) HttpLargeOnlyClient {
	return HttpLargeOnlyClient{c, baseAPIURL}
}

// HttpLargeOnlyClientService defines the operations for HttpLargeOnly
type HttpLargeOnlyClientService interface {
	AddHttpLargeCustomerOriginGroup(
		params AddHttpLargeCustomerOriginGroupParams,
	) (*CustomerOriginGroupHTTP, error)

	GetAllHttpLargeCustomerOriginGroups() ([]CustomerOriginGroupHTTP, error)

	GetHttpLargeCustomerOriginGroup(
		params GetHttpLargeCustomerOriginGroupParams,
	) (*CustomerOriginGroupHTTP, error)

	GetHttpLargeOriginShieldPops(
		params GetHttpLargeOriginShieldPopsParams,
	) ([]OriginShieldEdgeNode, error)

	UpdateHttpLargeCustomerOriginGroup(
		params UpdateHttpLargeCustomerOriginGroupParams,
	) (*CustomerOriginGroupHTTP, error)
}

// AddHttpLargeCustomerOriginGroupParams contains the parameters for AddHttpLargeCustomerOriginGroup
type AddHttpLargeCustomerOriginGroupParams struct {
	CustomerOriginGroupHTTPRequest CustomerOriginGroupHTTPRequest
}

// NewAddHttpLargeCustomerOriginGroupParams creates a new instance of AddHttpLargeCustomerOriginGroupParams
func NewAddHttpLargeCustomerOriginGroupParams() AddHttpLargeCustomerOriginGroupParams {
	return AddHttpLargeCustomerOriginGroupParams{}
}

// AddHttpLargeCustomerOriginGroup - Create new Http Large customer origin group
//
//	Create new Http Large Customer Origin Group
func (c HttpLargeOnlyClient) AddHttpLargeCustomerOriginGroup(
	params AddHttpLargeCustomerOriginGroupParams,
) (*CustomerOriginGroupHTTP, error) {
	req, err := buildAddHttpLargeCustomerOriginGroupRequest(params, c.baseAPIURL)
	if err != nil {
		return nil, err
	}

	parsedResponse := CustomerOriginGroupHTTP{}
	req.ParsedResponse = &parsedResponse

	_, err = c.apiClient.SubmitRequest(*req)

	if err != nil {
		return nil, fmt.Errorf("AddHttpLargeCustomerOriginGroup: %w", err)
	}

	return &parsedResponse, nil
}

func buildAddHttpLargeCustomerOriginGroupRequest(
	p AddHttpLargeCustomerOriginGroupParams,
	baseAPIURL string,
) (*ecclient.SubmitRequestParams, error) {
	req := ecclient.NewSubmitRequestParams()
	req.Path = path.Join(baseAPIURL, "/http-large/groups")
	errs := make([]error, 0)

	method, err := ecclient.ToHTTPMethod("Post")
	if err != nil {
		errs = append(errs, fmt.Errorf("AddHttpLargeCustomerOriginGroup: %w", err))
	}

	req.Method = method
	req.RawBody = p.CustomerOriginGroupHTTPRequest

	if len(errs) > 0 {
		return nil, errors.CompositeValidationError(errs...)
	}

	return &req, nil
}

// GetAllHttpLargeCustomerOriginGroups - Get Http Large customer origins groups
//
//	Get list of Http Large Customer Origin Groups
func (c HttpLargeOnlyClient) GetAllHttpLargeCustomerOriginGroups() ([]CustomerOriginGroupHTTP, error) {
	req, err := buildGetAllHttpLargeCustomerOriginGroupsRequest(c.baseAPIURL)
	if err != nil {
		return nil, err
	}

	parsedResponse := make([]CustomerOriginGroupHTTP, 0)
	req.ParsedResponse = &parsedResponse

	_, err = c.apiClient.SubmitRequest(*req)

	if err != nil {
		return nil, fmt.Errorf("GetAllHttpLargeCustomerOriginGroups: %w", err)
	}

	return parsedResponse, nil
}

func buildGetAllHttpLargeCustomerOriginGroupsRequest(
	baseAPIURL string,
) (*ecclient.SubmitRequestParams, error) {
	req := ecclient.NewSubmitRequestParams()
	req.Path = path.Join(baseAPIURL, "/http-large/groups")
	errs := make([]error, 0)

	method, err := ecclient.ToHTTPMethod("Get")
	if err != nil {
		errs = append(errs, fmt.Errorf("GetAllHttpLargeCustomerOriginGroups: %w", err))
	}

	req.Method = method

	if len(errs) > 0 {
		return nil, errors.CompositeValidationError(errs...)
	}

	return &req, nil
}

// GetHttpLargeCustomerOriginGroupParams contains the parameters for GetHttpLargeCustomerOriginGroup
type GetHttpLargeCustomerOriginGroupParams struct {
	// Customer Origin Group Id
	GroupId string
}

// NewGetHttpLargeCustomerOriginGroupParams creates a new instance of GetHttpLargeCustomerOriginGroupParams
func NewGetHttpLargeCustomerOriginGroupParams() GetHttpLargeCustomerOriginGroupParams {
	return GetHttpLargeCustomerOriginGroupParams{}
}

// GetHttpLargeCustomerOriginGroup - Get specific Http Large customer origin group by id
//
//	Get an individual Http Large Customer Origin Group
func (c HttpLargeOnlyClient) GetHttpLargeCustomerOriginGroup(
	params GetHttpLargeCustomerOriginGroupParams,
) (*CustomerOriginGroupHTTP, error) {
	req, err := buildGetHttpLargeCustomerOriginGroupRequest(params, c.baseAPIURL)
	if err != nil {
		return nil, err
	}

	parsedResponse := CustomerOriginGroupHTTP{}
	req.ParsedResponse = &parsedResponse

	_, err = c.apiClient.SubmitRequest(*req)

	if err != nil {
		return nil, fmt.Errorf("GetHttpLargeCustomerOriginGroup: %w", err)
	}

	return &parsedResponse, nil
}

func buildGetHttpLargeCustomerOriginGroupRequest(
	p GetHttpLargeCustomerOriginGroupParams,
	baseAPIURL string,
) (*ecclient.SubmitRequestParams, error) {
	req := ecclient.NewSubmitRequestParams()
	req.Path = path.Join(baseAPIURL, "/http-large/groups/{groupId}")
	errs := make([]error, 0)

	method, err := ecclient.ToHTTPMethod("Get")
	if err != nil {
		errs = append(errs, fmt.Errorf("GetHttpLargeCustomerOriginGroup: %w", err))
	}

	req.Method = method

	req.PathParams["groupId"] = p.GroupId

	if len(errs) > 0 {
		return nil, errors.CompositeValidationError(errs...)
	}

	return &req, nil
}

// GetHttpLargeOriginShieldPopsParams contains the parameters for GetHttpLargeOriginShieldPops
type GetHttpLargeOriginShieldPopsParams struct {
	Findcode string
}

// NewGetHttpLargeOriginShieldPopsParams creates a new instance of GetHttpLargeOriginShieldPopsParams
func NewGetHttpLargeOriginShieldPopsParams() GetHttpLargeOriginShieldPopsParams {
	return GetHttpLargeOriginShieldPopsParams{}
}

// GetHttpLargeOriginShieldPops - Get list of origin shield pops
//
//	Get list of Origin Shield Pops. This API should work only for http-large Origin
func (c HttpLargeOnlyClient) GetHttpLargeOriginShieldPops(
	params GetHttpLargeOriginShieldPopsParams,
) ([]OriginShieldEdgeNode, error) {
	req, err := buildGetHttpLargeOriginShieldPopsRequest(params, c.baseAPIURL)
	if err != nil {
		return nil, err
	}

	parsedResponse := make([]OriginShieldEdgeNode, 0)
	req.ParsedResponse = &parsedResponse

	_, err = c.apiClient.SubmitRequest(*req)

	if err != nil {
		return nil, fmt.Errorf("GetHttpLargeOriginShieldPops: %w", err)
	}

	return parsedResponse, nil
}

func buildGetHttpLargeOriginShieldPopsRequest(
	p GetHttpLargeOriginShieldPopsParams,
	baseAPIURL string,
) (*ecclient.SubmitRequestParams, error) {
	req := ecclient.NewSubmitRequestParams()
	req.Path = path.Join(baseAPIURL, "/http-large/shield-pops")
	errs := make([]error, 0)

	method, err := ecclient.ToHTTPMethod("Get")
	if err != nil {
		errs = append(errs, fmt.Errorf("GetHttpLargeOriginShieldPops: %w", err))
	}

	req.Method = method

	req.QueryParams["findcode"] = p.Findcode

	if len(errs) > 0 {
		return nil, errors.CompositeValidationError(errs...)
	}

	return &req, nil
}

// UpdateHttpLargeCustomerOriginGroupParams contains the parameters for UpdateHttpLargeCustomerOriginGroup
type UpdateHttpLargeCustomerOriginGroupParams struct {
	// Customer Origin Group Id
	GroupId string

	CustomerOriginGroupHTTPRequest CustomerOriginGroupHTTPRequest
}

// NewUpdateHttpLargeCustomerOriginGroupParams creates a new instance of UpdateHttpLargeCustomerOriginGroupParams
func NewUpdateHttpLargeCustomerOriginGroupParams() UpdateHttpLargeCustomerOriginGroupParams {
	return UpdateHttpLargeCustomerOriginGroupParams{}
}

// UpdateHttpLargeCustomerOriginGroup - Update Http Large customer origin group by id
//
//	Update an individual Http Large Customer Origin Group
func (c HttpLargeOnlyClient) UpdateHttpLargeCustomerOriginGroup(
	params UpdateHttpLargeCustomerOriginGroupParams,
) (*CustomerOriginGroupHTTP, error) {
	req, err := buildUpdateHttpLargeCustomerOriginGroupRequest(params, c.baseAPIURL)
	if err != nil {
		return nil, err
	}

	parsedResponse := CustomerOriginGroupHTTP{}
	req.ParsedResponse = &parsedResponse

	_, err = c.apiClient.SubmitRequest(*req)

	if err != nil {
		return nil, fmt.Errorf("UpdateHttpLargeCustomerOriginGroup: %w", err)
	}

	return &parsedResponse, nil
}

func buildUpdateHttpLargeCustomerOriginGroupRequest(
	p UpdateHttpLargeCustomerOriginGroupParams,
	baseAPIURL string,
) (*ecclient.SubmitRequestParams, error) {
	req := ecclient.NewSubmitRequestParams()
	req.Path = path.Join(baseAPIURL, "/http-large/groups/{groupId}")
	errs := make([]error, 0)

	method, err := ecclient.ToHTTPMethod("Put")
	if err != nil {
		errs = append(errs, fmt.Errorf("UpdateHttpLargeCustomerOriginGroup: %w", err))
	}

	req.Method = method

	req.PathParams["groupId"] = p.GroupId

	req.RawBody = p.CustomerOriginGroupHTTPRequest

	if len(errs) > 0 {
		return nil, errors.CompositeValidationError(errs...)
	}

	return &req, nil
}
