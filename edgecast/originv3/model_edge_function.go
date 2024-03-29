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
	"encoding/json"
)

// EdgeFunction
type EdgeFunction struct {
	EdgeFunctionName     string `json:"edge_function_name"`
	EdgeFunctionId       string `json:"edge_function_id"`
	EdgeFunctionHostname string `json:"edge_function_hostname"`
	AdditionalProperties map[string]interface{}
}

type _EdgeFunction EdgeFunction

// NewEdgeFunction instantiates a new EdgeFunction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeFunction(edgeFunctionName string, edgeFunctionId string, edgeFunctionHostname string) *EdgeFunction {
	this := EdgeFunction{}
	this.EdgeFunctionName = edgeFunctionName
	this.EdgeFunctionId = edgeFunctionId
	this.EdgeFunctionHostname = edgeFunctionHostname
	return &this
}

// NewEdgeFunctionWithDefaults instantiates a new EdgeFunction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeFunctionWithDefaults() *EdgeFunction {
	this := EdgeFunction{}
	return &this
}

// GetEdgeFunctionName returns the EdgeFunctionName field value
func (o *EdgeFunction) GetEdgeFunctionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EdgeFunctionName
}

// GetEdgeFunctionNameOk returns a tuple with the EdgeFunctionName field value
// and a boolean to check if the value has been set.
func (o *EdgeFunction) GetEdgeFunctionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdgeFunctionName, true
}

// SetEdgeFunctionName sets field value
func (o *EdgeFunction) SetEdgeFunctionName(v string) {
	o.EdgeFunctionName = v
}

// GetEdgeFunctionId returns the EdgeFunctionId field value
func (o *EdgeFunction) GetEdgeFunctionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EdgeFunctionId
}

// GetEdgeFunctionIdOk returns a tuple with the EdgeFunctionId field value
// and a boolean to check if the value has been set.
func (o *EdgeFunction) GetEdgeFunctionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdgeFunctionId, true
}

// SetEdgeFunctionId sets field value
func (o *EdgeFunction) SetEdgeFunctionId(v string) {
	o.EdgeFunctionId = v
}

// GetEdgeFunctionHostname returns the EdgeFunctionHostname field value
func (o *EdgeFunction) GetEdgeFunctionHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EdgeFunctionHostname
}

// GetEdgeFunctionHostnameOk returns a tuple with the EdgeFunctionHostname field value
// and a boolean to check if the value has been set.
func (o *EdgeFunction) GetEdgeFunctionHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdgeFunctionHostname, true
}

// SetEdgeFunctionHostname sets field value
func (o *EdgeFunction) SetEdgeFunctionHostname(v string) {
	o.EdgeFunctionHostname = v
}

func (o EdgeFunction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["edge_function_name"] = o.EdgeFunctionName
	}
	if true {
		toSerialize["edge_function_id"] = o.EdgeFunctionId
	}
	if true {
		toSerialize["edge_function_hostname"] = o.EdgeFunctionHostname
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EdgeFunction) UnmarshalJSON(bytes []byte) (err error) {
	varEdgeFunction := _EdgeFunction{}

	if err = json.Unmarshal(bytes, &varEdgeFunction); err == nil {
		*o = EdgeFunction(varEdgeFunction)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "edge_function_name")
		delete(additionalProperties, "edge_function_id")
		delete(additionalProperties, "edge_function_hostname")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEdgeFunction struct {
	value *EdgeFunction
	isSet bool
}

func (v NullableEdgeFunction) Get() *EdgeFunction {
	return v.value
}

func (v *NullableEdgeFunction) Set(val *EdgeFunction) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeFunction) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeFunction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeFunction(val EdgeFunction) NullableEdgeFunction {
	return NullableEdgeFunction{value: &val, isSet: true}
}

func (v NullableEdgeFunction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeFunction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
