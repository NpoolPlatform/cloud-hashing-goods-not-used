package goodinfo

import (
	"context"

	appgoodcrud "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/app-good"
	crud "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/good-info"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"

	"golang.org/x/xerrors"
)

func GetByApp(ctx context.Context, in *npool.GetGoodsByAppRequest) (*npool.GetGoodsByAppResponse, error) {
	resp, err := crud.GetAll(ctx, &npool.GetGoodsRequest{
		PageInfo: in.GetPageInfo(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get goods info: %v", err)
	}

	appGoods, err := appgoodcrud.GetByApp(ctx, &npool.GetAppGoodsRequest{
		AppID: in.GetAppID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get app good: %v", err)
	}

	infos := []*npool.GoodInfo{}
	for _, info := range resp.Infos {
		allowed := false
		for _, appGood := range appGoods.Infos {
			if info.ID == appGood.GoodID {
				allowed = true
				break
			}
		}

		if !allowed {
			continue
		}

		infos = append(infos, info)
	}

	return &npool.GetGoodsByAppResponse{
		Infos: infos,
	}, nil
}
