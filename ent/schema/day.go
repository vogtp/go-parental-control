package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Day holds the schema definition for the Day entity.
type Day struct {
	ent.Schema
}

// Fields of the User.
func (Day) Fields() []ent.Field {
	return []ent.Field{
		field.Int("year"),
		field.Int("month"),
		field.Int("day"),
	}
}

// Edges of the User.
func (Day) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("activity", Activity.Type),
	}
}
