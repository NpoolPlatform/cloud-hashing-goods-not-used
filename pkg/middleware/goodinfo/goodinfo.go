package goodinfo

import (
	"context"
	"sort"

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
	goodIndexes := map[string]uint32{}

	for _, info := range resp.Infos {
		allowed := false
		price := info.Price

		for _, appGood := range appGoods.Infos {
			if info.ID == appGood.GoodID && appGood.Visible {
				price = appGood.Price
				allowed = true
				goodIndexes[info.ID] = appGood.DisplayIndex
				break
			}
		}

		if !allowed {
			continue
		}

		info.Price = price
		infos = append(infos, info)
	}

	sort.Slice(infos, func(i, j int) bool {
		return goodIndexes[infos[i].ID] <= goodIndexes[infos[j].ID]
	})

	return &npool.GetGoodsByAppResponse{
		Infos: infos,
	}, nil
}
