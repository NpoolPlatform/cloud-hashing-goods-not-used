// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/app-good" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AuthorizeAppGood(ctx context.Context, in *npool.AuthorizeAppGoodRequest) (*npool.AuthorizeAppGoodResponse, error) {
	resp, err := appgood.Authorize(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("authorize app good error: %w", err)
		return &npool.AuthorizeAppGoodResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) SetAppGoodPrice(ctx context.Context, in *npool.SetAppGoodPriceRequest) (*npool.SetAppGoodPriceResponse, error) {
	resp, err := appgood.SetAppGoodPrice(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("authorize app good error: %w", err)
		return &npool.SetAppGoodPriceResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) CheckAppGood(ctx context.Context, in *npool.CheckAppGoodRequest) (*npool.CheckAppGoodResponse, error) {
	resp, err := appgood.Check(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("authorize app good error: %w", err)
		return &npool.CheckAppGoodResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) OnsaleAppGood(ctx context.Context, in *npool.OnsaleAppGoodRequest) (*npool.OnsaleAppGoodResponse, error) {
	resp, err := appgood.Onsale(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("authorize app good error: %w", err)
		return &npool.OnsaleAppGoodResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) OffsaleAppGood(ctx context.Context, in *npool.OffsaleAppGoodRequest) (*npool.OffsaleAppGoodResponse, error) {
	resp, err := appgood.Offsale(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("authorize app good error: %w", err)
		return &npool.OffsaleAppGoodResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UnauthorizeAppGood(ctx context.Context, in *npool.UnauthorizeAppGoodRequest) (*npool.UnauthorizeAppGoodResponse, error) {
	resp, err := appgood.Unauthorize(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("authorize app good error: %w", err)
		return &npool.UnauthorizeAppGoodResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
