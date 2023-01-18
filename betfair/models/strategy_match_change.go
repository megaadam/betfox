// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StrategyMatchChange strategy match change
//
// swagger:model StrategyMatchChange
type StrategyMatchChange struct {

	// Matched Backs - matched amounts by distinct matched price on the Back side for this strategy
	Mb [][]float64 `json:"mb"`

	// Matched Lays - matched amounts by distinct matched price on the Lay side for this strategy
	Ml [][]float64 `json:"ml"`
}

// Validate validates this strategy match change
func (m *StrategyMatchChange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this strategy match change based on context it is used
func (m *StrategyMatchChange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StrategyMatchChange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StrategyMatchChange) UnmarshalBinary(b []byte) error {
	var res StrategyMatchChange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}