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
	"time"
)

// Gateway struct for Gateway
type Gateway struct {
	SelectionLastRequested NullableTime   `json:"selection_last_requested,omitempty"`
	FollowRedirect         NullableBool   `json:"follow_redirect,omitempty"`
	PopLastUpdate          NullableTime   `json:"pop_last_update,omitempty"`
	SelectionError         NullableString `json:"selection_error,omitempty"`
	Selected               NullableBool   `json:"selected,omitempty"`
	GatewayReselection     NullableBool   `json:"gateway_reselection,omitempty"`
	Pops                   []AdnGateway   `json:"pops,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _Gateway Gateway

// NewGateway instantiates a new Gateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGateway() *Gateway {
	this := Gateway{}
	var followRedirect bool = false
	this.FollowRedirect = NewNullableBool(followRedirect)
	return &this
}

// NewGatewayWithDefaults instantiates a new Gateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayWithDefaults() *Gateway {
	this := Gateway{}
	var followRedirect bool = false
	this.FollowRedirect = NewNullableBool(followRedirect)
	return &this
}

// GetSelectionLastRequested returns the SelectionLastRequested field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Gateway) GetSelectionLastRequested() time.Time {
	if o == nil || o.SelectionLastRequested.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SelectionLastRequested.Get()
}

// GetSelectionLastRequestedOk returns a tuple with the SelectionLastRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gateway) GetSelectionLastRequestedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelectionLastRequested.Get(), o.SelectionLastRequested.IsSet()
}

// HasSelectionLastRequested returns a boolean if a field has been set.
func (o *Gateway) HasSelectionLastRequested() bool {
	if o != nil && o.SelectionLastRequested.IsSet() {
		return true
	}

	return false
}

// SetSelectionLastRequested gets a reference to the given NullableTime and assigns it to the SelectionLastRequested field.
func (o *Gateway) SetSelectionLastRequested(v time.Time) {
	o.SelectionLastRequested.Set(&v)
}

// SetSelectionLastRequestedNil sets the value for SelectionLastRequested to be an explicit nil
func (o *Gateway) SetSelectionLastRequestedNil() {
	o.SelectionLastRequested.Set(nil)
}

// UnsetSelectionLastRequested ensures that no value is present for SelectionLastRequested, not even an explicit nil
func (o *Gateway) UnsetSelectionLastRequested() {
	o.SelectionLastRequested.Unset()
}

// GetFollowRedirect returns the FollowRedirect field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Gateway) GetFollowRedirect() bool {
	if o == nil || o.FollowRedirect.Get() == nil {
		var ret bool
		return ret
	}
	return *o.FollowRedirect.Get()
}

// GetFollowRedirectOk returns a tuple with the FollowRedirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gateway) GetFollowRedirectOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowRedirect.Get(), o.FollowRedirect.IsSet()
}

// HasFollowRedirect returns a boolean if a field has been set.
func (o *Gateway) HasFollowRedirect() bool {
	if o != nil && o.FollowRedirect.IsSet() {
		return true
	}

	return false
}

// SetFollowRedirect gets a reference to the given NullableBool and assigns it to the FollowRedirect field.
func (o *Gateway) SetFollowRedirect(v bool) {
	o.FollowRedirect.Set(&v)
}

// SetFollowRedirectNil sets the value for FollowRedirect to be an explicit nil
func (o *Gateway) SetFollowRedirectNil() {
	o.FollowRedirect.Set(nil)
}

// UnsetFollowRedirect ensures that no value is present for FollowRedirect, not even an explicit nil
func (o *Gateway) UnsetFollowRedirect() {
	o.FollowRedirect.Unset()
}

// GetPopLastUpdate returns the PopLastUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Gateway) GetPopLastUpdate() time.Time {
	if o == nil || o.PopLastUpdate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.PopLastUpdate.Get()
}

// GetPopLastUpdateOk returns a tuple with the PopLastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gateway) GetPopLastUpdateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PopLastUpdate.Get(), o.PopLastUpdate.IsSet()
}

// HasPopLastUpdate returns a boolean if a field has been set.
func (o *Gateway) HasPopLastUpdate() bool {
	if o != nil && o.PopLastUpdate.IsSet() {
		return true
	}

	return false
}

// SetPopLastUpdate gets a reference to the given NullableTime and assigns it to the PopLastUpdate field.
func (o *Gateway) SetPopLastUpdate(v time.Time) {
	o.PopLastUpdate.Set(&v)
}

// SetPopLastUpdateNil sets the value for PopLastUpdate to be an explicit nil
func (o *Gateway) SetPopLastUpdateNil() {
	o.PopLastUpdate.Set(nil)
}

// UnsetPopLastUpdate ensures that no value is present for PopLastUpdate, not even an explicit nil
func (o *Gateway) UnsetPopLastUpdate() {
	o.PopLastUpdate.Unset()
}

// GetSelectionError returns the SelectionError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Gateway) GetSelectionError() string {
	if o == nil || o.SelectionError.Get() == nil {
		var ret string
		return ret
	}
	return *o.SelectionError.Get()
}

// GetSelectionErrorOk returns a tuple with the SelectionError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gateway) GetSelectionErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelectionError.Get(), o.SelectionError.IsSet()
}

// HasSelectionError returns a boolean if a field has been set.
func (o *Gateway) HasSelectionError() bool {
	if o != nil && o.SelectionError.IsSet() {
		return true
	}

	return false
}

// SetSelectionError gets a reference to the given NullableString and assigns it to the SelectionError field.
func (o *Gateway) SetSelectionError(v string) {
	o.SelectionError.Set(&v)
}

// SetSelectionErrorNil sets the value for SelectionError to be an explicit nil
func (o *Gateway) SetSelectionErrorNil() {
	o.SelectionError.Set(nil)
}

// UnsetSelectionError ensures that no value is present for SelectionError, not even an explicit nil
func (o *Gateway) UnsetSelectionError() {
	o.SelectionError.Unset()
}

// GetSelected returns the Selected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Gateway) GetSelected() bool {
	if o == nil || o.Selected.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Selected.Get()
}

// GetSelectedOk returns a tuple with the Selected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gateway) GetSelectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Selected.Get(), o.Selected.IsSet()
}

// HasSelected returns a boolean if a field has been set.
func (o *Gateway) HasSelected() bool {
	if o != nil && o.Selected.IsSet() {
		return true
	}

	return false
}

// SetSelected gets a reference to the given NullableBool and assigns it to the Selected field.
func (o *Gateway) SetSelected(v bool) {
	o.Selected.Set(&v)
}

// SetSelectedNil sets the value for Selected to be an explicit nil
func (o *Gateway) SetSelectedNil() {
	o.Selected.Set(nil)
}

// UnsetSelected ensures that no value is present for Selected, not even an explicit nil
func (o *Gateway) UnsetSelected() {
	o.Selected.Unset()
}

// GetGatewayReselection returns the GatewayReselection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Gateway) GetGatewayReselection() bool {
	if o == nil || o.GatewayReselection.Get() == nil {
		var ret bool
		return ret
	}
	return *o.GatewayReselection.Get()
}

// GetGatewayReselectionOk returns a tuple with the GatewayReselection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gateway) GetGatewayReselectionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.GatewayReselection.Get(), o.GatewayReselection.IsSet()
}

// HasGatewayReselection returns a boolean if a field has been set.
func (o *Gateway) HasGatewayReselection() bool {
	if o != nil && o.GatewayReselection.IsSet() {
		return true
	}

	return false
}

// SetGatewayReselection gets a reference to the given NullableBool and assigns it to the GatewayReselection field.
func (o *Gateway) SetGatewayReselection(v bool) {
	o.GatewayReselection.Set(&v)
}

// SetGatewayReselectionNil sets the value for GatewayReselection to be an explicit nil
func (o *Gateway) SetGatewayReselectionNil() {
	o.GatewayReselection.Set(nil)
}

// UnsetGatewayReselection ensures that no value is present for GatewayReselection, not even an explicit nil
func (o *Gateway) UnsetGatewayReselection() {
	o.GatewayReselection.Unset()
}

// GetPops returns the Pops field value if set, zero value otherwise.
func (o *Gateway) GetPops() []AdnGateway {
	if o == nil || o.Pops == nil {
		var ret []AdnGateway
		return ret
	}
	return o.Pops
}

// GetPopsOk returns a tuple with the Pops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetPopsOk() ([]AdnGateway, bool) {
	if o == nil || o.Pops == nil {
		return nil, false
	}
	return o.Pops, true
}

// HasPops returns a boolean if a field has been set.
func (o *Gateway) HasPops() bool {
	if o != nil && o.Pops != nil {
		return true
	}

	return false
}

// SetPops gets a reference to the given []AdnGateway and assigns it to the Pops field.
func (o *Gateway) SetPops(v []AdnGateway) {
	o.Pops = v
}

func (o Gateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SelectionLastRequested.IsSet() {
		toSerialize["selection_last_requested"] = o.SelectionLastRequested.Get()
	}
	if o.FollowRedirect.IsSet() {
		toSerialize["follow_redirect"] = o.FollowRedirect.Get()
	}
	if o.PopLastUpdate.IsSet() {
		toSerialize["pop_last_update"] = o.PopLastUpdate.Get()
	}
	if o.SelectionError.IsSet() {
		toSerialize["selection_error"] = o.SelectionError.Get()
	}
	if o.Selected.IsSet() {
		toSerialize["selected"] = o.Selected.Get()
	}
	if o.GatewayReselection.IsSet() {
		toSerialize["gateway_reselection"] = o.GatewayReselection.Get()
	}
	if o.Pops != nil {
		toSerialize["pops"] = o.Pops
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Gateway) UnmarshalJSON(bytes []byte) (err error) {
	varGateway := _Gateway{}

	if err = json.Unmarshal(bytes, &varGateway); err == nil {
		*o = Gateway(varGateway)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "selection_last_requested")
		delete(additionalProperties, "follow_redirect")
		delete(additionalProperties, "pop_last_update")
		delete(additionalProperties, "selection_error")
		delete(additionalProperties, "selected")
		delete(additionalProperties, "gateway_reselection")
		delete(additionalProperties, "pops")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGateway struct {
	value *Gateway
	isSet bool
}

func (v NullableGateway) Get() *Gateway {
	return v.value
}

func (v *NullableGateway) Set(val *Gateway) {
	v.value = val
	v.isSet = true
}

func (v NullableGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGateway(val Gateway) NullableGateway {
	return NullableGateway{value: &val, isSet: true}
}

func (v NullableGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
