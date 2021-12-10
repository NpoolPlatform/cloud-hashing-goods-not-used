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

func assertGoodRecommend(t *testing.T, actual, expected *npool.GoodRecommend) {
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.Content, expected.Content)
}

func TestCreateGoodRecommend(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	goodRecommend := npool.GoodRecommend{
		UserID:  uuid.New().String(),
		GoodID:  uuid.New().String(),
		Content: "test good recommend",
	}
	firstCreateInfo := npool.CreateGoodRecommendResponse{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateGoodRecommendRequest{
			Info: &goodRecommend,
		}).
		Post("http://localhost:50020/v1/create/good/recommend")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.UUID{})
			assertGoodRecommend(t, firstCreateInfo.Info, &goodRecommend)
		}
	}

	goodRecommend.ID = firstCreateInfo.Info.ID

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateGoodRecommendRequest{
			Info: &goodRecommend,
		}).
		Post("http://localhost:50020/v1/update/good/recommend")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := &npool.UpdateGoodRecommendResponse{}
		err := json.Unmarshal(resp.Body(), info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assertGoodRecommend(t, info.Info, &goodRecommend)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetGoodRecommendsByGoodRequest{
			GoodID: goodRecommend.GoodID,
		}).
		Post("http://localhost:50020/v1/get/good/recommends/by/good")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetGoodRecommendsByUserRequest{
			UserID:   goodRecommend.UserID,
			PageInfo: &npool.PageInfo{},
		}).
		Post("http://localhost:50020/v1/get/good/recommends/by/user")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
