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

// WiFiDetail wi fi detail
//
// swagger:model WiFiDetail
type WiFiDetail struct {

	// bcninterval
	// Read Only: true
	Bcninterval int32 `json:"bcninterval,omitempty"`

	// best cluster wi g l e qo s
	// Read Only: true
	BestClusterWiGLEQoS int32 `json:"bestClusterWiGLEQoS,omitempty"`

	// channel
	// Read Only: true
	Channel int32 `json:"channel,omitempty"`

	// comment
	// Read Only: true
	Comment string `json:"comment,omitempty"`

	// dhcp
	// Read Only: true
	Dhcp string `json:"dhcp,omitempty"`

	// encryption
	// Read Only: true
	Encryption string `json:"encryption,omitempty"`

	// first seen
	// Read Only: true
	// Format: date-time
	FirstSeen strfmt.DateTime `json:"firstSeen,omitempty"`

	// freenet
	// Read Only: true
	Freenet string `json:"freenet,omitempty"`

	// last seen
	// Read Only: true
	// Format: date-time
	LastSeen strfmt.DateTime `json:"lastSeen,omitempty"`

	// last update
	// Read Only: true
	// Format: date-time
	LastUpdate strfmt.DateTime `json:"lastUpdate,omitempty"`

	// location clusters
	LocationClusters []*WiFiLocationCluster `json:"locationClusters"`

	// name
	// Read Only: true
	Name string `json:"name,omitempty"`

	// network Id
	// Required: true
	// Read Only: true
	NetworkID string `json:"networkId"`

	// paynet
	// Read Only: true
	Paynet string `json:"paynet,omitempty"`

	// street address
	// Read Only: true
	StreetAddress *StreetAddress `json:"streetAddress,omitempty"`

	// trilaterated latitude
	// Read Only: true
	TrilateratedLatitude float64 `json:"trilateratedLatitude,omitempty"`

	// trilaterated longitude
	// Read Only: true
	TrilateratedLongitude float64 `json:"trilateratedLongitude,omitempty"`

	// type
	// Read Only: true
	Type string `json:"type,omitempty"`
}

// Validate validates this wi fi detail
func (m *WiFiDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirstSeen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastSeen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreetAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WiFiDetail) validateFirstSeen(formats strfmt.Registry) error {
	if swag.IsZero(m.FirstSeen) { // not required
		return nil
	}

	if err := validate.FormatOf("firstSeen", "body", "date-time", m.FirstSeen.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WiFiDetail) validateLastSeen(formats strfmt.Registry) error {
	if swag.IsZero(m.LastSeen) { // not required
		return nil
	}

	if err := validate.FormatOf("lastSeen", "body", "date-time", m.LastSeen.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WiFiDetail) validateLastUpdate(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdate) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdate", "body", "date-time", m.LastUpdate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WiFiDetail) validateLocationClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.LocationClusters) { // not required
		return nil
	}

	for i := 0; i < len(m.LocationClusters); i++ {
		if swag.IsZero(m.LocationClusters[i]) { // not required
			continue
		}

		if m.LocationClusters[i] != nil {
			if err := m.LocationClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("locationClusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("locationClusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WiFiDetail) validateNetworkID(formats strfmt.Registry) error {

	if err := validate.RequiredString("networkId", "body", m.NetworkID); err != nil {
		return err
	}

	return nil
}

func (m *WiFiDetail) validateStreetAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.StreetAddress) { // not required
		return nil
	}

	if m.StreetAddress != nil {
		if err := m.StreetAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("streetAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("streetAddress")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this wi fi detail based on the context it is used
func (m *WiFiDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBcninterval(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBestClusterWiGLEQoS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChannel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDhcp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEncryption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirstSeen(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFreenet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastSeen(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocationClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePaynet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStreetAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrilateratedLatitude(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrilateratedLongitude(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WiFiDetail) contextValidateBcninterval(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "bcninterval", "body", int32(m.Bcninterval)); err != nil {
		return err
	}

	return nil
}

func (m *WiFiDetail) contextValidateBestClusterWiGLEQoS(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "bestClusterWiGLEQoS", "body", int32(m.BestClusterWiGLEQoS)); err != nil {
		return err
	}

	return nil
}

func (m *WiFiDetail) contextValidateChannel(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "channel", "body", int32(m.Channel)); err != nil {
		return err
	}

	return nil
}

func (m *WiFiDetail) contextValidateComment(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "comment", "body", string(m.Comment)); err != nil {
		return err
	}

	return nil
}

func (m *WiFiDetail) contextValidateDhcp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dhcp", "body", string(m.Dhcp)); err != nil {
		return err
	}

	return nil
}

func (m *WiFiDetail) contextValidateEncryption(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "encryption", "body", string(m.Encryption)); err != nil {
		return err
	}

	return nil
}

func (m *WiFiDetail) contextValidateFirstSeen(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "firstSeen", "body", strfmt.DateTime(m.FirstSeen)); err != nil {
		return err
	}

	return nil
}

func (m *WiFiDetail) contextValidateFreenet(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "freenet", "body", string(m.Freenet)); err != nil {
		return err
	}

	return nil
}

func (m *WiFiDetail) contextValidateLastSeen(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastSeen", "body", strfmt.DateTime(m.LastSeen)); err != nil {
		return err
	}

	return nil
}

func (m *WiFiDetail) contextValidateLastUpdate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastUpdate", "body", strfmt.DateTime(m.LastUpdate)); err != nil {
		return err
	}

	return nil
}

func (m *WiFiDetail) contextValidateLocationClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LocationClusters); i++ {

		if m.LocationClusters[i] != nil {
			if err := m.LocationClusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("locationClusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("locationClusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WiFiDetail) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *WiFiDetail) contextValidateNetworkID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "networkId", "body", string(m.NetworkID)); err != nil {
		return err
	}

	return nil
}

func (m *WiFiDetail) contextValidatePaynet(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "paynet", "body", string(m.Paynet)); err != nil {
		return err
	}

	return nil
}

func (m *WiFiDetail) contextValidateStreetAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.StreetAddress != nil {
		if err := m.StreetAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("streetAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("streetAddress")
			}
			return err
		}
	}

	return nil
}

func (m *WiFiDetail) contextValidateTrilateratedLatitude(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "trilateratedLatitude", "body", float64(m.TrilateratedLatitude)); err != nil {
		return err
	}

	return nil
}

func (m *WiFiDetail) contextValidateTrilateratedLongitude(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "trilateratedLongitude", "body", float64(m.TrilateratedLongitude)); err != nil {
		return err
	}

	return nil
}

func (m *WiFiDetail) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WiFiDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WiFiDetail) UnmarshalBinary(b []byte) error {
	var res WiFiDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
