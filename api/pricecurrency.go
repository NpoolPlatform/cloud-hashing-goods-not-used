package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/price-currency" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreatePriceCurrency(ctx context.Context, in *npool.CreatePriceCurrencyRequest) (*npool.CreatePriceCurrencyResponse, error) {
	resp, err := pricecurrency.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create price currency error: %w", err)
		return &npool.CreatePriceCurrencyResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdatePriceCurrency(ctx context.Context, in *npool.UpdatePriceCurrencyRequest) (*npool.UpdatePriceCurrencyResponse, error) {
	resp, err := pricecurrency.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update price currency error: %w", err)
		return &npool.UpdatePriceCurrencyResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetPriceCurrency(ctx context.Context, in *npool.GetPriceCurrencyRequest) (*npool.GetPriceCurrencyResponse, error) {
	resp, err := pricecurrency.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get price currency error: %w", err)
		return &npool.GetPriceCurrencyResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetPriceCurrencys(ctx context.Context, in *npool.GetPriceCurrencysRequest) (*npool.GetPriceCurrencysResponse, error) {
	resp, err := pricecurrency.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get price currencys error: %w", err)
		return &npool.GetPriceCurrencysResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
