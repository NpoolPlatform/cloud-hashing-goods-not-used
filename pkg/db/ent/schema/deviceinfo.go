package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// DeviceInfo holds the schema definition for the DeviceInfo entity.
type DeviceInfo struct {
	ent.Schema
}

// Fields of the DeviceInfo.
func (DeviceInfo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("type").
			Default("").
			MaxLen(64),
		field.String("manufacturer").
			Default("").
			MaxLen(64),
		field.Int32("power_comsuption"),
		field.Int32("shipment_at"),
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

// Edges of the DeviceInfo.
func (DeviceInfo) Edges() []ent.Edge {
	return nil
}

func (DeviceInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type", "manufacturer", "shipment_at", "power_comsuption").
			Unique(),
	}
}
