// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Order order
//
// swagger:model Order
type Order struct {

	// Average Price Matched - the average price the order was matched at (null if the order is not matched). This value is not meaningful for activity on Line markets and is not guaranteed to be returned or maintained for these markets.
	Avp float64 `json:"avp,omitempty"`

	// BSP Liability - the BSP liability of the order (null if the order is not a BSP order)
	Bsp float64 `json:"bsp,omitempty"`

	// Cancelled Date - the date the order was cancelled (null if the order is not cancelled)
	Cd int64 `json:"cd,omitempty"`

	// Bet Id - the id of the order
	ID string `json:"id,omitempty"`

	// Lapsed Date - the date the order was lapsed (null if the order is not lapsed)
	Ld int64 `json:"ld,omitempty"`

	// Lapse Status Reason Code - the reason that some or all of this order has been lapsed (null if no portion of the order is lapsed
	Lsrc string `json:"lsrc,omitempty"`

	// Matched Date - the date the order was matched (null if the order is not matched)
	Md int64 `json:"md,omitempty"`

	// Order Type - the type of the order (L = LIMIT, MOC = MARKET_ON_CLOSE, LOC = LIMIT_ON_CLOSE)
	// Enum: [L LOC MOC]
	Ot string `json:"ot,omitempty"`

	// Price - the original placed price of the order. Line markets operate at even-money odds of 2.0. However, price for these markets refers to the line positions available as defined by the markets min-max range and interval steps
	P float64 `json:"p,omitempty"`

	// Placed Date - the date the order was placed
	Pd int64 `json:"pd,omitempty"`

	// Persistence Type - whether the order will persist at in play or not (L = LAPSE, P = PERSIST, MOC = Market On Close)
	// Enum: [L P MOC]
	Pt string `json:"pt,omitempty"`

	// Regulator Auth Code - the auth code returned by the regulator
	Rac string `json:"rac,omitempty"`

	// Regulator Code - the regulator of the order
	Rc string `json:"rc,omitempty"`

	// Order Reference - the customer's order reference for this order (empty string if one was not set)
	Rfo string `json:"rfo,omitempty"`

	// Strategy Reference - the customer's strategy reference for this order (empty string if one was not set)
	Rfs string `json:"rfs,omitempty"`

	// Size - the original placed size of the order
	S float64 `json:"s,omitempty"`

	// Size Cancelled - the amount of the order that has been cancelled
	Sc float64 `json:"sc,omitempty"`

	// Side - the side of the order. For Line markets a 'B' bet refers to a SELL line and an 'L' bet refers to a BUY line.
	// Enum: [B L]
	Side string `json:"side,omitempty"`

	// Size Lapsed - the amount of the order that has been lapsed
	Sl float64 `json:"sl,omitempty"`

	// Size Matched - the amount of the order that has been matched
	Sm float64 `json:"sm,omitempty"`

	// Size Remaining - the amount of the order that is remaining unmatched
	Sr float64 `json:"sr,omitempty"`

	// Status - the status of the order (E = EXECUTABLE, EC = EXECUTION_COMPLETE)
	// Enum: [E EC]
	Status string `json:"status,omitempty"`

	// Size Voided - the amount of the order that has been voided
	Sv float64 `json:"sv,omitempty"`
}

// Validate validates this order
func (m *Order) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSide(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var orderTypeOtPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["L","LOC","MOC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypeOtPropEnum = append(orderTypeOtPropEnum, v)
	}
}

const (

	// OrderOtL captures enum value "L"
	OrderOtL string = "L"

	// OrderOtLOC captures enum value "LOC"
	OrderOtLOC string = "LOC"

	// OrderOtMOC captures enum value "MOC"
	OrderOtMOC string = "MOC"
)

// prop value enum
func (m *Order) validateOtEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orderTypeOtPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Order) validateOt(formats strfmt.Registry) error {
	if swag.IsZero(m.Ot) { // not required
		return nil
	}

	// value enum
	if err := m.validateOtEnum("ot", "body", m.Ot); err != nil {
		return err
	}

	return nil
}

var orderTypePtPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["L","P","MOC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypePtPropEnum = append(orderTypePtPropEnum, v)
	}
}

const (

	// OrderPtL captures enum value "L"
	OrderPtL string = "L"

	// OrderPtP captures enum value "P"
	OrderPtP string = "P"

	// OrderPtMOC captures enum value "MOC"
	OrderPtMOC string = "MOC"
)

// prop value enum
func (m *Order) validatePtEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orderTypePtPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Order) validatePt(formats strfmt.Registry) error {
	if swag.IsZero(m.Pt) { // not required
		return nil
	}

	// value enum
	if err := m.validatePtEnum("pt", "body", m.Pt); err != nil {
		return err
	}

	return nil
}

var orderTypeSidePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["B","L"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypeSidePropEnum = append(orderTypeSidePropEnum, v)
	}
}

const (

	// OrderSideB captures enum value "B"
	OrderSideB string = "B"

	// OrderSideL captures enum value "L"
	OrderSideL string = "L"
)

// prop value enum
func (m *Order) validateSideEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orderTypeSidePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Order) validateSide(formats strfmt.Registry) error {
	if swag.IsZero(m.Side) { // not required
		return nil
	}

	// value enum
	if err := m.validateSideEnum("side", "body", m.Side); err != nil {
		return err
	}

	return nil
}

var orderTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["E","EC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypeStatusPropEnum = append(orderTypeStatusPropEnum, v)
	}
}

const (

	// OrderStatusE captures enum value "E"
	OrderStatusE string = "E"

	// OrderStatusEC captures enum value "EC"
	OrderStatusEC string = "EC"
)

// prop value enum
func (m *Order) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orderTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Order) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this order based on context it is used
func (m *Order) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Order) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Order) UnmarshalBinary(b []byte) error {
	var res Order
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
