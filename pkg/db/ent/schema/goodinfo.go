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
		field.Int64("gas_price").
			Positive(),
		field.Bool("separate_gas_fee"),
		field.Int32("unit_power").
			Positive(),
		field.Int32("duration_days").
			Positive(),
		field.UUID("coin_info_id", uuid.UUID{}),
		field.Bool("actuals"),
		field.Int32("delivery_at"),
		field.UUID("inherit_from_good_id", uuid.UUID{}),
		field.UUID("vendor_location_id", uuid.UUID{}),
		field.Int64("price").
			Positive(),
		field.Enum("benefit_type").
			Values("pool", "platform"),
		field.Bool("classic"),
		field.JSON("support_coin_type_ids", []uuid.UUID{}),
		field.Int32("total").
			Positive(),
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

// Edges of the GoodInfo.
func (GoodInfo) Edges() []ent.Edge {
	return nil
}
