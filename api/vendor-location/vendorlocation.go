package vendorlocation

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

func (s *Server) CreateVendorLocation(ctx context.Context, in *npool.CreateVendorLocationRequest) (*npool.CreateVendorLocationResponse, error) {
	return nil, nil
}

func (s *Server) UpdateVendorLocation(ctx context.Context, in *npool.UpdateVendorLocationRequest) (*npool.UpdateVendorLocationResponse, error) {
	return nil, nil
}

func (s *Server) GetVendorLocations(ctx context.Context, in *emptypb.Empty) (*npool.GetVendorLocationsResponse, error) {
	return nil, nil
}
