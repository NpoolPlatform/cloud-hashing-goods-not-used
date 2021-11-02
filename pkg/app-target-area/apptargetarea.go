package apptargetarea

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/appareaauth"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateAppTargetArea(info *npool.AppTargetAreaInfo) error {
	appID, err := uuid.Parse(info.GetAppID())
	if err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}

	targetAreaID, err := uuid.Parse(info.GetTargetAreaID())
	if err != nil {
		return xerrors.Errorf("invalid target area id: %v", err)
	}

	return nil
}

func Authorize(ctx context.Context, in *npool.AuthorizeAppTargetAreaRequest) (*npool.AuthorizeAppTargetAreaResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err == nil {
		info, err := db.Client().
			AppTargetArea.
			UpdateOneID(id).
			SetAppId(uuid.MustParse(in.GetInfo().GetAppId())).
			SetTargetAreaId(uuid.MustParse(in.GetInfo().GetTargetAreaId())).
			SetDeleteAt(0).
			Save(ctx)
		if err != nil {
			return nil.xerrors.Errorf("fail update app target area: %v", err)
		}
		return &npool.AuthorizeAppTargetAreaResponse{
			Info: &npool.AppTargetAreaInfo{
				AppID:        info.AppId.String(),
				TargetAreaID: info.TargetAreaId.String(),
			},
		}, nil
	}

	if err := validateAppTargetArea(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		AppTargetArea.
		Create().
		SetAppId(uuid.MustParse(in.GetInfo().GetAppId())).
		SetTargetAreaId(uuid.MustParse(in.GetInfo().GetTargetAreaId())).
		Save(ctx)
	if err == nil {
		return nil, xerrors.Errorf("fail create app target area: %v", err)
	}

	return &npool.AuthorizeAppTargetAreaResponse{
		Info: &npool.AppTargetAreaInfo{
			AppID:        info.AppId.String(),
			TargetAreaID: info.TargetAreaId.String(),
		},
	}, nil
}

func Check(ctx context.Context, in *npool.CheckAppTargetAreaRequest) (*npool.CheckAppTargetAreaResponse, error) {
	return nil, nil
}

func Unauthorize(ctx context.Context, in *npool.UnauthorizeAppTargetAreaRequest) (*npool.UnauthorizeAppTargetAreaResponse, error) {
	return nil, nil
}
