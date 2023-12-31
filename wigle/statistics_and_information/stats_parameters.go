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
	"github.com/go-openapi/swag"
)

// NewStatsParams creates a new StatsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStatsParams() *StatsParams {
	return &StatsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStatsParamsWithTimeout creates a new StatsParams object
// with the ability to set a timeout on a request.
func NewStatsParamsWithTimeout(timeout time.Duration) *StatsParams {
	return &StatsParams{
		timeout: timeout,
	}
}

// NewStatsParamsWithContext creates a new StatsParams object
// with the ability to set a context for a request.
func NewStatsParamsWithContext(ctx context.Context) *StatsParams {
	return &StatsParams{
		Context: ctx,
	}
}

// NewStatsParamsWithHTTPClient creates a new StatsParams object
// with the ability to set a custom HTTPClient for a request.
func NewStatsParamsWithHTTPClient(client *http.Client) *StatsParams {
	return &StatsParams{
		HTTPClient: client,
	}
}

/*
StatsParams contains all the parameters to send to the API endpoint

	for the stats operation.

	Typically these are written to a http.Request.
*/
type StatsParams struct {

	/* Pageend.

	   The last record to request according to the 'sort' parameter

	   Format: int64
	*/
	Pageend *int64

	/* Pagestart.

	   The first record to request according to the 'sort' parameter

	   Format: int64
	*/
	Pagestart *int64

	/* Sort.

	   The criteria by which to sort the results. Values are ['discovered', 'total', 'monthcount', 'prevmonthcount', 'gendisc', 'gentotal', 'firsttransid', 'lasttransid']

	   Default: "discovered"
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatsParams) WithDefaults() *StatsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatsParams) SetDefaults() {
	var (
		pagestartDefault = int64(0)

		sortDefault = string("discovered")
	)

	val := StatsParams{
		Pagestart: &pagestartDefault,
		Sort:      &sortDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the stats params
func (o *StatsParams) WithTimeout(timeout time.Duration) *StatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stats params
func (o *StatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stats params
func (o *StatsParams) WithContext(ctx context.Context) *StatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stats params
func (o *StatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stats params
func (o *StatsParams) WithHTTPClient(client *http.Client) *StatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stats params
func (o *StatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageend adds the pageend to the stats params
func (o *StatsParams) WithPageend(pageend *int64) *StatsParams {
	o.SetPageend(pageend)
	return o
}

// SetPageend adds the pageend to the stats params
func (o *StatsParams) SetPageend(pageend *int64) {
	o.Pageend = pageend
}

// WithPagestart adds the pagestart to the stats params
func (o *StatsParams) WithPagestart(pagestart *int64) *StatsParams {
	o.SetPagestart(pagestart)
	return o
}

// SetPagestart adds the pagestart to the stats params
func (o *StatsParams) SetPagestart(pagestart *int64) {
	o.Pagestart = pagestart
}

// WithSort adds the sort to the stats params
func (o *StatsParams) WithSort(sort *string) *StatsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the stats params
func (o *StatsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *StatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Pageend != nil {

		// query param pageend
		var qrPageend int64

		if o.Pageend != nil {
			qrPageend = *o.Pageend
		}
		qPageend := swag.FormatInt64(qrPageend)
		if qPageend != "" {

			if err := r.SetQueryParam("pageend", qPageend); err != nil {
				return err
			}
		}
	}

	if o.Pagestart != nil {

		// query param pagestart
		var qrPagestart int64

		if o.Pagestart != nil {
			qrPagestart = *o.Pagestart
		}
		qPagestart := swag.FormatInt64(qrPagestart)
		if qPagestart != "" {

			if err := r.SetQueryParam("pagestart", qPagestart); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
