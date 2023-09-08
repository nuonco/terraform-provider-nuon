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

// NewPostV1OrgsCurrentUserParams creates a new PostV1OrgsCurrentUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1OrgsCurrentUserParams() *PostV1OrgsCurrentUserParams {
	return &PostV1OrgsCurrentUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1OrgsCurrentUserParamsWithTimeout creates a new PostV1OrgsCurrentUserParams object
// with the ability to set a timeout on a request.
func NewPostV1OrgsCurrentUserParamsWithTimeout(timeout time.Duration) *PostV1OrgsCurrentUserParams {
	return &PostV1OrgsCurrentUserParams{
		timeout: timeout,
	}
}

// NewPostV1OrgsCurrentUserParamsWithContext creates a new PostV1OrgsCurrentUserParams object
// with the ability to set a context for a request.
func NewPostV1OrgsCurrentUserParamsWithContext(ctx context.Context) *PostV1OrgsCurrentUserParams {
	return &PostV1OrgsCurrentUserParams{
		Context: ctx,
	}
}

// NewPostV1OrgsCurrentUserParamsWithHTTPClient creates a new PostV1OrgsCurrentUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1OrgsCurrentUserParamsWithHTTPClient(client *http.Client) *PostV1OrgsCurrentUserParams {
	return &PostV1OrgsCurrentUserParams{
		HTTPClient: client,
	}
}

/*
PostV1OrgsCurrentUserParams contains all the parameters to send to the API endpoint

	for the post v1 orgs current user operation.

	Typically these are written to a http.Request.
*/
type PostV1OrgsCurrentUserParams struct {

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateOrgUserRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 orgs current user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1OrgsCurrentUserParams) WithDefaults() *PostV1OrgsCurrentUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 orgs current user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1OrgsCurrentUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 orgs current user params
func (o *PostV1OrgsCurrentUserParams) WithTimeout(timeout time.Duration) *PostV1OrgsCurrentUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 orgs current user params
func (o *PostV1OrgsCurrentUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 orgs current user params
func (o *PostV1OrgsCurrentUserParams) WithContext(ctx context.Context) *PostV1OrgsCurrentUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 orgs current user params
func (o *PostV1OrgsCurrentUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 orgs current user params
func (o *PostV1OrgsCurrentUserParams) WithHTTPClient(client *http.Client) *PostV1OrgsCurrentUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 orgs current user params
func (o *PostV1OrgsCurrentUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the post v1 orgs current user params
func (o *PostV1OrgsCurrentUserParams) WithReq(req *models.ServiceCreateOrgUserRequest) *PostV1OrgsCurrentUserParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the post v1 orgs current user params
func (o *PostV1OrgsCurrentUserParams) SetReq(req *models.ServiceCreateOrgUserRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1OrgsCurrentUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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