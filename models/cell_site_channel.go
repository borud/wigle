// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CellSiteChannel cell site channel
//
// swagger:model CellSiteChannel
type CellSiteChannel struct {

	// channel
	Channel int64 `json:"channel,omitempty"`

	// latitude
	Latitude float64 `json:"latitude,omitempty"`

	// longitude
	Longitude float64 `json:"longitude,omitempty"`

	// qos
	Qos int32 `json:"qos,omitempty"`
}

// Validate validates this cell site channel
func (m *CellSiteChannel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cell site channel based on context it is used
func (m *CellSiteChannel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CellSiteChannel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CellSiteChannel) UnmarshalBinary(b []byte) error {
	var res CellSiteChannel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
