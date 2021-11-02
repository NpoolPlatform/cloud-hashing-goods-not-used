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

func TestCreateTargetArea(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	nano := time.Now().Unix()
	continent := fmt.Sprintf("AsiaTargetAreaRestfulApiTest-%v", nano)
	country := fmt.Sprintf("ChinaTargetAreaRestfulApiTest-%v", nano)
	country1 := fmt.Sprintf("ChinaaaaTargetAreaRestfulApiTest-%v", nano)

	targetAreaInfo := npool.TargetAreaInfo{
		Continent: continent,
		Country:   country,
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
	assert.Nil(t, err)
	assert.NotEqual(t, 200, resp.StatusCode())

	targetAreaInfo.Country = country1
	targetAreaInfo.ID = firstCreateInfo.Info.ID
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateTargetAreaRequest{
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

	targetAreaInfo.Country = country
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

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.DeleteTargetAreaRequest{
			ID: firstCreateInfo.Info.ID,
		}).
		Post("http://localhost:33759/v1/delete/target-area")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.DeleteTargetAreaResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			if assert.NotNil(t, info.Info) {
				assert.Equal(t, firstCreateInfo.Info.ID, info.Info.ID)
				assert.Equal(t, info.Info.Continent, targetAreaInfo.Continent)
				assert.Equal(t, info.Info.Country, targetAreaInfo.Country)
			}
		}
	}

	nano = time.Now().Unix() + 1
	continent = fmt.Sprintf("AsiaTargetAreaRestfulApiTest-%v", nano)
	country = fmt.Sprintf("ChinaTargetAreaRestfulApiTest-%v", nano)

	targetAreaInfo = npool.TargetAreaInfo{
		Continent: continent,
		Country:   country,
	}

	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateTargetAreaRequest{
			Info: &targetAreaInfo,
		}).
		Post("http://localhost:33759/v1/create/target-area")
	assert.Nil(t, err)
	info1 := npool.CreateTargetAreaResponse{}
	err = json.Unmarshal(resp1.Body(), &info1)
	assert.Nil(t, err)

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetTargetAreaRequest{
			ID: info1.Info.ID,
		}).
		Post("http://localhost:33759/v1/get/target-area")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetTargetAreaResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			if assert.NotNil(t, info.Info) {
				assert.Equal(t, info1.Info.ID, info.Info.ID)
				assert.Equal(t, info.Info.Continent, targetAreaInfo.Continent)
				assert.Equal(t, info.Info.Country, targetAreaInfo.Country)
			}
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.DeleteTargetAreaByContinentCountryRequest{
			Continent: targetAreaInfo.Continent,
			Country:   targetAreaInfo.Country,
		}).
		Post("http://localhost:33759/v1/delete/target-area/continent/country")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.DeleteTargetAreaResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			if assert.NotNil(t, info.Info) {
				assert.Equal(t, info1.Info.ID, info.Info.ID)
				assert.Equal(t, info.Info.Continent, targetAreaInfo.Continent)
				assert.Equal(t, info.Info.Country, targetAreaInfo.Country)
			}
		}
	}
}

func TestGetTargetAreas(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()
	_, err := cli.R().
		Get("http://localhost:33759/v1/get/target-areas")
	assert.Nil(t, err)
}
