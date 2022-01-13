// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/good-review" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGoodReview(ctx context.Context, in *npool.CreateGoodReviewRequest) (*npool.CreateGoodReviewResponse, error) {
	resp, err := goodreview.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create good review error: %w", err)
		return &npool.CreateGoodReviewResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateGoodReview(ctx context.Context, in *npool.UpdateGoodReviewRequest) (*npool.UpdateGoodReviewResponse, error) {
	resp, err := goodreview.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update good review error: %w", err)
		return &npool.UpdateGoodReviewResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetGoodReview(ctx context.Context, in *npool.GetGoodReviewRequest) (*npool.GetGoodReviewResponse, error) {
	resp, err := goodreview.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good reviews error: %w", err)
		return &npool.GetGoodReviewResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
