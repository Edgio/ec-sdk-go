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

// CustomerOriginStatus struct for CustomerOriginStatus
type CustomerOriginStatus struct {
	State                string                 `json:"state"`
	PercentPropagated    float32                `json:"percent_propagated"`
	Pops                 []PopPropagationStatus `json:"pops,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomerOriginStatus CustomerOriginStatus

// NewCustomerOriginStatus instantiates a new CustomerOriginStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerOriginStatus(state string, percentPropagated float32) *CustomerOriginStatus {
	this := CustomerOriginStatus{}
	this.State = state
	this.PercentPropagated = percentPropagated
	return &this
}

// NewCustomerOriginStatusWithDefaults instantiates a new CustomerOriginStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerOriginStatusWithDefaults() *CustomerOriginStatus {
	this := CustomerOriginStatus{}
	return &this
}

// GetState returns the State field value
func (o *CustomerOriginStatus) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CustomerOriginStatus) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CustomerOriginStatus) SetState(v string) {
	o.State = v
}

// GetPercentPropagated returns the PercentPropagated field value
func (o *CustomerOriginStatus) GetPercentPropagated() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PercentPropagated
}

// GetPercentPropagatedOk returns a tuple with the PercentPropagated field value
// and a boolean to check if the value has been set.
func (o *CustomerOriginStatus) GetPercentPropagatedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentPropagated, true
}

// SetPercentPropagated sets field value
func (o *CustomerOriginStatus) SetPercentPropagated(v float32) {
	o.PercentPropagated = v
}

// GetPops returns the Pops field value if set, zero value otherwise.
func (o *CustomerOriginStatus) GetPops() []PopPropagationStatus {
	if o == nil || o.Pops == nil {
		var ret []PopPropagationStatus
		return ret
	}
	return o.Pops
}

// GetPopsOk returns a tuple with the Pops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginStatus) GetPopsOk() ([]PopPropagationStatus, bool) {
	if o == nil || o.Pops == nil {
		return nil, false
	}
	return o.Pops, true
}

// HasPops returns a boolean if a field has been set.
func (o *CustomerOriginStatus) HasPops() bool {
	if o != nil && o.Pops != nil {
		return true
	}

	return false
}

// SetPops gets a reference to the given []PopPropagationStatus and assigns it to the Pops field.
func (o *CustomerOriginStatus) SetPops(v []PopPropagationStatus) {
	o.Pops = v
}

func (o CustomerOriginStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["percent_propagated"] = o.PercentPropagated
	}
	if o.Pops != nil {
		toSerialize["pops"] = o.Pops
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CustomerOriginStatus) UnmarshalJSON(bytes []byte) (err error) {
	varCustomerOriginStatus := _CustomerOriginStatus{}

	if err = json.Unmarshal(bytes, &varCustomerOriginStatus); err == nil {
		*o = CustomerOriginStatus(varCustomerOriginStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "state")
		delete(additionalProperties, "percent_propagated")
		delete(additionalProperties, "pops")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomerOriginStatus struct {
	value *CustomerOriginStatus
	isSet bool
}

func (v NullableCustomerOriginStatus) Get() *CustomerOriginStatus {
	return v.value
}

func (v *NullableCustomerOriginStatus) Set(val *CustomerOriginStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerOriginStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerOriginStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerOriginStatus(val CustomerOriginStatus) NullableCustomerOriginStatus {
	return NullableCustomerOriginStatus{value: &val, isSet: true}
}

func (v NullableCustomerOriginStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerOriginStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}