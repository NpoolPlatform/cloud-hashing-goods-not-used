package schema

import (
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
			Default(uuid.New),
		field.UUID("device_info_id", uuid.UUID{}),
		field.Int("gas_price").
			Positive(),
		field.Bool("separate_gas_fee"),
		field.Float("unit_power"),
		field.Int("duration"),
		field.UUID("coin_info_id", uuid.UUID{}),
		field.Bool("actuals"),
		field.Time("delivery_time"),
		field.UUID("inherit_from_good_id", uuid.UUID{}),
		field.UUID("vendor_location_id", uuid.UUID{}),
		field.Int("price"),
		field.String("benefit_type"),
		field.Bool("classic"),
		field.JSON("support_coin_type_ids", []uuid.UUID{}),
		field.UUID("reviewer_id", uuid.UUID{}),
		field.String("state"),
		field.Int("Total"),
	}
}

// Edges of the GoodInfo.
func (GoodInfo) Edges() []ent.Edge {
	return nil
}
