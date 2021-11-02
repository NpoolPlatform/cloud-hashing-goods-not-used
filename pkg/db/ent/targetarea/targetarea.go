// Code generated by entc, DO NOT EDIT.

package targetarea

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the targetarea type in the database.
	Label = "target_area"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldContinent holds the string denoting the continent field in the database.
	FieldContinent = "continent"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the targetarea in the database.
	Table = "target_areas"
)

// Columns holds all SQL columns for targetarea fields.
var Columns = []string{
	FieldID,
	FieldContinent,
	FieldCountry,
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
	// DefaultContinent holds the default value on creation for the "continent" field.
	DefaultContinent string
	// ContinentValidator is a validator for the "continent" field. It is called by the builders before save.
	ContinentValidator func(string) error
	// DefaultCountry holds the default value on creation for the "country" field.
	DefaultCountry string
	// CountryValidator is a validator for the "country" field. It is called by the builders before save.
	CountryValidator func(string) error
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
