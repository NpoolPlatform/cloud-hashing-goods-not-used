package goodrecommend

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

func assertGoodRecommend(t *testing.T, actual, expected *npool.GoodRecommend) {
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.Content, expected.Content)
}

func TestGoodRecommendCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	goodRecommend := npool.GoodRecommend{
		UserID:  uuid.New().String(),
		GoodID:  uuid.New().String(),
		Content: "test good recommend",
	}

	resp, err := Create(context.Background(), &npool.CreateGoodRecommendRequest{
		Info: &goodRecommend,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertGoodRecommend(t, resp.Info, &goodRecommend)
	}

	goodRecommend.ID = resp.Info.ID
	goodRecommend.Content += "."

	resp1, err := Update(context.Background(), &npool.UpdateGoodRecommendRequest{
		Info: &goodRecommend,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertGoodRecommend(t, resp1.Info, &goodRecommend)
	}

	_, err = GetByGood(context.Background(), &npool.GetGoodRecommendsByGoodRequest{
		GoodID: goodRecommend.GoodID,
	})
	assert.Nil(t, err)

	_, err = GetByUser(context.Background(), &npool.GetGoodRecommendsByUserRequest{
		UserID: goodRecommend.UserID,
	})
	assert.Nil(t, err)
}
