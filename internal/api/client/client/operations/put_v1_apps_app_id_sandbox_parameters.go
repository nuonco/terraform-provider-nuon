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

	"github.com/nuonco/terraform-provider-nuon/internal/api/client/models"
)

// NewPutV1AppsAppIDSandboxParams creates a new PutV1AppsAppIDSandboxParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutV1AppsAppIDSandboxParams() *PutV1AppsAppIDSandboxParams {
	return &PutV1AppsAppIDSandboxParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutV1AppsAppIDSandboxParamsWithTimeout creates a new PutV1AppsAppIDSandboxParams object
// with the ability to set a timeout on a request.
func NewPutV1AppsAppIDSandboxParamsWithTimeout(timeout time.Duration) *PutV1AppsAppIDSandboxParams {
	return &PutV1AppsAppIDSandboxParams{
		timeout: timeout,
	}
}

// NewPutV1AppsAppIDSandboxParamsWithContext creates a new PutV1AppsAppIDSandboxParams object
// with the ability to set a context for a request.
func NewPutV1AppsAppIDSandboxParamsWithContext(ctx context.Context) *PutV1AppsAppIDSandboxParams {
	return &PutV1AppsAppIDSandboxParams{
		Context: ctx,
	}
}

// NewPutV1AppsAppIDSandboxParamsWithHTTPClient creates a new PutV1AppsAppIDSandboxParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutV1AppsAppIDSandboxParamsWithHTTPClient(client *http.Client) *PutV1AppsAppIDSandboxParams {
	return &PutV1AppsAppIDSandboxParams{
		HTTPClient: client,
	}
}

/*
PutV1AppsAppIDSandboxParams contains all the parameters to send to the API endpoint

	for the put v1 apps app ID sandbox operation.

	Typically these are written to a http.Request.
*/
type PutV1AppsAppIDSandboxParams struct {

	/* AppID.

	   app ID
	*/
	AppID string

	/* Req.

	   Input
	*/
	Req *models.ServiceUpdateAppSandboxRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put v1 apps app ID sandbox params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1AppsAppIDSandboxParams) WithDefaults() *PutV1AppsAppIDSandboxParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put v1 apps app ID sandbox params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1AppsAppIDSandboxParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put v1 apps app ID sandbox params
func (o *PutV1AppsAppIDSandboxParams) WithTimeout(timeout time.Duration) *PutV1AppsAppIDSandboxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put v1 apps app ID sandbox params
func (o *PutV1AppsAppIDSandboxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put v1 apps app ID sandbox params
func (o *PutV1AppsAppIDSandboxParams) WithContext(ctx context.Context) *PutV1AppsAppIDSandboxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put v1 apps app ID sandbox params
func (o *PutV1AppsAppIDSandboxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put v1 apps app ID sandbox params
func (o *PutV1AppsAppIDSandboxParams) WithHTTPClient(client *http.Client) *PutV1AppsAppIDSandboxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put v1 apps app ID sandbox params
func (o *PutV1AppsAppIDSandboxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the put v1 apps app ID sandbox params
func (o *PutV1AppsAppIDSandboxParams) WithAppID(appID string) *PutV1AppsAppIDSandboxParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the put v1 apps app ID sandbox params
func (o *PutV1AppsAppIDSandboxParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithReq adds the req to the put v1 apps app ID sandbox params
func (o *PutV1AppsAppIDSandboxParams) WithReq(req *models.ServiceUpdateAppSandboxRequest) *PutV1AppsAppIDSandboxParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the put v1 apps app ID sandbox params
func (o *PutV1AppsAppIDSandboxParams) SetReq(req *models.ServiceUpdateAppSandboxRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *PutV1AppsAppIDSandboxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_id
	if err := r.SetPathParam("app_id", o.AppID); err != nil {
		return err
	}
	if o.Req != nil {
		if err := r.SetBodyParam(o.Req); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}