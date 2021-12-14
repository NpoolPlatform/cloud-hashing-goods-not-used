package appgoodinfo

import (
	"context"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	appgood "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/app-good"
	goodinfo "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/good-info"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	// get existing good info
	appID := uuid.New()
	ctx := context.Background()
	respGoods, err := goodinfo.GetAll(ctx, &npool.GetGoodsRequest{
		PageInfo: &npool.PageInfo{
			PageIndex: 0,
			PageSize:  10,
		},
	})
	if err != nil {
		logger.Sugar().Fatal("insert goods first")
	}
	assert.Positive(t, len(respGoods.Infos))

	// create appgood
	_, err = appgood.Authorize(ctx, &npool.AuthorizeAppGoodRequest{
		Info: &npool.AppGoodInfo{
			AppID:            appID.String(),
			GoodID:           respGoods.Infos[0].ID,
			Price:            respGoods.Infos[0].Price,
			Authorized:       true,
			Online:           true,
			InitAreaStrategy: "all",
		},
	})
	assert.Nil(t, err)

	// test get
	respGet, err := Get(ctx, &npool.GetGoodsByAppRequest{
		AppID: appID.String(),
		PageInfo: &npool.PageInfo{
			PageIndex: 0,
			PageSize:  10,
		},
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, respGet.Infos)
		assert.Positive(t, respGet.Total)
		assert.NotEmpty(t, respGet.Infos[0].ID)
	}
}
