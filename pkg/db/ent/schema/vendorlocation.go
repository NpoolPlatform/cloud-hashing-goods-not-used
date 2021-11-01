package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// VendorLocation holds the schema definition for the VendorLocation entity.
type VendorLocation struct {
	ent.Schema
}

// Fields of the VendorLocation.
func (VendorLocation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("country").
			MaxLen(128),
		field.String("province").
			MaxLen(128),
		field.String("city").
			MaxLen(128),
		field.String("address").
			MaxLen(256),
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

// Edges of the VendorLocation.
func (VendorLocation) Edges() []ent.Edge {
	return nil
}

func (VendorLocation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("country", "province", "city", "address").
			Unique(),
	}
}
