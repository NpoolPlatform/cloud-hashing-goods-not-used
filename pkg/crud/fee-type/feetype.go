package feetype

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/feetype"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateFeeType(info *npool.FeeType) error {
	return nil
}

func Create(ctx context.Context, in *npool.CreateFeeTypeRequest) (*npool.CreateFeeTypeResponse, error) {
	if err := validateFeeType(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
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
		return nil, xerrors.Errorf("invalid fee type id: %v", err)
	}

	if err := validateFeeType(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
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
		return nil, xerrors.Errorf("invalid fee type id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		FeeType.
		Query().
		Where(
			feetype.Or(
				feetype.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query fee type: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty reply of fee type")
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
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		FeeType.
		Query().
		Where(
			feetype.And(
				feetype.DeleteAt(0),
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
