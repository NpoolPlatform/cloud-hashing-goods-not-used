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

func TestTargetAreaCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	targetAreaInfo := npool.TargetAreaInfo{
		Continent: "Asia",
		Country:   "China",
	}
	firstCreateInfo := npool.CreateTargetAreaResponse{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateTargetAreaRequest{
			Info: &targetAreaInfo,
		}).
		Post("http://localhost:33759/v1/create/target-area")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.New())
			assert.Equal(t, firstCreateInfo.Info.Continent, targetAreaInfo.Continent)
			assert.Equal(t, firstCreateInfo.Info.Country, targetAreaInfo.Country)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateTargetAreaRequest{
			Info: &targetAreaInfo,
		}).
		Post("http://localhost:33759/v1/create/target-area")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.CreateTargetAreaResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, firstCreateInfo.Info.ID, info.Info.ID)
			assert.Equal(t, info.Info.Continent, targetAreaInfo.Continent)
			assert.Equal(t, info.Info.Country, targetAreaInfo.Country)
		}
	}

	targetAreaInfo.Country = "Chinaa"
	targetAreaInfo.ID = firstCreateInfo.Info.ID
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateTargetAreaRequest{
			Info: &targetAreaInfo,
		}).
		Post("http://localhost:33759/v1/update/target-area")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.CreateTargetAreaResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			if assert.NotNil(t, info.Info) {
				assert.Equal(t, firstCreateInfo.Info.ID, info.Info.ID)
				assert.Equal(t, info.Info.Continent, targetAreaInfo.Continent)
				assert.Equal(t, info.Info.Country, targetAreaInfo.Country)
			}
		}
	}

	targetAreaInfo.Country = "China"
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateTargetAreaRequest{
			Info: &targetAreaInfo,
		}).
		Post("http://localhost:33759/v1/update/target-area")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.CreateTargetAreaResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			if assert.NotNil(t, info.Info) {
				assert.Equal(t, firstCreateInfo.Info.ID, info.Info.ID)
				assert.Equal(t, info.Info.Continent, targetAreaInfo.Continent)
				assert.Equal(t, info.Info.Country, targetAreaInfo.Country)
			}
		}
	}

	_, err = cli.R().
		Get("http://localhost:33759/v1/get/target-areas")
	assert.Nil(t, err)
}
