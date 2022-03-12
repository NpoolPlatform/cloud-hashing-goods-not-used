package recommend

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"

	appgoodcrud "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/app-good"
	crud "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/recommend"
	good "github.com/NpoolPlatform/cloud-hashing-goods/pkg/middleware/good-detail"

	"golang.org/x/xerrors"
)

func GetByApp(ctx context.Context, in *npool.GetRecommendGoodsByAppRequest) (*npool.GetRecommendGoodsByAppResponse, error) {
	myRecommends, err := crud.GetByApp(ctx, &npool.GetRecommendsByAppRequest{
		AppID: in.GetAppID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get recommends by app: %v", err)
	}

	appGoods, err := appgoodcrud.GetByApp(ctx, &npool.GetAppGoodsRequest{
		AppID: in.GetAppID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get app good: %v", err)
	}

	recommendGoods := []*npool.RecommendGood{}
	for _, myRecommend := range myRecommends.Infos {
		allowed := false
		price := 0.0

		for _, info := range appGoods.Infos {
			if myRecommend.GoodID == info.GoodID {
				price = info.Price
				allowed = true
				break
			}
		}

		if !allowed {
			continue
		}

		myGood, err := good.Get(ctx, &npool.GetGoodDetailRequest{
			ID: myRecommend.GoodID,
		})
		if err != nil {
			logger.Sugar().Errorf("fail get good: %v", err)
		}

		myGood.Info.Good.Price = price

		recommendGoods = append(recommendGoods, &npool.RecommendGood{
			Recommend: myRecommend,
			Good:      myGood.Info,
		})
	}

	return &npool.GetRecommendGoodsByAppResponse{
		Infos: recommendGoods,
	}, nil
}

func GetByRecommender(ctx context.Context, in *npool.GetRecommendGoodsByRecommenderRequest) (*npool.GetRecommendGoodsByRecommenderResponse, error) {
	myRecommends, err := crud.GetByRecommender(ctx, &npool.GetRecommendsByRecommenderRequest{
		RecommenderID: in.GetRecommenderID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get recommends by recommender: %v", err)
	}

	recommendGoods := []*npool.RecommendGood{}
	for _, myRecommend := range myRecommends.Infos {
		myGood, err := good.Get(ctx, &npool.GetGoodDetailRequest{
			ID: myRecommend.GoodID,
		})
		if err != nil {
			logger.Sugar().Errorf("fail get good: %v", err)
		}

		recommendGoods = append(recommendGoods, &npool.RecommendGood{
			Recommend: myRecommend,
			Good:      myGood.Info,
		})
	}

	return &npool.GetRecommendGoodsByRecommenderResponse{
		Infos: recommendGoods,
	}, nil
}
