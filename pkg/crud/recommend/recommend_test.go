package recommend

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

func assertRecommend(t *testing.T, actual, expected *npool.Recommend) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.RecommenderID, expected.RecommenderID)
	assert.Equal(t, actual.Message, expected.Message)
}

func TestCRUD(t *testing.T) {
	recommend := npool.Recommend{
		AppID:         uuid.New().String(),
		GoodID:        uuid.New().String(),
		RecommenderID: uuid.New().String(),
		Message:       "recommend test good",
	}

	resp, err := Create(context.Background(), &npool.CreateRecommendRequest{
		Info: &recommend,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertRecommend(t, resp.Info, &recommend)
	}

	recommend.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateRecommendRequest{
		Info: &recommend,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, recommend.ID)
		assertRecommend(t, resp1.Info, &recommend)
	}

	resp2, err := GetByApp(context.Background(), &npool.GetRecommendsByAppRequest{
		AppID: recommend.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp2.Infos), 1)
	}

	resp3, err := GetByRecommender(context.Background(), &npool.GetRecommendsByRecommenderRequest{
		RecommenderID: recommend.RecommenderID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp3.Infos), 1)
	}

	resp4, err := Delete(context.Background(), &npool.DeleteRecommendRequest{
		ID: recommend.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp4.Info.ID, recommend.ID)
		assertRecommend(t, resp4.Info, &recommend)
	}
}
