// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrganizationDetailDto organization detail dto
//
// swagger:model OrganizationDetailDto
type OrganizationDetailDto struct {

	// city
	City string `json:"city,omitempty"`

	// company address
	CompanyAddress string `json:"company_address,omitempty"`

	// company address2
	CompanyAddress2 string `json:"company_address2,omitempty"`

	// company name
	CompanyName string `json:"company_name,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// zip code
	ZipCode string `json:"zip_code,omitempty"`
}

// Validate validates this organization detail dto
func (m *OrganizationDetailDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this organization detail dto based on context it is used
func (m *OrganizationDetailDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationDetailDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationDetailDto) UnmarshalBinary(b []byte) error {
	var res OrganizationDetailDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}