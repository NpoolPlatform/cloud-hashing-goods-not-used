package appgoodpromotion

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/appgoodpromotion"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateAppGoodPromotion(info *npool.AppGoodPromotion) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invlaid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		return xerrors.Errorf("invlaid activity by: %v", err)
	}
	if info.GetEnd() <= info.GetStart() {
		return xerrors.Errorf("invalid expiration")
	}
	if info.GetStart() <= uint32(time.Now().Unix()) {
		return xerrors.Errorf("invalid start %v < %v", info.GetStart(), time.Now().Unix())
	}
	return nil
}

func dbRowToAppGoodPromotion(row *ent.AppGoodPromotion) *npool.AppGoodPromotion {
	return &npool.AppGoodPromotion{
		ID:      row.ID.String(),
		AppID:   row.AppID.String(),
		GoodID:  row.GoodID.String(),
		Message: row.Message,
		Start:   row.Start,
		End:     row.End,
		Price:   price.DBPriceToVisualPrice(row.Price),
	}
}

func Create(ctx context.Context, in *npool.CreateAppGoodPromotionRequest) (*npool.CreateAppGoodPromotionResponse, error) {
	if err := validateAppGoodPromotion(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppGoodPromotion.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetMessage(in.GetInfo().GetMessage()).
		SetStart(in.GetInfo().GetStart()).
		SetEnd(in.GetInfo().GetEnd()).
		SetPrice(price.VisualPriceToDBPrice(in.GetInfo().GetPrice())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create event promotion: %v", err)
	}

	return &npool.CreateAppGoodPromotionResponse{
		Info: dbRowToAppGoodPromotion(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppGoodPromotionRequest) (*npool.UpdateAppGoodPromotionResponse, error) {
	if err := validateAppGoodPromotion(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppGoodPromotion.
		UpdateOneID(id).
		SetMessage(in.GetInfo().GetMessage()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update event promotion: %v", err)
	}

	return &npool.UpdateAppGoodPromotionResponse{
		Info: dbRowToAppGoodPromotion(info),
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteAppGoodPromotionRequest) (*npool.DeleteAppGoodPromotionResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppGoodPromotion.
		UpdateOneID(id).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update event promotion: %v", err)
	}

	return &npool.DeleteAppGoodPromotionResponse{
		Info: dbRowToAppGoodPromotion(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppGoodPromotionRequest) (*npool.GetAppGoodPromotionResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppGoodPromotion.
		Query().
		Where(
			appgoodpromotion.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query event promotion: %v", err)
	}

	var promotion *npool.AppGoodPromotion
	for _, info := range infos {
		promotion = dbRowToAppGoodPromotion(info)
		break
	}

	return &npool.GetAppGoodPromotionResponse{
		Info: promotion,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAppGoodPromotionsByAppRequest) (*npool.GetAppGoodPromotionsByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppGoodPromotion.
		Query().
		Where(
			appgoodpromotion.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query event promotion: %v", err)
	}

	promotions := []*npool.AppGoodPromotion{}
	for _, info := range infos {
		promotions = append(promotions, dbRowToAppGoodPromotion(info))
	}

	return &npool.GetAppGoodPromotionsByAppResponse{
		Infos: promotions,
	}, nil
}

func GetByAppGoodStartEnd(ctx context.Context, in *npool.GetAppGoodPromotionByAppGoodStartEndRequest) (*npool.GetAppGoodPromotionByAppGoodStartEndResponse, error) { //nolint
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppGoodPromotion.
		Query().
		Where(
			appgoodpromotion.And(
				appgoodpromotion.AppID(appID),
				appgoodpromotion.GoodID(goodID),
				appgoodpromotion.Start(in.GetStart()),
				appgoodpromotion.End(in.GetEnd()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query event promotion: %v", err)
	}

	var promotion *npool.AppGoodPromotion
	for _, info := range infos {
		promotion = dbRowToAppGoodPromotion(info)
		break
	}

	return &npool.GetAppGoodPromotionByAppGoodStartEndResponse{
		Info: promotion,
	}, nil
}

func GetByAppGoodTimestamp(ctx context.Context, in *npool.GetAppGoodPromotionByAppGoodTimestampRequest) (*npool.GetAppGoodPromotionByAppGoodTimestampResponse, error) { //nolint
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppGoodPromotion.
		Query().
		Where(
			appgoodpromotion.And(
				appgoodpromotion.AppID(appID),
				appgoodpromotion.GoodID(goodID),
				appgoodpromotion.StartLT(in.GetTimestamp()),
				appgoodpromotion.EndGT(in.GetTimestamp()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query event promotion: %v", err)
	}

	var promotion *npool.AppGoodPromotion
	for _, info := range infos {
		promotion = dbRowToAppGoodPromotion(info)
		break
	}

	return &npool.GetAppGoodPromotionByAppGoodTimestampResponse{
		Info: promotion,
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetAppGoodPromotionsRequest) (*npool.GetAppGoodPromotionsResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppGoodPromotion.
		Query().
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query event promotion: %v", err)
	}

	promotions := []*npool.AppGoodPromotion{}
	for _, info := range infos {
		promotions = append(promotions, dbRowToAppGoodPromotion(info))
	}

	return &npool.GetAppGoodPromotionsResponse{
		Infos: promotions,
	}, nil
}

func GetByAppGood(ctx context.Context, in *npool.GetAppGoodPromotionsByAppGoodRequest) (*npool.GetAppGoodPromotionsByAppGoodResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppGoodPromotion.
		Query().
		Where(
			appgoodpromotion.And(
				appgoodpromotion.AppID(appID),
				appgoodpromotion.GoodID(goodID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query event promotion: %v", err)
	}

	promotions := []*npool.AppGoodPromotion{}
	for _, info := range infos {
		promotions = append(promotions, dbRowToAppGoodPromotion(info))
	}

	return &npool.GetAppGoodPromotionsByAppGoodResponse{
		Infos: promotions,
	}, nil
}
