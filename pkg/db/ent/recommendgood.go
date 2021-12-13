// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/recommendgood"
	"github.com/google/uuid"
)

// RecommendGood is the model entity for the RecommendGood schema.
type RecommendGood struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// RecommenderID holds the value of the "recommender_id" field.
	RecommenderID uuid.UUID `json:"recommender_id,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt int64 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt int64 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt int64 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RecommendGood) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case recommendgood.FieldCreateAt, recommendgood.FieldUpdateAt, recommendgood.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case recommendgood.FieldMessage:
			values[i] = new(sql.NullString)
		case recommendgood.FieldID, recommendgood.FieldRecommenderID, recommendgood.FieldGoodID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type RecommendGood", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RecommendGood fields.
func (rg *RecommendGood) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case recommendgood.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				rg.ID = *value
			}
		case recommendgood.FieldRecommenderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field recommender_id", values[i])
			} else if value != nil {
				rg.RecommenderID = *value
			}
		case recommendgood.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				rg.GoodID = *value
			}
		case recommendgood.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				rg.Message = value.String
			}
		case recommendgood.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				rg.CreateAt = value.Int64
			}
		case recommendgood.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				rg.UpdateAt = value.Int64
			}
		case recommendgood.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				rg.DeleteAt = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this RecommendGood.
// Note that you need to call RecommendGood.Unwrap() before calling this method if this RecommendGood
// was returned from a transaction, and the transaction was committed or rolled back.
func (rg *RecommendGood) Update() *RecommendGoodUpdateOne {
	return (&RecommendGoodClient{config: rg.config}).UpdateOne(rg)
}

// Unwrap unwraps the RecommendGood entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rg *RecommendGood) Unwrap() *RecommendGood {
	tx, ok := rg.config.driver.(*txDriver)
	if !ok {
		panic("ent: RecommendGood is not a transactional entity")
	}
	rg.config.driver = tx.drv
	return rg
}

// String implements the fmt.Stringer.
func (rg *RecommendGood) String() string {
	var builder strings.Builder
	builder.WriteString("RecommendGood(")
	builder.WriteString(fmt.Sprintf("id=%v", rg.ID))
	builder.WriteString(", recommender_id=")
	builder.WriteString(fmt.Sprintf("%v", rg.RecommenderID))
	builder.WriteString(", good_id=")
	builder.WriteString(fmt.Sprintf("%v", rg.GoodID))
	builder.WriteString(", message=")
	builder.WriteString(rg.Message)
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", rg.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", rg.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", rg.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// RecommendGoods is a parsable slice of RecommendGood.
type RecommendGoods []*RecommendGood

func (rg RecommendGoods) config(cfg config) {
	for _i := range rg {
		rg[_i].config = cfg
	}
}