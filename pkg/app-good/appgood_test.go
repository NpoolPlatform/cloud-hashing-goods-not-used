package appgood

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/test-init" //nolint

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
	assert.Equal(t, actual.Authorized, expected.Authorized)
	assert.Equal(t, actual.Online, expected.Online)
	assert.Equal(t, actual.InitAreaStrategy, expected.InitAreaStrategy)
	assert.Equal(t, actual.Price, expected.Price)
	assert.Equal(t, actual.GasPrice, expected.GasPrice)
}

func TestAppGoodCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appGoodInfo := npool.AppGoodInfo{
		AppID:  uuid.New().String(),
		GoodID: uuid.New().String(),
		// For result assert
		Authorized:       true,
		Online:           false,
		InitAreaStrategy: "none",
		Price:            0,
		GasPrice:         0,
	}
	resp, err := Authorize(context.Background(), &npool.AuthorizeAppGoodRequest{
		Info: &appGoodInfo,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertAppGood(t, resp.Info, &appGoodInfo)
	}
}
