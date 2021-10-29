package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) AuthorizeAppTargetArea(ctx context.Context, in *npool.AuthorizeAppTargetAreaRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *Server) UnauthorizeAppTargetArea(ctx context.Context, in *npool.UnauthorizeAppTargetAreaRequest) (*emptypb.Empty, error) {
	return nil, nil
}
