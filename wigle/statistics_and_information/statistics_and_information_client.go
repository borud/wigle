// Code generated by go-swagger; DO NOT EDIT.

package statistics_and_information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new statistics and information API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for statistics and information API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Countries(params *CountriesParams, opts ...ClientOption) (*CountriesOK, error)

	CountryRegion(params *CountryRegionParams, opts ...ClientOption) (*CountryRegionOK, error)

	GeneralStats(params *GeneralStatsParams, opts ...ClientOption) error

	GroupStats(params *GroupStatsParams, opts ...ClientOption) (*GroupStatsOK, error)

	SiteStats(params *SiteStatsParams, opts ...ClientOption) (*SiteStatsOK, error)

	Stats(params *StatsParams, opts ...ClientOption) error

	UserStatistics(params *UserStatisticsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserStatisticsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Countries gets statistics organized by country

Fetches all countries and basic stats
*/
func (a *Client) Countries(params *CountriesParams, opts ...ClientOption) (*CountriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCountriesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "countries",
		Method:             "GET",
		PathPattern:        "/api/v2/stats/countries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CountriesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CountriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for countries: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CountryRegion gets statistics for a specified country organized by region postal code and encryption
*/
func (a *Client) CountryRegion(params *CountryRegionParams, opts ...ClientOption) (*CountryRegionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCountryRegionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "countryRegion",
		Method:             "GET",
		PathPattern:        "/api/v2/stats/regions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CountryRegionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CountryRegionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for countryRegion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GeneralStats gets a named map of general upload statistics

Information condensed from site uploads about most popular OUIs, Manufacturers, and
*/
func (a *Client) GeneralStats(params *GeneralStatsParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGeneralStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "generalStats",
		Method:             "GET",
		PathPattern:        "/api/v2/stats/general",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GeneralStatsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GroupStats gets group standings

Fetches a list of all teams. Authenticated users receive additional group info for the group they administer.
*/
func (a *Client) GroupStats(params *GroupStatsParams, opts ...ClientOption) (*GroupStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGroupStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "groupStats",
		Method:             "GET",
		PathPattern:        "/api/v2/stats/group",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GroupStatsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GroupStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for groupStats: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SiteStats gets a named map of site level statistics

A big hash of short-named statistics used in providing site-wide information.
*/
func (a *Client) SiteStats(params *SiteStatsParams, opts ...ClientOption) (*SiteStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSiteStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "siteStats",
		Method:             "GET",
		PathPattern:        "/api/v2/stats/site",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SiteStatsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SiteStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for siteStats: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Stats gets user standings
*/
func (a *Client) Stats(params *StatsParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "stats",
		Method:             "GET",
		PathPattern:        "/api/v2/stats/standings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StatsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
UserStatistics gets user statistics

Get statistics and badge image for the authenticated user
*/
func (a *Client) UserStatistics(params *UserStatisticsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserStatisticsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserStatisticsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "userStatistics",
		Method:             "GET",
		PathPattern:        "/api/v2/stats/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserStatisticsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserStatisticsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for userStatistics: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
