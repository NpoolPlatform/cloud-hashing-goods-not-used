package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// AppGoodPromotion holds the schema definition for the AppGoodPromotion entity.
type AppGoodPromotion struct {
	ent.Schema
}

// Fields of the AppGoodPromotion.
func (AppGoodPromotion) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("good_id", uuid.UUID{}),
		field.String("message"),
		field.Uint32("start"),
		field.Uint32("end"),
		field.Uint64("price"),
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

// Edges of the AppGoodPromotion.
func (AppGoodPromotion) Edges() []ent.Edge {
	return nil
}

func (AppGoodPromotion) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("good_id", "app_id", "start", "end").
			Unique(),
	}
}
