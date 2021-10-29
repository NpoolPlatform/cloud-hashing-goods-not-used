package deviceinfo

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"google.golang.org/protobuf/types/known/emptypb"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	npool.UnimplementedCloudHashingGoodsServer
}

func (s *Server) CreateDeviceInfo(ctx context.Context, in *npool.CreateDeviceInfoRequest) (*npool.CreateDeviceInfoResponse, error) {
	return nil, nil
}

func (s *Server) UpdateDeviceInfo(ctx context.Context, in *npool.UpdateDeviceInfoRequest) (*npool.UpdateDeviceInfoResponse, error) {
	return nil, nil
}

func (s *Server) GetDeviceInfos(ctx context.Context, in *emptypb.Empty) (*npool.GetDeviceInfosResponse, error) {
	return nil, nil
}
