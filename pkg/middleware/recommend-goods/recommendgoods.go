package recommendgoods

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	goodinfo "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/good-info"
	recommendgood "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/recommend-good"
	"golang.org/x/xerrors"
	//nolint
)

func Get(ctx context.Context, in *npool.GetRecommendGoodsRequest) (*npool.GetRecommendGoodsResponse, error) {
	partialGoodsResponse, err := recommendgood.Get(ctx, in)
	if err != nil {
		return &npool.GetRecommendGoodsResponse{}, err
	}

	partialGoodsResponse.Infos = make([]*npool.GoodInfo, len(partialGoodsResponse.Recommends))
	for i := 0; i < len(partialGoodsResponse.Recommends); i++ {
		tmpResp, err := goodinfo.Get(ctx, &npool.GetGoodRequest{
			ID: partialGoodsResponse.Recommends[i].GoodID,
		})
		if err == nil {
			partialGoodsResponse.Infos = append(partialGoodsResponse.Infos, tmpResp.Info)
		}
	}

	if len(partialGoodsResponse.Infos) == 0 {
		return partialGoodsResponse, xerrors.Errorf("no record found: %v", err)
	}
	return partialGoodsResponse, nil
}

func GetByRecommender(ctx context.Context, in *npool.GetRecommendGoodsByRecommenderRequest) (*npool.GetRecommendGoodsByRecommenderResponse, error) {
	partialGoodsResponse, err := recommendgood.GetByRecommender(ctx, in)
	if err != nil {
		return &npool.GetRecommendGoodsByRecommenderResponse{}, err
	}
	partialGoodsResponse.Infos = make([]*npool.GoodInfo, len(partialGoodsResponse.Recommends))

	for i := 0; i < len(partialGoodsResponse.Recommends); i++ {
		tmpResp, err := goodinfo.Get(ctx, &npool.GetGoodRequest{
			ID: partialGoodsResponse.Recommends[i].GoodID,
		})
		if err == nil {
			partialGoodsResponse.Infos = append(partialGoodsResponse.Infos, tmpResp.Info)
		}
	}

	if len(partialGoodsResponse.Infos) == 0 {
		return partialGoodsResponse, xerrors.Errorf("no record found: %v", err)
	}
	return partialGoodsResponse, nil
}
