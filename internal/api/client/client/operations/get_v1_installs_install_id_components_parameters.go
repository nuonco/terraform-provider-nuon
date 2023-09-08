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

// NewGetV1InstallsInstallIDComponentsParams creates a new GetV1InstallsInstallIDComponentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1InstallsInstallIDComponentsParams() *GetV1InstallsInstallIDComponentsParams {
	return &GetV1InstallsInstallIDComponentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1InstallsInstallIDComponentsParamsWithTimeout creates a new GetV1InstallsInstallIDComponentsParams object
// with the ability to set a timeout on a request.
func NewGetV1InstallsInstallIDComponentsParamsWithTimeout(timeout time.Duration) *GetV1InstallsInstallIDComponentsParams {
	return &GetV1InstallsInstallIDComponentsParams{
		timeout: timeout,
	}
}

// NewGetV1InstallsInstallIDComponentsParamsWithContext creates a new GetV1InstallsInstallIDComponentsParams object
// with the ability to set a context for a request.
func NewGetV1InstallsInstallIDComponentsParamsWithContext(ctx context.Context) *GetV1InstallsInstallIDComponentsParams {
	return &GetV1InstallsInstallIDComponentsParams{
		Context: ctx,
	}
}

// NewGetV1InstallsInstallIDComponentsParamsWithHTTPClient creates a new GetV1InstallsInstallIDComponentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1InstallsInstallIDComponentsParamsWithHTTPClient(client *http.Client) *GetV1InstallsInstallIDComponentsParams {
	return &GetV1InstallsInstallIDComponentsParams{
		HTTPClient: client,
	}
}

/*
GetV1InstallsInstallIDComponentsParams contains all the parameters to send to the API endpoint

	for the get v1 installs install ID components operation.

	Typically these are written to a http.Request.
*/
type GetV1InstallsInstallIDComponentsParams struct {

	/* InstallID.

	   install ID
	*/
	InstallID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 installs install ID components params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1InstallsInstallIDComponentsParams) WithDefaults() *GetV1InstallsInstallIDComponentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 installs install ID components params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1InstallsInstallIDComponentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 installs install ID components params
func (o *GetV1InstallsInstallIDComponentsParams) WithTimeout(timeout time.Duration) *GetV1InstallsInstallIDComponentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 installs install ID components params
func (o *GetV1InstallsInstallIDComponentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 installs install ID components params
func (o *GetV1InstallsInstallIDComponentsParams) WithContext(ctx context.Context) *GetV1InstallsInstallIDComponentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 installs install ID components params
func (o *GetV1InstallsInstallIDComponentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 installs install ID components params
func (o *GetV1InstallsInstallIDComponentsParams) WithHTTPClient(client *http.Client) *GetV1InstallsInstallIDComponentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 installs install ID components params
func (o *GetV1InstallsInstallIDComponentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstallID adds the installID to the get v1 installs install ID components params
func (o *GetV1InstallsInstallIDComponentsParams) WithInstallID(installID string) *GetV1InstallsInstallIDComponentsParams {
	o.SetInstallID(installID)
	return o
}

// SetInstallID adds the installId to the get v1 installs install ID components params
func (o *GetV1InstallsInstallIDComponentsParams) SetInstallID(installID string) {
	o.InstallID = installID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1InstallsInstallIDComponentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param install_id
	if err := r.SetPathParam("install_id", o.InstallID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}