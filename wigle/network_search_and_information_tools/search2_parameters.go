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

// NewSearch2Params creates a new Search2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearch2Params() *Search2Params {
	return &Search2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearch2ParamsWithTimeout creates a new Search2Params object
// with the ability to set a timeout on a request.
func NewSearch2ParamsWithTimeout(timeout time.Duration) *Search2Params {
	return &Search2Params{
		timeout: timeout,
	}
}

// NewSearch2ParamsWithContext creates a new Search2Params object
// with the ability to set a context for a request.
func NewSearch2ParamsWithContext(ctx context.Context) *Search2Params {
	return &Search2Params{
		Context: ctx,
	}
}

// NewSearch2ParamsWithHTTPClient creates a new Search2Params object
// with the ability to set a custom HTTPClient for a request.
func NewSearch2ParamsWithHTTPClient(client *http.Client) *Search2Params {
	return &Search2Params{
		HTTPClient: client,
	}
}

/*
Search2Params contains all the parameters to send to the API endpoint

	for the search 2 operation.

	Typically these are written to a http.Request.
*/
type Search2Params struct {

	/* City.

	   Street address city
	*/
	City *string

	/* ClosestLat.

	   Latitude to order by closest network (requires closestLong)
	*/
	ClosestLat *float64

	/* ClosestLong.

	   Longitude to order by closest network (requires closestLat)
	*/
	ClosestLong *float64

	/* Country.

	   Street address country
	*/
	Country *string

	/* Encryption.

	   Encryption detected: 'None', 'WEP', 'WPA', 'WPA2', 'WPA3', 'Unknown'. Case insensitive.
	*/
	Encryption *string

	/* EndTransID.

	   Latest transid by which to bound (year-level precision only), format 'yyyyMMdd-00000'
	*/
	EndTransID *string

	/* Firsttime.

	   Filter points by when they were first created (more recent than this value), condensed date/time numeric string format 'yyyyMMdd[hhmm[ss]]'
	*/
	Firsttime *string

	/* Freenet.

	   Include only networks that have been marked as free access.
	*/
	Freenet *bool

	/* HouseNumber.

	   Street address house number
	*/
	HouseNumber *string

	/* Lasttime.

	   Filter points by how recently they've had data submitted (more recent than this value), condensed date/time numeric string format 'yyyyMMdd[hhmm[ss]]'
	*/
	Lasttime *string

	/* Lastupdt.

	   Filter points by how recently they've been updated (more recent than this value), condensed date/time numeric string format 'yyyyMMdd[hhmm[ss]]'
	*/
	Lastupdt *string

	/* Latrange1.

	   Lesser of two latitudes by which to bound the search (specify both)
	*/
	Latrange1 *float64

	/* Latrange2.

	   Greater of two latitudes by which to bound the search (specify both)
	*/
	Latrange2 *float64

	/* Longrange1.

	   Lesser of two longitudes by which to bound the search (specify both)
	*/
	Longrange1 *float64

	/* Longrange2.

	   Greater of two longitudes by which to bound the search (specify both)
	*/
	Longrange2 *float64

	/* MinQoS.

	   Minimum Quality of Signal (0-7).

	   Format: int32
	*/
	MinQoS *int32

	/* Netid.

	   Include only networks matching the string network BSSID, e.g. '0A:2C:EF:3D:25:1B' or '0A:2C:EF'. The first three octets are required.
	*/
	Netid *string

	/* Notmine.

	   Only search for networks first seen by other users
	*/
	Notmine *string

	/* Onlymine.

	   Search only for points first discovered by the current user. Use any string to set, leave unset for general search. Can't be used with COMMAPI auth, since these are points you have locally.

	   Default: "false"
	*/
	Onlymine *string

	/* Paynet.

	   Include only networks that have been marked as for-pay access.
	*/
	Paynet *bool

	/* PostalCode.

	   Street address postal code
	*/
	PostalCode *string

	/* Region.

	   Street address region
	*/
	Region *string

	/* ResultsPerPage.

	   How many results to return per request. Defaults to 25 for COMMAPI, 100 for site. Bounded at 1000 for COMMAPI, 100 for site.

	   Format: int64
	*/
	ResultsPerPage *int64

	/* Road.

	   Street address road
	*/
	Road *string

	/* SearchAfter.

	   Put in the previous page's searchAfter result to get the next page. Use this instead of 'first'
	*/
	SearchAfter *string

	/* Ssid.

	   Include only networks exactly matching the string network name.
	*/
	Ssid *string

	/* Ssidlike.

	   Include only networks matching the string network name, allowing wildcards '%' (any string) and '_' (any character).
	*/
	Ssidlike *string

	/* StartTransID.

	   Earliest transid by which to bound (year-level precision only), format 'yyyyMMdd-00000'
	*/
	StartTransID *string

	/* Variance.

	   How tightly to bound queries against the provided latitude/longitude box. Value must be between 0.001 and 0.2. Intended for use with non-exact decimals and geocoded bounds.
	*/
	Variance *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Search2Params) WithDefaults() *Search2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Search2Params) SetDefaults() {
	var (
		freenetDefault = bool(false)

		onlymineDefault = string("false")

		paynetDefault = bool(false)
	)

	val := Search2Params{
		Freenet:  &freenetDefault,
		Onlymine: &onlymineDefault,
		Paynet:   &paynetDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the search 2 params
func (o *Search2Params) WithTimeout(timeout time.Duration) *Search2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search 2 params
func (o *Search2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search 2 params
func (o *Search2Params) WithContext(ctx context.Context) *Search2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search 2 params
func (o *Search2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search 2 params
func (o *Search2Params) WithHTTPClient(client *http.Client) *Search2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search 2 params
func (o *Search2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCity adds the city to the search 2 params
func (o *Search2Params) WithCity(city *string) *Search2Params {
	o.SetCity(city)
	return o
}

// SetCity adds the city to the search 2 params
func (o *Search2Params) SetCity(city *string) {
	o.City = city
}

// WithClosestLat adds the closestLat to the search 2 params
func (o *Search2Params) WithClosestLat(closestLat *float64) *Search2Params {
	o.SetClosestLat(closestLat)
	return o
}

// SetClosestLat adds the closestLat to the search 2 params
func (o *Search2Params) SetClosestLat(closestLat *float64) {
	o.ClosestLat = closestLat
}

// WithClosestLong adds the closestLong to the search 2 params
func (o *Search2Params) WithClosestLong(closestLong *float64) *Search2Params {
	o.SetClosestLong(closestLong)
	return o
}

// SetClosestLong adds the closestLong to the search 2 params
func (o *Search2Params) SetClosestLong(closestLong *float64) {
	o.ClosestLong = closestLong
}

// WithCountry adds the country to the search 2 params
func (o *Search2Params) WithCountry(country *string) *Search2Params {
	o.SetCountry(country)
	return o
}

// SetCountry adds the country to the search 2 params
func (o *Search2Params) SetCountry(country *string) {
	o.Country = country
}

// WithEncryption adds the encryption to the search 2 params
func (o *Search2Params) WithEncryption(encryption *string) *Search2Params {
	o.SetEncryption(encryption)
	return o
}

// SetEncryption adds the encryption to the search 2 params
func (o *Search2Params) SetEncryption(encryption *string) {
	o.Encryption = encryption
}

// WithEndTransID adds the endTransID to the search 2 params
func (o *Search2Params) WithEndTransID(endTransID *string) *Search2Params {
	o.SetEndTransID(endTransID)
	return o
}

// SetEndTransID adds the endTransId to the search 2 params
func (o *Search2Params) SetEndTransID(endTransID *string) {
	o.EndTransID = endTransID
}

// WithFirsttime adds the firsttime to the search 2 params
func (o *Search2Params) WithFirsttime(firsttime *string) *Search2Params {
	o.SetFirsttime(firsttime)
	return o
}

// SetFirsttime adds the firsttime to the search 2 params
func (o *Search2Params) SetFirsttime(firsttime *string) {
	o.Firsttime = firsttime
}

// WithFreenet adds the freenet to the search 2 params
func (o *Search2Params) WithFreenet(freenet *bool) *Search2Params {
	o.SetFreenet(freenet)
	return o
}

// SetFreenet adds the freenet to the search 2 params
func (o *Search2Params) SetFreenet(freenet *bool) {
	o.Freenet = freenet
}

// WithHouseNumber adds the houseNumber to the search 2 params
func (o *Search2Params) WithHouseNumber(houseNumber *string) *Search2Params {
	o.SetHouseNumber(houseNumber)
	return o
}

// SetHouseNumber adds the houseNumber to the search 2 params
func (o *Search2Params) SetHouseNumber(houseNumber *string) {
	o.HouseNumber = houseNumber
}

// WithLasttime adds the lasttime to the search 2 params
func (o *Search2Params) WithLasttime(lasttime *string) *Search2Params {
	o.SetLasttime(lasttime)
	return o
}

// SetLasttime adds the lasttime to the search 2 params
func (o *Search2Params) SetLasttime(lasttime *string) {
	o.Lasttime = lasttime
}

// WithLastupdt adds the lastupdt to the search 2 params
func (o *Search2Params) WithLastupdt(lastupdt *string) *Search2Params {
	o.SetLastupdt(lastupdt)
	return o
}

// SetLastupdt adds the lastupdt to the search 2 params
func (o *Search2Params) SetLastupdt(lastupdt *string) {
	o.Lastupdt = lastupdt
}

// WithLatrange1 adds the latrange1 to the search 2 params
func (o *Search2Params) WithLatrange1(latrange1 *float64) *Search2Params {
	o.SetLatrange1(latrange1)
	return o
}

// SetLatrange1 adds the latrange1 to the search 2 params
func (o *Search2Params) SetLatrange1(latrange1 *float64) {
	o.Latrange1 = latrange1
}

// WithLatrange2 adds the latrange2 to the search 2 params
func (o *Search2Params) WithLatrange2(latrange2 *float64) *Search2Params {
	o.SetLatrange2(latrange2)
	return o
}

// SetLatrange2 adds the latrange2 to the search 2 params
func (o *Search2Params) SetLatrange2(latrange2 *float64) {
	o.Latrange2 = latrange2
}

// WithLongrange1 adds the longrange1 to the search 2 params
func (o *Search2Params) WithLongrange1(longrange1 *float64) *Search2Params {
	o.SetLongrange1(longrange1)
	return o
}

// SetLongrange1 adds the longrange1 to the search 2 params
func (o *Search2Params) SetLongrange1(longrange1 *float64) {
	o.Longrange1 = longrange1
}

// WithLongrange2 adds the longrange2 to the search 2 params
func (o *Search2Params) WithLongrange2(longrange2 *float64) *Search2Params {
	o.SetLongrange2(longrange2)
	return o
}

// SetLongrange2 adds the longrange2 to the search 2 params
func (o *Search2Params) SetLongrange2(longrange2 *float64) {
	o.Longrange2 = longrange2
}

// WithMinQoS adds the minQoS to the search 2 params
func (o *Search2Params) WithMinQoS(minQoS *int32) *Search2Params {
	o.SetMinQoS(minQoS)
	return o
}

// SetMinQoS adds the minQoS to the search 2 params
func (o *Search2Params) SetMinQoS(minQoS *int32) {
	o.MinQoS = minQoS
}

// WithNetid adds the netid to the search 2 params
func (o *Search2Params) WithNetid(netid *string) *Search2Params {
	o.SetNetid(netid)
	return o
}

// SetNetid adds the netid to the search 2 params
func (o *Search2Params) SetNetid(netid *string) {
	o.Netid = netid
}

// WithNotmine adds the notmine to the search 2 params
func (o *Search2Params) WithNotmine(notmine *string) *Search2Params {
	o.SetNotmine(notmine)
	return o
}

// SetNotmine adds the notmine to the search 2 params
func (o *Search2Params) SetNotmine(notmine *string) {
	o.Notmine = notmine
}

// WithOnlymine adds the onlymine to the search 2 params
func (o *Search2Params) WithOnlymine(onlymine *string) *Search2Params {
	o.SetOnlymine(onlymine)
	return o
}

// SetOnlymine adds the onlymine to the search 2 params
func (o *Search2Params) SetOnlymine(onlymine *string) {
	o.Onlymine = onlymine
}

// WithPaynet adds the paynet to the search 2 params
func (o *Search2Params) WithPaynet(paynet *bool) *Search2Params {
	o.SetPaynet(paynet)
	return o
}

// SetPaynet adds the paynet to the search 2 params
func (o *Search2Params) SetPaynet(paynet *bool) {
	o.Paynet = paynet
}

// WithPostalCode adds the postalCode to the search 2 params
func (o *Search2Params) WithPostalCode(postalCode *string) *Search2Params {
	o.SetPostalCode(postalCode)
	return o
}

// SetPostalCode adds the postalCode to the search 2 params
func (o *Search2Params) SetPostalCode(postalCode *string) {
	o.PostalCode = postalCode
}

// WithRegion adds the region to the search 2 params
func (o *Search2Params) WithRegion(region *string) *Search2Params {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the search 2 params
func (o *Search2Params) SetRegion(region *string) {
	o.Region = region
}

// WithResultsPerPage adds the resultsPerPage to the search 2 params
func (o *Search2Params) WithResultsPerPage(resultsPerPage *int64) *Search2Params {
	o.SetResultsPerPage(resultsPerPage)
	return o
}

// SetResultsPerPage adds the resultsPerPage to the search 2 params
func (o *Search2Params) SetResultsPerPage(resultsPerPage *int64) {
	o.ResultsPerPage = resultsPerPage
}

// WithRoad adds the road to the search 2 params
func (o *Search2Params) WithRoad(road *string) *Search2Params {
	o.SetRoad(road)
	return o
}

// SetRoad adds the road to the search 2 params
func (o *Search2Params) SetRoad(road *string) {
	o.Road = road
}

// WithSearchAfter adds the searchAfter to the search 2 params
func (o *Search2Params) WithSearchAfter(searchAfter *string) *Search2Params {
	o.SetSearchAfter(searchAfter)
	return o
}

// SetSearchAfter adds the searchAfter to the search 2 params
func (o *Search2Params) SetSearchAfter(searchAfter *string) {
	o.SearchAfter = searchAfter
}

// WithSsid adds the ssid to the search 2 params
func (o *Search2Params) WithSsid(ssid *string) *Search2Params {
	o.SetSsid(ssid)
	return o
}

// SetSsid adds the ssid to the search 2 params
func (o *Search2Params) SetSsid(ssid *string) {
	o.Ssid = ssid
}

// WithSsidlike adds the ssidlike to the search 2 params
func (o *Search2Params) WithSsidlike(ssidlike *string) *Search2Params {
	o.SetSsidlike(ssidlike)
	return o
}

// SetSsidlike adds the ssidlike to the search 2 params
func (o *Search2Params) SetSsidlike(ssidlike *string) {
	o.Ssidlike = ssidlike
}

// WithStartTransID adds the startTransID to the search 2 params
func (o *Search2Params) WithStartTransID(startTransID *string) *Search2Params {
	o.SetStartTransID(startTransID)
	return o
}

// SetStartTransID adds the startTransId to the search 2 params
func (o *Search2Params) SetStartTransID(startTransID *string) {
	o.StartTransID = startTransID
}

// WithVariance adds the variance to the search 2 params
func (o *Search2Params) WithVariance(variance *float64) *Search2Params {
	o.SetVariance(variance)
	return o
}

// SetVariance adds the variance to the search 2 params
func (o *Search2Params) SetVariance(variance *float64) {
	o.Variance = variance
}

// WriteToRequest writes these params to a swagger request
func (o *Search2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.City != nil {

		// query param city
		var qrCity string

		if o.City != nil {
			qrCity = *o.City
		}
		qCity := qrCity
		if qCity != "" {

			if err := r.SetQueryParam("city", qCity); err != nil {
				return err
			}
		}
	}

	if o.ClosestLat != nil {

		// query param closestLat
		var qrClosestLat float64

		if o.ClosestLat != nil {
			qrClosestLat = *o.ClosestLat
		}
		qClosestLat := swag.FormatFloat64(qrClosestLat)
		if qClosestLat != "" {

			if err := r.SetQueryParam("closestLat", qClosestLat); err != nil {
				return err
			}
		}
	}

	if o.ClosestLong != nil {

		// query param closestLong
		var qrClosestLong float64

		if o.ClosestLong != nil {
			qrClosestLong = *o.ClosestLong
		}
		qClosestLong := swag.FormatFloat64(qrClosestLong)
		if qClosestLong != "" {

			if err := r.SetQueryParam("closestLong", qClosestLong); err != nil {
				return err
			}
		}
	}

	if o.Country != nil {

		// query param country
		var qrCountry string

		if o.Country != nil {
			qrCountry = *o.Country
		}
		qCountry := qrCountry
		if qCountry != "" {

			if err := r.SetQueryParam("country", qCountry); err != nil {
				return err
			}
		}
	}

	if o.Encryption != nil {

		// query param encryption
		var qrEncryption string

		if o.Encryption != nil {
			qrEncryption = *o.Encryption
		}
		qEncryption := qrEncryption
		if qEncryption != "" {

			if err := r.SetQueryParam("encryption", qEncryption); err != nil {
				return err
			}
		}
	}

	if o.EndTransID != nil {

		// query param endTransID
		var qrEndTransID string

		if o.EndTransID != nil {
			qrEndTransID = *o.EndTransID
		}
		qEndTransID := qrEndTransID
		if qEndTransID != "" {

			if err := r.SetQueryParam("endTransID", qEndTransID); err != nil {
				return err
			}
		}
	}

	if o.Firsttime != nil {

		// query param firsttime
		var qrFirsttime string

		if o.Firsttime != nil {
			qrFirsttime = *o.Firsttime
		}
		qFirsttime := qrFirsttime
		if qFirsttime != "" {

			if err := r.SetQueryParam("firsttime", qFirsttime); err != nil {
				return err
			}
		}
	}

	if o.Freenet != nil {

		// query param freenet
		var qrFreenet bool

		if o.Freenet != nil {
			qrFreenet = *o.Freenet
		}
		qFreenet := swag.FormatBool(qrFreenet)
		if qFreenet != "" {

			if err := r.SetQueryParam("freenet", qFreenet); err != nil {
				return err
			}
		}
	}

	if o.HouseNumber != nil {

		// query param houseNumber
		var qrHouseNumber string

		if o.HouseNumber != nil {
			qrHouseNumber = *o.HouseNumber
		}
		qHouseNumber := qrHouseNumber
		if qHouseNumber != "" {

			if err := r.SetQueryParam("houseNumber", qHouseNumber); err != nil {
				return err
			}
		}
	}

	if o.Lasttime != nil {

		// query param lasttime
		var qrLasttime string

		if o.Lasttime != nil {
			qrLasttime = *o.Lasttime
		}
		qLasttime := qrLasttime
		if qLasttime != "" {

			if err := r.SetQueryParam("lasttime", qLasttime); err != nil {
				return err
			}
		}
	}

	if o.Lastupdt != nil {

		// query param lastupdt
		var qrLastupdt string

		if o.Lastupdt != nil {
			qrLastupdt = *o.Lastupdt
		}
		qLastupdt := qrLastupdt
		if qLastupdt != "" {

			if err := r.SetQueryParam("lastupdt", qLastupdt); err != nil {
				return err
			}
		}
	}

	if o.Latrange1 != nil {

		// query param latrange1
		var qrLatrange1 float64

		if o.Latrange1 != nil {
			qrLatrange1 = *o.Latrange1
		}
		qLatrange1 := swag.FormatFloat64(qrLatrange1)
		if qLatrange1 != "" {

			if err := r.SetQueryParam("latrange1", qLatrange1); err != nil {
				return err
			}
		}
	}

	if o.Latrange2 != nil {

		// query param latrange2
		var qrLatrange2 float64

		if o.Latrange2 != nil {
			qrLatrange2 = *o.Latrange2
		}
		qLatrange2 := swag.FormatFloat64(qrLatrange2)
		if qLatrange2 != "" {

			if err := r.SetQueryParam("latrange2", qLatrange2); err != nil {
				return err
			}
		}
	}

	if o.Longrange1 != nil {

		// query param longrange1
		var qrLongrange1 float64

		if o.Longrange1 != nil {
			qrLongrange1 = *o.Longrange1
		}
		qLongrange1 := swag.FormatFloat64(qrLongrange1)
		if qLongrange1 != "" {

			if err := r.SetQueryParam("longrange1", qLongrange1); err != nil {
				return err
			}
		}
	}

	if o.Longrange2 != nil {

		// query param longrange2
		var qrLongrange2 float64

		if o.Longrange2 != nil {
			qrLongrange2 = *o.Longrange2
		}
		qLongrange2 := swag.FormatFloat64(qrLongrange2)
		if qLongrange2 != "" {

			if err := r.SetQueryParam("longrange2", qLongrange2); err != nil {
				return err
			}
		}
	}

	if o.MinQoS != nil {

		// query param minQoS
		var qrMinQoS int32

		if o.MinQoS != nil {
			qrMinQoS = *o.MinQoS
		}
		qMinQoS := swag.FormatInt32(qrMinQoS)
		if qMinQoS != "" {

			if err := r.SetQueryParam("minQoS", qMinQoS); err != nil {
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

	if o.Notmine != nil {

		// query param notmine
		var qrNotmine string

		if o.Notmine != nil {
			qrNotmine = *o.Notmine
		}
		qNotmine := qrNotmine
		if qNotmine != "" {

			if err := r.SetQueryParam("notmine", qNotmine); err != nil {
				return err
			}
		}
	}

	if o.Onlymine != nil {

		// query param onlymine
		var qrOnlymine string

		if o.Onlymine != nil {
			qrOnlymine = *o.Onlymine
		}
		qOnlymine := qrOnlymine
		if qOnlymine != "" {

			if err := r.SetQueryParam("onlymine", qOnlymine); err != nil {
				return err
			}
		}
	}

	if o.Paynet != nil {

		// query param paynet
		var qrPaynet bool

		if o.Paynet != nil {
			qrPaynet = *o.Paynet
		}
		qPaynet := swag.FormatBool(qrPaynet)
		if qPaynet != "" {

			if err := r.SetQueryParam("paynet", qPaynet); err != nil {
				return err
			}
		}
	}

	if o.PostalCode != nil {

		// query param postalCode
		var qrPostalCode string

		if o.PostalCode != nil {
			qrPostalCode = *o.PostalCode
		}
		qPostalCode := qrPostalCode
		if qPostalCode != "" {

			if err := r.SetQueryParam("postalCode", qPostalCode); err != nil {
				return err
			}
		}
	}

	if o.Region != nil {

		// query param region
		var qrRegion string

		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {

			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}
	}

	if o.ResultsPerPage != nil {

		// query param resultsPerPage
		var qrResultsPerPage int64

		if o.ResultsPerPage != nil {
			qrResultsPerPage = *o.ResultsPerPage
		}
		qResultsPerPage := swag.FormatInt64(qrResultsPerPage)
		if qResultsPerPage != "" {

			if err := r.SetQueryParam("resultsPerPage", qResultsPerPage); err != nil {
				return err
			}
		}
	}

	if o.Road != nil {

		// query param road
		var qrRoad string

		if o.Road != nil {
			qrRoad = *o.Road
		}
		qRoad := qrRoad
		if qRoad != "" {

			if err := r.SetQueryParam("road", qRoad); err != nil {
				return err
			}
		}
	}

	if o.SearchAfter != nil {

		// query param searchAfter
		var qrSearchAfter string

		if o.SearchAfter != nil {
			qrSearchAfter = *o.SearchAfter
		}
		qSearchAfter := qrSearchAfter
		if qSearchAfter != "" {

			if err := r.SetQueryParam("searchAfter", qSearchAfter); err != nil {
				return err
			}
		}
	}

	if o.Ssid != nil {

		// query param ssid
		var qrSsid string

		if o.Ssid != nil {
			qrSsid = *o.Ssid
		}
		qSsid := qrSsid
		if qSsid != "" {

			if err := r.SetQueryParam("ssid", qSsid); err != nil {
				return err
			}
		}
	}

	if o.Ssidlike != nil {

		// query param ssidlike
		var qrSsidlike string

		if o.Ssidlike != nil {
			qrSsidlike = *o.Ssidlike
		}
		qSsidlike := qrSsidlike
		if qSsidlike != "" {

			if err := r.SetQueryParam("ssidlike", qSsidlike); err != nil {
				return err
			}
		}
	}

	if o.StartTransID != nil {

		// query param startTransID
		var qrStartTransID string

		if o.StartTransID != nil {
			qrStartTransID = *o.StartTransID
		}
		qStartTransID := qrStartTransID
		if qStartTransID != "" {

			if err := r.SetQueryParam("startTransID", qStartTransID); err != nil {
				return err
			}
		}
	}

	if o.Variance != nil {

		// query param variance
		var qrVariance float64

		if o.Variance != nil {
			qrVariance = *o.Variance
		}
		qVariance := swag.FormatFloat64(qrVariance)
		if qVariance != "" {

			if err := r.SetQueryParam("variance", qVariance); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
