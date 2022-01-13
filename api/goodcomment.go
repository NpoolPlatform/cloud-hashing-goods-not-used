// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/good-comment" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGoodComment(ctx context.Context, in *npool.CreateGoodCommentRequest) (*npool.CreateGoodCommentResponse, error) {
	resp, err := goodcomment.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create good comment error: %w", err)
		return &npool.CreateGoodCommentResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateGoodComment(ctx context.Context, in *npool.UpdateGoodCommentRequest) (*npool.UpdateGoodCommentResponse, error) {
	resp, err := goodcomment.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update good comment error: %w", err)
		return &npool.UpdateGoodCommentResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetGoodComments(ctx context.Context, in *npool.GetGoodCommentsRequest) (*npool.GetGoodCommentsResponse, error) {
	resp, err := goodcomment.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good comments error: %w", err)
		return &npool.GetGoodCommentsResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
