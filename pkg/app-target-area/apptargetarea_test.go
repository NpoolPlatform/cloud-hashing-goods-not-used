package apptargetarea

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

func assertAppTargetAreaInfo(t *testing.T, actual, expected *npool.AppTargetAreaInfo) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.TargetAreaID, expected.TargetAreaID)
}

func TestAppTargetAreaCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appTargetArea := npool.AppTargetAreaInfo{
		AppID:        uuid.New().String(),
		TargetAreaID: uuid.New().String(),
	}

	resp, err := Authorize(context.Background(), &npool.AuthorizeAppTargetAreaRequest{
		Info: &appTargetArea,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertAppTargetAreaInfo(t, resp.Info, &appTargetArea)
	}

	resp1, err := Check(context.Background(), &npool.CheckAppTargetAreaRequest{
		Info: &appTargetArea,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp.Info.ID, resp1.Info.ID)
		assertAppTargetAreaInfo(t, resp1.Info, &appTargetArea)
		assert.True(t, resp1.Authorized)
	}

	resp2, err := Unauthorize(context.Background(), &npool.UnauthorizeAppTargetAreaRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp.Info.ID, resp2.Info.ID)
		assertAppTargetAreaInfo(t, resp2.Info, &appTargetArea)
		assert.False(t, resp2.Authorized)
	}

	resp3, err := Check(context.Background(), &npool.CheckAppTargetAreaRequest{
		Info: &appTargetArea,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, uuid.UUID{}.String())
		assertAppTargetAreaInfo(t, resp3.Info, &appTargetArea)
	}

	appTargetArea.ID = resp.Info.ID
	resp4, err := Authorize(context.Background(), &npool.AuthorizeAppTargetAreaRequest{
		Info: &appTargetArea,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp4.Info.ID, resp.Info.ID)
		assertAppTargetAreaInfo(t, resp4.Info, &appTargetArea)
	}

	resp5, err := Check(context.Background(), &npool.CheckAppTargetAreaRequest{
		Info: &appTargetArea,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp.Info.ID, resp5.Info.ID)
		assertAppTargetAreaInfo(t, resp5.Info, &appTargetArea)
		assert.True(t, resp5.Authorized)
	}

	resp6, err := Unauthorize(context.Background(), &npool.UnauthorizeAppTargetAreaRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp.Info.ID, resp6.Info.ID)
		assertAppTargetAreaInfo(t, resp6.Info, &appTargetArea)
		assert.False(t, resp6.Authorized)
	}
}
