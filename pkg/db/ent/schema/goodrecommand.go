package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// GoodRecommand holds the schema definition for the GoodRecommand entity.
type GoodRecommand struct {
	ent.Schema
}

// Fields of the GoodRecommand.
func (GoodRecommand) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("good_id", uuid.UUID{}),
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

// Indexes of the GoodRecommand.
func (GoodRecommand) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("good_id"),
	}
}
