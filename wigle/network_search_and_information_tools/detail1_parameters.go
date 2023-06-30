// Code generated by go-swagger; DO NOT EDIT.

package network_search_and_information_tools

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

// NewDetail1Params creates a new Detail1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDetail1Params() *Detail1Params {
	return &Detail1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDetail1ParamsWithTimeout creates a new Detail1Params object
// with the ability to set a timeout on a request.
func NewDetail1ParamsWithTimeout(timeout time.Duration) *Detail1Params {
	return &Detail1Params{
		timeout: timeout,
	}
}

// NewDetail1ParamsWithContext creates a new Detail1Params object
// with the ability to set a context for a request.
func NewDetail1ParamsWithContext(ctx context.Context) *Detail1Params {
	return &Detail1Params{
		Context: ctx,
	}
}

// NewDetail1ParamsWithHTTPClient creates a new Detail1Params object
// with the ability to set a custom HTTPClient for a request.
func NewDetail1ParamsWithHTTPClient(client *http.Client) *Detail1Params {
	return &Detail1Params{
		HTTPClient: client,
	}
}

/*
Detail1Params contains all the parameters to send to the API endpoint

	for the detail 1 operation.

	Typically these are written to a http.Request.
*/
type Detail1Params struct {

	/* Basestation.

	   CDMA Base Station ID

	   Format: int64
	*/
	Basestation *int64

	/* Cid.

	   GSM/LTE/WCDMA/5G NR Cell ID/NIR

	   Format: int64
	*/
	Cid *int64

	/* Lac.

	   GSM/LTE/WCDMA/5G NR Location Area Code

	   Format: int64
	*/
	Lac *int64

	/* Netid.

	   The WiFi Network BSSID to search
	*/
	Netid *string

	/* Network.

	   CDMA Network ID

	   Format: int64
	*/
	Network *int64

	/* Operator.

	   GSM/LTE/WCDMA/5G NR Operator ID

	   Format: int64
	*/
	Operator *int64

	/* System.

	   CDMA System ID

	   Format: int64
	*/
	System *int64

	/* Type.

	   Network Type: CDMA/GSM/LTE/WCDMA/NR/WIFI
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the detail 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Detail1Params) WithDefaults() *Detail1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the detail 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Detail1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the detail 1 params
func (o *Detail1Params) WithTimeout(timeout time.Duration) *Detail1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the detail 1 params
func (o *Detail1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the detail 1 params
func (o *Detail1Params) WithContext(ctx context.Context) *Detail1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the detail 1 params
func (o *Detail1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the detail 1 params
func (o *Detail1Params) WithHTTPClient(client *http.Client) *Detail1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the detail 1 params
func (o *Detail1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBasestation adds the basestation to the detail 1 params
func (o *Detail1Params) WithBasestation(basestation *int64) *Detail1Params {
	o.SetBasestation(basestation)
	return o
}

// SetBasestation adds the basestation to the detail 1 params
func (o *Detail1Params) SetBasestation(basestation *int64) {
	o.Basestation = basestation
}

// WithCid adds the cid to the detail 1 params
func (o *Detail1Params) WithCid(cid *int64) *Detail1Params {
	o.SetCid(cid)
	return o
}

// SetCid adds the cid to the detail 1 params
func (o *Detail1Params) SetCid(cid *int64) {
	o.Cid = cid
}

// WithLac adds the lac to the detail 1 params
func (o *Detail1Params) WithLac(lac *int64) *Detail1Params {
	o.SetLac(lac)
	return o
}

// SetLac adds the lac to the detail 1 params
func (o *Detail1Params) SetLac(lac *int64) {
	o.Lac = lac
}

// WithNetid adds the netid to the detail 1 params
func (o *Detail1Params) WithNetid(netid *string) *Detail1Params {
	o.SetNetid(netid)
	return o
}

// SetNetid adds the netid to the detail 1 params
func (o *Detail1Params) SetNetid(netid *string) {
	o.Netid = netid
}

// WithNetwork adds the network to the detail 1 params
func (o *Detail1Params) WithNetwork(network *int64) *Detail1Params {
	o.SetNetwork(network)
	return o
}

// SetNetwork adds the network to the detail 1 params
func (o *Detail1Params) SetNetwork(network *int64) {
	o.Network = network
}

// WithOperator adds the operator to the detail 1 params
func (o *Detail1Params) WithOperator(operator *int64) *Detail1Params {
	o.SetOperator(operator)
	return o
}

// SetOperator adds the operator to the detail 1 params
func (o *Detail1Params) SetOperator(operator *int64) {
	o.Operator = operator
}

// WithSystem adds the system to the detail 1 params
func (o *Detail1Params) WithSystem(system *int64) *Detail1Params {
	o.SetSystem(system)
	return o
}

// SetSystem adds the system to the detail 1 params
func (o *Detail1Params) SetSystem(system *int64) {
	o.System = system
}

// WithType adds the typeVar to the detail 1 params
func (o *Detail1Params) WithType(typeVar *string) *Detail1Params {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the detail 1 params
func (o *Detail1Params) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *Detail1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Basestation != nil {

		// query param basestation
		var qrBasestation int64

		if o.Basestation != nil {
			qrBasestation = *o.Basestation
		}
		qBasestation := swag.FormatInt64(qrBasestation)
		if qBasestation != "" {

			if err := r.SetQueryParam("basestation", qBasestation); err != nil {
				return err
			}
		}
	}

	if o.Cid != nil {

		// query param cid
		var qrCid int64

		if o.Cid != nil {
			qrCid = *o.Cid
		}
		qCid := swag.FormatInt64(qrCid)
		if qCid != "" {

			if err := r.SetQueryParam("cid", qCid); err != nil {
				return err
			}
		}
	}

	if o.Lac != nil {

		// query param lac
		var qrLac int64

		if o.Lac != nil {
			qrLac = *o.Lac
		}
		qLac := swag.FormatInt64(qrLac)
		if qLac != "" {

			if err := r.SetQueryParam("lac", qLac); err != nil {
				return err
			}
		}
	}

	if o.Netid != nil {

		// query param netid
		var qrNetid string

		if o.Netid != nil {
			qrNetid = *o.Netid
		}
		qNetid := qrNetid
		if qNetid != "" {

			if err := r.SetQueryParam("netid", qNetid); err != nil {
				return err
			}
		}
	}

	if o.Network != nil {

		// query param network
		var qrNetwork int64

		if o.Network != nil {
			qrNetwork = *o.Network
		}
		qNetwork := swag.FormatInt64(qrNetwork)
		if qNetwork != "" {

			if err := r.SetQueryParam("network", qNetwork); err != nil {
				return err
			}
		}
	}

	if o.Operator != nil {

		// query param operator
		var qrOperator int64

		if o.Operator != nil {
			qrOperator = *o.Operator
		}
		qOperator := swag.FormatInt64(qrOperator)
		if qOperator != "" {

			if err := r.SetQueryParam("operator", qOperator); err != nil {
				return err
			}
		}
	}

	if o.System != nil {

		// query param system
		var qrSystem int64

		if o.System != nil {
			qrSystem = *o.System
		}
		qSystem := swag.FormatInt64(qrSystem)
		if qSystem != "" {

			if err := r.SetQueryParam("system", qSystem); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}