// Code generated by go-swagger; DO NOT EDIT.

package rtldmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RtldSplunkSettingDto rtld splunk setting dto
//
// swagger:model RtldSplunkSettingDto
type RtldSplunkSettingDto struct {

	// masked token
	MaskedToken string `json:"masked_token,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this rtld splunk setting dto
func (m *RtldSplunkSettingDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rtld splunk setting dto based on context it is used
func (m *RtldSplunkSettingDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RtldSplunkSettingDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RtldSplunkSettingDto) UnmarshalBinary(b []byte) error {
	var res RtldSplunkSettingDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}