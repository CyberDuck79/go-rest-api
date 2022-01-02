package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Cat holds the schema definition for the Cat entity.
type Cat struct {
	ent.Schema
}

// Fields of the Cat.
func (Cat) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			NonNegative().
			Optional(),
		field.String("name").
			Unique(),
		field.Time("created_at").
			Default(time.Now),
		field.Enum("color").
			Values("black", "white", "ginger", "stripped"),
		field.Bool("feral").
			Default(false),
	}
}

// Edges of the Cat.
func (Cat) Edges() []ent.Edge {
	// human master ?
	return nil
}
