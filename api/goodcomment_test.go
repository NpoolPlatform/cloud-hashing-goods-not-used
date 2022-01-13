package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"
)

func assertGoodComment(t *testing.T, actual, expected *npool.GoodComment) {
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.OrderID, expected.OrderID)
	assert.Equal(t, actual.Content, expected.Content)
}

func TestCreateGoodComment(t *testing.T) { //nolint
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
	firstCreateInfo := npool.CreateGoodCommentResponse{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateGoodCommentRequest{
			Comment: &goodComment,
		}).
		Post("http://localhost:50020/v1/create/good/comment")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Comment.ID, uuid.UUID{})
			assert.Equal(t, firstCreateInfo.Comment.ReplyToID, uuid.UUID{}.String())
			assertGoodComment(t, firstCreateInfo.Comment, &goodComment)
		}
	}

	goodComment.ID = firstCreateInfo.Comment.ID

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateGoodCommentRequest{
			Comment: &goodComment,
		}).
		Post("http://localhost:50020/v1/update/good/comment")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.UpdateGoodCommentResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Comment.ID, firstCreateInfo.Comment.ID)
			assertGoodComment(t, info.Comment, &goodComment)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetGoodCommentsRequest{
			GoodID: goodComment.GoodID,
		}).
		Post("http://localhost:50020/v1/get/good/comments")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
