package recommend

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	crud "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/recommend"         //nolint
	good "github.com/NpoolPlatform/cloud-hashing-goods/pkg/middleware/good-detail" //nolint

	"golang.org/x/xerrors"
)

func GetByApp(ctx context.Context, in *npool.GetRecommendGoodsByAppRequest) (*npool.GetRecommendGoodsByAppResponse, error) {
	myRecommends, err := crud.GetByApp(ctx, &npool.GetRecommendsByAppRequest{
		AppID: in.GetAppID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get recommends by app: %v", err)
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
			Good:      myGood.Detail,
		})
	}

	return &npool.GetRecommendGoodsByAppResponse{
		Infos: recommendGoods,
	}, nil
}

func GetByRecommender(ctx context.Context, in *npool.GetRecommendGoodsByRecommenderRequest) (*npool.GetRecommendGoodsByRecommenderResponse, error) {
	myRecommends, err := crud.GetByRecommender(ctx, &npool.GetRecommendsByRecommenderRequest{
		UserID: in.GetUserID(),
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
			Good:      myGood.Detail,
		})
	}

	return &npool.GetRecommendGoodsByRecommenderResponse{
		Infos: recommendGoods,
	}, nil
}
