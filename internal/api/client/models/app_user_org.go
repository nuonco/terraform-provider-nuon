// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppUserOrg app user org
//
// swagger:model app.UserOrg
type AppUserOrg struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// org
	Org *AppOrg `json:"org,omitempty"`

	// parent relationship
	OrgID string `json:"orgID,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// user ID
	UserID string `json:"userID,omitempty"`
}

// Validate validates this app user org
func (m *AppUserOrg) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrg(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppUserOrg) validateOrg(formats strfmt.Registry) error {
	if swag.IsZero(m.Org) { // not required
		return nil
	}

	if m.Org != nil {
		if err := m.Org.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("org")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("org")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this app user org based on the context it is used
func (m *AppUserOrg) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrg(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppUserOrg) contextValidateOrg(ctx context.Context, formats strfmt.Registry) error {

	if m.Org != nil {

		if swag.IsZero(m.Org) { // not required
			return nil
		}

		if err := m.Org.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("org")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("org")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppUserOrg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppUserOrg) UnmarshalBinary(b []byte) error {
	var res AppUserOrg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}