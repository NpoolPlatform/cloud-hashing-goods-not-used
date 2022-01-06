package recommend

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	// crud "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/recommend" //nolint
)

func GetByApp(ctx context.Context, in *npool.GetRecommendGoodsByAppRequest) (*npool.GetRecommendGoodsByAppResponse, error) {
	return nil, nil
}

func GetByRecommender(ctx context.Context, in *npool.GetRecommendGoodsByRecommenderRequest) (*npool.GetRecommendGoodsByRecommenderResponse, error) {
	return nil, nil
}
