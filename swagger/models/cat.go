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

// Cat cat
//
// swagger:model cat
type Cat struct {

	// age
	Age int32 `json:"age,omitempty"`

	// color
	// Required: true
	// Enum: [black white ginger stripped]
	Color *string `json:"color"`

	// created at
	// Read Only: true
	CreatedAt string `json:"created_at,omitempty"`

	// feral
	Feral bool `json:"feral,omitempty"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// name
	// Required: true
	// Min Length: 2
	Name *string `json:"name"`
}

// Validate validates this cat
func (m *Cat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var catTypeColorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["black","white","ginger","stripped"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		catTypeColorPropEnum = append(catTypeColorPropEnum, v)
	}
}

const (

	// CatColorBlack captures enum value "black"
	CatColorBlack string = "black"

	// CatColorWhite captures enum value "white"
	CatColorWhite string = "white"

	// CatColorGinger captures enum value "ginger"
	CatColorGinger string = "ginger"

	// CatColorStripped captures enum value "stripped"
	CatColorStripped string = "stripped"
)

// prop value enum
func (m *Cat) validateColorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, catTypeColorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Cat) validateColor(formats strfmt.Registry) error {

	if err := validate.Required("color", "body", m.Color); err != nil {
		return err
	}

	// value enum
	if err := m.validateColorEnum("color", "body", *m.Color); err != nil {
		return err
	}

	return nil
}

func (m *Cat) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 2); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cat based on the context it is used
func (m *Cat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cat) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", string(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Cat) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cat) UnmarshalBinary(b []byte) error {
	var res Cat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
