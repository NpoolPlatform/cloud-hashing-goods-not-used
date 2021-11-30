//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateFee(ctx context.Context, in *npool.CreateFeeRequest) (*npool.CreateFeeResponse, error) {
	resp, err := fee.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create fee error: %v", err)
		return &npool.CreateFeeResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetFee(ctx context.Context, in *npool.GetFeeRequest) (*npool.GetFeeResponse, error) {
	resp, err := fee.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get fee error: %v", err)
		return &npool.GetFeeResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
