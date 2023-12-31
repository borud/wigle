// Code generated by go-swagger; DO NOT EDIT.

package statistics_and_information

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

// NewSiteStatsParams creates a new SiteStatsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSiteStatsParams() *SiteStatsParams {
	return &SiteStatsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSiteStatsParamsWithTimeout creates a new SiteStatsParams object
// with the ability to set a timeout on a request.
func NewSiteStatsParamsWithTimeout(timeout time.Duration) *SiteStatsParams {
	return &SiteStatsParams{
		timeout: timeout,
	}
}

// NewSiteStatsParamsWithContext creates a new SiteStatsParams object
// with the ability to set a context for a request.
func NewSiteStatsParamsWithContext(ctx context.Context) *SiteStatsParams {
	return &SiteStatsParams{
		Context: ctx,
	}
}

// NewSiteStatsParamsWithHTTPClient creates a new SiteStatsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSiteStatsParamsWithHTTPClient(client *http.Client) *SiteStatsParams {
	return &SiteStatsParams{
		HTTPClient: client,
	}
}

/*
SiteStatsParams contains all the parameters to send to the API endpoint

	for the site stats operation.

	Typically these are written to a http.Request.
*/
type SiteStatsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the site stats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SiteStatsParams) WithDefaults() *SiteStatsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the site stats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SiteStatsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the site stats params
func (o *SiteStatsParams) WithTimeout(timeout time.Duration) *SiteStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the site stats params
func (o *SiteStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the site stats params
func (o *SiteStatsParams) WithContext(ctx context.Context) *SiteStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the site stats params
func (o *SiteStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the site stats params
func (o *SiteStatsParams) WithHTTPClient(client *http.Client) *SiteStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the site stats params
func (o *SiteStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SiteStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
