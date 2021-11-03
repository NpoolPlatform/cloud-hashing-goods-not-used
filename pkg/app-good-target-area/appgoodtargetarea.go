package appgoodtargetarea

import (
	"context"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/appgoodtargetarea"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateAppGoodTargetArea(info *npool.AppGoodTargetAreaInfo) error {
	_, err := uuid.Parse(info.GetAppID())
	if err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}

	_, err = uuid.Parse(info.GetGoodID())
	if err != nil {
		return xerrors.Errorf("invalid good id: %v", err)
	}

	_, err = uuid.Parse(info.GetTargetAreaID())
	if err != nil {
		return xerrors.Errorf("invalid target area id: %v", err)
	}

	return nil
}

func dbRowToAppGoodTargetArea(info *ent.AppGoodTargetArea) *npool.AppGoodTargetAreaInfo {
	return &npool.AppGoodTargetAreaInfo{
		ID:           info.ID.String(),
		AppID:        info.AppID.String(),
		GoodID:       info.GoodID.String(),
		TargetAreaID: info.TargetAreaID.String(),
	}
}

func Authorize(ctx context.Context, in *npool.AuthorizeAppGoodTargetAreaRequest) (*npool.AuthorizeAppGoodTargetAreaResponse, error) {
	if err := validateAppGoodTargetArea(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err == nil {
		info, err := db.Client().
			AppGoodTargetArea.
			UpdateOneID(id).
			SetDeleteAt(0).
			Save(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail authorize app good target area: %v", err)
		}
		return &npool.AuthorizeAppGoodTargetAreaResponse{
			Info: dbRowToAppGoodTargetArea(info),
		}, nil
	}

	info, err := db.Client().
		AppGoodTargetArea.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetTargetAreaID(uuid.MustParse(in.GetInfo().GetTargetAreaID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app good target area: %v", err)
	}

	return &npool.AuthorizeAppGoodTargetAreaResponse{
		Info: dbRowToAppGoodTargetArea(info),
	}, nil
}

func Check(ctx context.Context, in *npool.CheckAppGoodTargetAreaRequest) (*npool.CheckAppGoodTargetAreaResponse, error) {
	if err := validateAppGoodTargetArea(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	infos, err := db.Client().
		AppGoodTargetArea.
		Query().
		Where(
			appgoodtargetarea.And(
				appgoodtargetarea.AppID(uuid.MustParse(in.GetInfo().GetAppID())),
				appgoodtargetarea.GoodID(uuid.MustParse(in.GetInfo().GetGoodID())),
				appgoodtargetarea.TargetAreaID(uuid.MustParse(in.GetInfo().GetTargetAreaID())),
				appgoodtargetarea.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app good target area: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty app good target area")
	}

	return &npool.CheckAppGoodTargetAreaResponse{
		Info: dbRowToAppGoodTargetArea(infos[0]),
	}, nil
}

func Unauthorize(ctx context.Context, in *npool.UnauthorizeAppGoodTargetAreaRequest) (*npool.UnauthorizeAppGoodTargetAreaResponse, error) {
	if err := validateAppGoodTargetArea(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app good target area id: %v", err)
	}

	info, err := db.Client().
		AppGoodTargetArea.
		UpdateOneID(id).
		SetDeleteAt(time.Now().UnixNano()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail unauthorize app good target area: %v", err)
	}

	return &npool.UnauthorizeAppGoodTargetAreaResponse{
		Info: dbRowToAppGoodTargetArea(info),
	}, nil
}
