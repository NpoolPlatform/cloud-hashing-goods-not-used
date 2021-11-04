package gooddetail

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/test-init" //nolint

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/device-info"     //nolint
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/good-info"       //nolint
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/vendor-location" //nolint

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

func TestGet(t *testing.T) {
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

	goodInfo := npool.GoodInfo{
		DeviceInfoID:       deviceResp.Info.ID,
		GasPrice:           0.13,
		SeparateGasFee:     true,
		UnitPower:          10,
		DurationDays:       180,
		CoinInfoID:         uuid.New().String(),
		Actuals:            true,
		DeliveryAt:         int32(time.Now().Unix()),
		VendorLocationID:   vendorLocationResp.Info.ID,
		InheritFromGoodID:  uuid.UUID{}.String(),
		Price:              13.0,
		BenefitType:        "platform",
		Classic:            true,
		SupportCoinTypeIDs: []string{uuid.New().String(), uuid.New().String()},
		Total:              100,
	}
	goodInfoResp, err := goodinfo.Create(context.Background(), &npool.CreateGoodRequest{
		Info: &goodInfo,
	})
	assert.Nil(t, err)

	resp, err := Get(context.Background(), &npool.GetGoodDetailRequest{
		ID: goodInfoResp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp.Detail.ID, goodInfoResp.Info.ID)
		assert.Equal(t, resp.Detail.DeviceInfo.ID, deviceResp.Info.ID)
		assert.Equal(t, resp.Detail.DeviceInfo.Type, deviceResp.Info.Type)
		assert.Equal(t, resp.Detail.DeviceInfo.Manufacturer, deviceResp.Info.Manufacturer)
		assert.Equal(t, resp.Detail.DeviceInfo.PowerComsuption, deviceResp.Info.PowerComsuption)
		assert.Equal(t, resp.Detail.DeviceInfo.ShipmentAt, deviceResp.Info.ShipmentAt)
		assert.Equal(t, resp.Detail.GasPrice, goodInfoResp.Info.GasPrice)
		assert.Equal(t, resp.Detail.SeparateGasFee, goodInfoResp.Info.SeparateGasFee)
		assert.Equal(t, resp.Detail.UnitPower, goodInfoResp.Info.UnitPower)
		assert.Equal(t, resp.Detail.DurationDays, goodInfoResp.Info.DurationDays)
		assert.Equal(t, resp.Detail.CoinInfoID, goodInfoResp.Info.CoinInfoID)
		assert.Equal(t, resp.Detail.Actuals, goodInfoResp.Info.Actuals)
		assert.Equal(t, resp.Detail.DeliveryAt, goodInfoResp.Info.DeliveryAt)
		assert.Nil(t, resp.Detail.InheritFromGood)
		assert.Equal(t, resp.Detail.VendorLocation.ID, vendorLocationResp.Info.ID)
		assert.Equal(t, resp.Detail.VendorLocation.Country, vendorLocationResp.Info.Country)
		assert.Equal(t, resp.Detail.VendorLocation.Province, vendorLocationResp.Info.Province)
		assert.Equal(t, resp.Detail.VendorLocation.City, vendorLocationResp.Info.City)
		assert.Equal(t, resp.Detail.VendorLocation.Address, vendorLocationResp.Info.Address)
		assert.Equal(t, resp.Detail.Price, goodInfoResp.Info.Price)
		assert.Equal(t, resp.Detail.BenefitType, goodInfoResp.Info.BenefitType)
		assert.Equal(t, resp.Detail.Classic, goodInfoResp.Info.Classic)
		assert.Equal(t, resp.Detail.SupportCoinTypeIDs, goodInfoResp.Info.SupportCoinTypeIDs)
		assert.Equal(t, resp.Detail.Total, goodInfoResp.Info.Total)
	}
}
