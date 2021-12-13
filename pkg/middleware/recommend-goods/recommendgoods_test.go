package recommendgoods

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	testinit "github.com/NpoolPlatform/cloud-hashing-goods/pkg/test-init" //nolint
	"github.com/stretchr/testify/assert"
)

func TestRecommendGoodsCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}

	resp1, err := Get(context.Background(), &npool.GetRecommendGoodsRequest{
		PageInfo: &npool.PageInfo{
			PageIndex: 0,
			PageSize:  10,
		},
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp1)
		assert.Equal(t, len(resp1.Infos), len(resp1.Recommends))
	}

	resp2, err := GetByRecommender(context.Background(), &npool.GetRecommendGoodsByRecommenderRequest{
		PageInfo: &npool.PageInfo{
			PageIndex: 0,
			PageSize:  10,
		},
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp2)
		assert.Equal(t, len(resp2.Infos), len(resp2.Recommends))
	}
}
