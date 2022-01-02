// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"rest_api/ent/cat"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Cat is the model entity for the Cat schema.
type Cat struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Color holds the value of the "color" field.
	Color cat.Color `json:"color,omitempty"`
	// Feral holds the value of the "feral" field.
	Feral bool `json:"feral,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cat) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case cat.FieldFeral:
			values[i] = new(sql.NullBool)
		case cat.FieldID, cat.FieldAge:
			values[i] = new(sql.NullInt64)
		case cat.FieldName, cat.FieldColor:
			values[i] = new(sql.NullString)
		case cat.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Cat", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cat fields.
func (c *Cat) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cat.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case cat.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				c.Age = int(value.Int64)
			}
		case cat.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case cat.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case cat.FieldColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field color", values[i])
			} else if value.Valid {
				c.Color = cat.Color(value.String)
			}
		case cat.FieldFeral:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field feral", values[i])
			} else if value.Valid {
				c.Feral = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Cat.
// Note that you need to call Cat.Unwrap() before calling this method if this Cat
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cat) Update() *CatUpdateOne {
	return (&CatClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Cat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cat) Unwrap() *Cat {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cat is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cat) String() string {
	var builder strings.Builder
	builder.WriteString("Cat(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", c.Age))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", color=")
	builder.WriteString(fmt.Sprintf("%v", c.Color))
	builder.WriteString(", feral=")
	builder.WriteString(fmt.Sprintf("%v", c.Feral))
	builder.WriteByte(')')
	return builder.String()
}

// Cats is a parsable slice of Cat.
type Cats []*Cat

func (c Cats) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
