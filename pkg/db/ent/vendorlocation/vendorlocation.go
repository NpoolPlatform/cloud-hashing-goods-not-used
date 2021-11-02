// Code generated by entc, DO NOT EDIT.

package vendorlocation

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the vendorlocation type in the database.
	Label = "vendor_location"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldProvince holds the string denoting the province field in the database.
	FieldProvince = "province"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the vendorlocation in the database.
	Table = "vendor_locations"
)

// Columns holds all SQL columns for vendorlocation fields.
var Columns = []string{
	FieldID,
	FieldCountry,
	FieldProvince,
	FieldCity,
	FieldAddress,
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
	// DefaultCountry holds the default value on creation for the "country" field.
	DefaultCountry string
	// CountryValidator is a validator for the "country" field. It is called by the builders before save.
	CountryValidator func(string) error
	// DefaultProvince holds the default value on creation for the "province" field.
	DefaultProvince string
	// ProvinceValidator is a validator for the "province" field. It is called by the builders before save.
	ProvinceValidator func(string) error
	// DefaultCity holds the default value on creation for the "city" field.
	DefaultCity string
	// CityValidator is a validator for the "city" field. It is called by the builders before save.
	CityValidator func(string) error
	// DefaultAddress holds the default value on creation for the "address" field.
	DefaultAddress string
	// AddressValidator is a validator for the "address" field. It is called by the builders before save.
	AddressValidator func(string) error
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
