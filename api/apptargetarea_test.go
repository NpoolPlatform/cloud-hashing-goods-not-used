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

func assertAppTargetAreaInfo(t *testing.T, actual, expected *npool.AppTargetAreaInfo) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.TargetAreaID, expected.TargetAreaID)
}

func TestAppTargetAreaCRUD(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appTargetArea := npool.AppTargetAreaInfo{
		AppID:        uuid.New().String(),
		TargetAreaID: uuid.New().String(),
	}
	firstCreateInfo := npool.AuthorizeAppTargetAreaRequest{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.AuthorizeAppTargetAreaRequest{
			Info: &appTargetArea,
		}).
		Post("http://localhost:50020/v1/authorize/app/target-area")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.UUID{})
			assertAppTargetAreaInfo(t, firstCreateInfo.Info, &appTargetArea)
		}
	}

	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CheckAppTargetAreaRequest{
			Info: &appTargetArea,
		}).
		Post("http://localhost:50020/v1/check/app/target-area")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp1.StatusCode())
		info := npool.CheckAppTargetAreaRequest{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assertAppTargetAreaInfo(t, info.Info, &appTargetArea)
		}
	}

	resp2, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UnauthorizeAppTargetAreaRequest{
			ID: firstCreateInfo.Info.ID,
		}).
		Post("http://localhost:50020/v1/unauthorize/app/target-area")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp2.StatusCode())
		info := npool.CheckAppTargetAreaRequest{}
		err := json.Unmarshal(resp2.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.Info.ID, uuid.UUID{})
			assertAppTargetAreaInfo(t, info.Info, &appTargetArea)
		}
	}
}
