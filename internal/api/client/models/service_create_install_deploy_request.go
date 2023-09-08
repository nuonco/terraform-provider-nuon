// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceCreateInstallDeployRequest service create install deploy request
//
// swagger:model service.CreateInstallDeployRequest
type ServiceCreateInstallDeployRequest struct {

	// build id
	BuildID string `json:"build_id,omitempty"`
}

// Validate validates this service create install deploy request
func (m *ServiceCreateInstallDeployRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service create install deploy request based on context it is used
func (m *ServiceCreateInstallDeployRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCreateInstallDeployRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCreateInstallDeployRequest) UnmarshalBinary(b []byte) error {
	var res ServiceCreateInstallDeployRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}