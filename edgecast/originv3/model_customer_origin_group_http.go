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

// CustomerOriginGroupHTTP
type CustomerOriginGroupHTTP struct {
	Id                   *int32       `json:"id,omitempty"`
	Name                 *string      `json:"name,omitempty"`
	HostHeader           *string      `json:"host_header,omitempty"`
	ShieldPops           []string     `json:"shield_pops,omitempty"`
	NetworkTypeId        *int32       `json:"network_type_id,omitempty"`
	StrictPciCertified   *bool        `json:"strict_pci_certified,omitempty"`
	TlsSettings          *TlsSettings `json:"tls_settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomerOriginGroupHTTP CustomerOriginGroupHTTP

// NewCustomerOriginGroupHTTP instantiates a new CustomerOriginGroupHTTP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerOriginGroupHTTP() *CustomerOriginGroupHTTP {
	this := CustomerOriginGroupHTTP{}
	return &this
}

// NewCustomerOriginGroupHTTPWithDefaults instantiates a new CustomerOriginGroupHTTP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerOriginGroupHTTPWithDefaults() *CustomerOriginGroupHTTP {
	this := CustomerOriginGroupHTTP{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerOriginGroupHTTP) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginGroupHTTP) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerOriginGroupHTTP) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CustomerOriginGroupHTTP) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerOriginGroupHTTP) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginGroupHTTP) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerOriginGroupHTTP) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerOriginGroupHTTP) SetName(v string) {
	o.Name = &v
}

// GetHostHeader returns the HostHeader field value if set, zero value otherwise.
func (o *CustomerOriginGroupHTTP) GetHostHeader() string {
	if o == nil || o.HostHeader == nil {
		var ret string
		return ret
	}
	return *o.HostHeader
}

// GetHostHeaderOk returns a tuple with the HostHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginGroupHTTP) GetHostHeaderOk() (*string, bool) {
	if o == nil || o.HostHeader == nil {
		return nil, false
	}
	return o.HostHeader, true
}

// HasHostHeader returns a boolean if a field has been set.
func (o *CustomerOriginGroupHTTP) HasHostHeader() bool {
	if o != nil && o.HostHeader != nil {
		return true
	}

	return false
}

// SetHostHeader gets a reference to the given string and assigns it to the HostHeader field.
func (o *CustomerOriginGroupHTTP) SetHostHeader(v string) {
	o.HostHeader = &v
}

// GetShieldPops returns the ShieldPops field value if set, zero value otherwise.
func (o *CustomerOriginGroupHTTP) GetShieldPops() []string {
	if o == nil || o.ShieldPops == nil {
		var ret []string
		return ret
	}
	return o.ShieldPops
}

// GetShieldPopsOk returns a tuple with the ShieldPops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginGroupHTTP) GetShieldPopsOk() ([]string, bool) {
	if o == nil || o.ShieldPops == nil {
		return nil, false
	}
	return o.ShieldPops, true
}

// HasShieldPops returns a boolean if a field has been set.
func (o *CustomerOriginGroupHTTP) HasShieldPops() bool {
	if o != nil && o.ShieldPops != nil {
		return true
	}

	return false
}

// SetShieldPops gets a reference to the given []string and assigns it to the ShieldPops field.
func (o *CustomerOriginGroupHTTP) SetShieldPops(v []string) {
	o.ShieldPops = v
}

// GetNetworkTypeId returns the NetworkTypeId field value if set, zero value otherwise.
func (o *CustomerOriginGroupHTTP) GetNetworkTypeId() int32 {
	if o == nil || o.NetworkTypeId == nil {
		var ret int32
		return ret
	}
	return *o.NetworkTypeId
}

// GetNetworkTypeIdOk returns a tuple with the NetworkTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginGroupHTTP) GetNetworkTypeIdOk() (*int32, bool) {
	if o == nil || o.NetworkTypeId == nil {
		return nil, false
	}
	return o.NetworkTypeId, true
}

// HasNetworkTypeId returns a boolean if a field has been set.
func (o *CustomerOriginGroupHTTP) HasNetworkTypeId() bool {
	if o != nil && o.NetworkTypeId != nil {
		return true
	}

	return false
}

// SetNetworkTypeId gets a reference to the given int32 and assigns it to the NetworkTypeId field.
func (o *CustomerOriginGroupHTTP) SetNetworkTypeId(v int32) {
	o.NetworkTypeId = &v
}

// GetStrictPciCertified returns the StrictPciCertified field value if set, zero value otherwise.
func (o *CustomerOriginGroupHTTP) GetStrictPciCertified() bool {
	if o == nil || o.StrictPciCertified == nil {
		var ret bool
		return ret
	}
	return *o.StrictPciCertified
}

// GetStrictPciCertifiedOk returns a tuple with the StrictPciCertified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginGroupHTTP) GetStrictPciCertifiedOk() (*bool, bool) {
	if o == nil || o.StrictPciCertified == nil {
		return nil, false
	}
	return o.StrictPciCertified, true
}

// HasStrictPciCertified returns a boolean if a field has been set.
func (o *CustomerOriginGroupHTTP) HasStrictPciCertified() bool {
	if o != nil && o.StrictPciCertified != nil {
		return true
	}

	return false
}

// SetStrictPciCertified gets a reference to the given bool and assigns it to the StrictPciCertified field.
func (o *CustomerOriginGroupHTTP) SetStrictPciCertified(v bool) {
	o.StrictPciCertified = &v
}

// GetTlsSettings returns the TlsSettings field value if set, zero value otherwise.
func (o *CustomerOriginGroupHTTP) GetTlsSettings() TlsSettings {
	if o == nil || o.TlsSettings == nil {
		var ret TlsSettings
		return ret
	}
	return *o.TlsSettings
}

// GetTlsSettingsOk returns a tuple with the TlsSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginGroupHTTP) GetTlsSettingsOk() (*TlsSettings, bool) {
	if o == nil || o.TlsSettings == nil {
		return nil, false
	}
	return o.TlsSettings, true
}

// HasTlsSettings returns a boolean if a field has been set.
func (o *CustomerOriginGroupHTTP) HasTlsSettings() bool {
	if o != nil && o.TlsSettings != nil {
		return true
	}

	return false
}

// SetTlsSettings gets a reference to the given TlsSettings and assigns it to the TlsSettings field.
func (o *CustomerOriginGroupHTTP) SetTlsSettings(v TlsSettings) {
	o.TlsSettings = &v
}

func (o CustomerOriginGroupHTTP) MarshalJSON() ([]byte, error) {
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
	if o.ShieldPops != nil {
		toSerialize["shield_pops"] = o.ShieldPops
	}
	if o.NetworkTypeId != nil {
		toSerialize["network_type_id"] = o.NetworkTypeId
	}
	if o.StrictPciCertified != nil {
		toSerialize["strict_pci_certified"] = o.StrictPciCertified
	}
	if o.TlsSettings != nil {
		toSerialize["tls_settings"] = o.TlsSettings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CustomerOriginGroupHTTP) UnmarshalJSON(bytes []byte) (err error) {
	varCustomerOriginGroupHTTP := _CustomerOriginGroupHTTP{}

	if err = json.Unmarshal(bytes, &varCustomerOriginGroupHTTP); err == nil {
		*o = CustomerOriginGroupHTTP(varCustomerOriginGroupHTTP)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "host_header")
		delete(additionalProperties, "shield_pops")
		delete(additionalProperties, "network_type_id")
		delete(additionalProperties, "strict_pci_certified")
		delete(additionalProperties, "tls_settings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomerOriginGroupHTTP struct {
	value *CustomerOriginGroupHTTP
	isSet bool
}

func (v NullableCustomerOriginGroupHTTP) Get() *CustomerOriginGroupHTTP {
	return v.value
}

func (v *NullableCustomerOriginGroupHTTP) Set(val *CustomerOriginGroupHTTP) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerOriginGroupHTTP) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerOriginGroupHTTP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerOriginGroupHTTP(val CustomerOriginGroupHTTP) NullableCustomerOriginGroupHTTP {
	return NullableCustomerOriginGroupHTTP{value: &val, isSet: true}
}

func (v NullableCustomerOriginGroupHTTP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerOriginGroupHTTP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
