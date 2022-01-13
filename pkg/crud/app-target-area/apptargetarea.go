package apptargetarea

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/apptargetarea"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateAppTargetArea(info *npool.AppTargetAreaInfo) error {
	_, err := uuid.Parse(info.GetAppID())
	if err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}

	_, err = uuid.Parse(info.GetTargetAreaID())
	if err != nil {
		return xerrors.Errorf("invalid target area id: %v", err)
	}

	return nil
}

func Authorize(ctx context.Context, in *npool.AuthorizeAppTargetAreaRequest) (*npool.AuthorizeAppTargetAreaResponse, error) {
	if err := validateAppTargetArea(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err == nil {
		info, err := cli.
			AppTargetArea.
			UpdateOneID(id).
			SetDeleteAt(0).
			Save(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail update app target area: %v", err)
		}
		return &npool.AuthorizeAppTargetAreaResponse{
			Info: &npool.AppTargetAreaInfo{
				ID:           info.ID.String(),
				AppID:        info.AppID.String(),
				TargetAreaID: info.TargetAreaID.String(),
			},
		}, nil
	}

	info, err := cli.
		AppTargetArea.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetTargetAreaID(uuid.MustParse(in.GetInfo().GetTargetAreaID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app target area: %v", err)
	}

	return &npool.AuthorizeAppTargetAreaResponse{
		Info: &npool.AppTargetAreaInfo{
			ID:           info.ID.String(),
			AppID:        info.AppID.String(),
			TargetAreaID: info.TargetAreaID.String(),
		},
	}, nil
}

func Check(ctx context.Context, in *npool.CheckAppTargetAreaRequest) (*npool.CheckAppTargetAreaResponse, error) {
	if err := validateAppTargetArea(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppTargetArea.
		Query().
		Where(
			apptargetarea.And(
				apptargetarea.AppID(uuid.MustParse(in.GetInfo().GetAppID())),
				apptargetarea.TargetAreaID(uuid.MustParse(in.GetInfo().GetTargetAreaID())),
				apptargetarea.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app target area: %v", err)
	}

	authorized := true
	id := uuid.UUID{}

	if len(infos) == 0 {
		authorized = false
	} else {
		id = infos[0].ID
	}

	return &npool.CheckAppTargetAreaResponse{
		Info: &npool.AppTargetAreaInfo{
			ID:           id.String(),
			AppID:        in.GetInfo().GetAppID(),
			TargetAreaID: in.GetInfo().GetTargetAreaID(),
		},
		Authorized: authorized,
	}, nil
}

func Unauthorize(ctx context.Context, in *npool.UnauthorizeAppTargetAreaRequest) (*npool.UnauthorizeAppTargetAreaResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("fail unauthorize app target area: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppTargetArea.
		UpdateOneID(id).
		SetDeleteAt(time.Now().UnixNano()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail unauthorize app target area: %v", err)
	}
	return &npool.UnauthorizeAppTargetAreaResponse{
		Info: &npool.AppTargetAreaInfo{
			ID:           info.ID.String(),
			AppID:        info.AppID.String(),
			TargetAreaID: info.TargetAreaID.String(),
		},
	}, nil
}
