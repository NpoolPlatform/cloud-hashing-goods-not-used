package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// Recommend holds the schema definition for the Recommend entity.
type Recommend struct {
	ent.Schema
}

// Fields of the Recommend.
func (Recommend) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("good_id", uuid.UUID{}),
		field.UUID("recommender_id", uuid.UUID{}),
		field.String("message").
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

// Indexes of the Recommend.
func (Recommend) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "good_id"),
	}
}
