//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/middleware/recommend-goods" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetRecommendGoods(ctx context.Context, in *npool.GetRecommendGoodsRequest) (*npool.GetRecommendGoodsResponse, error) {
	resp, err := recommendgoods.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get recommend goods error: %w", err)
		return &npool.GetRecommendGoodsResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetRecommendGoodsByRecommender(ctx context.Context, in *npool.GetRecommendGoodsByRecommenderRequest) (*npool.GetRecommendGoodsByRecommenderResponse, error) {
	resp, err := recommendgoods.GetByRecommender(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get recommend goods error: %w", err)
		return &npool.GetRecommendGoodsByRecommenderResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
