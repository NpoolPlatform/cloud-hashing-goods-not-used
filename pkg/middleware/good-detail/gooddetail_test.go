package gooddetail

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/test-init" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/device-info"      //nolint
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/fee"              //nolint
	feetype "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/fee-type" //nolint
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/good-info"        //nolint
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/price-currency"   //nolint
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/vendor-location"  //nolint

	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func assertGoodDetail(t *testing.T, actual *npool.GoodDetail, expectGoodInfo *npool.GoodInfo, expectDevice *npool.DeviceInfo, expectVendorLocation *npool.VendorLocationInfo) {
	assert.Equal(t, actual.DeviceInfo.ID, expectDevice.ID)
	assert.Equal(t, actual.DeviceInfo.Type, expectDevice.Type)
	assert.Equal(t, actual.DeviceInfo.Manufacturer, expectDevice.Manufacturer)
	assert.Equal(t, actual.DeviceInfo.PowerComsuption, expectDevice.PowerComsuption)
	assert.Equal(t, actual.DeviceInfo.ShipmentAt, expectDevice.ShipmentAt)
	assert.Equal(t, actual.Good.SeparateFee, expectGoodInfo.SeparateFee)
	assert.Equal(t, actual.Good.UnitPower, expectGoodInfo.UnitPower)
	assert.Equal(t, actual.Good.DurationDays, expectGoodInfo.DurationDays)
	assert.Equal(t, actual.Good.CoinInfoID, expectGoodInfo.CoinInfoID)
	assert.Equal(t, actual.Good.Actuals, expectGoodInfo.Actuals)
	assert.Equal(t, actual.Good.DeliveryAt, expectGoodInfo.DeliveryAt)
	assert.Equal(t, actual.VendorLocation.ID, expectVendorLocation.ID)
	assert.Equal(t, actual.VendorLocation.Country, expectVendorLocation.Country)
	assert.Equal(t, actual.VendorLocation.Province, expectVendorLocation.Province)
	assert.Equal(t, actual.VendorLocation.City, expectVendorLocation.City)
	assert.Equal(t, actual.VendorLocation.Address, expectVendorLocation.Address)
	assert.Equal(t, actual.Good.Price, expectGoodInfo.Price)
	assert.Equal(t, actual.Good.BenefitType, expectGoodInfo.BenefitType)
	assert.Equal(t, actual.Good.Classic, expectGoodInfo.Classic)
	assert.Equal(t, actual.Good.SupportCoinTypeIDs, expectGoodInfo.SupportCoinTypeIDs)
}

func TestGet(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	nano := time.Now().UnixNano()

	deviceInfo := npool.DeviceInfo{
		Type:            fmt.Sprintf("S19-%v", nano),
		Manufacturer:    fmt.Sprintf("Ant-%v", nano),
		PowerComsuption: 120,
		ShipmentAt:      int32(time.Now().Unix()),
	}

	deviceResp, err := deviceinfo.Create(context.Background(), &npool.CreateDeviceInfoRequest{
		Info: &deviceInfo,
	})
	assert.Nil(t, err)

	vendorLocation := npool.VendorLocationInfo{
		Country:  fmt.Sprintf("China-%v", nano),
		Province: fmt.Sprintf("Shanghai-%v", nano),
		City:     fmt.Sprintf("Shanghai-%v", nano),
		Address:  fmt.Sprintf("Shanghai-%v", nano),
	}
	vendorLocationResp, err := vendorlocation.Create(context.Background(), &npool.CreateVendorLocationRequest{
		Info: &vendorLocation,
	})
	assert.Nil(t, err)

	currency := npool.PriceCurrency{
		Name:   fmt.Sprintf("USDT-%v", nano),
		Unit:   "USDT",
		Symbol: "$",
	}

	currencyResp, err := pricecurrency.Create(context.Background(), &npool.CreatePriceCurrencyRequest{
		Info: &currency,
	})
	assert.Nil(t, err)

	feeType1 := npool.FeeType{
		FeeType: fmt.Sprintf("Gas Fee-%v", nano),
		PayType: "amount",
	}
	feeTypeResp1, err := feetype.Create(context.Background(), &npool.CreateFeeTypeRequest{
		Info: &feeType1,
	})
	assert.Nil(t, err)

	fee1 := npool.Fee{
		AppID:     uuid.New().String(),
		FeeTypeID: feeTypeResp1.Info.ID,
		Value:     20,
	}
	feeResp1, err := fee.Create(context.Background(), &npool.CreateFeeRequest{
		Info: &fee1,
	})
	assert.Nil(t, err)

	feeType2 := npool.FeeType{
		FeeType: fmt.Sprintf("Gas Fee 1-%v", nano),
		PayType: "amount",
	}
	feeTypeResp2, err := feetype.Create(context.Background(), &npool.CreateFeeTypeRequest{
		Info: &feeType2,
	})
	assert.Nil(t, err)

	fee2 := npool.Fee{
		AppID:     fee1.AppID,
		FeeTypeID: feeTypeResp2.Info.ID,
		Value:     20,
	}
	feeResp2, err := fee.Create(context.Background(), &npool.CreateFeeRequest{
		Info: &fee2,
	})
	assert.Nil(t, err)

	goodInfo := npool.GoodInfo{
		Title:              "Ant Miner S19 Pro",
		Unit:               "TH/s",
		DeviceInfoID:       deviceResp.Info.ID,
		SeparateFee:        true,
		UnitPower:          10,
		DurationDays:       180,
		CoinInfoID:         uuid.New().String(),
		Actuals:            true,
		DeliveryAt:         uint32(time.Now().Unix()),
		VendorLocationID:   vendorLocationResp.Info.ID,
		InheritFromGoodID:  uuid.UUID{}.String(),
		Price:              13.0,
		PriceCurrency:      currencyResp.Info.ID,
		BenefitType:        "platform",
		Classic:            true,
		SupportCoinTypeIDs: []string{uuid.New().String(), uuid.New().String()},
		FeeIDs:             []string{feeResp1.Info.ID, feeResp2.Info.ID},
	}
	goodInfoResp, err := goodinfo.Create(context.Background(), &npool.CreateGoodRequest{
		Info: &goodInfo,
	})
	assert.Nil(t, err)

	resp, err := Get(context.Background(), &npool.GetGoodDetailRequest{
		ID: goodInfoResp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp.Info.Good.ID, goodInfoResp.Info.ID)
		assert.Nil(t, resp.Info.InheritFromGood)
		assertGoodDetail(t, resp.Info, goodInfoResp.Info, deviceResp.Info, vendorLocationResp.Info)
	}

	goodInfo.InheritFromGoodID = goodInfoResp.Info.ID
	goodInfoResp1, err := goodinfo.Create(context.Background(), &npool.CreateGoodRequest{
		Info: &goodInfo,
	})
	assert.Nil(t, err)

	resp1, err := Get(context.Background(), &npool.GetGoodDetailRequest{
		ID: goodInfoResp1.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.Good.ID, goodInfoResp1.Info.ID)
		assert.Equal(t, resp1.Info.InheritFromGood.ID, goodInfoResp.Info.ID)
		assertGoodDetail(t, resp1.Info, goodInfoResp.Info, deviceResp.Info, vendorLocationResp.Info)
	}

	resp2, err := GetAll(context.Background(), &npool.GetGoodsDetailRequest{})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp2.Infos)
	}
}
