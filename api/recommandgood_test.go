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

func assertRecommendGood(t *testing.T, actual, expected *npool.RecommendGood) {
	assert.Equal(t, actual.RecommenderID, expected.RecommenderID)
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.Message, expected.Message)
}

func TestRecommendGoodCRUD(t *testing.T) {
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
			assertRecommendGood(t, firstCreateInfo.Info, &Recommendgood)
		}
	}

	Recommendgood.ID = firstCreateInfo.Info.ID

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateRecommendGoodRequest{
			Info: &Recommendgood,
		}).
		Post("http://localhost:50020/v1/update/recommend/good")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := &npool.UpdateRecommendGoodResponse{}
		err := json.Unmarshal(resp.Body(), info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assertRecommendGood(t, info.Info, &Recommendgood)
		}
	}
}
