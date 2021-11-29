// Code generated by entc, DO NOT EDIT.

package goodinfo

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the goodinfo type in the database.
	Label = "good_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDeviceInfoID holds the string denoting the device_info_id field in the database.
	FieldDeviceInfoID = "device_info_id"
	// FieldSeparateFee holds the string denoting the separate_fee field in the database.
	FieldSeparateFee = "separate_fee"
	// FieldUnitPower holds the string denoting the unit_power field in the database.
	FieldUnitPower = "unit_power"
	// FieldDurationDays holds the string denoting the duration_days field in the database.
	FieldDurationDays = "duration_days"
	// FieldCoinInfoID holds the string denoting the coin_info_id field in the database.
	FieldCoinInfoID = "coin_info_id"
	// FieldActuals holds the string denoting the actuals field in the database.
	FieldActuals = "actuals"
	// FieldDeliveryAt holds the string denoting the delivery_at field in the database.
	FieldDeliveryAt = "delivery_at"
	// FieldInheritFromGoodID holds the string denoting the inherit_from_good_id field in the database.
	FieldInheritFromGoodID = "inherit_from_good_id"
	// FieldVendorLocationID holds the string denoting the vendor_location_id field in the database.
	FieldVendorLocationID = "vendor_location_id"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldPriceCurrency holds the string denoting the price_currency field in the database.
	FieldPriceCurrency = "price_currency"
	// FieldBenefitType holds the string denoting the benefit_type field in the database.
	FieldBenefitType = "benefit_type"
	// FieldClassic holds the string denoting the classic field in the database.
	FieldClassic = "classic"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldUnit holds the string denoting the unit field in the database.
	FieldUnit = "unit"
	// FieldFeeIds holds the string denoting the fee_ids field in the database.
	FieldFeeIds = "fee_ids"
	// FieldSupportCoinTypeIds holds the string denoting the support_coin_type_ids field in the database.
	FieldSupportCoinTypeIds = "support_coin_type_ids"
	// FieldTotal holds the string denoting the total field in the database.
	FieldTotal = "total"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the goodinfo in the database.
	Table = "good_infos"
)

// Columns holds all SQL columns for goodinfo fields.
var Columns = []string{
	FieldID,
	FieldDeviceInfoID,
	FieldSeparateFee,
	FieldUnitPower,
	FieldDurationDays,
	FieldCoinInfoID,
	FieldActuals,
	FieldDeliveryAt,
	FieldInheritFromGoodID,
	FieldVendorLocationID,
	FieldPrice,
	FieldPriceCurrency,
	FieldBenefitType,
	FieldClassic,
	FieldTitle,
	FieldUnit,
	FieldFeeIds,
	FieldSupportCoinTypeIds,
	FieldTotal,
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
	// UnitPowerValidator is a validator for the "unit_power" field. It is called by the builders before save.
	UnitPowerValidator func(int32) error
	// DurationDaysValidator is a validator for the "duration_days" field. It is called by the builders before save.
	DurationDaysValidator func(int32) error
	// PriceValidator is a validator for the "price" field. It is called by the builders before save.
	PriceValidator func(uint64) error
	// TotalValidator is a validator for the "total" field. It is called by the builders before save.
	TotalValidator func(int32) error
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() uint32
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// BenefitType defines the type for the "benefit_type" enum field.
type BenefitType string

// BenefitType values.
const (
	BenefitTypePool     BenefitType = "pool"
	BenefitTypePlatform BenefitType = "platform"
)

func (bt BenefitType) String() string {
	return string(bt)
}

// BenefitTypeValidator is a validator for the "benefit_type" field enum values. It is called by the builders before save.
func BenefitTypeValidator(bt BenefitType) error {
	switch bt {
	case BenefitTypePool, BenefitTypePlatform:
		return nil
	default:
		return fmt.Errorf("goodinfo: invalid enum value for benefit_type field: %q", bt)
	}
}
