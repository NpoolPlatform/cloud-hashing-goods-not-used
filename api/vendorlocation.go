package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
)

func (s *Server) CreateVendorLocation(ctx context.Context, in *npool.CreateVendorLocationRequest) (*npool.CreateVendorLocationResponse, error) {
	return nil, nil
}

func (s *Server) UpdateVendorLocation(ctx context.Context, in *npool.UpdateVendorLocationRequest) (*npool.UpdateVendorLocationResponse, error) {
	return nil, nil
}

func (s *Server) GetVendorLocations(ctx context.Context, in *npool.GetVendorLocationsRequest) (*npool.GetVendorLocationsResponse, error) {
	return &npool.GetVendorLocationsResponse{}, nil
}
