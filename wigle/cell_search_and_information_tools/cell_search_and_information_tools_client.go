// Code generated by go-swagger; DO NOT EDIT.

package cell_search_and_information_tools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cell search and information tools API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cell search and information tools API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	MccMnc(params *MccMncParams, opts ...ClientOption) (*MccMncOK, error)

	Search1(params *Search1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Search1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
MccMnc gets m c c and m n c codes for cellular networks

Get metadata for cell networks - optionally filter by country and network codes
*/
func (a *Client) MccMnc(params *MccMncParams, opts ...ClientOption) (*MccMncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMccMncParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "mccMnc",
		Method:             "GET",
		PathPattern:        "/api/v2/cell/mccMnc",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MccMncReader{formats: a.formats},
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
	success, ok := result.(*MccMncOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for mccMnc: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Search1 searches the wi g l e cell database

Query the WiGLE database for paginated results based on multiple criteria. API and session authentication default to a page size of 100 results/page. COMMAPI defaults to a page size of 25 with a maximum of 1000 results per return. Number of daily queries allowed per user are throttled based on history and participation.  Search endpoints are the only feature included in COMMAPI subscriptions at this time.<br><br>To get next page of results put the previous request's searchAfter value as a parameter in the new request's searchAfter.
*/
func (a *Client) Search1(params *Search1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Search1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearch1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "search_1",
		Method:             "GET",
		PathPattern:        "/api/v2/cell/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Search1Reader{formats: a.formats},
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
	success, ok := result.(*Search1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for search_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
