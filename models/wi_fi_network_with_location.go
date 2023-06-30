// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WiFiNetworkWithLocation wi fi network with location
//
// swagger:model WiFiNetworkWithLocation
type WiFiNetworkWithLocation struct {

	// bcninterval
	Bcninterval int32 `json:"bcninterval,omitempty"`

	// channel
	Channel int32 `json:"channel,omitempty"`

	// city
	City string `json:"city,omitempty"`

	// comment
	Comment string `json:"comment,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// dhcp
	Dhcp string `json:"dhcp,omitempty"`

	// encryption
	Encryption string `json:"encryption,omitempty"`

	// firsttime
	// Format: date-time
	Firsttime strfmt.DateTime `json:"firsttime,omitempty"`

	// freenet
	Freenet string `json:"freenet,omitempty"`

	// housenumber
	Housenumber string `json:"housenumber,omitempty"`

	// lasttime
	// Format: date-time
	Lasttime strfmt.DateTime `json:"lasttime,omitempty"`

	// lastupdt
	// Format: date-time
	Lastupdt strfmt.DateTime `json:"lastupdt,omitempty"`

	// location data
	LocationData []*WiFiLocation `json:"locationData"`

	// name
	Name string `json:"name,omitempty"`

	// netid
	Netid string `json:"netid,omitempty"`

	// paynet
	Paynet string `json:"paynet,omitempty"`

	// postalcode
	Postalcode string `json:"postalcode,omitempty"`

	// qos
	Qos int32 `json:"qos,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// road
	Road string `json:"road,omitempty"`

	// ssid
	Ssid string `json:"ssid,omitempty"`

	// transid
	Transid string `json:"transid,omitempty"`

	// trilat
	Trilat float64 `json:"trilat,omitempty"`

	// trilong
	Trilong float64 `json:"trilong,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// userfound
	Userfound bool `json:"userfound,omitempty"`

	// wep
	Wep string `json:"wep,omitempty"`
}

// Validate validates this wi fi network with location
func (m *WiFiNetworkWithLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirsttime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLasttime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastupdt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WiFiNetworkWithLocation) validateFirsttime(formats strfmt.Registry) error {
	if swag.IsZero(m.Firsttime) { // not required
		return nil
	}

	if err := validate.FormatOf("firsttime", "body", "date-time", m.Firsttime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WiFiNetworkWithLocation) validateLasttime(formats strfmt.Registry) error {
	if swag.IsZero(m.Lasttime) { // not required
		return nil
	}

	if err := validate.FormatOf("lasttime", "body", "date-time", m.Lasttime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WiFiNetworkWithLocation) validateLastupdt(formats strfmt.Registry) error {
	if swag.IsZero(m.Lastupdt) { // not required
		return nil
	}

	if err := validate.FormatOf("lastupdt", "body", "date-time", m.Lastupdt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WiFiNetworkWithLocation) validateLocationData(formats strfmt.Registry) error {
	if swag.IsZero(m.LocationData) { // not required
		return nil
	}

	for i := 0; i < len(m.LocationData); i++ {
		if swag.IsZero(m.LocationData[i]) { // not required
			continue
		}

		if m.LocationData[i] != nil {
			if err := m.LocationData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("locationData" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("locationData" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this wi fi network with location based on the context it is used
func (m *WiFiNetworkWithLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocationData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WiFiNetworkWithLocation) contextValidateLocationData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LocationData); i++ {

		if m.LocationData[i] != nil {
			if err := m.LocationData[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("locationData" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("locationData" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WiFiNetworkWithLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WiFiNetworkWithLocation) UnmarshalBinary(b []byte) error {
	var res WiFiNetworkWithLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}