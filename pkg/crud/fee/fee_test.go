package fee

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"
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

	myFee := npool.Fee{
		AppID:     uuid.New().String(),
		FeeTypeID: uuid.New().String(),
		Value:     10,
	}

	resp, err := Create(context.Background(), &npool.CreateFeeRequest{
		Info: &myFee,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assert.Equal(t, resp.Info.AppID, myFee.AppID)
		assert.Equal(t, resp.Info.FeeTypeID, myFee.FeeTypeID)
		assert.Equal(t, resp.Info.Value, myFee.Value)
	}

	resp1, err := Get(context.Background(), &npool.GetFeeRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assert.Equal(t, resp.Info.AppID, myFee.AppID)
		assert.Equal(t, resp.Info.FeeTypeID, myFee.FeeTypeID)
		assert.Equal(t, resp.Info.Value, myFee.Value)
	}
}
