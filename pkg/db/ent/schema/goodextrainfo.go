package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// GoodExtraInfo holds the schema definition for the GoodExtraInfo entity.
type GoodExtraInfo struct {
	ent.Schema
}

// Fields of the GoodExtraInfo.
func (GoodExtraInfo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("good_id", uuid.UUID{}).
			Unique(),
		field.JSON("posters", []string{}),
		field.JSON("labels", []string{}),
		field.Bool("out_sale"),
		field.Bool("pre_sale"),
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

// Edges of the GoodExtraInfo.
func (GoodExtraInfo) Edges() []ent.Edge {
	return nil
}
