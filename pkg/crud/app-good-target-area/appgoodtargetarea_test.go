package appgoodtargetarea

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

func assertAppGoodTargetArea(t *testing.T, actual, expected *npool.AppGoodTargetAreaInfo) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.TargetAreaID, expected.TargetAreaID)
}

func TestAppGoodTargetAreaCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appGoodInfo := npool.AppGoodTargetAreaInfo{
		AppID:        uuid.New().String(),
		GoodID:       uuid.New().String(),
		TargetAreaID: uuid.New().String(),
	}
	resp, err := Authorize(context.Background(), &npool.AuthorizeAppGoodTargetAreaRequest{
		Info: &appGoodInfo,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertAppGoodTargetArea(t, resp.Info, &appGoodInfo)
	}

	resp1, err := Check(context.Background(), &npool.CheckAppGoodTargetAreaRequest{
		Info: &appGoodInfo,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertAppGoodTargetArea(t, resp1.Info, &appGoodInfo)
	}

	appGoodInfo.ID = resp.Info.ID

	resp2, err := Unauthorize(context.Background(), &npool.UnauthorizeAppGoodTargetAreaRequest{
		Info: &appGoodInfo,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertAppGoodTargetArea(t, resp2.Info, &appGoodInfo)
	}

	_, err = Check(context.Background(), &npool.CheckAppGoodTargetAreaRequest{
		Info: &appGoodInfo,
	})
	assert.NotNil(t, err)
}
