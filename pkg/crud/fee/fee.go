package fee

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/fee"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateFee(info *npool.Fee) error {
	if _, err := uuid.Parse(info.GetFeeTypeID()); err != nil {
		return xerrors.Errorf("invalid fee type id: %v", err)
	}
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}

	return nil
}

func Create(ctx context.Context, in *npool.CreateFeeRequest) (*npool.CreateFeeResponse, error) {
	if err := validateFee(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		Fee.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetFeeTypeID(uuid.MustParse(in.GetInfo().GetFeeTypeID())).
		SetValue(price.VisualPriceToDBPrice(in.GetInfo().GetValue())).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &npool.CreateFeeResponse{
		Info: &npool.Fee{
			ID:        info.ID.String(),
			AppID:     info.AppID.String(),
			FeeTypeID: info.FeeTypeID.String(),
			Value:     price.DBPriceToVisualPrice(info.Value),
		},
	}, nil
}

func Get(ctx context.Context, in *npool.GetFeeRequest) (*npool.GetFeeResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid fee id: %v", err)
	}

	infos, err := db.Client().
		Fee.
		Query().
		Where(
			fee.Or(
				fee.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query fee: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty reply of fee")
	}

	return &npool.GetFeeResponse{
		Info: &npool.Fee{
			ID:        infos[0].ID.String(),
			AppID:     infos[0].AppID.String(),
			FeeTypeID: infos[0].FeeTypeID.String(),
			Value:     price.DBPriceToVisualPrice(infos[0].Value),
		},
	}, nil
}
