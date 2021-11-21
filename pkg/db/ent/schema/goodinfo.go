package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// GoodInfo holds the schema definition for the GoodInfo entity.
type GoodInfo struct {
	ent.Schema
}

// Fields of the GoodInfo.
func (GoodInfo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("device_info_id", uuid.UUID{}),
		field.Bool("separate_fee"),
		field.Int32("unit_power").
			Positive(),
		field.Int32("duration_days").
			Positive(),
		field.UUID("coin_info_id", uuid.UUID{}),
		field.Bool("actuals"),
		field.Int32("delivery_at"),
		field.UUID("inherit_from_good_id", uuid.UUID{}),
		field.UUID("vendor_location_id", uuid.UUID{}),
		field.Uint64("price").
			Positive(),
		field.String("price_unit"),
		field.String("price_currency"),
		field.String("price_symbol"),
		field.Enum("benefit_type").
			Values("pool", "platform"),
		field.Bool("classic"),
		field.String("title"),
		field.String("unit"),
		field.Uint32("start"),
		field.JSON("support_coin_type_ids", []uuid.UUID{}),
		field.Int32("total").
			Positive(),
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

// Edges of the GoodInfo.
func (GoodInfo) Edges() []ent.Edge {
	return nil
}
