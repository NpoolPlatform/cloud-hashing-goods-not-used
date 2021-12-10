// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodrecommend"
	"github.com/google/uuid"
)

// GoodRecommend is the model entity for the GoodRecommend schema.
type GoodRecommend struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt int64 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt int64 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt int64 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GoodRecommend) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case goodrecommend.FieldCreateAt, goodrecommend.FieldUpdateAt, goodrecommend.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case goodrecommend.FieldContent:
			values[i] = new(sql.NullString)
		case goodrecommend.FieldID, goodrecommend.FieldAppID, goodrecommend.FieldUserID, goodrecommend.FieldGoodID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GoodRecommend", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GoodRecommend fields.
func (gr *GoodRecommend) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case goodrecommend.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				gr.ID = *value
			}
		case goodrecommend.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				gr.AppID = *value
			}
		case goodrecommend.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				gr.UserID = *value
			}
		case goodrecommend.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				gr.GoodID = *value
			}
		case goodrecommend.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				gr.Content = value.String
			}
		case goodrecommend.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				gr.CreateAt = value.Int64
			}
		case goodrecommend.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				gr.UpdateAt = value.Int64
			}
		case goodrecommend.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				gr.DeleteAt = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this GoodRecommend.
// Note that you need to call GoodRecommend.Unwrap() before calling this method if this GoodRecommend
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *GoodRecommend) Update() *GoodRecommendUpdateOne {
	return (&GoodRecommendClient{config: gr.config}).UpdateOne(gr)
}

// Unwrap unwraps the GoodRecommend entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gr *GoodRecommend) Unwrap() *GoodRecommend {
	tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: GoodRecommend is not a transactional entity")
	}
	gr.config.driver = tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *GoodRecommend) String() string {
	var builder strings.Builder
	builder.WriteString("GoodRecommend(")
	builder.WriteString(fmt.Sprintf("id=%v", gr.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", gr.AppID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", gr.UserID))
	builder.WriteString(", good_id=")
	builder.WriteString(fmt.Sprintf("%v", gr.GoodID))
	builder.WriteString(", content=")
	builder.WriteString(gr.Content)
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", gr.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", gr.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", gr.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// GoodRecommends is a parsable slice of GoodRecommend.
type GoodRecommends []*GoodRecommend

func (gr GoodRecommends) config(cfg config) {
	for _i := range gr {
		gr[_i].config = cfg
	}
}