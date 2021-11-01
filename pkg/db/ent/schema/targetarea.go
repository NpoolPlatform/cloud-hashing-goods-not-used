package schema

import (
	"time"

	"entgo.io/ent"
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
			Default("").
			MaxLen(128),
		field.String("country").
			Default("").
			MaxLen(128),
		field.Int64("create_at").
			DefaultFunc(func() int64 {
				return time.Now().UnixNano()
			}),
		field.Int64("update_at").
			DefaultFunc(func() int64 {
				return time.Now().UnixNano()
			}).
			UpdateDefault(func() int64 {
				return time.Now().UnixNano()
			}),
		field.Int64("delete_at").
			DefaultFunc(func() int64 {
				return 0
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
