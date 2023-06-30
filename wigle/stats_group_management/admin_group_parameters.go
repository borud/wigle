// Code generated by go-swagger; DO NOT EDIT.

package stats_group_management

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

// NewAdminGroupParams creates a new AdminGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminGroupParams() *AdminGroupParams {
	return &AdminGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminGroupParamsWithTimeout creates a new AdminGroupParams object
// with the ability to set a timeout on a request.
func NewAdminGroupParamsWithTimeout(timeout time.Duration) *AdminGroupParams {
	return &AdminGroupParams{
		timeout: timeout,
	}
}

// NewAdminGroupParamsWithContext creates a new AdminGroupParams object
// with the ability to set a context for a request.
func NewAdminGroupParamsWithContext(ctx context.Context) *AdminGroupParams {
	return &AdminGroupParams{
		Context: ctx,
	}
}

// NewAdminGroupParamsWithHTTPClient creates a new AdminGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminGroupParamsWithHTTPClient(client *http.Client) *AdminGroupParams {
	return &AdminGroupParams{
		HTTPClient: client,
	}
}

/*
AdminGroupParams contains all the parameters to send to the API endpoint

	for the admin group operation.

	Typically these are written to a http.Request.
*/
type AdminGroupParams struct {

	/* Groupid.

	   The unique string key of the group
	*/
	Groupid *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminGroupParams) WithDefaults() *AdminGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin group params
func (o *AdminGroupParams) WithTimeout(timeout time.Duration) *AdminGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin group params
func (o *AdminGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin group params
func (o *AdminGroupParams) WithContext(ctx context.Context) *AdminGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin group params
func (o *AdminGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin group params
func (o *AdminGroupParams) WithHTTPClient(client *http.Client) *AdminGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin group params
func (o *AdminGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupid adds the groupid to the admin group params
func (o *AdminGroupParams) WithGroupid(groupid *string) *AdminGroupParams {
	o.SetGroupid(groupid)
	return o
}

// SetGroupid adds the groupid to the admin group params
func (o *AdminGroupParams) SetGroupid(groupid *string) {
	o.Groupid = groupid
}

// WriteToRequest writes these params to a swagger request
func (o *AdminGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Groupid != nil {

		// query param groupid
		var qrGroupid string

		if o.Groupid != nil {
			qrGroupid = *o.Groupid
		}
		qGroupid := qrGroupid
		if qGroupid != "" {

			if err := r.SetQueryParam("groupid", qGroupid); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
