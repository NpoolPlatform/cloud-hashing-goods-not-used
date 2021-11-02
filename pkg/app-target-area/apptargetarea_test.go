package appareaauth

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

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
		AppID:        uuid.New(),
		TargetAreaID: uuid.New(),
	}

	resp, err := Authorize(context.Background(), &appTargetArea)
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertAppTargetAreaInfo(t, resp.Info, &appTargetArea)
	}
}
