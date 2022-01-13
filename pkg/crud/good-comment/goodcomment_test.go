package goodcomment

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/test-init" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"

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

func assertGoodComment(t *testing.T, actual, expected *npool.GoodComment) {
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.OrderID, expected.OrderID)
	assert.Equal(t, actual.Content, expected.Content)
}

func TestGoodCommentCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	goodComment := npool.GoodComment{
		UserID:  uuid.New().String(),
		AppID:   uuid.New().String(),
		GoodID:  uuid.New().String(),
		OrderID: uuid.New().String(),
		Content: "test good comment",
	}

	resp, err := Create(context.Background(), &npool.CreateGoodCommentRequest{
		Comment: &goodComment,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Comment.ID, uuid.UUID{})
		assertGoodComment(t, resp.Comment, &goodComment)
	}

	goodComment.ReplyToID = resp.Comment.ID
	resp, err = Create(context.Background(), &npool.CreateGoodCommentRequest{
		Comment: &goodComment,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Comment.ID, uuid.UUID{})
		assertGoodComment(t, resp.Comment, &goodComment)
	}

	goodComment.ID = resp.Comment.ID

	resp1, err := Update(context.Background(), &npool.UpdateGoodCommentRequest{
		Comment: &goodComment,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Comment.ID, resp.Comment.ID)
		assertGoodComment(t, resp1.Comment, &goodComment)
	}

	_, err = GetAll(context.Background(), &npool.GetGoodCommentsRequest{
		GoodID: goodComment.GoodID,
	})
	assert.Nil(t, err)
}
