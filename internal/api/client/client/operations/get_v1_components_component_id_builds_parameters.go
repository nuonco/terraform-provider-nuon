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

// NewGetV1ComponentsComponentIDBuildsParams creates a new GetV1ComponentsComponentIDBuildsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ComponentsComponentIDBuildsParams() *GetV1ComponentsComponentIDBuildsParams {
	return &GetV1ComponentsComponentIDBuildsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ComponentsComponentIDBuildsParamsWithTimeout creates a new GetV1ComponentsComponentIDBuildsParams object
// with the ability to set a timeout on a request.
func NewGetV1ComponentsComponentIDBuildsParamsWithTimeout(timeout time.Duration) *GetV1ComponentsComponentIDBuildsParams {
	return &GetV1ComponentsComponentIDBuildsParams{
		timeout: timeout,
	}
}

// NewGetV1ComponentsComponentIDBuildsParamsWithContext creates a new GetV1ComponentsComponentIDBuildsParams object
// with the ability to set a context for a request.
func NewGetV1ComponentsComponentIDBuildsParamsWithContext(ctx context.Context) *GetV1ComponentsComponentIDBuildsParams {
	return &GetV1ComponentsComponentIDBuildsParams{
		Context: ctx,
	}
}

// NewGetV1ComponentsComponentIDBuildsParamsWithHTTPClient creates a new GetV1ComponentsComponentIDBuildsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ComponentsComponentIDBuildsParamsWithHTTPClient(client *http.Client) *GetV1ComponentsComponentIDBuildsParams {
	return &GetV1ComponentsComponentIDBuildsParams{
		HTTPClient: client,
	}
}

/*
GetV1ComponentsComponentIDBuildsParams contains all the parameters to send to the API endpoint

	for the get v1 components component ID builds operation.

	Typically these are written to a http.Request.
*/
type GetV1ComponentsComponentIDBuildsParams struct {

	/* ComponentID.

	   component ID
	*/
	ComponentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 components component ID builds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ComponentsComponentIDBuildsParams) WithDefaults() *GetV1ComponentsComponentIDBuildsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 components component ID builds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ComponentsComponentIDBuildsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 components component ID builds params
func (o *GetV1ComponentsComponentIDBuildsParams) WithTimeout(timeout time.Duration) *GetV1ComponentsComponentIDBuildsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 components component ID builds params
func (o *GetV1ComponentsComponentIDBuildsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 components component ID builds params
func (o *GetV1ComponentsComponentIDBuildsParams) WithContext(ctx context.Context) *GetV1ComponentsComponentIDBuildsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 components component ID builds params
func (o *GetV1ComponentsComponentIDBuildsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 components component ID builds params
func (o *GetV1ComponentsComponentIDBuildsParams) WithHTTPClient(client *http.Client) *GetV1ComponentsComponentIDBuildsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 components component ID builds params
func (o *GetV1ComponentsComponentIDBuildsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComponentID adds the componentID to the get v1 components component ID builds params
func (o *GetV1ComponentsComponentIDBuildsParams) WithComponentID(componentID string) *GetV1ComponentsComponentIDBuildsParams {
	o.SetComponentID(componentID)
	return o
}

// SetComponentID adds the componentId to the get v1 components component ID builds params
func (o *GetV1ComponentsComponentIDBuildsParams) SetComponentID(componentID string) {
	o.ComponentID = componentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ComponentsComponentIDBuildsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param component_id
	if err := r.SetPathParam("component_id", o.ComponentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}