// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ResponseMessage response message
//
// swagger:discriminator ResponseMessage op
type ResponseMessage interface {
	runtime.Validatable
	runtime.ContextValidatable

	// Client generated unique id to link request with response (like json rpc)
	ID() int32
	SetID(int32)

	// The operation type
	Op() string
	SetOp(string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type responseMessage struct {
	idField int32

	opField string
}

// ID gets the id of this polymorphic type
func (m *responseMessage) ID() int32 {
	return m.idField
}

// SetID sets the id of this polymorphic type
func (m *responseMessage) SetID(val int32) {
	m.idField = val
}

// Op gets the op of this polymorphic type
func (m *responseMessage) Op() string {
	return "ResponseMessage"
}

// SetOp sets the op of this polymorphic type
func (m *responseMessage) SetOp(val string) {
}

// UnmarshalResponseMessageSlice unmarshals polymorphic slices of ResponseMessage
func UnmarshalResponseMessageSlice(reader io.Reader, consumer runtime.Consumer) ([]ResponseMessage, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []ResponseMessage
	for _, element := range elements {
		obj, err := unmarshalResponseMessage(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalResponseMessage unmarshals polymorphic ResponseMessage
func UnmarshalResponseMessage(reader io.Reader, consumer runtime.Consumer) (ResponseMessage, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalResponseMessage(data, consumer)
}

func unmarshalResponseMessage(data []byte, consumer runtime.Consumer) (ResponseMessage, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the op property.
	var getType struct {
		Op string `json:"op"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("op", "body", getType.Op); err != nil {
		return nil, err
	}

	// The value of op is used to determine which type to create and unmarshal the data into
	switch getType.Op {
	case "ConnectionMessage":
		var result ConnectionMessage
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "MarketChangeMessage":
		var result MarketChangeMessage
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "OrderChangeMessage":
		var result OrderChangeMessage
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "ResponseMessage":
		var result responseMessage
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "StatusMessage":
		var result StatusMessage
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid op value: %q", getType.Op)
}

// Validate validates this response message
func (m *responseMessage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this response message based on context it is used
func (m *responseMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
