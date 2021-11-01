package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

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
			Default(uuid.New).
			Unique(),
		field.String("continent").
			MaxLen(32),
		field.String("country").
			MaxLen(64),
		field.Time("create_at").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}),
		field.Time("update_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}),
	}
}

// Edges of the TargetArea.
func (TargetArea) Edges() []ent.Edge {
	return nil
}

func (TargetArea) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("continent", "country").
			Unique(),
	}
}
