//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/fee-type" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateFeeType(ctx context.Context, in *npool.CreateFeeTypeRequest) (*npool.CreateFeeTypeResponse, error) {
	resp, err := feetype.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create good fee error: %w", err)
		return &npool.CreateFeeTypeResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateFeeType(ctx context.Context, in *npool.UpdateFeeTypeRequest) (*npool.UpdateFeeTypeResponse, error) {
	resp, err := feetype.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update good fee error: %w", err)
		return &npool.UpdateFeeTypeResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetFeeType(ctx context.Context, in *npool.GetFeeTypeRequest) (*npool.GetFeeTypeResponse, error) {
	resp, err := feetype.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("delete good fee error: %w", err)
		return &npool.GetFeeTypeResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetFeeTypes(ctx context.Context, in *npool.GetFeeTypesRequest) (*npool.GetFeeTypesResponse, error) {
	resp, err := feetype.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good fees error: %w", err)
		return &npool.GetFeeTypesResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
