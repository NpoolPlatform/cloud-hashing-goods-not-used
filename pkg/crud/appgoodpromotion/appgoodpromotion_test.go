package appgoodpromotion

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/test-init" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func assertAppGoodPromotion(t *testing.T, actual, expected *npool.AppGoodPromotion) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.Start, expected.Start)
	assert.Equal(t, actual.End, expected.End)
	assert.Equal(t, actual.Message, expected.Message)
	assert.Equal(t, actual.Price, expected.Price)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	promotion := npool.AppGoodPromotion{
		AppID:   uuid.New().String(),
		GoodID:  uuid.New().String(),
		Start:   uint32(time.Now().Unix() + 1000),
		End:     uint32(time.Now().Unix() + 2000),
		Message: "sharing",
		Price:   100.0,
	}

	resp, err := Create(context.Background(), &npool.CreateAppGoodPromotionRequest{
		Info: &promotion,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertAppGoodPromotion(t, resp.Info, &promotion)
	}

	promotion.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateAppGoodPromotionRequest{
		Info: &promotion,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertAppGoodPromotion(t, resp1.Info, &promotion)
	}

	resp2, err := Get(context.Background(), &npool.GetAppGoodPromotionRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertAppGoodPromotion(t, resp2.Info, &promotion)
	}

	resp3, err := GetByApp(context.Background(), &npool.GetAppGoodPromotionsByAppRequest{
		AppID: resp.Info.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp3.Infos), 1)
	}

	resp4, err := GetByAppGoodStartEnd(context.Background(), &npool.GetAppGoodPromotionByAppGoodStartEndRequest{
		AppID:  resp.Info.AppID,
		GoodID: resp.Info.GoodID,
		Start:  resp.Info.Start,
		End:    resp.Info.End,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp4.Info.ID, resp.Info.ID)
		assertAppGoodPromotion(t, resp4.Info, &promotion)
	}
}
