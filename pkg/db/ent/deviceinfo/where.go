// Code generated by entc, DO NOT EDIT.

package deviceinfo

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// Manufacturer applies equality check predicate on the "manufacturer" field. It's identical to ManufacturerEQ.
func Manufacturer(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldManufacturer), v))
	})
}

// PowerComsuption applies equality check predicate on the "power_comsuption" field. It's identical to PowerComsuptionEQ.
func PowerComsuption(v int) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPowerComsuption), v))
	})
}

// ShipmentDate applies equality check predicate on the "shipment_date" field. It's identical to ShipmentDateEQ.
func ShipmentDate(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShipmentDate), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// ManufacturerEQ applies the EQ predicate on the "manufacturer" field.
func ManufacturerEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldManufacturer), v))
	})
}

// ManufacturerNEQ applies the NEQ predicate on the "manufacturer" field.
func ManufacturerNEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldManufacturer), v))
	})
}

// ManufacturerIn applies the In predicate on the "manufacturer" field.
func ManufacturerIn(vs ...string) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldManufacturer), v...))
	})
}

// ManufacturerNotIn applies the NotIn predicate on the "manufacturer" field.
func ManufacturerNotIn(vs ...string) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldManufacturer), v...))
	})
}

// ManufacturerGT applies the GT predicate on the "manufacturer" field.
func ManufacturerGT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldManufacturer), v))
	})
}

// ManufacturerGTE applies the GTE predicate on the "manufacturer" field.
func ManufacturerGTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldManufacturer), v))
	})
}

// ManufacturerLT applies the LT predicate on the "manufacturer" field.
func ManufacturerLT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldManufacturer), v))
	})
}

// ManufacturerLTE applies the LTE predicate on the "manufacturer" field.
func ManufacturerLTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldManufacturer), v))
	})
}

// ManufacturerContains applies the Contains predicate on the "manufacturer" field.
func ManufacturerContains(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldManufacturer), v))
	})
}

// ManufacturerHasPrefix applies the HasPrefix predicate on the "manufacturer" field.
func ManufacturerHasPrefix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldManufacturer), v))
	})
}

// ManufacturerHasSuffix applies the HasSuffix predicate on the "manufacturer" field.
func ManufacturerHasSuffix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldManufacturer), v))
	})
}

// ManufacturerEqualFold applies the EqualFold predicate on the "manufacturer" field.
func ManufacturerEqualFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldManufacturer), v))
	})
}

// ManufacturerContainsFold applies the ContainsFold predicate on the "manufacturer" field.
func ManufacturerContainsFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldManufacturer), v))
	})
}

// PowerComsuptionEQ applies the EQ predicate on the "power_comsuption" field.
func PowerComsuptionEQ(v int) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPowerComsuption), v))
	})
}

// PowerComsuptionNEQ applies the NEQ predicate on the "power_comsuption" field.
func PowerComsuptionNEQ(v int) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPowerComsuption), v))
	})
}

// PowerComsuptionIn applies the In predicate on the "power_comsuption" field.
func PowerComsuptionIn(vs ...int) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPowerComsuption), v...))
	})
}

// PowerComsuptionNotIn applies the NotIn predicate on the "power_comsuption" field.
func PowerComsuptionNotIn(vs ...int) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPowerComsuption), v...))
	})
}

// PowerComsuptionGT applies the GT predicate on the "power_comsuption" field.
func PowerComsuptionGT(v int) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPowerComsuption), v))
	})
}

// PowerComsuptionGTE applies the GTE predicate on the "power_comsuption" field.
func PowerComsuptionGTE(v int) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPowerComsuption), v))
	})
}

// PowerComsuptionLT applies the LT predicate on the "power_comsuption" field.
func PowerComsuptionLT(v int) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPowerComsuption), v))
	})
}

// PowerComsuptionLTE applies the LTE predicate on the "power_comsuption" field.
func PowerComsuptionLTE(v int) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPowerComsuption), v))
	})
}

// ShipmentDateEQ applies the EQ predicate on the "shipment_date" field.
func ShipmentDateEQ(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShipmentDate), v))
	})
}

// ShipmentDateNEQ applies the NEQ predicate on the "shipment_date" field.
func ShipmentDateNEQ(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldShipmentDate), v))
	})
}

// ShipmentDateIn applies the In predicate on the "shipment_date" field.
func ShipmentDateIn(vs ...int64) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldShipmentDate), v...))
	})
}

// ShipmentDateNotIn applies the NotIn predicate on the "shipment_date" field.
func ShipmentDateNotIn(vs ...int64) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldShipmentDate), v...))
	})
}

// ShipmentDateGT applies the GT predicate on the "shipment_date" field.
func ShipmentDateGT(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldShipmentDate), v))
	})
}

// ShipmentDateGTE applies the GTE predicate on the "shipment_date" field.
func ShipmentDateGTE(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldShipmentDate), v))
	})
}

// ShipmentDateLT applies the LT predicate on the "shipment_date" field.
func ShipmentDateLT(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldShipmentDate), v))
	})
}

// ShipmentDateLTE applies the LTE predicate on the "shipment_date" field.
func ShipmentDateLTE(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldShipmentDate), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...int64) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...int64) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...int64) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...int64) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...int64) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...int64) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v int64) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DeviceInfo) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DeviceInfo) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DeviceInfo) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		p(s.Not())
	})
}
