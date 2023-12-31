// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GroupMember group member
//
// swagger:model GroupMember
type GroupMember struct {

	// discovered
	Discovered int64 `json:"discovered,omitempty"`

	// first transid
	FirstTransid string `json:"firstTransid,omitempty"`

	// gen disc
	GenDisc int64 `json:"genDisc,omitempty"`

	// group Id
	GroupID string `json:"groupId,omitempty"`

	// last transid
	LastTransid string `json:"lastTransid,omitempty"`

	// month count
	MonthCount int64 `json:"monthCount,omitempty"`

	// prev month count
	PrevMonthCount int64 `json:"prevMonthCount,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this group member
func (m *GroupMember) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this group member based on context it is used
func (m *GroupMember) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GroupMember) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupMember) UnmarshalBinary(b []byte) error {
	var res GroupMember
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
