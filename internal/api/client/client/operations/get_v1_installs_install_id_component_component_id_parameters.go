// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetV1InstallsInstallIDComponentComponentIDParams creates a new GetV1InstallsInstallIDComponentComponentIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1InstallsInstallIDComponentComponentIDParams() *GetV1InstallsInstallIDComponentComponentIDParams {
	return &GetV1InstallsInstallIDComponentComponentIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1InstallsInstallIDComponentComponentIDParamsWithTimeout creates a new GetV1InstallsInstallIDComponentComponentIDParams object
// with the ability to set a timeout on a request.
func NewGetV1InstallsInstallIDComponentComponentIDParamsWithTimeout(timeout time.Duration) *GetV1InstallsInstallIDComponentComponentIDParams {
	return &GetV1InstallsInstallIDComponentComponentIDParams{
		timeout: timeout,
	}
}

// NewGetV1InstallsInstallIDComponentComponentIDParamsWithContext creates a new GetV1InstallsInstallIDComponentComponentIDParams object
// with the ability to set a context for a request.
func NewGetV1InstallsInstallIDComponentComponentIDParamsWithContext(ctx context.Context) *GetV1InstallsInstallIDComponentComponentIDParams {
	return &GetV1InstallsInstallIDComponentComponentIDParams{
		Context: ctx,
	}
}

// NewGetV1InstallsInstallIDComponentComponentIDParamsWithHTTPClient creates a new GetV1InstallsInstallIDComponentComponentIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1InstallsInstallIDComponentComponentIDParamsWithHTTPClient(client *http.Client) *GetV1InstallsInstallIDComponentComponentIDParams {
	return &GetV1InstallsInstallIDComponentComponentIDParams{
		HTTPClient: client,
	}
}

/*
GetV1InstallsInstallIDComponentComponentIDParams contains all the parameters to send to the API endpoint

	for the get v1 installs install ID component component ID operation.

	Typically these are written to a http.Request.
*/
type GetV1InstallsInstallIDComponentComponentIDParams struct {

	/* ComponentID.

	   component ID
	*/
	ComponentID string

	/* InstallID.

	   install ID
	*/
	InstallID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 installs install ID component component ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1InstallsInstallIDComponentComponentIDParams) WithDefaults() *GetV1InstallsInstallIDComponentComponentIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 installs install ID component component ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1InstallsInstallIDComponentComponentIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 installs install ID component component ID params
func (o *GetV1InstallsInstallIDComponentComponentIDParams) WithTimeout(timeout time.Duration) *GetV1InstallsInstallIDComponentComponentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 installs install ID component component ID params
func (o *GetV1InstallsInstallIDComponentComponentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 installs install ID component component ID params
func (o *GetV1InstallsInstallIDComponentComponentIDParams) WithContext(ctx context.Context) *GetV1InstallsInstallIDComponentComponentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 installs install ID component component ID params
func (o *GetV1InstallsInstallIDComponentComponentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 installs install ID component component ID params
func (o *GetV1InstallsInstallIDComponentComponentIDParams) WithHTTPClient(client *http.Client) *GetV1InstallsInstallIDComponentComponentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 installs install ID component component ID params
func (o *GetV1InstallsInstallIDComponentComponentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComponentID adds the componentID to the get v1 installs install ID component component ID params
func (o *GetV1InstallsInstallIDComponentComponentIDParams) WithComponentID(componentID string) *GetV1InstallsInstallIDComponentComponentIDParams {
	o.SetComponentID(componentID)
	return o
}

// SetComponentID adds the componentId to the get v1 installs install ID component component ID params
func (o *GetV1InstallsInstallIDComponentComponentIDParams) SetComponentID(componentID string) {
	o.ComponentID = componentID
}

// WithInstallID adds the installID to the get v1 installs install ID component component ID params
func (o *GetV1InstallsInstallIDComponentComponentIDParams) WithInstallID(installID string) *GetV1InstallsInstallIDComponentComponentIDParams {
	o.SetInstallID(installID)
	return o
}

// SetInstallID adds the installId to the get v1 installs install ID component component ID params
func (o *GetV1InstallsInstallIDComponentComponentIDParams) SetInstallID(installID string) {
	o.InstallID = installID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1InstallsInstallIDComponentComponentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param component_id
	if err := r.SetPathParam("component_id", o.ComponentID); err != nil {
		return err
	}

	// path param install_id
	if err := r.SetPathParam("install_id", o.InstallID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}