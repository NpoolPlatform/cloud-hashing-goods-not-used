// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/app-target-area" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AuthorizeAppTargetArea(ctx context.Context, in *npool.AuthorizeAppTargetAreaRequest) (*npool.AuthorizeAppTargetAreaResponse, error) {
	resp, err := apptargetarea.Authorize(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create target area error: %w", err)
		return &npool.AuthorizeAppTargetAreaResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) CheckAppTargetArea(ctx context.Context, in *npool.CheckAppTargetAreaRequest) (*npool.CheckAppTargetAreaResponse, error) {
	resp, err := apptargetarea.Check(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create target area error: %w", err)
		return &npool.CheckAppTargetAreaResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UnauthorizeAppTargetArea(ctx context.Context, in *npool.UnauthorizeAppTargetAreaRequest) (*npool.UnauthorizeAppTargetAreaResponse, error) {
	resp, err := apptargetarea.Unauthorize(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create target area error: %w", err)
		return &npool.UnauthorizeAppTargetAreaResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
