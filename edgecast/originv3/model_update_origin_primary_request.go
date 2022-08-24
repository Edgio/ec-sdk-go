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

// UpdateOriginPrimaryRequest
type UpdateOriginPrimaryRequest struct {
	IsPrimary            bool `json:"is_primary"`
	AdditionalProperties map[string]interface{}
}

type _UpdateOriginPrimaryRequest UpdateOriginPrimaryRequest

// NewUpdateOriginPrimaryRequest instantiates a new UpdateOriginPrimaryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOriginPrimaryRequest(isPrimary bool) *UpdateOriginPrimaryRequest {
	this := UpdateOriginPrimaryRequest{}
	this.IsPrimary = isPrimary
	return &this
}

// NewUpdateOriginPrimaryRequestWithDefaults instantiates a new UpdateOriginPrimaryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOriginPrimaryRequestWithDefaults() *UpdateOriginPrimaryRequest {
	this := UpdateOriginPrimaryRequest{}
	return &this
}

// GetIsPrimary returns the IsPrimary field value
func (o *UpdateOriginPrimaryRequest) GetIsPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value
// and a boolean to check if the value has been set.
func (o *UpdateOriginPrimaryRequest) GetIsPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPrimary, true
}

// SetIsPrimary sets field value
func (o *UpdateOriginPrimaryRequest) SetIsPrimary(v bool) {
	o.IsPrimary = v
}

func (o UpdateOriginPrimaryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["is_primary"] = o.IsPrimary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateOriginPrimaryRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateOriginPrimaryRequest := _UpdateOriginPrimaryRequest{}

	if err = json.Unmarshal(bytes, &varUpdateOriginPrimaryRequest); err == nil {
		*o = UpdateOriginPrimaryRequest(varUpdateOriginPrimaryRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "is_primary")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateOriginPrimaryRequest struct {
	value *UpdateOriginPrimaryRequest
	isSet bool
}

func (v NullableUpdateOriginPrimaryRequest) Get() *UpdateOriginPrimaryRequest {
	return v.value
}

func (v *NullableUpdateOriginPrimaryRequest) Set(val *UpdateOriginPrimaryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOriginPrimaryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOriginPrimaryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOriginPrimaryRequest(val UpdateOriginPrimaryRequest) NullableUpdateOriginPrimaryRequest {
	return NullableUpdateOriginPrimaryRequest{value: &val, isSet: true}
}

func (v NullableUpdateOriginPrimaryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOriginPrimaryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
