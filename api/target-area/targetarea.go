package targetarea

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

func (s *Server) CreateTargetArea(ctx context.Context, in *npool.CreateTargetAreaRequest) (*npool.CreateTargetAreaResponse, error) {
	return nil, nil
}

func (s *Server) UpdateTargetArea(ctx context.Context, in *npool.UpdateTargetAreaRequest) (*npool.UpdateTargetAreaResponse, error) {
	return nil, nil
}

func (s *Server) GetTargetAreas(ctx context.Context, in *emptypb.Empty) (*npool.GetTargetAreasResponse, error) {
	return nil, nil
}
