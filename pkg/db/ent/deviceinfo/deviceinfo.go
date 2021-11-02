// Code generated by entc, DO NOT EDIT.

package deviceinfo

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the deviceinfo type in the database.
	Label = "device_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldManufacturer holds the string denoting the manufacturer field in the database.
	FieldManufacturer = "manufacturer"
	// FieldPowerComsuption holds the string denoting the power_comsuption field in the database.
	FieldPowerComsuption = "power_comsuption"
	// FieldShipmentDate holds the string denoting the shipment_date field in the database.
	FieldShipmentDate = "shipment_date"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the deviceinfo in the database.
	Table = "device_infos"
)

// Columns holds all SQL columns for deviceinfo fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldManufacturer,
	FieldPowerComsuption,
	FieldShipmentDate,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType string
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// DefaultManufacturer holds the default value on creation for the "manufacturer" field.
	DefaultManufacturer string
	// ManufacturerValidator is a validator for the "manufacturer" field. It is called by the builders before save.
	ManufacturerValidator func(string) error
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() int64
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() int64
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() int64
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
