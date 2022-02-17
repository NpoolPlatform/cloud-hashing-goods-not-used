package appgood

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

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

func assertAppGood(t *testing.T, actual, expected *npool.AppGoodInfo) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.Online, expected.Online)
	assert.Equal(t, actual.InitAreaStrategy, expected.InitAreaStrategy)
	assert.Equal(t, actual.Price, expected.Price)
}

func TestAppGoodCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appGoodInfo := npool.AppGoodInfo{
		AppID:  uuid.New().String(),
		GoodID: uuid.New().String(),
		// For result assert
		Online:           false,
		InitAreaStrategy: "none",
		Price:            0,
	}
	resp, err := Authorize(context.Background(), &npool.AuthorizeAppGoodRequest{
		Info: &appGoodInfo,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertAppGood(t, resp.Info, &appGoodInfo)
	}

	resp1, err := Check(context.Background(), &npool.CheckAppGoodRequest{
		Info: &appGoodInfo,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertAppGood(t, resp1.Info, &appGoodInfo)
	}

	appGoodInfo.ID = resp.Info.ID

	resp2, err := Unauthorize(context.Background(), &npool.UnauthorizeAppGoodRequest{
		ID: resp1.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertAppGood(t, resp2.Info, &appGoodInfo)
	}

	resp3, err := Check(context.Background(), &npool.CheckAppGoodRequest{
		Info: &appGoodInfo,
	})
	assert.Nil(t, err)
	assert.Nil(t, resp3.Info)

	resp4, err := Authorize(context.Background(), &npool.AuthorizeAppGoodRequest{
		Info: &appGoodInfo,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp.Info.ID, resp4.Info.ID)
		assertAppGood(t, resp4.Info, &appGoodInfo)
	}

	resp5, err := Check(context.Background(), &npool.CheckAppGoodRequest{
		Info: &appGoodInfo,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp5.Info.ID, resp.Info.ID)
		assertAppGood(t, resp5.Info, &appGoodInfo)
	}

	appGoodInfo.Price = 0.13

	resp6, err := SetAppGoodPrice(context.Background(), &npool.SetAppGoodPriceRequest{
		Info: &appGoodInfo,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp6.Info.ID, resp.Info.ID)
		assertAppGood(t, resp6.Info, &appGoodInfo)
	}

	appGoodInfo.Online = true

	resp7, err := Onsale(context.Background(), &npool.OnsaleAppGoodRequest{
		Info: &appGoodInfo,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp7.Info.ID, resp.Info.ID)
		assertAppGood(t, resp7.Info, &appGoodInfo)
	}

	appGoodInfo.Online = false

	resp8, err := Offsale(context.Background(), &npool.OffsaleAppGoodRequest{
		Info: &appGoodInfo,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp8.Info.ID, resp.Info.ID)
		assertAppGood(t, resp8.Info, &appGoodInfo)
	}

	resp9, err := Unauthorize(context.Background(), &npool.UnauthorizeAppGoodRequest{
		ID: resp8.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp9.Info.ID, resp.Info.ID)
		assertAppGood(t, resp9.Info, &appGoodInfo)
	}
}
