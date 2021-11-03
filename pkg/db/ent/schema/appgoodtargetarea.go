package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// AppGoodTargetArea holds the schema definition for the AppGoodTargetArea entity.
type AppGoodTargetArea struct {
	ent.Schema
}

// Fields of the AppGoodTargetArea.
func (AppGoodTargetArea) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("good_id", uuid.UUID{}),
		field.UUID("target_area_id", uuid.UUID{}),
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

// Edges of the AppGoodTargetArea.
func (AppGoodTargetArea) Edges() []ent.Edge {
	return nil
}

func (AppGoodTargetArea) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("good_id", "app_id", "target_area_id").
			Unique(),
	}
}
