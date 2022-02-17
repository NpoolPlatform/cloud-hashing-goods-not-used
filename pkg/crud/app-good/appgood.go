package appgood

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent"
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

func dbRowToAppGood(info *ent.AppGood) *npool.AppGoodInfo {
	return &npool.AppGoodInfo{
		ID:               info.ID.String(),
		AppID:            info.AppID.String(),
		GoodID:           info.GoodID.String(),
		Online:           info.Online,
		InitAreaStrategy: string(info.InitAreaStrategy),
		Price:            price.DBPriceToVisualPrice(info.Price),
	}
}

func Authorize(ctx context.Context, in *npool.AuthorizeAppGoodRequest) (*npool.AuthorizeAppGoodResponse, error) {
	if err := validateAppGood(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err == nil {
		info, err := cli.
			AppGood.
			UpdateOneID(id).
			SetDeleteAt(0).
			Save(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail authorize app good: %v", err)
		}
		return &npool.AuthorizeAppGoodResponse{
			Info: dbRowToAppGood(info),
		}, nil
	}

	initAreaStrategy := in.GetInfo().GetInitAreaStrategy()
	if in.GetInfo().GetInitAreaStrategy() == "" {
		initAreaStrategy = "all"
	}

	info, err := cli.
		AppGood.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetOnline(false).
		SetInitAreaStrategy(appgood.InitAreaStrategy(initAreaStrategy)).
		SetPrice(0).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app good: %v", err)
	}

	return &npool.AuthorizeAppGoodResponse{
		Info: dbRowToAppGood(info),
	}, nil
}

func Check(ctx context.Context, in *npool.CheckAppGoodRequest) (*npool.CheckAppGoodResponse, error) {
	if err := validateAppGood(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppGood.
		Query().
		Where(
			appgood.And(
				appgood.AppID(uuid.MustParse(in.GetInfo().GetAppID())),
				appgood.GoodID(uuid.MustParse(in.GetInfo().GetGoodID())),
				appgood.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app good: %v", err)
	}

	var appGood *npool.AppGoodInfo
	for _, info := range infos {
		appGood = dbRowToAppGood(info)
	}

	return &npool.CheckAppGoodResponse{
		Info: appGood,
	}, nil
}

func SetAppGoodPrice(ctx context.Context, in *npool.SetAppGoodPriceRequest) (*npool.SetAppGoodPriceResponse, error) {
	if err := validateAppGood(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app good id: %v", err)
	}

	info1, err := Check(ctx, &npool.CheckAppGoodRequest{
		Info: in.GetInfo(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail check app good: %v", err)
	}

	if info1.Info.Online {
		return nil, xerrors.Errorf("cannot set price to online good")
	}

	if price.VisualPriceToDBPrice(in.GetInfo().GetPrice()) == 0 {
		return nil, xerrors.Errorf("price should be greater than 0")
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppGood.
		UpdateOneID(id).
		SetPrice(price.VisualPriceToDBPrice(in.GetInfo().GetPrice())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail set price app good: %v", err)
	}

	return &npool.SetAppGoodPriceResponse{
		Info: dbRowToAppGood(info),
	}, nil
}

func Onsale(ctx context.Context, in *npool.OnsaleAppGoodRequest) (*npool.OnsaleAppGoodResponse, error) {
	if err := validateAppGood(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app good id: %v", err)
	}

	_, err = Check(ctx, &npool.CheckAppGoodRequest{
		Info: in.GetInfo(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail onsale app good: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppGood.
		UpdateOneID(id).
		SetOnline(true).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail onsale app good: %v", err)
	}

	return &npool.OnsaleAppGoodResponse{
		Info: dbRowToAppGood(info),
	}, nil
}

func Offsale(ctx context.Context, in *npool.OffsaleAppGoodRequest) (*npool.OffsaleAppGoodResponse, error) {
	if err := validateAppGood(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app good id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppGood.
		UpdateOneID(id).
		SetOnline(false).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail offsale app good: %v", err)
	}

	return &npool.OffsaleAppGoodResponse{
		Info: dbRowToAppGood(info),
	}, nil
}

func Unauthorize(ctx context.Context, in *npool.UnauthorizeAppGoodRequest) (*npool.UnauthorizeAppGoodResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app good id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppGood.
		UpdateOneID(id).
		SetOnline(false).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail unauthorize app good: %v", err)
	}

	return &npool.UnauthorizeAppGoodResponse{
		Info: dbRowToAppGood(info),
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAppGoodsRequest) (*npool.GetAppGoodsResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppGood.
		Query().
		Where(
			appgood.And(
				appgood.AppID(appID),
				appgood.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app good: %v", err)
	}

	appGoods := []*npool.AppGoodInfo{}
	for _, info := range infos {
		appGoods = append(appGoods, dbRowToAppGood(info))
	}

	return &npool.GetAppGoodsResponse{
		Infos: appGoods,
	}, nil
}
