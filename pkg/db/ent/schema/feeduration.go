package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// FeeDuration holds the schema definition for the FeeDuration entity.
type FeeDuration struct {
	ent.Schema
}

// Fields of the FeeDuration.
func (FeeDuration) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("fee_type_id", uuid.UUID{}),
		field.Int32("duration").Comment("n days"),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

// Indexes of the FeeDuration.
func (FeeDuration) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("fee_type_id"),
		index.Fields("fee_type_id", "duration").Unique(),
	}
}
