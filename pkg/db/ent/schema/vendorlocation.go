package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
		field.String("country"),
		field.String("province"),
		field.String("city"),
		field.String("address"),
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
