package goodextrainfo

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodextrainfo"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func Create(ctx context.Context, in *npool.CreateGoodExtraInfoRequest) (*npool.CreateGoodExtraInfoResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	info, err := db.Client().
		GoodExtraInfo.
		Create().
		SetGoodID(id).
		SetPosters(in.GetInfo().GetPosters()).
		SetLabels(in.GetInfo().GetLabels()).
		SetPreSale(in.GetInfo().GetPreSale()).
		SetOutSale(in.GetInfo().GetOutSale()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &npool.CreateGoodExtraInfoResponse{
		Info: &npool.GoodExtraInfo{
			ID:      info.ID.String(),
			GoodID:  info.GoodID.String(),
			Posters: info.Posters,
			Labels:  info.Labels,
			OutSale: info.OutSale,
			PreSale: info.PreSale,
		},
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateGoodExtraInfoRequest) (*npool.UpdateGoodExtraInfoResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good extra info id: %v", err)
	}

	goodID, err := uuid.Parse(in.GetInfo().GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	_, err = db.Client().
		GoodExtraInfo.
		Query().
		Where(
			goodextrainfo.And(
				goodextrainfo.ID(id),
				goodextrainfo.GoodID(goodID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("mismatch good extra info: %v", err)
	}

	info, err := db.Client().
		GoodExtraInfo.
		UpdateOneID(id).
		SetPosters(in.GetInfo().GetPosters()).
		SetLabels(in.GetInfo().GetLabels()).
		SetOutSale(in.GetInfo().GetOutSale()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail set poster to good extra info: %v", err)
	}
	return &npool.UpdateGoodExtraInfoResponse{
		Info: &npool.GoodExtraInfo{
			ID:      info.ID.String(),
			GoodID:  info.GoodID.String(),
			Posters: info.Posters,
			Labels:  info.Labels,
			OutSale: info.OutSale,
			PreSale: info.PreSale,
		},
	}, nil
}

func Get(ctx context.Context, in *npool.GetGoodExtraInfoRequest) (*npool.GetGoodExtraInfoResponse, error) {
	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good extra info id: %v", err)
	}

	infos, err := db.Client().
		GoodExtraInfo.
		Query().
		Where(
			goodextrainfo.Or(
				goodextrainfo.GoodID(goodID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query good extra info: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty reply of good extra info")
	}

	return &npool.GetGoodExtraInfoResponse{
		Info: &npool.GoodExtraInfo{
			ID:      infos[0].ID.String(),
			GoodID:  infos[0].GoodID.String(),
			Posters: infos[0].Posters,
			Labels:  infos[0].Labels,
			OutSale: infos[0].OutSale,
			PreSale: infos[0].PreSale,
		},
	}, nil
}
