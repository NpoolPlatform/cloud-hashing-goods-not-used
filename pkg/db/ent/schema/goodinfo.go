package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
		field.Int("gas_price").
			Positive(),
		field.Bool("separate_gas_fee"),
		field.Float("unit_power").
			Positive(),
		field.Int("duration").
			Positive(),
		field.UUID("coin_info_id", uuid.UUID{}),
		field.Bool("actuals"),
		field.Int("delivery_time"),
		field.UUID("inherit_from_good_id", uuid.UUID{}),
		field.UUID("vendor_location_id", uuid.UUID{}),
		field.Int("price").
			Positive(),
		field.Enum("benefit_type").
			Values("pool", "platform"),
		field.Bool("classic"),
		field.JSON("support_coin_type_ids", []uuid.UUID{}),
		field.UUID("reviewer_id", uuid.UUID{}),
		field.Enum("review_state").
			Values("passed", "rejected"),
		field.Int("total").
			Positive(),
		field.Time("create_at").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}),
		field.Time("update_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}),
	}
}

// Edges of the GoodInfo.
func (GoodInfo) Edges() []ent.Edge {
	return nil
}
