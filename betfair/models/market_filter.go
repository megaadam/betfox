// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MarketFilter market filter
//
// swagger:model MarketFilter
type MarketFilter struct {

	// betting types
	BettingTypes []string `json:"bettingTypes,omitempty"`

	// bsp market
	BspMarket bool `json:"bspMarket,omitempty"`

	// country codes
	CountryCodes []string `json:"countryCodes,omitempty""`

	// event ids
	EventIds []string `json:"eventIds,omitempty""`

	// event type ids
	EventTypeIds []string `json:"eventTypeIds,omitempty""`

	// market ids
	MarketIds []string `json:"marketIds"`

	// market types
	MarketTypes []string `json:"marketTypes,omitempty""`

	// race types
	RaceTypes []string `json:"raceTypes,omitempty""`

	// turn in play enabled
	TurnInPlayEnabled bool `json:"turnInPlayEnabled,omitempty"`

	// venues
	Venues []string `json:"venues,omitempty""`
}

// Validate validates this market filter
func (m *MarketFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBettingTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var marketFilterBettingTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ODDS","LINE","RANGE","ASIAN_HANDICAP_DOUBLE_LINE","ASIAN_HANDICAP_SINGLE_LINE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		marketFilterBettingTypesItemsEnum = append(marketFilterBettingTypesItemsEnum, v)
	}
}

func (m *MarketFilter) validateBettingTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, marketFilterBettingTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MarketFilter) validateBettingTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.BettingTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.BettingTypes); i++ {

		// value enum
		if err := m.validateBettingTypesItemsEnum("bettingTypes"+"."+strconv.Itoa(i), "body", m.BettingTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this market filter based on context it is used
func (m *MarketFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MarketFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarketFilter) UnmarshalBinary(b []byte) error {
	var res MarketFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
