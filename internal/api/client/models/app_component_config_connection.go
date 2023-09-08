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

// AppComponentConfigConnection app component config connection
//
// swagger:model app.ComponentConfigConnection
type AppComponentConfigConnection struct {

	// component id
	ComponentID string `json:"component_id,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// docker build
	DockerBuild *AppDockerBuildComponentConfig `json:"docker_build,omitempty"`

	// external image
	ExternalImage *AppExternalImageComponentConfig `json:"external_image,omitempty"`

	// helm
	Helm *AppHelmComponentConfig `json:"helm,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// terraform module
	TerraformModule *AppTerraformModuleComponentConfig `json:"terraform_module,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app component config connection
func (m *AppComponentConfigConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDockerBuild(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHelm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerraformModule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppComponentConfigConnection) validateDockerBuild(formats strfmt.Registry) error {
	if swag.IsZero(m.DockerBuild) { // not required
		return nil
	}

	if m.DockerBuild != nil {
		if err := m.DockerBuild.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("docker_build")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("docker_build")
			}
			return err
		}
	}

	return nil
}

func (m *AppComponentConfigConnection) validateExternalImage(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalImage) { // not required
		return nil
	}

	if m.ExternalImage != nil {
		if err := m.ExternalImage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external_image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("external_image")
			}
			return err
		}
	}

	return nil
}

func (m *AppComponentConfigConnection) validateHelm(formats strfmt.Registry) error {
	if swag.IsZero(m.Helm) { // not required
		return nil
	}

	if m.Helm != nil {
		if err := m.Helm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("helm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("helm")
			}
			return err
		}
	}

	return nil
}

func (m *AppComponentConfigConnection) validateTerraformModule(formats strfmt.Registry) error {
	if swag.IsZero(m.TerraformModule) { // not required
		return nil
	}

	if m.TerraformModule != nil {
		if err := m.TerraformModule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terraform_module")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terraform_module")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this app component config connection based on the context it is used
func (m *AppComponentConfigConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDockerBuild(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHelm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTerraformModule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppComponentConfigConnection) contextValidateDockerBuild(ctx context.Context, formats strfmt.Registry) error {

	if m.DockerBuild != nil {

		if swag.IsZero(m.DockerBuild) { // not required
			return nil
		}

		if err := m.DockerBuild.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("docker_build")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("docker_build")
			}
			return err
		}
	}

	return nil
}

func (m *AppComponentConfigConnection) contextValidateExternalImage(ctx context.Context, formats strfmt.Registry) error {

	if m.ExternalImage != nil {

		if swag.IsZero(m.ExternalImage) { // not required
			return nil
		}

		if err := m.ExternalImage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external_image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("external_image")
			}
			return err
		}
	}

	return nil
}

func (m *AppComponentConfigConnection) contextValidateHelm(ctx context.Context, formats strfmt.Registry) error {

	if m.Helm != nil {

		if swag.IsZero(m.Helm) { // not required
			return nil
		}

		if err := m.Helm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("helm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("helm")
			}
			return err
		}
	}

	return nil
}

func (m *AppComponentConfigConnection) contextValidateTerraformModule(ctx context.Context, formats strfmt.Registry) error {

	if m.TerraformModule != nil {

		if swag.IsZero(m.TerraformModule) { // not required
			return nil
		}

		if err := m.TerraformModule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terraform_module")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terraform_module")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppComponentConfigConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppComponentConfigConnection) UnmarshalBinary(b []byte) error {
	var res AppComponentConfigConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}