package vendorlocation

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

func TestVendorLocationCRUD(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	nano := time.Now().UnixNano()
	country := fmt.Sprintf("ChinaVendorLocationPackageTest-%v", nano)
	province := fmt.Sprintf("ShanghaiVendorLocationPackageTest-%v", nano)
	city := fmt.Sprintf("ShanghaiVendorLocationPackageTest-%v", nano)
	address := fmt.Sprintf("YangpuVendorLocationPakcageTest-%v", nano)

	resp, err := Create(context.Background(), &npool.CreateVendorLocationRequest{
		Info: &npool.VendorLocationInfo{
			Country:  country,
			Province: province,
			City:     city,
			Address:  address,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp.Info.Country, country)
		assert.Equal(t, resp.Info.Province, province)
		assert.Equal(t, resp.Info.City, city)
		assert.Equal(t, resp.Info.Address, address)
	}

	province = fmt.Sprintf("ShanghaiVendorLocationPackageTest-%v-1", nano)
	resp1, err := Update(context.Background(), &npool.UpdateVendorLocationRequest{
		Info: &npool.VendorLocationInfo{
			ID:       resp.Info.ID,
			Country:  country,
			Province: province,
			City:     city,
			Address:  address,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.Country, country)
		assert.Equal(t, resp1.Info.Province, province)
		assert.Equal(t, resp1.Info.City, city)
		assert.Equal(t, resp1.Info.Address, address)
	}

	resp2, err := Get(context.Background(), &npool.GetVendorLocationRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.Country, country)
		assert.Equal(t, resp2.Info.Province, province)
		assert.Equal(t, resp2.Info.City, city)
		assert.Equal(t, resp2.Info.Address, address)
	}

	resp3, err := Delete(context.Background(), &npool.DeleteVendorLocationRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.Country, country)
		assert.Equal(t, resp3.Info.Province, province)
		assert.Equal(t, resp3.Info.City, city)
		assert.Equal(t, resp3.Info.Address, address)
	}
}
