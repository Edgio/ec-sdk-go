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

// CustomerOriginGroupADN
type CustomerOriginGroupADN struct {
	Id                   *int32       `json:"id,omitempty"`
	Name                 *string      `json:"name,omitempty"`
	HostHeader           *string      `json:"host_header,omitempty"`
	ValidationPath       *string      `json:"validation_path,omitempty"`
	NetworkTypeId        *int32       `json:"network_type_id,omitempty"`
	Gateway              *Gateway     `json:"gateway,omitempty"`
	TlsSettings          *TlsSettings `json:"tls_settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomerOriginGroupADN CustomerOriginGroupADN

// NewCustomerOriginGroupADN instantiates a new CustomerOriginGroupADN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerOriginGroupADN() *CustomerOriginGroupADN {
	this := CustomerOriginGroupADN{}
	return &this
}

// NewCustomerOriginGroupADNWithDefaults instantiates a new CustomerOriginGroupADN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerOriginGroupADNWithDefaults() *CustomerOriginGroupADN {
	this := CustomerOriginGroupADN{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerOriginGroupADN) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginGroupADN) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerOriginGroupADN) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CustomerOriginGroupADN) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerOriginGroupADN) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginGroupADN) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerOriginGroupADN) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerOriginGroupADN) SetName(v string) {
	o.Name = &v
}

// GetHostHeader returns the HostHeader field value if set, zero value otherwise.
func (o *CustomerOriginGroupADN) GetHostHeader() string {
	if o == nil || o.HostHeader == nil {
		var ret string
		return ret
	}
	return *o.HostHeader
}

// GetHostHeaderOk returns a tuple with the HostHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginGroupADN) GetHostHeaderOk() (*string, bool) {
	if o == nil || o.HostHeader == nil {
		return nil, false
	}
	return o.HostHeader, true
}

// HasHostHeader returns a boolean if a field has been set.
func (o *CustomerOriginGroupADN) HasHostHeader() bool {
	if o != nil && o.HostHeader != nil {
		return true
	}

	return false
}

// SetHostHeader gets a reference to the given string and assigns it to the HostHeader field.
func (o *CustomerOriginGroupADN) SetHostHeader(v string) {
	o.HostHeader = &v
}

// GetValidationPath returns the ValidationPath field value if set, zero value otherwise.
func (o *CustomerOriginGroupADN) GetValidationPath() string {
	if o == nil || o.ValidationPath == nil {
		var ret string
		return ret
	}
	return *o.ValidationPath
}

// GetValidationPathOk returns a tuple with the ValidationPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginGroupADN) GetValidationPathOk() (*string, bool) {
	if o == nil || o.ValidationPath == nil {
		return nil, false
	}
	return o.ValidationPath, true
}

// HasValidationPath returns a boolean if a field has been set.
func (o *CustomerOriginGroupADN) HasValidationPath() bool {
	if o != nil && o.ValidationPath != nil {
		return true
	}

	return false
}

// SetValidationPath gets a reference to the given string and assigns it to the ValidationPath field.
func (o *CustomerOriginGroupADN) SetValidationPath(v string) {
	o.ValidationPath = &v
}

// GetNetworkTypeId returns the NetworkTypeId field value if set, zero value otherwise.
func (o *CustomerOriginGroupADN) GetNetworkTypeId() int32 {
	if o == nil || o.NetworkTypeId == nil {
		var ret int32
		return ret
	}
	return *o.NetworkTypeId
}

// GetNetworkTypeIdOk returns a tuple with the NetworkTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginGroupADN) GetNetworkTypeIdOk() (*int32, bool) {
	if o == nil || o.NetworkTypeId == nil {
		return nil, false
	}
	return o.NetworkTypeId, true
}

// HasNetworkTypeId returns a boolean if a field has been set.
func (o *CustomerOriginGroupADN) HasNetworkTypeId() bool {
	if o != nil && o.NetworkTypeId != nil {
		return true
	}

	return false
}

// SetNetworkTypeId gets a reference to the given int32 and assigns it to the NetworkTypeId field.
func (o *CustomerOriginGroupADN) SetNetworkTypeId(v int32) {
	o.NetworkTypeId = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *CustomerOriginGroupADN) GetGateway() Gateway {
	if o == nil || o.Gateway == nil {
		var ret Gateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginGroupADN) GetGatewayOk() (*Gateway, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *CustomerOriginGroupADN) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given Gateway and assigns it to the Gateway field.
func (o *CustomerOriginGroupADN) SetGateway(v Gateway) {
	o.Gateway = &v
}

// GetTlsSettings returns the TlsSettings field value if set, zero value otherwise.
func (o *CustomerOriginGroupADN) GetTlsSettings() TlsSettings {
	if o == nil || o.TlsSettings == nil {
		var ret TlsSettings
		return ret
	}
	return *o.TlsSettings
}

// GetTlsSettingsOk returns a tuple with the TlsSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginGroupADN) GetTlsSettingsOk() (*TlsSettings, bool) {
	if o == nil || o.TlsSettings == nil {
		return nil, false
	}
	return o.TlsSettings, true
}

// HasTlsSettings returns a boolean if a field has been set.
func (o *CustomerOriginGroupADN) HasTlsSettings() bool {
	if o != nil && o.TlsSettings != nil {
		return true
	}

	return false
}

// SetTlsSettings gets a reference to the given TlsSettings and assigns it to the TlsSettings field.
func (o *CustomerOriginGroupADN) SetTlsSettings(v TlsSettings) {
	o.TlsSettings = &v
}

func (o CustomerOriginGroupADN) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.HostHeader != nil {
		toSerialize["host_header"] = o.HostHeader
	}
	if o.ValidationPath != nil {
		toSerialize["validation_path"] = o.ValidationPath
	}
	if o.NetworkTypeId != nil {
		toSerialize["network_type_id"] = o.NetworkTypeId
	}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
	}
	if o.TlsSettings != nil {
		toSerialize["tls_settings"] = o.TlsSettings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CustomerOriginGroupADN) UnmarshalJSON(bytes []byte) (err error) {
	varCustomerOriginGroupADN := _CustomerOriginGroupADN{}

	if err = json.Unmarshal(bytes, &varCustomerOriginGroupADN); err == nil {
		*o = CustomerOriginGroupADN(varCustomerOriginGroupADN)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "host_header")
		delete(additionalProperties, "validation_path")
		delete(additionalProperties, "network_type_id")
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "tls_settings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomerOriginGroupADN struct {
	value *CustomerOriginGroupADN
	isSet bool
}

func (v NullableCustomerOriginGroupADN) Get() *CustomerOriginGroupADN {
	return v.value
}

func (v *NullableCustomerOriginGroupADN) Set(val *CustomerOriginGroupADN) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerOriginGroupADN) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerOriginGroupADN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerOriginGroupADN(val CustomerOriginGroupADN) NullableCustomerOriginGroupADN {
	return NullableCustomerOriginGroupADN{value: &val, isSet: true}
}

func (v NullableCustomerOriginGroupADN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerOriginGroupADN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}