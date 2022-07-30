package models

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Country Code
//
// model CountryCode
type CountryCode struct {

	// at id
	AtID string `json:"@id,omitempty"`

	// at type
	AtType string `json:"@type,omitempty"`

	// items
	Items []*CountryCodeItem `json:"items"`

	// total items
	TotalItems int32 `json:"total_items,omitempty"`
}

// CountryCode item
//
// model CountryCodeItem
type CountryCodeItem struct {

	// country
	Country string `json:"country"`

	// two letter code
	TwoLetterCode string `json:"two_letter_code"`
}

// Validate validates this country code
func (m *CountryCode) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this country code based on context it is used
func (m *CountryCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CountryCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CountryCode) UnmarshalBinary(b []byte) error {
	var res CountryCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Validate validates this country code item
func (m *CountryCodeItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this country code item based on context it is used
func (m *CountryCodeItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CountryCodeItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CountryCodeItem) UnmarshalBinary(b []byte) error {
	var res CountryCodeItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
