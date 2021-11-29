package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// GoodFee holds the schema definition for the GoodFee entity.
type GoodFee struct {
	ent.Schema
}

// Fields of the GoodFee.
func (GoodFee) Fields() []ent.Field {
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

// Edges of the GoodFee.
func (GoodFee) Edges() []ent.Edge {
	return nil
}
