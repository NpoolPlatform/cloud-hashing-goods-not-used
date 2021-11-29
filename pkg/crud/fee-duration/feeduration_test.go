package feeduration

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

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	myFeeDuration := npool.FeeDuration{
		ID:        uuid.New().String(),
		FeeTypeID: uuid.New().String(),
		Duration:  30,
	}

	resp, err := Create(context.Background(), &npool.CreateFeeDurationRequest{
		Info: &myFeeDuration,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assert.Equal(t, resp.Info.FeeTypeID, myFeeDuration.FeeTypeID)
		assert.Equal(t, resp.Info.Duration, myFeeDuration.Duration)
	}

	resp1, err := Get(context.Background(), &npool.GetFeeDurationRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assert.Equal(t, resp.Info.FeeTypeID, myFeeDuration.FeeTypeID)
		assert.Equal(t, resp.Info.Duration, myFeeDuration.Duration)
	}
}
