package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Activity holds the schema definition for the Activity entity.
type Activity struct {
	ent.Schema
}

// Fields of the Activity.
func (Activity) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.Int64("duration").Default(0).StorageKey("activity"),
	}
}

// Edges of the Activity.
func (Activity) Edges() []ent.Edge {
	return nil
}
