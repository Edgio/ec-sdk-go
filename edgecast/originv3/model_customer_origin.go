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

// CustomerOrigin struct for CustomerOrigin
type CustomerOrigin struct {
	Id                   *int32  `json:"id,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Host                 *string `json:"host,omitempty"`
	Port                 *int32  `json:"port,omitempty"`
	IsPrimary            *bool   `json:"is_primary,omitempty"`
	StorageTypeId        *int32  `json:"storage_type_id,omitempty"`
	ProtocolTypeId       *int32  `json:"protocol_type_id,omitempty"`
	GroupId              *int32  `json:"group_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomerOrigin CustomerOrigin

// NewCustomerOrigin instantiates a new CustomerOrigin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerOrigin() *CustomerOrigin {
	this := CustomerOrigin{}
	return &this
}

// NewCustomerOriginWithDefaults instantiates a new CustomerOrigin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerOriginWithDefaults() *CustomerOrigin {
	this := CustomerOrigin{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerOrigin) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOrigin) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerOrigin) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CustomerOrigin) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerOrigin) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOrigin) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerOrigin) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerOrigin) SetName(v string) {
	o.Name = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *CustomerOrigin) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOrigin) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *CustomerOrigin) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *CustomerOrigin) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *CustomerOrigin) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOrigin) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *CustomerOrigin) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *CustomerOrigin) SetPort(v int32) {
	o.Port = &v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *CustomerOrigin) GetIsPrimary() bool {
	if o == nil || o.IsPrimary == nil {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOrigin) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || o.IsPrimary == nil {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *CustomerOrigin) HasIsPrimary() bool {
	if o != nil && o.IsPrimary != nil {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *CustomerOrigin) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetStorageTypeId returns the StorageTypeId field value if set, zero value otherwise.
func (o *CustomerOrigin) GetStorageTypeId() int32 {
	if o == nil || o.StorageTypeId == nil {
		var ret int32
		return ret
	}
	return *o.StorageTypeId
}

// GetStorageTypeIdOk returns a tuple with the StorageTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOrigin) GetStorageTypeIdOk() (*int32, bool) {
	if o == nil || o.StorageTypeId == nil {
		return nil, false
	}
	return o.StorageTypeId, true
}

// HasStorageTypeId returns a boolean if a field has been set.
func (o *CustomerOrigin) HasStorageTypeId() bool {
	if o != nil && o.StorageTypeId != nil {
		return true
	}

	return false
}

// SetStorageTypeId gets a reference to the given int32 and assigns it to the StorageTypeId field.
func (o *CustomerOrigin) SetStorageTypeId(v int32) {
	o.StorageTypeId = &v
}

// GetProtocolTypeId returns the ProtocolTypeId field value if set, zero value otherwise.
func (o *CustomerOrigin) GetProtocolTypeId() int32 {
	if o == nil || o.ProtocolTypeId == nil {
		var ret int32
		return ret
	}
	return *o.ProtocolTypeId
}

// GetProtocolTypeIdOk returns a tuple with the ProtocolTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOrigin) GetProtocolTypeIdOk() (*int32, bool) {
	if o == nil || o.ProtocolTypeId == nil {
		return nil, false
	}
	return o.ProtocolTypeId, true
}

// HasProtocolTypeId returns a boolean if a field has been set.
func (o *CustomerOrigin) HasProtocolTypeId() bool {
	if o != nil && o.ProtocolTypeId != nil {
		return true
	}

	return false
}

// SetProtocolTypeId gets a reference to the given int32 and assigns it to the ProtocolTypeId field.
func (o *CustomerOrigin) SetProtocolTypeId(v int32) {
	o.ProtocolTypeId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *CustomerOrigin) GetGroupId() int32 {
	if o == nil || o.GroupId == nil {
		var ret int32
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOrigin) GetGroupIdOk() (*int32, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *CustomerOrigin) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given int32 and assigns it to the GroupId field.
func (o *CustomerOrigin) SetGroupId(v int32) {
	o.GroupId = &v
}

func (o CustomerOrigin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.IsPrimary != nil {
		toSerialize["is_primary"] = o.IsPrimary
	}
	if o.StorageTypeId != nil {
		toSerialize["storage_type_id"] = o.StorageTypeId
	}
	if o.ProtocolTypeId != nil {
		toSerialize["protocol_type_id"] = o.ProtocolTypeId
	}
	if o.GroupId != nil {
		toSerialize["group_id"] = o.GroupId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CustomerOrigin) UnmarshalJSON(bytes []byte) (err error) {
	varCustomerOrigin := _CustomerOrigin{}

	if err = json.Unmarshal(bytes, &varCustomerOrigin); err == nil {
		*o = CustomerOrigin(varCustomerOrigin)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "host")
		delete(additionalProperties, "port")
		delete(additionalProperties, "is_primary")
		delete(additionalProperties, "storage_type_id")
		delete(additionalProperties, "protocol_type_id")
		delete(additionalProperties, "group_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomerOrigin struct {
	value *CustomerOrigin
	isSet bool
}

func (v NullableCustomerOrigin) Get() *CustomerOrigin {
	return v.value
}

func (v *NullableCustomerOrigin) Set(val *CustomerOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerOrigin(val CustomerOrigin) NullableCustomerOrigin {
	return NullableCustomerOrigin{value: &val, isSet: true}
}

func (v NullableCustomerOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
