// Code generated by go-swagger; DO NOT EDIT.

package bluetooth_search_and_information_tools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new bluetooth search and information tools API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bluetooth search and information tools API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Detail(params *DetailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailOK, error)

	Search(params *SearchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Detail gets details and observation records for a single network

Provide unique information for a Bluetooth network. API and session authentication default to a page size of 100 results/page. Number of daily queries allowed per user are throttled based on history and participation. Detail endpoints are NOT included in COMMAPI subscriptions at this time.
*/
func (a *Client) Detail(params *DetailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "detail",
		Method:             "GET",
		PathPattern:        "/api/v2/bluetooth/detail",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DetailReader{formats: a.formats},
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
	success, ok := result.(*DetailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for detail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Search searches the wi g l e bluetooth database

Query the WiGLE database for paginated results based on multiple criteria. API and session authentication default to a page size of 100 results/page. COMMAPI defaults to a page size of 25 with a maximum of 1000 results per return. Number of daily queries allowed per user are throttled based on history and participation. Search endpoints are the only feature included in COMMAPI subscriptions at this time.<br><br>To get next page of results put the previous request's searchAfter value as a parameter in the new request's searchAfter.
*/
func (a *Client) Search(params *SearchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "search",
		Method:             "GET",
		PathPattern:        "/api/v2/bluetooth/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchReader{formats: a.formats},
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
	success, ok := result.(*SearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
