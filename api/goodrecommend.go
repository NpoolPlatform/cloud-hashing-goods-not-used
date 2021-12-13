//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/recommend-good" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateRecommendGood(ctx context.Context, in *npool.CreateRecommendGoodRequest) (*npool.CreateRecommendGoodResponse, error) {
	resp, err := recommendgood.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create recommend good error: %w", err)
		return &npool.CreateRecommendGoodResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateRecommendGood(ctx context.Context, in *npool.UpdateRecommendGoodRequest) (*npool.UpdateRecommendGoodResponse, error) {
	resp, err := recommendgood.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update recommend good error: %w", err)
		return &npool.UpdateRecommendGoodResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetRecommendGoods(ctx context.Context, in *npool.GetRecommendGoodsRequest) (*npool.GetRecommendGoodsResponse, error) {
	resp, err := recommendgood.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get recommend goods error: %w", err)
		return &npool.GetRecommendGoodsResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetRecommendGoodsByRecommender(ctx context.Context, in *npool.GetRecommendGoodsByRecommenderRequest) (*npool.GetRecommendGoodsByRecommenderResponse, error) {
	resp, err := recommendgood.GetByRecommender(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get recommend goods error: %w", err)
		return &npool.GetRecommendGoodsByRecommenderResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
