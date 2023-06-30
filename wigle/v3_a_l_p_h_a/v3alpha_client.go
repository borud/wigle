// Code generated by go-swagger; DO NOT EDIT.

package v3_a_l_p_h_a

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new v3 a l p h a API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v3 a l p h a API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Bt(params *BtParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BtOK, error)

	CdmaCell(params *CdmaCellParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CdmaCellOK, error)

	Cell(params *CellParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CellOK, error)

	CellChannel(params *CellChannelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CellChannelOK, error)

	Wifi(params *WifiParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WifiOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Bt requests detail for a bluetooth network

Location and detail properties for bluetooth networks. Detail endpoints are NOT included in COMMAPI subscriptions at this time.  Daily query limit subject to user DETAIL limit.
*/
func (a *Client) Bt(params *BtParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BtOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBtParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "bt",
		Method:             "GET",
		PathPattern:        "/api/v3/detail/bt/{btNetworkId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BtReader{formats: a.formats},
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
	success, ok := result.(*BtOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for bt: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CdmaCell requests detail for a c d m a network

Location and detail properties for CDMA networks. Detail endpoints are NOT included in COMMAPI subscriptions at this time.  Daily query limit subject to user DETAIL limit.
*/
func (a *Client) CdmaCell(params *CdmaCellParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CdmaCellOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCdmaCellParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cdmaCell",
		Method:             "GET",
		PathPattern:        "/api/v3/detail/cell/CDMA/{sid}/{nid}/{bsid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CdmaCellReader{formats: a.formats},
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
	success, ok := result.(*CdmaCellOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cdmaCell: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Cell requests detail for a g s m l t e w c d m a or 5 g n r network

Location and detail properties for non-CDMA cell types. Note that the type parameter is sensitive to current type in the WiGLE database - this means that LTE and WCDMA networks may have been reported as GSM in some packages, a fallback query to GSM is recommended if the LTE/WCDMA network appears to be absent. Detail endpoints are NOT included in COMMAPI subscriptions at this time.  Daily query limit subject to user DETAIL limit.
*/
func (a *Client) Cell(params *CellParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CellOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCellParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cell",
		Method:             "GET",
		PathPattern:        "/api/v3/detail/cell/{type}/{operator}/{lac}/{cid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CellReader{formats: a.formats},
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
	success, ok := result.(*CellOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cell: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CellChannel requests list of channels for g s m l t e w c d m a or 5 g n r transmitters in specified region

List of known channels and QoS values for any of the non-CDMA cell types in a bounding box. Note that the type parameter is sensitive to current type in the WiGLE database - this means that LTE and WCDMA networks may have been reported as GSM in some packages, a fallback query to GSM is recommended if the LTE/WCDMA network appears to be absent. Channel endpoints are NOT included in COMMAPI subscriptions at this time. Region cannot be large than 99 square miles (~256 sq km). Daily query limit subject to user DETAIL limit. Designed to support cell-site simulator detection.
*/
func (a *Client) CellChannel(params *CellChannelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CellChannelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCellChannelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cellChannel",
		Method:             "GET",
		PathPattern:        "/api/v3/cellChannel/{type}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CellChannelReader{formats: a.formats},
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
	success, ok := result.(*CellChannelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cellChannel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Wifi requests detail for a wi fi network

Location and detail properties for WiFi networks. Detail endpoints are NOT included in COMMAPI subscriptions at this time.  Daily query limit subject to user DETAIL limit.
*/
func (a *Client) Wifi(params *WifiParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WifiOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWifiParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "wifi",
		Method:             "GET",
		PathPattern:        "/api/v3/detail/wifi/{wifiNetworkId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WifiReader{formats: a.formats},
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
	success, ok := result.(*WifiOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for wifi: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}