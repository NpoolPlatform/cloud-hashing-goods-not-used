// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/appgood"
	"github.com/google/uuid"
)

// AppGood is the model entity for the AppGood schema.
type AppGood struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// Authorized holds the value of the "authorized" field.
	Authorized bool `json:"authorized,omitempty"`
	// Online holds the value of the "online" field.
	Online bool `json:"online,omitempty"`
	// InitAreaStrategy holds the value of the "init_area_strategy" field.
	InitAreaStrategy appgood.InitAreaStrategy `json:"init_area_strategy,omitempty"`
	// Price holds the value of the "price" field.
	Price uint64 `json:"price,omitempty"`
	// InvitationOnly holds the value of the "invitation_only" field.
	InvitationOnly bool `json:"invitation_only,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppGood) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appgood.FieldAuthorized, appgood.FieldOnline, appgood.FieldInvitationOnly:
			values[i] = new(sql.NullBool)
		case appgood.FieldPrice, appgood.FieldCreateAt, appgood.FieldUpdateAt, appgood.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case appgood.FieldInitAreaStrategy:
			values[i] = new(sql.NullString)
		case appgood.FieldID, appgood.FieldAppID, appgood.FieldGoodID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppGood", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppGood fields.
func (ag *AppGood) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appgood.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ag.ID = *value
			}
		case appgood.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ag.AppID = *value
			}
		case appgood.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				ag.GoodID = *value
			}
		case appgood.FieldAuthorized:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field authorized", values[i])
			} else if value.Valid {
				ag.Authorized = value.Bool
			}
		case appgood.FieldOnline:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field online", values[i])
			} else if value.Valid {
				ag.Online = value.Bool
			}
		case appgood.FieldInitAreaStrategy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field init_area_strategy", values[i])
			} else if value.Valid {
				ag.InitAreaStrategy = appgood.InitAreaStrategy(value.String)
			}
		case appgood.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				ag.Price = uint64(value.Int64)
			}
		case appgood.FieldInvitationOnly:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field invitation_only", values[i])
			} else if value.Valid {
				ag.InvitationOnly = value.Bool
			}
		case appgood.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				ag.CreateAt = uint32(value.Int64)
			}
		case appgood.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				ag.UpdateAt = uint32(value.Int64)
			}
		case appgood.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				ag.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppGood.
// Note that you need to call AppGood.Unwrap() before calling this method if this AppGood
// was returned from a transaction, and the transaction was committed or rolled back.
func (ag *AppGood) Update() *AppGoodUpdateOne {
	return (&AppGoodClient{config: ag.config}).UpdateOne(ag)
}

// Unwrap unwraps the AppGood entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ag *AppGood) Unwrap() *AppGood {
	tx, ok := ag.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppGood is not a transactional entity")
	}
	ag.config.driver = tx.drv
	return ag
}

// String implements the fmt.Stringer.
func (ag *AppGood) String() string {
	var builder strings.Builder
	builder.WriteString("AppGood(")
	builder.WriteString(fmt.Sprintf("id=%v", ag.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", ag.AppID))
	builder.WriteString(", good_id=")
	builder.WriteString(fmt.Sprintf("%v", ag.GoodID))
	builder.WriteString(", authorized=")
	builder.WriteString(fmt.Sprintf("%v", ag.Authorized))
	builder.WriteString(", online=")
	builder.WriteString(fmt.Sprintf("%v", ag.Online))
	builder.WriteString(", init_area_strategy=")
	builder.WriteString(fmt.Sprintf("%v", ag.InitAreaStrategy))
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", ag.Price))
	builder.WriteString(", invitation_only=")
	builder.WriteString(fmt.Sprintf("%v", ag.InvitationOnly))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", ag.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", ag.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", ag.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// AppGoods is a parsable slice of AppGood.
type AppGoods []*AppGood

func (ag AppGoods) config(cfg config) {
	for _i := range ag {
		ag[_i].config = cfg
	}
}
