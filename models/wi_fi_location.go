// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WiFiLocation wi fi location
//
// swagger:model WiFiLocation
type WiFiLocation struct {

	// accuracy
	Accuracy float32 `json:"accuracy,omitempty"`

	// alt
	Alt int32 `json:"alt,omitempty"`

	// channel
	Channel int32 `json:"channel,omitempty"`

	// encryption value
	EncryptionValue string `json:"encryptionValue,omitempty"`

	// lastupdt
	// Format: date-time
	Lastupdt strfmt.DateTime `json:"lastupdt,omitempty"`

	// latitude
	Latitude float64 `json:"latitude,omitempty"`

	// longitude
	Longitude float64 `json:"longitude,omitempty"`

	// month
	Month string `json:"month,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// net Id
	NetID string `json:"netId,omitempty"`

	// noise
	Noise float32 `json:"noise,omitempty"`

	// signal
	Signal float32 `json:"signal,omitempty"`

	// snr
	Snr float32 `json:"snr,omitempty"`

	// ssid
	Ssid string `json:"ssid,omitempty"`

	// time
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`

	// wep
	Wep string `json:"wep,omitempty"`
}

// Validate validates this wi fi location
func (m *WiFiLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastupdt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WiFiLocation) validateLastupdt(formats strfmt.Registry) error {
	if swag.IsZero(m.Lastupdt) { // not required
		return nil
	}

	if err := validate.FormatOf("lastupdt", "body", "date-time", m.Lastupdt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WiFiLocation) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this wi fi location based on context it is used
func (m *WiFiLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WiFiLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WiFiLocation) UnmarshalBinary(b []byte) error {
	var res WiFiLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
