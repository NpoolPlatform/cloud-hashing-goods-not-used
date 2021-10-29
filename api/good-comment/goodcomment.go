package goodcomment

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	npool.UnimplementedCloudHashingGoodsServer
}

func (s *Server) CreateGoodComment(ctx context.Context, in *npool.CreateGoodCommentRequest) (*npool.CreateGoodCommentResponse, error) {
	return nil, nil
}

func (s *Server) UpdateGoodComment(ctx context.Context, in *npool.UpdateGoodCommentRequest) (*npool.UpdateGoodCommentResponse, error) {
	return nil, nil
}

func (s *Server) GetGoodComments(ctx context.Context, in *npool.GetGoodCommentsRequest) (*npool.GetGoodCommentsResponse, error) {
	return nil, nil
}
