package api

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
)

func TestCreateDeviceInfo(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	nano := time.Now().Unix()
	deviceType := fmt.Sprintf("ChinaDeviceInfoRestfulApiTest-%v", nano)
	manufacturer := fmt.Sprintf("ChinaDeviceInfoRestfulApiTest-%v", nano)
	var powerComsuption int32 = 120    //nolint
	var shipmentAt int32 = int32(nano) //nolint

	vendorLocationInfo := npool.DeviceInfo{
		Type:            deviceType,
		Manufacturer:    manufacturer,
		PowerComsuption: powerComsuption,
		ShipmentAt:      shipmentAt,
	}
	firstCreateInfo := npool.CreateDeviceInfoResponse{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateDeviceInfoRequest{
			Info: &vendorLocationInfo,
		}).
		Post("http://localhost:50020/v1/create/device")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.New())
			assert.Equal(t, firstCreateInfo.Info.Type, vendorLocationInfo.Type)
			assert.Equal(t, firstCreateInfo.Info.Manufacturer, vendorLocationInfo.Manufacturer)
			assert.Equal(t, firstCreateInfo.Info.PowerComsuption, vendorLocationInfo.PowerComsuption)
			assert.Equal(t, firstCreateInfo.Info.ShipmentAt, vendorLocationInfo.ShipmentAt)
		}
	}

	vendorLocationInfo.Manufacturer = fmt.Sprintf("ChinaDeviceInfoRestfulApiTest-%v-1", nano)
	vendorLocationInfo.ID = firstCreateInfo.Info.ID

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateDeviceInfoRequest{
			Info: &vendorLocationInfo,
		}).
		Post("http://localhost:50020/v1/update/device")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.UpdateDeviceInfoResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assert.Equal(t, info.Info.Type, vendorLocationInfo.Type)
			assert.Equal(t, info.Info.Manufacturer, vendorLocationInfo.Manufacturer)
			assert.Equal(t, info.Info.PowerComsuption, vendorLocationInfo.PowerComsuption)
			assert.Equal(t, info.Info.ShipmentAt, vendorLocationInfo.ShipmentAt)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetDeviceInfoRequest{
			ID: firstCreateInfo.Info.ID,
		}).
		Post("http://localhost:50020/v1/get/device")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetDeviceInfoResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assert.Equal(t, info.Info.Type, vendorLocationInfo.Type)
			assert.Equal(t, info.Info.Manufacturer, vendorLocationInfo.Manufacturer)
			assert.Equal(t, info.Info.PowerComsuption, vendorLocationInfo.PowerComsuption)
			assert.Equal(t, info.Info.ShipmentAt, vendorLocationInfo.ShipmentAt)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		Post("http://localhost:50020/v1/get/devices")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.DeleteDeviceInfoRequest{
			ID: firstCreateInfo.Info.ID,
		}).
		Post("http://localhost:50020/v1/delete/device")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.DeleteDeviceInfoResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assert.Equal(t, info.Info.Type, vendorLocationInfo.Type)
			assert.Equal(t, info.Info.Manufacturer, vendorLocationInfo.Manufacturer)
			assert.Equal(t, info.Info.PowerComsuption, vendorLocationInfo.PowerComsuption)
			assert.Equal(t, info.Info.ShipmentAt, vendorLocationInfo.ShipmentAt)
		}
	}
}
