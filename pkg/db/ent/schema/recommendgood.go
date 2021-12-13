package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// RecommendGood holds the schema definition for the RecommendGood entity.
type RecommendGood struct {
	ent.Schema
}

// Fields of the RecommendGood.
func (RecommendGood) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("recommender_id", uuid.UUID{}),
		field.UUID("good_id", uuid.UUID{}),
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

// Indexes of the RecommendGood.
func (RecommendGood) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("recommender_id"),
		index.Fields("good_id"),
	}
}
