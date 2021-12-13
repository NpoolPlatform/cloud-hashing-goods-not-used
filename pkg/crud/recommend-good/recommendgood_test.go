package recommendgood

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/test-init" //nolint

	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

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

	resp, err := Create(context.Background(), &npool.CreateRecommendGoodRequest{
		Info: &Recommendgood,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertRecommendGood(t, resp.Info, &Recommendgood)
	}

	Recommendgood.ID = resp.Info.ID
	Recommendgood.Message += "."

	resp1, err := Update(context.Background(), &npool.UpdateRecommendGoodRequest{
		Info: &Recommendgood,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertRecommendGood(t, resp1.Info, &Recommendgood)
	}

	_, err = GetByRecommender(context.Background(), &npool.GetRecommendGoodsByRecommenderRequest{
		RecommenderID: Recommendgood.RecommenderID,
		PageInfo:      &npool.PageInfo{},
	})
	assert.Nil(t, err)

	_, err = Get(context.Background(), &npool.GetRecommendGoodsRequest{
		PageInfo: &npool.PageInfo{},
	})
	assert.Nil(t, err)
}
