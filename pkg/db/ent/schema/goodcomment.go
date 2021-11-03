package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// GoodComment holds the schema definition for the GoodComment entity.
type GoodComment struct {
	ent.Schema
}

// Fields of the GoodComment.
func (GoodComment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("reply_to_id", uuid.UUID{}).
			Optional(),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("good_id", uuid.UUID{}),
		field.UUID("order_id", uuid.UUID{}),
		field.String("content").
			Default(""),
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

// Edges of the GoodComment.
func (GoodComment) Edges() []ent.Edge {
	return nil
}
