// Code generated by entc, DO NOT EDIT.

package cat

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the cat type in the database.
	Label = "cat"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldColor holds the string denoting the color field in the database.
	FieldColor = "color"
	// FieldFeral holds the string denoting the feral field in the database.
	FieldFeral = "feral"
	// Table holds the table name of the cat in the database.
	Table = "cats"
)

// Columns holds all SQL columns for cat fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldName,
	FieldCreatedAt,
	FieldColor,
	FieldFeral,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator func(int) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultFeral holds the default value on creation for the "feral" field.
	DefaultFeral bool
)

// Color defines the type for the "color" enum field.
type Color string

// Color values.
const (
	ColorBlack    Color = "black"
	ColorWhite    Color = "white"
	ColorGinger   Color = "ginger"
	ColorStripped Color = "stripped"
)

func (c Color) String() string {
	return string(c)
}

// ColorValidator is a validator for the "color" field enum values. It is called by the builders before save.
func ColorValidator(c Color) error {
	switch c {
	case ColorBlack, ColorWhite, ColorGinger, ColorStripped:
		return nil
	default:
		return fmt.Errorf("cat: invalid enum value for color field: %q", c)
	}
}
