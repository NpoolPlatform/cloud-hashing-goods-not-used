package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
)

func assertAppGoodTargetAreaInfo(t *testing.T, actual, expected *npool.AppGoodTargetAreaInfo) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.TargetAreaID, expected.TargetAreaID)
}

func TestAppGoodTargetAreaCRUD(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appGoodTargetArea := npool.AppGoodTargetAreaInfo{
		AppID:        uuid.New().String(),
		GoodID:       uuid.New().String(),
		TargetAreaID: uuid.New().String(),
	}
	firstCreateInfo := npool.AuthorizeAppGoodTargetAreaRequest{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.AuthorizeAppGoodTargetAreaRequest{
			Info: &appGoodTargetArea,
		}).
		Post("http://localhost:33759/v1/authorize/app/good/target-area")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.UUID{})
			assertAppGoodTargetAreaInfo(t, firstCreateInfo.Info, &appGoodTargetArea)
		}
	}

	appGoodTargetArea.ID = firstCreateInfo.Info.ID

	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CheckAppGoodTargetAreaRequest{
			Info: &appGoodTargetArea,
		}).
		Post("http://localhost:33759/v1/check/app/good/target-area")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp1.StatusCode())
		info := npool.CheckAppGoodTargetAreaRequest{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assertAppGoodTargetAreaInfo(t, info.Info, &appGoodTargetArea)
		}
	}

	resp2, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UnauthorizeAppGoodTargetAreaRequest{
			Info: &appGoodTargetArea,
		}).
		Post("http://localhost:33759/v1/unauthorize/app/good/target-area")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp2.StatusCode())
		info := npool.UnauthorizeAppGoodTargetAreaRequest{}
		err := json.Unmarshal(resp2.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assertAppGoodTargetAreaInfo(t, info.Info, &appGoodTargetArea)
		}
	}
}
