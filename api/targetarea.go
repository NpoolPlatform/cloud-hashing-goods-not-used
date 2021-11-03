package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/target-area" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateTargetArea(ctx context.Context, in *npool.CreateTargetAreaRequest) (*npool.CreateTargetAreaResponse, error) {
	resp, err := targetarea.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create target area error: %w", err)
		return &npool.CreateTargetAreaResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateTargetArea(ctx context.Context, in *npool.UpdateTargetAreaRequest) (*npool.UpdateTargetAreaResponse, error) {
	resp, err := targetarea.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update target area error: %w", err)
		return &npool.UpdateTargetAreaResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetTargetArea(ctx context.Context, in *npool.GetTargetAreaRequest) (*npool.GetTargetAreaResponse, error) {
	resp, err := targetarea.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get target area error: %w", err)
		return &npool.GetTargetAreaResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) DeleteTargetArea(ctx context.Context, in *npool.DeleteTargetAreaRequest) (*npool.DeleteTargetAreaResponse, error) {
	resp, err := targetarea.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("delete target area error: %w", err)
		return &npool.DeleteTargetAreaResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) DeleteTargetAreaByContinentCountry(ctx context.Context, in *npool.DeleteTargetAreaByContinentCountryRequest) (*npool.DeleteTargetAreaByContinentCountryResponse, error) {
	resp, err := targetarea.DeleteByContinentCountry(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("delete target area by continent '%v' and country '%v' error: %w", in.GetContinent(), in.GetCountry(), err)
		return &npool.DeleteTargetAreaByContinentCountryResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetTargetAreas(ctx context.Context, in *npool.GetTargetAreasRequest) (*npool.GetTargetAreasResponse, error) {
	resp, err := targetarea.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get target areas error: %w", err)
		return &npool.GetTargetAreasResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
