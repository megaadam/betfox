// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StatusMessage status message
//
// swagger:model StatusMessage
type StatusMessage struct {
	idField int32

	// Is the connection now closed
	ConnectionClosed bool `json:"connectionClosed,omitempty"`

	// The connection id
	ConnectionID string `json:"connectionId,omitempty"`

	// The number of connections available for this account at this moment in time. Present on responses to Authentication messages only.
	ConnectionsAvailable int32 `json:"connectionsAvailable,omitempty"`

	// The type of error in case of a failure
	// Enum: [NO_APP_KEY INVALID_APP_KEY NO_SESSION INVALID_SESSION_INFORMATION NOT_AUTHORIZED INVALID_INPUT INVALID_CLOCK UNEXPECTED_ERROR TIMEOUT SUBSCRIPTION_LIMIT_EXCEEDED INVALID_REQUEST CONNECTION_FAILED MAX_CONNECTION_LIMIT_EXCEEDED TOO_MANY_REQUESTS]
	ErrorCode string `json:"errorCode,omitempty"`

	// Additional message in case of a failure
	ErrorMessage string `json:"errorMessage,omitempty"`

	// The status of the last request
	// Enum: [SUCCESS FAILURE]
	StatusCode string `json:"statusCode,omitempty"`
}

// ID gets the id of this subtype
func (m *StatusMessage) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *StatusMessage) SetID(val int32) {
	m.idField = val
}

// Op gets the op of this subtype
func (m *StatusMessage) Op() string {
	return "StatusMessage"
}

// SetOp sets the op of this subtype
func (m *StatusMessage) SetOp(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *StatusMessage) UnmarshalJSON(raw []byte) error {
	var data struct {

		// Is the connection now closed
		ConnectionClosed bool `json:"connectionClosed,omitempty"`

		// The connection id
		ConnectionID string `json:"connectionId,omitempty"`

		// The number of connections available for this account at this moment in time. Present on responses to Authentication messages only.
		ConnectionsAvailable int32 `json:"connectionsAvailable,omitempty"`

		// The type of error in case of a failure
		// Enum: [NO_APP_KEY INVALID_APP_KEY NO_SESSION INVALID_SESSION_INFORMATION NOT_AUTHORIZED INVALID_INPUT INVALID_CLOCK UNEXPECTED_ERROR TIMEOUT SUBSCRIPTION_LIMIT_EXCEEDED INVALID_REQUEST CONNECTION_FAILED MAX_CONNECTION_LIMIT_EXCEEDED TOO_MANY_REQUESTS]
		ErrorCode string `json:"errorCode,omitempty"`

		// Additional message in case of a failure
		ErrorMessage string `json:"errorMessage,omitempty"`

		// The status of the last request
		// Enum: [SUCCESS FAILURE]
		StatusCode string `json:"statusCode,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		ID int32 `json:"id,omitempty"`

		Op string `json:"op,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result StatusMessage

	result.idField = base.ID

	if base.Op != result.Op() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid op value: %q", base.Op)
	}

	result.ConnectionClosed = data.ConnectionClosed
	result.ConnectionID = data.ConnectionID
	result.ConnectionsAvailable = data.ConnectionsAvailable
	result.ErrorCode = data.ErrorCode
	result.ErrorMessage = data.ErrorMessage
	result.StatusCode = data.StatusCode

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m StatusMessage) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// Is the connection now closed
		ConnectionClosed bool `json:"connectionClosed,omitempty"`

		// The connection id
		ConnectionID string `json:"connectionId,omitempty"`

		// The number of connections available for this account at this moment in time. Present on responses to Authentication messages only.
		ConnectionsAvailable int32 `json:"connectionsAvailable,omitempty"`

		// The type of error in case of a failure
		// Enum: [NO_APP_KEY INVALID_APP_KEY NO_SESSION INVALID_SESSION_INFORMATION NOT_AUTHORIZED INVALID_INPUT INVALID_CLOCK UNEXPECTED_ERROR TIMEOUT SUBSCRIPTION_LIMIT_EXCEEDED INVALID_REQUEST CONNECTION_FAILED MAX_CONNECTION_LIMIT_EXCEEDED TOO_MANY_REQUESTS]
		ErrorCode string `json:"errorCode,omitempty"`

		// Additional message in case of a failure
		ErrorMessage string `json:"errorMessage,omitempty"`

		// The status of the last request
		// Enum: [SUCCESS FAILURE]
		StatusCode string `json:"statusCode,omitempty"`
	}{

		ConnectionClosed: m.ConnectionClosed,

		ConnectionID: m.ConnectionID,

		ConnectionsAvailable: m.ConnectionsAvailable,

		ErrorCode: m.ErrorCode,

		ErrorMessage: m.ErrorMessage,

		StatusCode: m.StatusCode,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		ID int32 `json:"id,omitempty"`

		Op string `json:"op,omitempty"`
	}{

		ID: m.ID(),

		Op: m.Op(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this status message
func (m *StatusMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var statusMessageTypeErrorCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NO_APP_KEY","INVALID_APP_KEY","NO_SESSION","INVALID_SESSION_INFORMATION","NOT_AUTHORIZED","INVALID_INPUT","INVALID_CLOCK","UNEXPECTED_ERROR","TIMEOUT","SUBSCRIPTION_LIMIT_EXCEEDED","INVALID_REQUEST","CONNECTION_FAILED","MAX_CONNECTION_LIMIT_EXCEEDED","TOO_MANY_REQUESTS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusMessageTypeErrorCodePropEnum = append(statusMessageTypeErrorCodePropEnum, v)
	}
}

// property enum
func (m *StatusMessage) validateErrorCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, statusMessageTypeErrorCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StatusMessage) validateErrorCode(formats strfmt.Registry) error {

	if swag.IsZero(m.ErrorCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateErrorCodeEnum("errorCode", "body", m.ErrorCode); err != nil {
		return err
	}

	return nil
}

var statusMessageTypeStatusCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUCCESS","FAILURE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusMessageTypeStatusCodePropEnum = append(statusMessageTypeStatusCodePropEnum, v)
	}
}

// property enum
func (m *StatusMessage) validateStatusCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, statusMessageTypeStatusCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StatusMessage) validateStatusCode(formats strfmt.Registry) error {

	if swag.IsZero(m.StatusCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusCodeEnum("statusCode", "body", m.StatusCode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this status message based on the context it is used
func (m *StatusMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StatusMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusMessage) UnmarshalBinary(b []byte) error {
	var res StatusMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
