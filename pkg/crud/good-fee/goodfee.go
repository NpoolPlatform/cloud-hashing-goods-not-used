package goodfee

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodfee"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateGoodFee(info *npool.GoodFee) error {
	return nil
}

func Create(ctx context.Context, in *npool.CreateGoodFeeRequest) (*npool.CreateGoodFeeResponse, error) {
	if err := validateGoodFee(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		GoodFee.
		Create().
		SetFeeType(in.GetInfo().GetFeeType()).
		SetFeeDescription(in.GetInfo().GetFeeDescription()).
		SetPayType(goodfee.PayType(in.GetInfo().GetPayType())).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &npool.CreateGoodFeeResponse{
		Info: &npool.GoodFee{
			ID:             info.ID.String(),
			FeeType:        info.FeeType,
			FeeDescription: info.FeeDescription,
			PayType:        string(info.PayType),
		},
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateGoodFeeRequest) (*npool.UpdateGoodFeeResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good fee id: %v", err)
	}

	if err := validateGoodFee(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		GoodFee.
		UpdateOneID(id).
		SetFeeType(in.GetInfo().GetFeeType()).
		SetFeeDescription(in.GetInfo().GetFeeDescription()).
		SetPayType(goodfee.PayType(in.GetInfo().GetPayType())).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &npool.UpdateGoodFeeResponse{
		Info: &npool.GoodFee{
			ID:             info.ID.String(),
			FeeType:        info.FeeType,
			FeeDescription: info.FeeDescription,
			PayType:        string(info.PayType),
		},
	}, nil
}

func Get(ctx context.Context, in *npool.GetGoodFeeRequest) (*npool.GetGoodFeeResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good fee id: %v", err)
	}

	infos, err := db.Client().
		GoodFee.
		Query().
		Where(
			goodfee.Or(
				goodfee.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query good fee: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty reply of good fee")
	}

	return &npool.GetGoodFeeResponse{
		Info: &npool.GoodFee{
			ID:             infos[0].ID.String(),
			FeeType:        infos[0].FeeType,
			FeeDescription: infos[0].FeeDescription,
			PayType:        string(infos[0].PayType),
		},
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetGoodFeesRequest) (*npool.GetGoodFeesResponse, error) {
	infos, err := db.Client().
		GoodFee.
		Query().
		Where(
			goodfee.And(
				goodfee.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	fees := []*npool.GoodFee{}
	for _, info := range infos {
		fees = append(fees, &npool.GoodFee{
			ID:             info.ID.String(),
			FeeType:        info.FeeType,
			FeeDescription: info.FeeDescription,
			PayType:        string(info.PayType),
		})
	}

	return &npool.GetGoodFeesResponse{
		Infos: fees,
	}, nil
}
