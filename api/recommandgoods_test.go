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

func TestRecommendGoodsCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	Recommendgood := npool.RecommendGood{
		RecommenderID: uuid.New().String(),
		GoodID:        uuid.New().String(),
		Message:       "test recommend good",
	}
	firstCreateInfo := npool.CreateRecommendGoodResponse{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateRecommendGoodRequest{
			Info: &Recommendgood,
		}).
		Post("http://localhost:50020/v1/create/recommend/good")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.UUID{})
			assert.Equal(t, firstCreateInfo.Info.RecommenderID, Recommendgood.RecommenderID)
			assert.Equal(t, firstCreateInfo.Info.GoodID, Recommendgood.GoodID)
			assert.Equal(t, firstCreateInfo.Info.Message, Recommendgood.Message)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetRecommendGoodsRequest{
			PageInfo: &npool.PageInfo{},
		}).
		Post("http://localhost:50020/v1/get/recommend/goods")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		resp3 := &npool.GetRecommendGoodsResponse{}
		err = json.Unmarshal(resp.Body(), resp3)
		if assert.Nil(t, err) {
			assert.NotEmpty(t, resp3.Infos)
			assert.NotEmpty(t, resp3.Recommends)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetRecommendGoodsByRecommenderRequest{
			RecommenderID: Recommendgood.RecommenderID,
			PageInfo:      &npool.PageInfo{},
		}).
		Post("http://localhost:50020/v1/get/recommend/goods/by/recommender")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		resp4 := &npool.GetRecommendGoodsByRecommenderResponse{}
		err = json.Unmarshal(resp.Body(), resp4)
		if assert.Nil(t, err) {
			assert.NotEmpty(t, resp4.Infos)
			assert.NotEmpty(t, resp4.Recommends)
		}
	}
}
