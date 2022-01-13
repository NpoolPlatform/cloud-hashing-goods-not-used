package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"
)

func assertAppGoodInfo(t *testing.T, actual, expected *npool.AppGoodInfo) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.Authorized, expected.Authorized)
	assert.Equal(t, actual.Online, expected.Online)
	assert.Equal(t, actual.InitAreaStrategy, expected.InitAreaStrategy)
	assert.Equal(t, actual.Price, expected.Price)
}

func TestAppGoodCRUD(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appGood := npool.AppGoodInfo{
		AppID:  uuid.New().String(),
		GoodID: uuid.New().String(),
		// For result assert
		Authorized:       true,
		Online:           false,
		InitAreaStrategy: "none",
		Price:            0,
	}
	firstCreateInfo := npool.AuthorizeAppGoodRequest{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.AuthorizeAppGoodRequest{
			Info: &appGood,
		}).
		Post("http://localhost:50020/v1/authorize/app/good")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.UUID{})
			assertAppGoodInfo(t, firstCreateInfo.Info, &appGood)
		}
	}

	appGood.ID = firstCreateInfo.Info.ID

	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CheckAppGoodRequest{
			Info: &appGood,
		}).
		Post("http://localhost:50020/v1/check/app/good")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp1.StatusCode())
		info := npool.CheckAppGoodRequest{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assertAppGoodInfo(t, info.Info, &appGood)
		}
	}

	appGood.Price = 0.13

	resp3, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.SetAppGoodPriceRequest{
			Info: &appGood,
		}).
		Post("http://localhost:50020/v1/set/app/good/price")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp3.StatusCode())
		info := npool.SetAppGoodPriceRequest{}
		err := json.Unmarshal(resp3.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.Info.ID, uuid.UUID{})
			assertAppGoodInfo(t, info.Info, &appGood)
		}
	}

	appGood.Online = true

	resp4, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.OnsaleAppGoodRequest{
			Info: &appGood,
		}).
		Post("http://localhost:50020/v1/onsale/app/good")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp4.StatusCode())
		info := npool.OnsaleAppGoodRequest{}
		err := json.Unmarshal(resp4.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.Info.ID, uuid.UUID{})
			assertAppGoodInfo(t, info.Info, &appGood)
		}
	}

	appGood.Online = false

	resp5, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.OffsaleAppGoodRequest{
			Info: &appGood,
		}).
		Post("http://localhost:50020/v1/offsale/app/good")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp5.StatusCode())
		info := npool.OffsaleAppGoodRequest{}
		err := json.Unmarshal(resp5.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.Info.ID, uuid.UUID{})
			assertAppGoodInfo(t, info.Info, &appGood)
		}
	}

	appGood.Authorized = false

	resp6, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UnauthorizeAppGoodRequest{
			Info: &appGood,
		}).
		Post("http://localhost:50020/v1/unauthorize/app/good")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp6.StatusCode())
		info := npool.UnauthorizeAppGoodRequest{}
		err := json.Unmarshal(resp6.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.Info.ID, uuid.UUID{})
			assertAppGoodInfo(t, info.Info, &appGood)
		}
	}
}
