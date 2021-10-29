package appgood

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	npool.UnimplementedCloudHashingGoodsServer
}

func (s *Server) AuthorizeAppGood(ctx context.Context, in *npool.AuthorizeAppGoodRequest) (*npool.AuthorizeAppGoodResponse, error) {
	return nil, nil
}

func (s *Server) UnauthorizeAppGood(ctx context.Context, in *npool.UnauthorizeAppGoodRequest) (*npool.UnauthorizeAppGoodResponse, error) {
	return nil, nil
}
