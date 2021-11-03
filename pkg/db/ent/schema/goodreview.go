package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// GoodReview holds the schema definition for the GoodReview entity.
type GoodReview struct {
	ent.Schema
}

// Fields of the GoodReview.
func (GoodReview) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Enum("type").
			Values("good", "appgood", "apptargetarea", "appgoodtargetarea"),
		field.UUID("reviewed_id", uuid.UUID{}),
		field.UUID("reviewer_id", uuid.UUID{}),
		field.Enum("state").
			Default("none").
			Values("approved", "rejected", "none"),
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

// Edges of the GoodReview.
func (GoodReview) Edges() []ent.Edge {
	return nil
}

func (GoodReview) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type", "reviewed_id").
			Unique(),
	}
}
