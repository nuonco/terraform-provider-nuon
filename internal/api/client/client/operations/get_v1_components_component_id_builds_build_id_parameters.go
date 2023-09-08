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

// NewGetV1ComponentsComponentIDBuildsBuildIDParams creates a new GetV1ComponentsComponentIDBuildsBuildIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ComponentsComponentIDBuildsBuildIDParams() *GetV1ComponentsComponentIDBuildsBuildIDParams {
	return &GetV1ComponentsComponentIDBuildsBuildIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ComponentsComponentIDBuildsBuildIDParamsWithTimeout creates a new GetV1ComponentsComponentIDBuildsBuildIDParams object
// with the ability to set a timeout on a request.
func NewGetV1ComponentsComponentIDBuildsBuildIDParamsWithTimeout(timeout time.Duration) *GetV1ComponentsComponentIDBuildsBuildIDParams {
	return &GetV1ComponentsComponentIDBuildsBuildIDParams{
		timeout: timeout,
	}
}

// NewGetV1ComponentsComponentIDBuildsBuildIDParamsWithContext creates a new GetV1ComponentsComponentIDBuildsBuildIDParams object
// with the ability to set a context for a request.
func NewGetV1ComponentsComponentIDBuildsBuildIDParamsWithContext(ctx context.Context) *GetV1ComponentsComponentIDBuildsBuildIDParams {
	return &GetV1ComponentsComponentIDBuildsBuildIDParams{
		Context: ctx,
	}
}

// NewGetV1ComponentsComponentIDBuildsBuildIDParamsWithHTTPClient creates a new GetV1ComponentsComponentIDBuildsBuildIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ComponentsComponentIDBuildsBuildIDParamsWithHTTPClient(client *http.Client) *GetV1ComponentsComponentIDBuildsBuildIDParams {
	return &GetV1ComponentsComponentIDBuildsBuildIDParams{
		HTTPClient: client,
	}
}

/*
GetV1ComponentsComponentIDBuildsBuildIDParams contains all the parameters to send to the API endpoint

	for the get v1 components component ID builds build ID operation.

	Typically these are written to a http.Request.
*/
type GetV1ComponentsComponentIDBuildsBuildIDParams struct {

	/* BuildID.

	   build ID
	*/
	BuildID string

	/* ComponentID.

	   component ID
	*/
	ComponentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 components component ID builds build ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ComponentsComponentIDBuildsBuildIDParams) WithDefaults() *GetV1ComponentsComponentIDBuildsBuildIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 components component ID builds build ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ComponentsComponentIDBuildsBuildIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 components component ID builds build ID params
func (o *GetV1ComponentsComponentIDBuildsBuildIDParams) WithTimeout(timeout time.Duration) *GetV1ComponentsComponentIDBuildsBuildIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 components component ID builds build ID params
func (o *GetV1ComponentsComponentIDBuildsBuildIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 components component ID builds build ID params
func (o *GetV1ComponentsComponentIDBuildsBuildIDParams) WithContext(ctx context.Context) *GetV1ComponentsComponentIDBuildsBuildIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 components component ID builds build ID params
func (o *GetV1ComponentsComponentIDBuildsBuildIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 components component ID builds build ID params
func (o *GetV1ComponentsComponentIDBuildsBuildIDParams) WithHTTPClient(client *http.Client) *GetV1ComponentsComponentIDBuildsBuildIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 components component ID builds build ID params
func (o *GetV1ComponentsComponentIDBuildsBuildIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildID adds the buildID to the get v1 components component ID builds build ID params
func (o *GetV1ComponentsComponentIDBuildsBuildIDParams) WithBuildID(buildID string) *GetV1ComponentsComponentIDBuildsBuildIDParams {
	o.SetBuildID(buildID)
	return o
}

// SetBuildID adds the buildId to the get v1 components component ID builds build ID params
func (o *GetV1ComponentsComponentIDBuildsBuildIDParams) SetBuildID(buildID string) {
	o.BuildID = buildID
}

// WithComponentID adds the componentID to the get v1 components component ID builds build ID params
func (o *GetV1ComponentsComponentIDBuildsBuildIDParams) WithComponentID(componentID string) *GetV1ComponentsComponentIDBuildsBuildIDParams {
	o.SetComponentID(componentID)
	return o
}

// SetComponentID adds the componentId to the get v1 components component ID builds build ID params
func (o *GetV1ComponentsComponentIDBuildsBuildIDParams) SetComponentID(componentID string) {
	o.ComponentID = componentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ComponentsComponentIDBuildsBuildIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param build_id
	if err := r.SetPathParam("build_id", o.BuildID); err != nil {
		return err
	}

	// path param component_id
	if err := r.SetPathParam("component_id", o.ComponentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}