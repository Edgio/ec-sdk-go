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

// OriginShieldPop OriginShieldPop object
type OriginShieldPop struct {
	Id                   *int32  `json:"id,omitempty"`
	Code                 *string `json:"code,omitempty"`
	City                 *string `json:"city,omitempty"`
	IsPciCertified       *bool   `json:"is_pci_certified,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OriginShieldPop OriginShieldPop

// NewOriginShieldPop instantiates a new OriginShieldPop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginShieldPop() *OriginShieldPop {
	this := OriginShieldPop{}
	return &this
}

// NewOriginShieldPopWithDefaults instantiates a new OriginShieldPop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginShieldPopWithDefaults() *OriginShieldPop {
	this := OriginShieldPop{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OriginShieldPop) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginShieldPop) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OriginShieldPop) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OriginShieldPop) SetId(v int32) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OriginShieldPop) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginShieldPop) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OriginShieldPop) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OriginShieldPop) SetCode(v string) {
	o.Code = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *OriginShieldPop) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginShieldPop) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *OriginShieldPop) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *OriginShieldPop) SetCity(v string) {
	o.City = &v
}

// GetIsPciCertified returns the IsPciCertified field value if set, zero value otherwise.
func (o *OriginShieldPop) GetIsPciCertified() bool {
	if o == nil || o.IsPciCertified == nil {
		var ret bool
		return ret
	}
	return *o.IsPciCertified
}

// GetIsPciCertifiedOk returns a tuple with the IsPciCertified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginShieldPop) GetIsPciCertifiedOk() (*bool, bool) {
	if o == nil || o.IsPciCertified == nil {
		return nil, false
	}
	return o.IsPciCertified, true
}

// HasIsPciCertified returns a boolean if a field has been set.
func (o *OriginShieldPop) HasIsPciCertified() bool {
	if o != nil && o.IsPciCertified != nil {
		return true
	}

	return false
}

// SetIsPciCertified gets a reference to the given bool and assigns it to the IsPciCertified field.
func (o *OriginShieldPop) SetIsPciCertified(v bool) {
	o.IsPciCertified = &v
}

func (o OriginShieldPop) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.IsPciCertified != nil {
		toSerialize["is_pci_certified"] = o.IsPciCertified
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OriginShieldPop) UnmarshalJSON(bytes []byte) (err error) {
	varOriginShieldPop := _OriginShieldPop{}

	if err = json.Unmarshal(bytes, &varOriginShieldPop); err == nil {
		*o = OriginShieldPop(varOriginShieldPop)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "code")
		delete(additionalProperties, "city")
		delete(additionalProperties, "is_pci_certified")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOriginShieldPop struct {
	value *OriginShieldPop
	isSet bool
}

func (v NullableOriginShieldPop) Get() *OriginShieldPop {
	return v.value
}

func (v *NullableOriginShieldPop) Set(val *OriginShieldPop) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginShieldPop) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginShieldPop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginShieldPop(val OriginShieldPop) NullableOriginShieldPop {
	return NullableOriginShieldPop{value: &val, isSet: true}
}

func (v NullableOriginShieldPop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginShieldPop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
