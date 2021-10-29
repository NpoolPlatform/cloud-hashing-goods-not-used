package goodinfo

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	npool.UnimplementedCloudHashingGoodsServer
}

func (s *Server) CreateGood(ctx context.Context, in *npool.CreateGoodRequest) (*npool.CreateGoodResponse, error) {
	return nil, nil
}

func (s *Server) UpdateGood(ctx context.Context, in *npool.UpdateGoodRequest) (*npool.UpdateGoodResponse, error) {
	return nil, nil
}

func (s *Server) GetAllGoods(ctx context.Context, in *npool.GetAllGoodsRequest) (*npool.GetAllGoodsResponse, error) {
	return nil, nil
}

func (s *Server) GetGoods(ctx context.Context, in *npool.GetGoodsRequest) (*npool.GetGoodsResponse, error) {
	return nil, nil
}
