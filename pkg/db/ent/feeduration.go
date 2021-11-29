// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/feeduration"
	"github.com/google/uuid"
)

// FeeDuration is the model entity for the FeeDuration schema.
type FeeDuration struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// FeeTypeID holds the value of the "fee_type_id" field.
	FeeTypeID uuid.UUID `json:"fee_type_id,omitempty"`
	// Duration holds the value of the "duration" field.
	// n days
	Duration int32 `json:"duration,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FeeDuration) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case feeduration.FieldDuration, feeduration.FieldCreateAt, feeduration.FieldUpdateAt, feeduration.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case feeduration.FieldID, feeduration.FieldFeeTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FeeDuration", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FeeDuration fields.
func (fd *FeeDuration) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case feeduration.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				fd.ID = *value
			}
		case feeduration.FieldFeeTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field fee_type_id", values[i])
			} else if value != nil {
				fd.FeeTypeID = *value
			}
		case feeduration.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				fd.Duration = int32(value.Int64)
			}
		case feeduration.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				fd.CreateAt = uint32(value.Int64)
			}
		case feeduration.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				fd.UpdateAt = uint32(value.Int64)
			}
		case feeduration.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				fd.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this FeeDuration.
// Note that you need to call FeeDuration.Unwrap() before calling this method if this FeeDuration
// was returned from a transaction, and the transaction was committed or rolled back.
func (fd *FeeDuration) Update() *FeeDurationUpdateOne {
	return (&FeeDurationClient{config: fd.config}).UpdateOne(fd)
}

// Unwrap unwraps the FeeDuration entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fd *FeeDuration) Unwrap() *FeeDuration {
	tx, ok := fd.config.driver.(*txDriver)
	if !ok {
		panic("ent: FeeDuration is not a transactional entity")
	}
	fd.config.driver = tx.drv
	return fd
}

// String implements the fmt.Stringer.
func (fd *FeeDuration) String() string {
	var builder strings.Builder
	builder.WriteString("FeeDuration(")
	builder.WriteString(fmt.Sprintf("id=%v", fd.ID))
	builder.WriteString(", fee_type_id=")
	builder.WriteString(fmt.Sprintf("%v", fd.FeeTypeID))
	builder.WriteString(", duration=")
	builder.WriteString(fmt.Sprintf("%v", fd.Duration))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", fd.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", fd.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", fd.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// FeeDurations is a parsable slice of FeeDuration.
type FeeDurations []*FeeDuration

func (fd FeeDurations) config(cfg config) {
	for _i := range fd {
		fd[_i].config = cfg
	}
}
