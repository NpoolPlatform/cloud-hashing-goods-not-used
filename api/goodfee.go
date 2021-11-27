// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/good-fee" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGoodFee(ctx context.Context, in *npool.CreateGoodFeeRequest) (*npool.CreateGoodFeeResponse, error) {
	resp, err := goodfee.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create good fee error: %w", err)
		return &npool.CreateGoodFeeResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateGoodFee(ctx context.Context, in *npool.UpdateGoodFeeRequest) (*npool.UpdateGoodFeeResponse, error) {
	resp, err := goodfee.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update good fee error: %w", err)
		return &npool.UpdateGoodFeeResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetGoodFee(ctx context.Context, in *npool.GetGoodFeeRequest) (*npool.GetGoodFeeResponse, error) {
	resp, err := goodfee.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("delete good fee error: %w", err)
		return &npool.GetGoodFeeResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetGoodFees(ctx context.Context, in *npool.GetGoodFeesRequest) (*npool.GetGoodFeesResponse, error) {
	resp, err := goodfee.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good fees error: %w", err)
		return &npool.GetGoodFeesResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
