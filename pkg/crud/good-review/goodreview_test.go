package goodreview

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

func assertGoodReview(t *testing.T, actual, expected *npool.GoodReviewInfo) {
	assert.Equal(t, actual.Type, expected.Type)
	assert.Equal(t, actual.ReviewerID, expected.ReviewerID)
	assert.Equal(t, actual.ReviewedID, expected.ReviewedID)
	assert.Equal(t, actual.State, expected.State)
	assert.Equal(t, actual.Message, expected.Message)
}

func TestGoodReviewCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	goodReview := npool.GoodReviewInfo{
		Type:       "good",
		ReviewerID: uuid.New().String(),
		ReviewedID: uuid.New().String(),
		State:      "none",
		Message:    "for review test",
	}

	resp, err := Create(context.Background(), &npool.CreateGoodReviewRequest{
		Info: &goodReview,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertGoodReview(t, resp.Info, &goodReview)
	}

	goodReview.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateGoodReviewRequest{
		Info: &goodReview,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertGoodReview(t, resp1.Info, &goodReview)
	}

	resp2, err := Get(context.Background(), &npool.GetGoodReviewRequest{
		Info: &goodReview,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertGoodReview(t, resp2.Info, &goodReview)
	}
}
