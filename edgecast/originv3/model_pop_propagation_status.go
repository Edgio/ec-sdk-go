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

// PopPropagationStatus struct for PopPropagationStatus
type PopPropagationStatus struct {
	Name                 *string  `json:"name,omitempty"`
	PercentagePropagated *float32 `json:"percentage_propagated,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PopPropagationStatus PopPropagationStatus

// NewPopPropagationStatus instantiates a new PopPropagationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPopPropagationStatus() *PopPropagationStatus {
	this := PopPropagationStatus{}
	return &this
}

// NewPopPropagationStatusWithDefaults instantiates a new PopPropagationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPopPropagationStatusWithDefaults() *PopPropagationStatus {
	this := PopPropagationStatus{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PopPropagationStatus) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PopPropagationStatus) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PopPropagationStatus) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PopPropagationStatus) SetName(v string) {
	o.Name = &v
}

// GetPercentagePropagated returns the PercentagePropagated field value if set, zero value otherwise.
func (o *PopPropagationStatus) GetPercentagePropagated() float32 {
	if o == nil || o.PercentagePropagated == nil {
		var ret float32
		return ret
	}
	return *o.PercentagePropagated
}

// GetPercentagePropagatedOk returns a tuple with the PercentagePropagated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PopPropagationStatus) GetPercentagePropagatedOk() (*float32, bool) {
	if o == nil || o.PercentagePropagated == nil {
		return nil, false
	}
	return o.PercentagePropagated, true
}

// HasPercentagePropagated returns a boolean if a field has been set.
func (o *PopPropagationStatus) HasPercentagePropagated() bool {
	if o != nil && o.PercentagePropagated != nil {
		return true
	}

	return false
}

// SetPercentagePropagated gets a reference to the given float32 and assigns it to the PercentagePropagated field.
func (o *PopPropagationStatus) SetPercentagePropagated(v float32) {
	o.PercentagePropagated = &v
}

func (o PopPropagationStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PercentagePropagated != nil {
		toSerialize["percentage_propagated"] = o.PercentagePropagated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PopPropagationStatus) UnmarshalJSON(bytes []byte) (err error) {
	varPopPropagationStatus := _PopPropagationStatus{}

	if err = json.Unmarshal(bytes, &varPopPropagationStatus); err == nil {
		*o = PopPropagationStatus(varPopPropagationStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "percentage_propagated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePopPropagationStatus struct {
	value *PopPropagationStatus
	isSet bool
}

func (v NullablePopPropagationStatus) Get() *PopPropagationStatus {
	return v.value
}

func (v *NullablePopPropagationStatus) Set(val *PopPropagationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePopPropagationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePopPropagationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePopPropagationStatus(val PopPropagationStatus) NullablePopPropagationStatus {
	return NullablePopPropagationStatus{value: &val, isSet: true}
}

func (v NullablePopPropagationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePopPropagationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
