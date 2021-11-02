package deviceinfo

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/test-init" //nolint

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestDeviceInfoCRUD(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	nano := time.Now().Unix()
	deviceType := fmt.Sprintf("ChinaDeviceInfoPackageTest-%v", nano)
	var powerComsuption int32 = 120
	manufacturer := fmt.Sprintf("ShanghaiDeviceInfoPackageTest-%v", nano)
	var shipmentAt int32 = int32(nano)

	resp, err := Create(context.Background(), &npool.CreateDeviceInfoRequest{
		Info: &npool.DeviceInfo{
			Type:            deviceType,
			PowerComsuption: powerComsuption,
			Manufacturer:    manufacturer,
			ShipmentAt:      shipmentAt,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp.Info.Type, deviceType)
		assert.Equal(t, resp.Info.PowerComsuption, powerComsuption)
		assert.Equal(t, resp.Info.Manufacturer, manufacturer)
		assert.Equal(t, resp.Info.ShipmentAt, shipmentAt)
	}

	powerComsuption = 130
	resp1, err := Update(context.Background(), &npool.UpdateDeviceInfoRequest{
		Info: &npool.DeviceInfo{
			ID:              resp.Info.ID,
			Type:            deviceType,
			PowerComsuption: powerComsuption,
			Manufacturer:    manufacturer,
			ShipmentAt:      shipmentAt,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.Type, deviceType)
		assert.Equal(t, resp1.Info.PowerComsuption, powerComsuption)
		assert.Equal(t, resp1.Info.Manufacturer, manufacturer)
		assert.Equal(t, resp1.Info.ShipmentAt, shipmentAt)
	}

	resp2, err := Get(context.Background(), &npool.GetDeviceInfoRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.Type, deviceType)
		assert.Equal(t, resp2.Info.PowerComsuption, powerComsuption)
		assert.Equal(t, resp2.Info.Manufacturer, manufacturer)
		assert.Equal(t, resp2.Info.ShipmentAt, shipmentAt)
	}

	resp3, err := Delete(context.Background(), &npool.DeleteDeviceInfoRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.Type, deviceType)
		assert.Equal(t, resp3.Info.PowerComsuption, powerComsuption)
		assert.Equal(t, resp3.Info.Manufacturer, manufacturer)
		assert.Equal(t, resp3.Info.ShipmentAt, shipmentAt)
	}
}

func TestGetAll(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	_, err := GetAll(context.Background(), &npool.GetDeviceInfosRequest{})
	assert.Nil(t, err)
}
