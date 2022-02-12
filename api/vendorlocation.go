// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/vendor-location" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateVendorLocation(ctx context.Context, in *npool.CreateVendorLocationRequest) (*npool.CreateVendorLocationResponse, error) {
	resp, err := vendorlocation.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create vendor location error: %v", err)
		return &npool.CreateVendorLocationResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateVendorLocation(ctx context.Context, in *npool.UpdateVendorLocationRequest) (*npool.UpdateVendorLocationResponse, error) {
	resp, err := vendorlocation.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update vendor location error: %v", err)
		return &npool.UpdateVendorLocationResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetVendorLocation(ctx context.Context, in *npool.GetVendorLocationRequest) (*npool.GetVendorLocationResponse, error) {
	resp, err := vendorlocation.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("delete vendor location error: %v", err)
		return &npool.GetVendorLocationResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) DeleteVendorLocation(ctx context.Context, in *npool.DeleteVendorLocationRequest) (*npool.DeleteVendorLocationResponse, error) {
	resp, err := vendorlocation.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("delete vendor location error: %v", err)
		return &npool.DeleteVendorLocationResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetVendorLocations(ctx context.Context, in *npool.GetVendorLocationsRequest) (*npool.GetVendorLocationsResponse, error) {
	resp, err := vendorlocation.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get vendor locations error: %v", err)
		return &npool.GetVendorLocationsResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
