package feetype

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/feetype"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateFeeType(info *npool.FeeType) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}

	return nil
}

func Create(ctx context.Context, in *npool.CreateFeeTypeRequest) (*npool.CreateFeeTypeResponse, error) {
	if err := validateFeeType(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		FeeType.
		Create().
		SetFeeType(in.GetInfo().GetFeeType()).
		SetFeeDescription(in.GetInfo().GetFeeDescription()).
		SetPayType(feetype.PayType(in.GetInfo().GetPayType())).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &npool.CreateFeeTypeResponse{
		Info: &npool.FeeType{
			ID:             info.ID.String(),
			FeeType:        info.FeeType,
			FeeDescription: info.FeeDescription,
			PayType:        string(info.PayType),
		},
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateFeeTypeRequest) (*npool.UpdateFeeTypeResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good fee id: %v", err)
	}

	if err := validateFeeType(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		FeeType.
		UpdateOneID(id).
		SetFeeType(in.GetInfo().GetFeeType()).
		SetFeeDescription(in.GetInfo().GetFeeDescription()).
		SetPayType(feetype.PayType(in.GetInfo().GetPayType())).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &npool.UpdateFeeTypeResponse{
		Info: &npool.FeeType{
			ID:             info.ID.String(),
			FeeType:        info.FeeType,
			FeeDescription: info.FeeDescription,
			PayType:        string(info.PayType),
		},
	}, nil
}

func Get(ctx context.Context, in *npool.GetFeeTypeRequest) (*npool.GetFeeTypeResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good fee id: %v", err)
	}

	infos, err := db.Client().
		FeeType.
		Query().
		Where(
			FeeType.Or(
				FeeType.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query good fee: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty reply of good fee")
	}

	return &npool.GetFeeTypeResponse{
		Info: &npool.FeeType{
			ID:             infos[0].ID.String(),
			FeeType:        infos[0].FeeType,
			FeeDescription: infos[0].FeeDescription,
			PayType:        string(infos[0].PayType),
		},
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetFeeTypesRequest) (*npool.GetFeeTypesResponse, error) {
	infos, err := db.Client().
		FeeType.
		Query().
		Where(
			FeeType.And(
				FeeType.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	fees := []*npool.FeeType{}
	for _, info := range infos {
		fees = append(fees, &npool.FeeType{
			ID:             info.ID.String(),
			FeeType:        info.FeeType,
			FeeDescription: info.FeeDescription,
			PayType:        string(info.PayType),
		})
	}

	return &npool.GetFeeTypesResponse{
		Infos: fees,
	}, nil
}
