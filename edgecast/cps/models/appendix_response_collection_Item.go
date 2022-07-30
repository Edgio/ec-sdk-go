package models

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppendixCollectionItem appendix collection item
//
// model AppendixCollectionItem
type AppendixCollectionItem struct {

	// name
	Name string `json:"name"`

	// id
	ID int32 `json:"id"`
}

// Validate validates this collection item
func (m *AppendixCollectionItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this collection item based on context it is used
func (m *AppendixCollectionItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppendixCollectionItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppendixCollectionItem) UnmarshalBinary(b []byte) error {
	var res AppendixCollectionItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
