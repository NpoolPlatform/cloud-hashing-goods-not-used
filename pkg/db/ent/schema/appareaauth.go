package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// AppAreaAuth holds the schema definition for the AppAreaAuth entity.
type AppAreaAuth struct {
	ent.Schema
}

// Fields of the AppAreaAuth.
func (AppAreaAuth) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("target_area_id", uuid.UUID{}),
		field.UUID("app_id", uuid.UUID{}),
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

// Edges of the AppAreaAuth.
func (AppAreaAuth) Edges() []ent.Edge {
	return nil
}

func (AppAreaAuth) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("target_area_id", "app_id").
			Unique(),
	}
}
