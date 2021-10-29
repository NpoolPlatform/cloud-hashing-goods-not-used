package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) AuthorizeAppGoodTargetArea(ctx context.Context, in *npool.AuthorizeAppGoodTargetAreaRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *Server) UnauthorizeAppGoodTargetArea(ctx context.Context, in *npool.UnauthorizeAppGoodTargetAreaRequest) (*emptypb.Empty, error) {
	return nil, nil
}
