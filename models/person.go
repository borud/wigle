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

// Person person
//
// swagger:model Person
type Person struct {

	// donate
	Donate string `json:"donate,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// flags
	Flags int64 `json:"flags,omitempty"`

	// joindate
	// Format: date-time
	Joindate strfmt.DateTime `json:"joindate,omitempty"`

	// lastlogin
	// Format: date-time
	Lastlogin strfmt.DateTime `json:"lastlogin,omitempty"`

	// session
	Session string `json:"session,omitempty"`

	// success
	Success string `json:"success,omitempty"`

	// userid
	Userid string `json:"userid,omitempty"`
}

// Validate validates this person
func (m *Person) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJoindate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastlogin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Person) validateJoindate(formats strfmt.Registry) error {
	if swag.IsZero(m.Joindate) { // not required
		return nil
	}

	if err := validate.FormatOf("joindate", "body", "date-time", m.Joindate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Person) validateLastlogin(formats strfmt.Registry) error {
	if swag.IsZero(m.Lastlogin) { // not required
		return nil
	}

	if err := validate.FormatOf("lastlogin", "body", "date-time", m.Lastlogin.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this person based on context it is used
func (m *Person) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Person) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Person) UnmarshalBinary(b []byte) error {
	var res Person
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
