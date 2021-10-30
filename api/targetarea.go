package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/target-area" //nolint
)

func (s *Server) CreateTargetArea(ctx context.Context, in *npool.CreateTargetAreaRequest) (*npool.CreateTargetAreaResponse, error) {
	return targetarea.Create(in)
}

func (s *Server) UpdateTargetArea(ctx context.Context, in *npool.UpdateTargetAreaRequest) (*npool.UpdateTargetAreaResponse, error) {
	return targetarea.Update(in)
}

func (s *Server) GetTargetAreas(ctx context.Context, in *npool.GetTargetAreasRequest) (*npool.GetTargetAreasResponse, error) {
	return targetarea.GetAll(in)
}
