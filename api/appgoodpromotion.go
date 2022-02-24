package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	crud "github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/appgoodpromotion"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppGoodPromotion(ctx context.Context, in *npool.CreateAppGoodPromotionRequest) (*npool.CreateAppGoodPromotionResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create app good promotion error: %v", err)
		return &npool.CreateAppGoodPromotionResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppGoodPromotionForOtherApp(ctx context.Context, in *npool.CreateAppGoodPromotionForOtherAppRequest) (*npool.CreateAppGoodPromotionForOtherAppResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()

	resp, err := crud.Create(ctx, &npool.CreateAppGoodPromotionRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorf("create app good promotion error: %v", err)
		return &npool.CreateAppGoodPromotionForOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.CreateAppGoodPromotionForOtherAppResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) UpdateAppGoodPromotion(ctx context.Context, in *npool.UpdateAppGoodPromotionRequest) (*npool.UpdateAppGoodPromotionResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update app good promotion error: %v", err)
		return &npool.UpdateAppGoodPromotionResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) DeleteAppGoodPromotion(ctx context.Context, in *npool.DeleteAppGoodPromotionRequest) (*npool.DeleteAppGoodPromotionResponse, error) {
	resp, err := crud.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update app good promotion error: %v", err)
		return &npool.DeleteAppGoodPromotionResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppGoodPromotion(ctx context.Context, in *npool.GetAppGoodPromotionRequest) (*npool.GetAppGoodPromotionResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app good promotion error: %v", err)
		return &npool.GetAppGoodPromotionResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppGoodPromotions(ctx context.Context, in *npool.GetAppGoodPromotionsRequest) (*npool.GetAppGoodPromotionsResponse, error) {
	resp, err := crud.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app good promotions error: %v", err)
		return &npool.GetAppGoodPromotionsResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppGoodPromotionByAppGoodStartEnd(ctx context.Context, in *npool.GetAppGoodPromotionByAppGoodStartEndRequest) (*npool.GetAppGoodPromotionByAppGoodStartEndResponse, error) {
	resp, err := crud.GetByAppGoodStartEnd(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app good promotions error: %v", err)
		return &npool.GetAppGoodPromotionByAppGoodStartEndResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppGoodPromotionByAppGoodTimestamp(ctx context.Context, in *npool.GetAppGoodPromotionByAppGoodTimestampRequest) (*npool.GetAppGoodPromotionByAppGoodTimestampResponse, error) {
	resp, err := crud.GetByAppGoodTimestamp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app good promotions error: %v", err)
		return &npool.GetAppGoodPromotionByAppGoodTimestampResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppGoodPromotionsByAppGood(ctx context.Context, in *npool.GetAppGoodPromotionsByAppGoodRequest) (*npool.GetAppGoodPromotionsByAppGoodResponse, error) {
	resp, err := crud.GetByAppGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app good promotions error: %v", err)
		return &npool.GetAppGoodPromotionsByAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppGoodPromotionsByOtherAppGood(ctx context.Context, in *npool.GetAppGoodPromotionsByOtherAppGoodRequest) (*npool.GetAppGoodPromotionsByOtherAppGoodResponse, error) {
	resp, err := crud.GetByAppGood(ctx, &npool.GetAppGoodPromotionsByAppGoodRequest{
		AppID:  in.GetTargetAppID(),
		GoodID: in.GetGoodID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get app good promotions error: %v", err)
		return &npool.GetAppGoodPromotionsByOtherAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAppGoodPromotionsByOtherAppGoodResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetAppGoodPromotionsByApp(ctx context.Context, in *npool.GetAppGoodPromotionsByAppRequest) (*npool.GetAppGoodPromotionsByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app good promotions error: %v", err)
		return &npool.GetAppGoodPromotionsByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppGoodPromotionsByOtherApp(ctx context.Context, in *npool.GetAppGoodPromotionsByOtherAppRequest) (*npool.GetAppGoodPromotionsByOtherAppResponse, error) {
	resp, err := crud.GetByApp(ctx, &npool.GetAppGoodPromotionsByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get app good promotions error: %v", err)
		return &npool.GetAppGoodPromotionsByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAppGoodPromotionsByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}
