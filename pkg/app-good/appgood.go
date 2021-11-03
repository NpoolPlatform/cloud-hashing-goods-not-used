package appgood

import (
	"context"
	_ "time" //nolint

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/appgood"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateAppGood(info *npool.AppGoodInfo) error {
	_, err := uuid.Parse(info.GetAppID())
	if err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}

	_, err = uuid.Parse(info.GetGoodID())
	if err != nil {
		return xerrors.Errorf("invalid good id: %v", err)
	}

	return nil
}

func Authorize(ctx context.Context, in *npool.AuthorizeAppGoodRequest) (*npool.AuthorizeAppGoodResponse, error) {
	if err := validateAppGood(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err == nil {
		info, err := db.Client().
			AppGood.
			UpdateOneID(id).
			SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
			SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
			SetAuthorized(true).
			SetDeleteAt(0).
			Save(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail update app good: %v", err)
		}
		return &npool.AuthorizeAppGoodResponse{
			Info: &npool.AppGoodInfo{
				ID:               info.ID.String(),
				AppID:            info.AppID.String(),
				GoodID:           info.GoodID.String(),
				Authorized:       info.Authorized,
				Online:           info.Online,
				InitAreaStrategy: string(info.InitAreaStrategy),
				Price:            price.DBPriceToVisualPrice(info.Price),
				GasPrice:         price.DBPriceToVisualPrice(info.GasPrice),
			},
		}, nil
	}

	info, err := db.Client().
		AppGood.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetAuthorized(true).
		SetOnline(false).
		SetInitAreaStrategy(appgood.InitAreaStrategy(in.GetInfo().GetInitAreaStrategy())).
		SetPrice(0).
		SetGasPrice(0).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app good: %v", err)
	}

	return &npool.AuthorizeAppGoodResponse{
		Info: &npool.AppGoodInfo{
			ID:               info.ID.String(),
			AppID:            info.AppID.String(),
			GoodID:           info.GoodID.String(),
			Authorized:       info.Authorized,
			Online:           info.Online,
			InitAreaStrategy: string(info.InitAreaStrategy),
			Price:            price.DBPriceToVisualPrice(info.Price),
			GasPrice:         price.DBPriceToVisualPrice(info.GasPrice),
		},
	}, nil
}

func SetAppGoodPrice(ctx context.Context, in *npool.SetAppGoodPriceRequest) (*npool.SetAppGoodPriceResponse, error) {
	return nil, nil
}

func Check(ctx context.Context, in *npool.CheckAppGoodRequest) (*npool.CheckAppGoodResponse, error) {
	return nil, nil
}

func Onsale(ctx context.Context, in *npool.OnsaleAppGoodRequest) (*npool.OnsaleAppGoodResponse, error) {
	return nil, nil
}

func Offsale(ctx context.Context, in *npool.OffsaleAppGoodRequest) (*npool.OffsaleAppGoodResponse, error) {
	return nil, nil
}

func Unauthorize(ctx context.Context, in *npool.UnauthorizeAppGoodRequest) (*npool.UnauthorizeAppGoodResponse, error) {
	return nil, nil
}
