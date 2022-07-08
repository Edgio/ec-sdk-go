// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CertificateCreate certificate create
//
// swagger:model CertificateCreate
type CertificateCreate struct {

	// auto renew
	AutoRenew bool `json:"auto_renew,omitempty"`

	// certificate authority
	// Enum: [DigiCert LetsEncrypt]
	CertificateAuthority string `json:"certificate_authority,omitempty"`

	// certificate label
	CertificateLabel string `json:"certificate_label,omitempty"`

	// dcv method
	// Enum: [Email DnsCnameToken DnsTxtToken]
	DcvMethod string `json:"dcv_method,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// domains
	Domains []*DomainCreateUpdate `json:"domains"`

	// organization
	Organization *OrganizationDetail `json:"organization,omitempty"`

	// validation type
	// Enum: [None DV OV EV]
	ValidationType string `json:"validation_type,omitempty"`
}

// Validate validates this certificate create
func (m *CertificateCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificateAuthority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDcvMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomains(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var certificateCreateTypeCertificateAuthorityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DigiCert","LetsEncrypt"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificateCreateTypeCertificateAuthorityPropEnum = append(certificateCreateTypeCertificateAuthorityPropEnum, v)
	}
}

const (

	// CertificateCreateCertificateAuthorityDigiCert captures enum value "DigiCert"
	CertificateCreateCertificateAuthorityDigiCert string = "DigiCert"

	// CertificateCreateCertificateAuthorityLetsEncrypt captures enum value "LetsEncrypt"
	CertificateCreateCertificateAuthorityLetsEncrypt string = "LetsEncrypt"
)

// prop value enum
func (m *CertificateCreate) validateCertificateAuthorityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificateCreateTypeCertificateAuthorityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificateCreate) validateCertificateAuthority(formats strfmt.Registry) error {
	if swag.IsZero(m.CertificateAuthority) { // not required
		return nil
	}

	// value enum
	if err := m.validateCertificateAuthorityEnum("certificate_authority", "body", m.CertificateAuthority); err != nil {
		return err
	}

	return nil
}

var certificateCreateTypeDcvMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Email","DnsCnameToken","DnsTxtToken"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificateCreateTypeDcvMethodPropEnum = append(certificateCreateTypeDcvMethodPropEnum, v)
	}
}

const (

	// CertificateCreateDcvMethodEmail captures enum value "Email"
	CertificateCreateDcvMethodEmail string = "Email"

	// CertificateCreateDcvMethodDNSCnameToken captures enum value "DnsCnameToken"
	CertificateCreateDcvMethodDNSCnameToken string = "DnsCnameToken"

	// CertificateCreateDcvMethodDNSTxtToken captures enum value "DnsTxtToken"
	CertificateCreateDcvMethodDNSTxtToken string = "DnsTxtToken"
)

// prop value enum
func (m *CertificateCreate) validateDcvMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificateCreateTypeDcvMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificateCreate) validateDcvMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.DcvMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateDcvMethodEnum("dcv_method", "body", m.DcvMethod); err != nil {
		return err
	}

	return nil
}

func (m *CertificateCreate) validateDomains(formats strfmt.Registry) error {
	if swag.IsZero(m.Domains) { // not required
		return nil
	}

	for i := 0; i < len(m.Domains); i++ {
		if swag.IsZero(m.Domains[i]) { // not required
			continue
		}

		if m.Domains[i] != nil {
			if err := m.Domains[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("domains" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("domains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CertificateCreate) validateOrganization(formats strfmt.Registry) error {
	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organization")
			}
			return err
		}
	}

	return nil
}

var certificateCreateTypeValidationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","DV","OV","EV"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificateCreateTypeValidationTypePropEnum = append(certificateCreateTypeValidationTypePropEnum, v)
	}
}

const (

	// CertificateCreateValidationTypeNone captures enum value "None"
	CertificateCreateValidationTypeNone string = "None"

	// CertificateCreateValidationTypeDV captures enum value "DV"
	CertificateCreateValidationTypeDV string = "DV"

	// CertificateCreateValidationTypeOV captures enum value "OV"
	CertificateCreateValidationTypeOV string = "OV"

	// CertificateCreateValidationTypeEV captures enum value "EV"
	CertificateCreateValidationTypeEV string = "EV"
)

// prop value enum
func (m *CertificateCreate) validateValidationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificateCreateTypeValidationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificateCreate) validateValidationType(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateValidationTypeEnum("validation_type", "body", m.ValidationType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this certificate create based on the context it is used
func (m *CertificateCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDomains(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrganization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateCreate) contextValidateDomains(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Domains); i++ {

		if m.Domains[i] != nil {
			if err := m.Domains[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("domains" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("domains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CertificateCreate) contextValidateOrganization(ctx context.Context, formats strfmt.Registry) error {

	if m.Organization != nil {
		if err := m.Organization.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organization")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificateCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateCreate) UnmarshalBinary(b []byte) error {
	var res CertificateCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}