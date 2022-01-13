// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/app-good-target-area" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AuthorizeAppGoodTargetArea(ctx context.Context, in *npool.AuthorizeAppGoodTargetAreaRequest) (*npool.AuthorizeAppGoodTargetAreaResponse, error) {
	resp, err := appgoodtargetarea.Authorize(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("authorize app good target area error: %w", err)
		return &npool.AuthorizeAppGoodTargetAreaResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) CheckAppGoodTargetArea(ctx context.Context, in *npool.CheckAppGoodTargetAreaRequest) (*npool.CheckAppGoodTargetAreaResponse, error) {
	resp, err := appgoodtargetarea.Check(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("check app good target area error: %w", err)
		return &npool.CheckAppGoodTargetAreaResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UnauthorizeAppGoodTargetArea(ctx context.Context, in *npool.UnauthorizeAppGoodTargetAreaRequest) (*npool.UnauthorizeAppGoodTargetAreaResponse, error) {
	resp, err := appgoodtargetarea.Unauthorize(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("unauthorize app good target area error: %w", err)
		return &npool.UnauthorizeAppGoodTargetAreaResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
