package appareaauth

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

func Authorize(ctx context.Context, in *npool.AuthorizeAppTargetAreaRequest) (*npool.AuthorizeAppTargetAreaResponse, error) {
	info, err := db.Client().
		AppArea.
		Create().
		SetContinent(in.GetInfo().GetContinent()).
		SetCountry(in.GetInfo().GetCountry()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &npool.CreateAppAreaResponse{
		Info: &npool.AppAreaInfo{
			ID:        info.ID.String(),
			Continent: info.Continent,
			Country:   info.Country,
		},
	}, nil
}

func Unauthorize(ctx context.Context, in *npool.UpdateAppAreaRequest) (*npool.UpdateTargetAreaResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid target area id: %v", err)
	}

	info, err := db.Client().
		AppArea.
		UpdateOneID(id).
		SetContinent(in.GetInfo().GetContinent()).
		SetCountry(in.GetInfo().GetCountry()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
}

func Check(ctx context.Context, in *npool.CheckAppTargetAreaRequest)
