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

func assertGoodReview(t *testing.T, actual, expected *npool.GoodReviewInfo) {
	assert.Equal(t, actual.Type, expected.Type)
	assert.Equal(t, actual.ReviewerID, expected.ReviewerID)
	assert.Equal(t, actual.ReviewedID, expected.ReviewedID)
	assert.Equal(t, actual.State, expected.State)
	assert.Equal(t, actual.Message, expected.Message)
}

func TestCreateGoodReview(t *testing.T) { //nolint
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
	firstCreateInfo := npool.CreateGoodReviewResponse{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateGoodReviewRequest{
			Info: &goodReview,
		}).
		Post("http://localhost:50020/v1/create/good/review")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.New())
			assertGoodReview(t, firstCreateInfo.Info, &goodReview)
		}
	}

	goodReview.ID = firstCreateInfo.Info.ID

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateGoodReviewRequest{
			Info: &goodReview,
		}).
		Post("http://localhost:50020/v1/update/good/review")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.UpdateGoodReviewResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assertGoodReview(t, info.Info, &goodReview)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetGoodReviewRequest{
			Info: &goodReview,
		}).
		Post("http://localhost:50020/v1/get/good/review")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetGoodReviewResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assertGoodReview(t, info.Info, &goodReview)
		}
	}
}
