// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodinfo"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/schema"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/targetarea"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/vendorlocation"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	goodinfoFields := schema.GoodInfo{}.Fields()
	_ = goodinfoFields
	// goodinfoDescGasPrice is the schema descriptor for gas_price field.
	goodinfoDescGasPrice := goodinfoFields[2].Descriptor()
	// goodinfo.GasPriceValidator is a validator for the "gas_price" field. It is called by the builders before save.
	goodinfo.GasPriceValidator = goodinfoDescGasPrice.Validators[0].(func(int) error)
	// goodinfoDescUnitPower is the schema descriptor for unit_power field.
	goodinfoDescUnitPower := goodinfoFields[4].Descriptor()
	// goodinfo.UnitPowerValidator is a validator for the "unit_power" field. It is called by the builders before save.
	goodinfo.UnitPowerValidator = goodinfoDescUnitPower.Validators[0].(func(float64) error)
	// goodinfoDescDuration is the schema descriptor for duration field.
	goodinfoDescDuration := goodinfoFields[5].Descriptor()
	// goodinfo.DurationValidator is a validator for the "duration" field. It is called by the builders before save.
	goodinfo.DurationValidator = goodinfoDescDuration.Validators[0].(func(int) error)
	// goodinfoDescPrice is the schema descriptor for price field.
	goodinfoDescPrice := goodinfoFields[11].Descriptor()
	// goodinfo.PriceValidator is a validator for the "price" field. It is called by the builders before save.
	goodinfo.PriceValidator = goodinfoDescPrice.Validators[0].(func(int) error)
	// goodinfoDescTotal is the schema descriptor for total field.
	goodinfoDescTotal := goodinfoFields[17].Descriptor()
	// goodinfo.TotalValidator is a validator for the "total" field. It is called by the builders before save.
	goodinfo.TotalValidator = goodinfoDescTotal.Validators[0].(func(int) error)
	// goodinfoDescCreateAt is the schema descriptor for create_at field.
	goodinfoDescCreateAt := goodinfoFields[18].Descriptor()
	// goodinfo.DefaultCreateAt holds the default value on creation for the create_at field.
	goodinfo.DefaultCreateAt = goodinfoDescCreateAt.Default.(func() time.Time)
	// goodinfoDescUpdateAt is the schema descriptor for update_at field.
	goodinfoDescUpdateAt := goodinfoFields[19].Descriptor()
	// goodinfo.DefaultUpdateAt holds the default value on creation for the update_at field.
	goodinfo.DefaultUpdateAt = goodinfoDescUpdateAt.Default.(func() time.Time)
	// goodinfo.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	goodinfo.UpdateDefaultUpdateAt = goodinfoDescUpdateAt.UpdateDefault.(func() time.Time)
	// goodinfoDescID is the schema descriptor for id field.
	goodinfoDescID := goodinfoFields[0].Descriptor()
	// goodinfo.DefaultID holds the default value on creation for the id field.
	goodinfo.DefaultID = goodinfoDescID.Default.(func() uuid.UUID)
	targetareaFields := schema.TargetArea{}.Fields()
	_ = targetareaFields
	// targetareaDescContinent is the schema descriptor for continent field.
	targetareaDescContinent := targetareaFields[1].Descriptor()
	// targetarea.ContinentValidator is a validator for the "continent" field. It is called by the builders before save.
	targetarea.ContinentValidator = targetareaDescContinent.Validators[0].(func(string) error)
	// targetareaDescCountry is the schema descriptor for country field.
	targetareaDescCountry := targetareaFields[2].Descriptor()
	// targetarea.CountryValidator is a validator for the "country" field. It is called by the builders before save.
	targetarea.CountryValidator = targetareaDescCountry.Validators[0].(func(string) error)
	// targetareaDescCreateAt is the schema descriptor for create_at field.
	targetareaDescCreateAt := targetareaFields[3].Descriptor()
	// targetarea.DefaultCreateAt holds the default value on creation for the create_at field.
	targetarea.DefaultCreateAt = targetareaDescCreateAt.Default.(func() time.Time)
	// targetareaDescUpdateAt is the schema descriptor for update_at field.
	targetareaDescUpdateAt := targetareaFields[4].Descriptor()
	// targetarea.DefaultUpdateAt holds the default value on creation for the update_at field.
	targetarea.DefaultUpdateAt = targetareaDescUpdateAt.Default.(func() time.Time)
	// targetarea.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	targetarea.UpdateDefaultUpdateAt = targetareaDescUpdateAt.UpdateDefault.(func() time.Time)
	// targetareaDescID is the schema descriptor for id field.
	targetareaDescID := targetareaFields[0].Descriptor()
	// targetarea.DefaultID holds the default value on creation for the id field.
	targetarea.DefaultID = targetareaDescID.Default.(func() uuid.UUID)
	vendorlocationFields := schema.VendorLocation{}.Fields()
	_ = vendorlocationFields
	// vendorlocationDescCountry is the schema descriptor for country field.
	vendorlocationDescCountry := vendorlocationFields[1].Descriptor()
	// vendorlocation.CountryValidator is a validator for the "country" field. It is called by the builders before save.
	vendorlocation.CountryValidator = vendorlocationDescCountry.Validators[0].(func(string) error)
	// vendorlocationDescProvince is the schema descriptor for province field.
	vendorlocationDescProvince := vendorlocationFields[2].Descriptor()
	// vendorlocation.ProvinceValidator is a validator for the "province" field. It is called by the builders before save.
	vendorlocation.ProvinceValidator = vendorlocationDescProvince.Validators[0].(func(string) error)
	// vendorlocationDescCity is the schema descriptor for city field.
	vendorlocationDescCity := vendorlocationFields[3].Descriptor()
	// vendorlocation.CityValidator is a validator for the "city" field. It is called by the builders before save.
	vendorlocation.CityValidator = vendorlocationDescCity.Validators[0].(func(string) error)
	// vendorlocationDescAddress is the schema descriptor for address field.
	vendorlocationDescAddress := vendorlocationFields[4].Descriptor()
	// vendorlocation.AddressValidator is a validator for the "address" field. It is called by the builders before save.
	vendorlocation.AddressValidator = vendorlocationDescAddress.Validators[0].(func(string) error)
	// vendorlocationDescCreateAt is the schema descriptor for create_at field.
	vendorlocationDescCreateAt := vendorlocationFields[5].Descriptor()
	// vendorlocation.DefaultCreateAt holds the default value on creation for the create_at field.
	vendorlocation.DefaultCreateAt = vendorlocationDescCreateAt.Default.(func() time.Time)
	// vendorlocationDescUpdateAt is the schema descriptor for update_at field.
	vendorlocationDescUpdateAt := vendorlocationFields[6].Descriptor()
	// vendorlocation.DefaultUpdateAt holds the default value on creation for the update_at field.
	vendorlocation.DefaultUpdateAt = vendorlocationDescUpdateAt.Default.(func() time.Time)
	// vendorlocation.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	vendorlocation.UpdateDefaultUpdateAt = vendorlocationDescUpdateAt.UpdateDefault.(func() time.Time)
	// vendorlocationDescID is the schema descriptor for id field.
	vendorlocationDescID := vendorlocationFields[0].Descriptor()
	// vendorlocation.DefaultID holds the default value on creation for the id field.
	vendorlocation.DefaultID = vendorlocationDescID.Default.(func() uuid.UUID)
}
