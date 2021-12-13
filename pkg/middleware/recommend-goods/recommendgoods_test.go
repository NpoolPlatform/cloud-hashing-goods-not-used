package recommendgoods

import (
	"context"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	"github.com/stretchr/testify/assert"
)

func TestRecommendGoodsCRUD(t *testing.T) {
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
