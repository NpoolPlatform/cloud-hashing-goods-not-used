//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/good-recommend" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGoodRecommend(ctx context.Context, in *npool.CreateGoodRecommendRequest) (*npool.CreateGoodRecommendResponse, error) {
	resp, err := goodrecommend.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create good recommend error: %w", err)
		return &npool.CreateGoodRecommendResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateGoodRecommend(ctx context.Context, in *npool.UpdateGoodRecommendRequest) (*npool.UpdateGoodRecommendResponse, error) {
	resp, err := goodrecommend.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update good recommend error: %w", err)
		return &npool.UpdateGoodRecommendResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetGoodRecommendsByGood(ctx context.Context, in *npool.GetGoodRecommendsByGoodRequest) (*npool.GetGoodRecommendsByGoodResponse, error) {
	resp, err := goodrecommend.GetByGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good recommends error: %w", err)
		return &npool.GetGoodRecommendsByGoodResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetGoodRecommendsByUser(ctx context.Context, in *npool.GetGoodRecommendsByUserRequest) (*npool.GetGoodRecommendsByUserResponse, error) {
	resp, err := goodrecommend.GetByUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good recommends error: %w", err)
		return &npool.GetGoodRecommendsByUserResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
