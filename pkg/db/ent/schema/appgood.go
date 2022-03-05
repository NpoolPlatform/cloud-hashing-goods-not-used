package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// AppGood holds the schema definition for the AppGood entity.
type AppGood struct {
	ent.Schema
}

// Fields of the AppGood.
func (AppGood) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("good_id", uuid.UUID{}),
		field.Bool("online").
			Default(false),
		field.Enum("init_area_strategy").
			Values("all", "none"),
		field.Uint64("price"),
		field.Uint32("display_index"),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

// Edges of the AppGood.
func (AppGood) Edges() []ent.Edge {
	return nil
}

func (AppGood) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("good_id", "app_id").
			Unique(),
	}
}
