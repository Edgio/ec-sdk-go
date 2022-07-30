package models

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppendixCollection appendix collection
//
// model AppendixCollection
type AppendixCollection struct {

	// at id
	AtID string `json:"@id,omitempty"`

	// at type
	AtType string `json:"@type,omitempty"`

	// items
	Items []*AppendixCollectionItem `json:"items"`

	// total items
	TotalItems int32 `json:"total_items,omitempty"`
}

// Validate validates this collection
func (m *AppendixCollection) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this collection based on context it is used
func (m *AppendixCollection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppendixCollection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppendixCollection) UnmarshalBinary(b []byte) error {
	var res AppendixCollection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
