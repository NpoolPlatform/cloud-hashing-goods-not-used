package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// TargetArea holds the schema definition for the TargetArea entity.
type TargetArea struct {
	ent.Schema
}

// Fields of the TargetArea.
func (TargetArea) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("continent"),
		field.String("country"),
	}
}

// Edges of the TargetArea.
func (TargetArea) Edges() []ent.Edge {
	return nil
}
