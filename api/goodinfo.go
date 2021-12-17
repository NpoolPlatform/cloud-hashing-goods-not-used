//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/good-info"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/middleware/app-good-info"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGood(ctx context.Context, in *npool.CreateGoodRequest) (*npool.CreateGoodResponse, error) {
	resp, err := goodinfo.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update target area error: %w", err)
		return &npool.CreateGoodResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateGood(ctx context.Context, in *npool.UpdateGoodRequest) (*npool.UpdateGoodResponse, error) {
	resp, err := goodinfo.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update target area error: %w", err)
		return &npool.UpdateGoodResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetGood(ctx context.Context, in *npool.GetGoodRequest) (*npool.GetGoodResponse, error) {
	resp, err := goodinfo.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update target area error: %w", err)
		return &npool.GetGoodResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) DeleteGood(ctx context.Context, in *npool.DeleteGoodRequest) (*npool.DeleteGoodResponse, error) {
	resp, err := goodinfo.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update target area error: %w", err)
		return &npool.DeleteGoodResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetGoods(ctx context.Context, in *npool.GetGoodsRequest) (*npool.GetGoodsResponse, error) {
	resp, err := goodinfo.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update target area error: %w", err)
		return &npool.GetGoodsResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetGoodsByApp(ctx context.Context, in *npool.GetGoodsByAppRequest) (*npool.GetGoodsByAppResponse, error) {
	resp, err := appgoodinfo.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get goods by app error: %w", err)
		return &npool.GetGoodsByAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
