//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/fee-duration" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateFeeDuration(ctx context.Context, in *npool.CreateFeeDurationRequest) (*npool.CreateFeeDurationResponse, error) {
	resp, err := feeduration.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create good fee error: %w", err)
		return &npool.CreateFeeDurationResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateFeeDuration(ctx context.Context, in *npool.UpdateFeeDurationRequest) (*npool.UpdateFeeDurationResponse, error) {
	resp, err := feeduration.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update good fee error: %w", err)
		return &npool.UpdateFeeDurationResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetFeeDuration(ctx context.Context, in *npool.GetFeeDurationRequest) (*npool.GetFeeDurationResponse, error) {
	resp, err := feeduration.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("delete good fee error: %w", err)
		return &npool.GetFeeDurationResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetFeeDurationsByFeeType(ctx context.Context, in *npool.GetFeeDurationsByFeeTypeRequest) (*npool.GetFeeDurationsByFeeTypeResponse, error) {
	resp, err := feeduration.GetByFeeType(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good fees error: %w", err)
		return &npool.GetFeeDurationsByFeeTypeResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) DeleteFeeDuration(ctx context.Context, in *npool.DeleteFeeDurationRequest) (*npool.DeleteFeeDurationResponse, error) {
	resp, err := feeduration.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good fees error: %w", err)
		return &npool.DeleteFeeDurationResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
