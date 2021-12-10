package goodinfo

import (
	"context"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/appgood"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodinfo"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func dbRowToInfo(row *ent.GoodInfo) *npool.GoodInfo {
	ids := []string{}
	for _, id := range row.SupportCoinTypeIds {
		ids = append(ids, id.String())
	}

	feeIDs := []string{}
	for _, id := range row.FeeIds {
		feeIDs = append(feeIDs, id.String())
	}

	return &npool.GoodInfo{
		ID:                 row.ID.String(),
		Title:              row.Title,
		DeviceInfoID:       row.DeviceInfoID.String(),
		SeparateFee:        row.SeparateFee,
		UnitPower:          row.UnitPower,
		DurationDays:       row.DurationDays,
		CoinInfoID:         row.CoinInfoID.String(),
		Actuals:            row.Actuals,
		DeliveryAt:         row.DeliveryAt,
		InheritFromGoodID:  row.InheritFromGoodID.String(),
		VendorLocationID:   row.VendorLocationID.String(),
		Price:              price.DBPriceToVisualPrice(row.Price),
		PriceCurrency:      row.PriceCurrency.String(),
		BenefitType:        string(row.BenefitType),
		Classic:            row.Classic,
		SupportCoinTypeIDs: ids,
		Total:              row.Total,
		Unit:               row.Unit,
		FeeIDs:             feeIDs,
	}
}

func validateRequestGoodInfo(info *npool.GoodInfo) error {
	_, err := uuid.Parse(info.GetDeviceInfoID())
	if err != nil {
		return xerrors.Errorf("invalid device id: %v", err)
	}
	_, err = uuid.Parse(info.GetCoinInfoID())
	if err != nil {
		return xerrors.Errorf("invalid coin info id: %v", err)
	}
	_, err = uuid.Parse(info.GetInheritFromGoodID())
	if err != nil {
		return xerrors.Errorf("invalid inherit from good id: %v", err)
	}
	_, err = uuid.Parse(info.GetVendorLocationID())
	if err != nil {
		return xerrors.Errorf("invalid vendor location id: %v", err)
	}
	_, err = uuid.Parse(info.GetPriceCurrency())
	if err != nil {
		return xerrors.Errorf("invalid price currency id: %v", err)
	}
	for _, id := range info.GetSupportCoinTypeIDs() {
		if _, err = uuid.Parse(id); err != nil {
			return xerrors.Errorf("invalid support coin type id: %v", err)
		}
	}
	return nil
}

func Create(ctx context.Context, in *npool.CreateGoodRequest) (*npool.CreateGoodResponse, error) {
	if err := validateRequestGoodInfo(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid request good info: %v", err)
	}

	ids := []uuid.UUID{}
	for _, id := range in.GetInfo().GetSupportCoinTypeIDs() {
		ids = append(ids, uuid.MustParse(id))
	}

	feeIDs := []uuid.UUID{}
	for _, id := range in.GetInfo().GetFeeIDs() {
		feeIDs = append(feeIDs, uuid.MustParse(id))
	}

	info, err := db.Client().
		GoodInfo.
		Create().
		SetTitle(in.GetInfo().GetTitle()).
		SetDeviceInfoID(uuid.MustParse(in.GetInfo().GetDeviceInfoID())).
		SetSeparateFee(in.GetInfo().GetSeparateFee()).
		SetUnitPower(in.GetInfo().GetUnitPower()).
		SetDurationDays(in.GetInfo().GetDurationDays()).
		SetCoinInfoID(uuid.MustParse(in.GetInfo().GetCoinInfoID())).
		SetActuals(in.GetInfo().GetActuals()).
		SetDeliveryAt(in.GetInfo().GetDeliveryAt()).
		SetInheritFromGoodID(uuid.MustParse(in.GetInfo().GetInheritFromGoodID())).
		SetVendorLocationID(uuid.MustParse(in.GetInfo().GetVendorLocationID())).
		SetPrice(price.VisualPriceToDBPrice(in.GetInfo().GetPrice())).
		SetPriceCurrency(uuid.MustParse(in.GetInfo().GetPriceCurrency())).
		SetBenefitType(goodinfo.BenefitType(in.GetInfo().GetBenefitType())).
		SetClassic(in.GetInfo().GetClassic()).
		SetSupportCoinTypeIds(ids).
		SetTotal(in.GetInfo().GetTotal()).
		SetUnit(in.GetInfo().GetUnit()).
		SetFeeIds(feeIDs).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to create good: %v", err)
	}

	return &npool.CreateGoodResponse{
		Info: dbRowToInfo(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateGoodRequest) (*npool.UpdateGoodResponse, error) {
	goodID, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	if err := validateRequestGoodInfo(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid request good info: %v", err)
	}

	ids := []uuid.UUID{}
	for _, id := range in.GetInfo().GetSupportCoinTypeIDs() {
		ids = append(ids, uuid.MustParse(id))
	}

	info, err := db.Client().
		GoodInfo.
		UpdateOneID(goodID).
		SetTitle(in.GetInfo().GetTitle()).
		SetDeviceInfoID(uuid.MustParse(in.GetInfo().GetDeviceInfoID())).
		SetSeparateFee(in.GetInfo().GetSeparateFee()).
		SetUnitPower(in.GetInfo().GetUnitPower()).
		SetDurationDays(in.GetInfo().GetDurationDays()).
		SetCoinInfoID(uuid.MustParse(in.GetInfo().GetCoinInfoID())).
		SetActuals(in.GetInfo().GetActuals()).
		SetDeliveryAt(in.GetInfo().GetDeliveryAt()).
		SetInheritFromGoodID(uuid.MustParse(in.GetInfo().GetInheritFromGoodID())).
		SetVendorLocationID(uuid.MustParse(in.GetInfo().GetVendorLocationID())).
		SetPrice(price.VisualPriceToDBPrice(in.GetInfo().GetPrice())).
		SetBenefitType(goodinfo.BenefitType(in.GetInfo().GetBenefitType())).
		SetClassic(in.GetInfo().GetClassic()).
		SetSupportCoinTypeIds(ids).
		SetTotal(in.GetInfo().GetTotal()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to update good info: %v", err)
	}

	return &npool.UpdateGoodResponse{
		Info: dbRowToInfo(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetGoodRequest) (*npool.GetGoodResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	infos, err := db.Client().
		GoodInfo.
		Query().
		Where(
			goodinfo.Or(
				goodinfo.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query good info: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty result of good info")
	}

	return &npool.GetGoodResponse{
		Info: dbRowToInfo(infos[0]),
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteGoodRequest) (*npool.DeleteGoodResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	info, err := db.Client().
		GoodInfo.
		UpdateOneID(id).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to delete good info: %v", err)
	}

	return &npool.DeleteGoodResponse{
		Info: dbRowToInfo(info),
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetGoodsRequest) (*npool.GetGoodsResponse, error) {
	rows, err := db.Client().
		GoodInfo.
		Query().
		Where(
			goodinfo.Or(
				goodinfo.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query good info: %v", err)
	}

	infos := []*npool.GoodInfo{}
	for _, row := range rows {
		infos = append(infos, dbRowToInfo(row))
	}

	return &npool.GetGoodsResponse{
		Infos: infos,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetGoodsByAppRequest) (*npool.GetGoodsByAppResponse, error) {
	appUUID, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("cannot parse %v as uuid", in.AppID)
	}

	goodIDs, err := appID2GoodIDs(ctx, appUUID)
	if err != nil {
		return nil, err
	}
	if len(goodIDs) == 0 {
		return &npool.GetGoodsByAppResponse{
			Infos: []*npool.GoodInfo{},
			Total: 0,
		}, nil
	}

	respEnt, err := db.Client().GoodInfo.Query().
		Where(goodinfo.And(
			goodinfo.DeleteAt(0),
			goodinfo.IDIn(goodIDs...),
		)).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query goodinfo %v", err)
	}

	infos := []*npool.GoodInfo{}
	for _, row := range respEnt {
		infos = append(infos, dbRowToInfo(row))
	}

	return &npool.GetGoodsByAppResponse{
		Infos: infos,
		Total: 0,
	}, nil
}

func appID2GoodIDs(ctx context.Context, appID uuid.UUID) ([]uuid.UUID, error) {
	infos, err := db.Client().
		AppGood.
		Query().
		Where(
			appgood.And(
				appgood.AppID(appID),
				appgood.DeleteAt(0),
				appgood.Authorized(true),
				appgood.Online(true),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("failed to query app good %v", err)
	}

	goodIDs := make([]uuid.UUID, len(infos))
	for i, v := range infos {
		goodIDs[i] = v.GoodID
	}
	return goodIDs, nil
}
