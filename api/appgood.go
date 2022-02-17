// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/crud/app-good" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AuthorizeAppGood(ctx context.Context, in *npool.AuthorizeAppGoodRequest) (*npool.AuthorizeAppGoodResponse, error) {
	resp, err := appgood.Authorize(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("authorize app good error: %v", err)
		return &npool.AuthorizeAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) AuthorizeAppGoodForOtherApp(ctx context.Context, in *npool.AuthorizeAppGoodForOtherAppRequest) (*npool.AuthorizeAppGoodForOtherAppResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()

	resp, err := appgood.Authorize(ctx, &npool.AuthorizeAppGoodRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorf("authorize app good error: %v", err)
		return &npool.AuthorizeAppGoodForOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.AuthorizeAppGoodForOtherAppResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) SetAppGoodPrice(ctx context.Context, in *npool.SetAppGoodPriceRequest) (*npool.SetAppGoodPriceResponse, error) {
	resp, err := appgood.SetAppGoodPrice(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("authorize app good error: %v", err)
		return &npool.SetAppGoodPriceResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) SetAppGoodPriceForOtherApp(ctx context.Context, in *npool.SetAppGoodPriceForOtherAppRequest) (*npool.SetAppGoodPriceForOtherAppResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()

	resp, err := appgood.SetAppGoodPrice(ctx, &npool.SetAppGoodPriceRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorf("authorize app good error: %v", err)
		return &npool.SetAppGoodPriceForOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.SetAppGoodPriceForOtherAppResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) CheckAppGood(ctx context.Context, in *npool.CheckAppGoodRequest) (*npool.CheckAppGoodResponse, error) {
	resp, err := appgood.Check(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("authorize app good error: %v", err)
		return &npool.CheckAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) OnsaleAppGood(ctx context.Context, in *npool.OnsaleAppGoodRequest) (*npool.OnsaleAppGoodResponse, error) {
	resp, err := appgood.Onsale(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("authorize app good error: %v", err)
		return &npool.OnsaleAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) OnsaleAppGoodForOtherApp(ctx context.Context, in *npool.OnsaleAppGoodForOtherAppRequest) (*npool.OnsaleAppGoodForOtherAppResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()

	resp, err := appgood.Onsale(ctx, &npool.OnsaleAppGoodRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorf("authorize app good error: %v", err)
		return &npool.OnsaleAppGoodForOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.OnsaleAppGoodForOtherAppResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) OffsaleAppGood(ctx context.Context, in *npool.OffsaleAppGoodRequest) (*npool.OffsaleAppGoodResponse, error) {
	resp, err := appgood.Offsale(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("authorize app good error: %v", err)
		return &npool.OffsaleAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) OffsaleAppGoodForOtherApp(ctx context.Context, in *npool.OffsaleAppGoodForOtherAppRequest) (*npool.OffsaleAppGoodForOtherAppResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()

	resp, err := appgood.Offsale(ctx, &npool.OffsaleAppGoodRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorf("authorize app good error: %v", err)
		return &npool.OffsaleAppGoodForOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.OffsaleAppGoodForOtherAppResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) UnauthorizeAppGood(ctx context.Context, in *npool.UnauthorizeAppGoodRequest) (*npool.UnauthorizeAppGoodResponse, error) {
	resp, err := appgood.Unauthorize(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("authorize app good error: %v", err)
		return &npool.UnauthorizeAppGoodResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppGoods(ctx context.Context, in *npool.GetAppGoodsRequest) (*npool.GetAppGoodsResponse, error) {
	resp, err := appgood.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app goods error: %v", err)
		return &npool.GetAppGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppGoodsByOtherApp(ctx context.Context, in *npool.GetAppGoodsByOtherAppRequest) (*npool.GetAppGoodsByOtherAppResponse, error) {
	resp, err := appgood.GetByApp(ctx, &npool.GetAppGoodsRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get app goods error: %v", err)
		return &npool.GetAppGoodsByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAppGoodsByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}
