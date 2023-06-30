// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserStatsResponse user stats response
//
// swagger:model UserStatsResponse
type UserStatsResponse struct {

	// image badge Url
	ImageBadgeURL string `json:"imageBadgeUrl,omitempty"`

	// month rank
	MonthRank int64 `json:"monthRank,omitempty"`

	// rank
	Rank int64 `json:"rank,omitempty"`

	// statistics
	Statistics *UserStandings `json:"statistics,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// user
	User string `json:"user,omitempty"`
}

// Validate validates this user stats response
func (m *UserStatsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatistics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserStatsResponse) validateStatistics(formats strfmt.Registry) error {
	if swag.IsZero(m.Statistics) { // not required
		return nil
	}

	if m.Statistics != nil {
		if err := m.Statistics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statistics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statistics")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this user stats response based on the context it is used
func (m *UserStatsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatistics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserStatsResponse) contextValidateStatistics(ctx context.Context, formats strfmt.Registry) error {

	if m.Statistics != nil {
		if err := m.Statistics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statistics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statistics")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserStatsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserStatsResponse) UnmarshalBinary(b []byte) error {
	var res UserStatsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
