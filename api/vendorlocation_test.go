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

func TestCreateVendorLocation(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	nano := time.Now().Unix()
	country := fmt.Sprintf("ChinaVendorLocationRestfulApiTest-%v", nano)
	province := fmt.Sprintf("ChinaVendorLocationRestfulApiTest-%v", nano)
	city := fmt.Sprintf("ChinaVendorLocationRestfulApiTest-%v", nano)
	address := fmt.Sprintf("ChinaVendorLocationRestfulApiTest-%v", nano)

	vendorLocationInfo := npool.VendorLocationInfo{
		Country:  country,
		Province: province,
		City:     city,
		Address:  address,
	}
	firstCreateInfo := npool.CreateVendorLocationResponse{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateVendorLocationRequest{
			Info: &vendorLocationInfo,
		}).
		Post("http://localhost:33759/v1/create/vendor-location")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.New())
			assert.Equal(t, firstCreateInfo.Info.Country, vendorLocationInfo.Country)
			assert.Equal(t, firstCreateInfo.Info.Province, vendorLocationInfo.Province)
			assert.Equal(t, firstCreateInfo.Info.City, vendorLocationInfo.City)
			assert.Equal(t, firstCreateInfo.Info.Address, vendorLocationInfo.Address)
		}
	}

	vendorLocationInfo.Province = fmt.Sprintf("ChinaVendorLocationRestfulApiTest-%v-1", nano)
	vendorLocationInfo.ID = firstCreateInfo.Info.ID

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateVendorLocationRequest{
			Info: &vendorLocationInfo,
		}).
		Post("http://localhost:33759/v1/update/vendor-location")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.UpdateVendorLocationResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assert.Equal(t, info.Info.Country, vendorLocationInfo.Country)
			assert.Equal(t, info.Info.Province, vendorLocationInfo.Province)
			assert.Equal(t, info.Info.City, vendorLocationInfo.City)
			assert.Equal(t, info.Info.Address, vendorLocationInfo.Address)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetVendorLocationRequest{
			ID: firstCreateInfo.Info.ID,
		}).
		Post("http://localhost:33759/v1/get/vendor-location")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetVendorLocationResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assert.Equal(t, info.Info.Country, vendorLocationInfo.Country)
			assert.Equal(t, info.Info.Province, vendorLocationInfo.Province)
			assert.Equal(t, info.Info.City, vendorLocationInfo.City)
			assert.Equal(t, info.Info.Address, vendorLocationInfo.Address)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetVendorLocationsRequest{}).
		Post("http://localhost:33759/v1/get/vendor-locations")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.DeleteVendorLocationRequest{
			ID: firstCreateInfo.Info.ID,
		}).
		Post("http://localhost:33759/v1/delete/vendor-location")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.DeleteVendorLocationResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assert.Equal(t, info.Info.Country, vendorLocationInfo.Country)
			assert.Equal(t, info.Info.Province, vendorLocationInfo.Province)
			assert.Equal(t, info.Info.City, vendorLocationInfo.City)
			assert.Equal(t, info.Info.Address, vendorLocationInfo.Address)
		}
	}
}
