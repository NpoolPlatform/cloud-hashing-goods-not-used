package goodinfo

import (
	"context"
	_ "fmt"
	_ "time"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodinfo"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func dbRowToInfo(row *ent.GoodInfo) *npool.GoodInfo {
	ids := []string{}

	for _, id := range row.SupportCoinTypeIds {
		ids = append(ids, id.String())
	}

	return &npool.GoodInfo{
		ID:                 row.ID.String(),
		DeviceInfoID:       row.DeviceInfoID.String(),
		GasPrice:           row.GasPrice,
		SeparateGasFee:     row.SeparateGasFee,
		UnitPower:          row.UnitPower,
		DurationDays:       row.DurationDays,
		CoinInfoID:         row.CoinInfoID.String(),
		Actuals:            row.Actuals,
		DeliveryAt:         row.DeliveryAt,
		InheritFromGoodID:  row.InheritFromGoodID.String(),
		VendorLocationID:   row.VendorLocationID.String(),
		Price:              row.Price,
		BenefitType:        string(row.BenefitType),
		Classic:            row.Classic,
		SupportCoinTypeIDs: ids,
		Total:              row.Total,
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

	info, err := db.Client().
		GoodInfo.
		Create().
		SetDeviceInfoID(uuid.MustParse(in.GetInfo().GetDeviceInfoID())).
		SetGasPrice(in.GetInfo().GetGasPrice()).
		SetSeparateGasFee(in.GetInfo().GetSeparateGasFee()).
		SetUnitPower(in.GetInfo().GetUnitPower()).
		SetDurationDays(in.GetInfo().GetDurationDays()).
		SetCoinInfoID(uuid.MustParse(in.GetInfo().GetCoinInfoID())).
		SetActuals(in.GetInfo().GetActuals()).
		SetDeliveryAt(in.GetInfo().GetDeliveryAt()).
		SetInheritFromGoodID(uuid.MustParse(in.GetInfo().GetInheritFromGoodID())).
		SetVendorLocationID(uuid.MustParse(in.GetInfo().GetVendorLocationID())).
		SetPrice(in.GetInfo().GetPrice()).
		SetBenefitType(goodinfo.BenefitType(in.GetInfo().GetBenefitType())).
		SetClassic(in.GetInfo().GetClassic()).
		SetSupportCoinTypeIds(ids).
		SetTotal(in.GetInfo().GetTotal()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to create good: %v", err)
	}

	return &npool.CreateGoodResponse{
		Info: dbRowToInfo(info),
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetGoodsRequest) (*npool.GetGoodsResponse, error) {
	return nil, nil
}
