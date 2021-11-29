package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// FeeType holds the schema definition for the FeeType entity.
type FeeType struct {
	ent.Schema
}

// Fields of the FeeType.
func (FeeType) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("fee_type").
			Unique(),
		field.String("fee_description").
			MaxLen(256),
		field.Enum("pay_type").
			Values("percent", "amount"),
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

// Edges of the FeeType.
func (FeeType) Edges() []ent.Edge {
	return nil
}
