// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/device-info" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateDeviceInfo(ctx context.Context, in *npool.CreateDeviceInfoRequest) (*npool.CreateDeviceInfoResponse, error) {
	resp, err := deviceinfo.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create device info error: %w", err)
		return &npool.CreateDeviceInfoResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateDeviceInfo(ctx context.Context, in *npool.UpdateDeviceInfoRequest) (*npool.UpdateDeviceInfoResponse, error) {
	resp, err := deviceinfo.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update device info error: %w", err)
		return &npool.UpdateDeviceInfoResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetDeviceInfo(ctx context.Context, in *npool.GetDeviceInfoRequest) (*npool.GetDeviceInfoResponse, error) {
	resp, err := deviceinfo.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("delete device info error: %w", err)
		return &npool.GetDeviceInfoResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) DeleteDeviceInfo(ctx context.Context, in *npool.DeleteDeviceInfoRequest) (*npool.DeleteDeviceInfoResponse, error) {
	resp, err := deviceinfo.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("delete device info error: %w", err)
		return &npool.DeleteDeviceInfoResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetDeviceInfos(ctx context.Context, in *npool.GetDeviceInfosRequest) (*npool.GetDeviceInfosResponse, error) {
	resp, err := deviceinfo.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get device infos error: %w", err)
		return &npool.GetDeviceInfosResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
