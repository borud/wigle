// Code generated by go-swagger; DO NOT EDIT.

package network_search_and_information_tools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new network search and information tools API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for network search and information tools API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Comment(params *CommentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CommentOK, error)

	Detail1(params *Detail1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Detail1OK, error)

	Geocode(params *GeocodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GeocodeOK, error)

	Search2(params *Search2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Search2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Comment adds a comment to a network

provide custom information regarding a single network
*/
func (a *Client) Comment(params *CommentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CommentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCommentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "comment",
		Method:             "POST",
		PathPattern:        "/api/v2/network/comment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CommentReader{formats: a.formats},
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
	success, ok := result.(*CommentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for comment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Detail1 gets details and observation records for a single wifi or cell network

Provide unique information for a WiFi or cell network to request detailed information. Providing a netId value searches WiFi, operator searches GSM, and system searches CDMA. Detail endpoints are NOT included in COMMAPI subscriptions at this time.
*/
func (a *Client) Detail1(params *Detail1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Detail1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetail1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "detail_1",
		Method:             "GET",
		PathPattern:        "/api/v2/network/detail",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Detail1Reader{formats: a.formats},
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
	success, ok := result.(*Detail1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for detail_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Geocode gets coordinates for an address for use in searching

Relies on OpenStreetMap nominatim
*/
func (a *Client) Geocode(params *GeocodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GeocodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGeocodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "geocode",
		Method:             "GET",
		PathPattern:        "/api/v2/network/geocode",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GeocodeReader{formats: a.formats},
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
	success, ok := result.(*GeocodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for geocode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Search2 searches the wi g l e wifi database

Query the WiGLE database for paginated results based on multiple criteria. API and session authentication default to a page size of 100 results/page. COMMAPI defaults to a page size of 25 with a maximum of 1000 results per return. Number of daily queries allowed per user are throttled based on history and participation.  Search endpoints are the only feature included in COMMAPI subscriptions at this time.<br><br>To get next page of results put the previous request's searchAfter value as a parameter in the new request's searchAfter.
*/
func (a *Client) Search2(params *Search2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Search2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearch2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "search_2",
		Method:             "GET",
		PathPattern:        "/api/v2/network/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Search2Reader{formats: a.formats},
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
	success, ok := result.(*Search2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for search_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}