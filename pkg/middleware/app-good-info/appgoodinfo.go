package appgoodinfo

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	appgood "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/app-good"
	goodinfo "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/good-info"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func Get(ctx context.Context, in *npool.GetGoodsByAppRequest) (*npool.GetGoodsByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	appGoods, err := appgood.GetAppGoodInfosByApp(ctx, &npool.GetAppGoodInfosByAppRequest{
		AppID: appID.String(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get app good: %v", err)
	}

	goodIDs := []string{}
	numStart, numEnd := in.PageInfo.PageIndex*in.PageInfo.PageSize, (in.PageInfo.PageIndex+1)*in.PageInfo.PageSize
	for i, v := range appGoods.Infos {
		if int32(i) < numStart || int32(i) > numEnd {
			continue
		}
		goodIDs = append(goodIDs, v.GoodID)
	}
	goodInfos, err := goodinfo.GetByIDs(ctx, &npool.GetGoodsByIDsRequest{
		IDs: goodIDs,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get good info: %v", err)
	}

	return &npool.GetGoodsByAppResponse{
		Infos: goodInfos.Infos,
		Total: int32(len(appGoods.Infos)),
	}, nil
}
