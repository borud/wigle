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

// NewGroupMembersParams creates a new GroupMembersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupMembersParams() *GroupMembersParams {
	return &GroupMembersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupMembersParamsWithTimeout creates a new GroupMembersParams object
// with the ability to set a timeout on a request.
func NewGroupMembersParamsWithTimeout(timeout time.Duration) *GroupMembersParams {
	return &GroupMembersParams{
		timeout: timeout,
	}
}

// NewGroupMembersParamsWithContext creates a new GroupMembersParams object
// with the ability to set a context for a request.
func NewGroupMembersParamsWithContext(ctx context.Context) *GroupMembersParams {
	return &GroupMembersParams{
		Context: ctx,
	}
}

// NewGroupMembersParamsWithHTTPClient creates a new GroupMembersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupMembersParamsWithHTTPClient(client *http.Client) *GroupMembersParams {
	return &GroupMembersParams{
		HTTPClient: client,
	}
}

/*
GroupMembersParams contains all the parameters to send to the API endpoint

	for the group members operation.

	Typically these are written to a http.Request.
*/
type GroupMembersParams struct {

	/* Groupid.

	   The unique numeric ID of the group
	*/
	Groupid *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the group members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupMembersParams) WithDefaults() *GroupMembersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the group members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupMembersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the group members params
func (o *GroupMembersParams) WithTimeout(timeout time.Duration) *GroupMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group members params
func (o *GroupMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group members params
func (o *GroupMembersParams) WithContext(ctx context.Context) *GroupMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group members params
func (o *GroupMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group members params
func (o *GroupMembersParams) WithHTTPClient(client *http.Client) *GroupMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group members params
func (o *GroupMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupid adds the groupid to the group members params
func (o *GroupMembersParams) WithGroupid(groupid *string) *GroupMembersParams {
	o.SetGroupid(groupid)
	return o
}

// SetGroupid adds the groupid to the group members params
func (o *GroupMembersParams) SetGroupid(groupid *string) {
	o.Groupid = groupid
}

// WriteToRequest writes these params to a swagger request
func (o *GroupMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
