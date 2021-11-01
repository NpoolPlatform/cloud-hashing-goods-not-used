package vendorlocation

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/vendorlocation"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func Create(ctx context.Context, in *npool.CreateTargetAreaRequest) (*npool.CreateTargetAreaResponse, error) {
	info, err := db.Client().
		TargetArea.
		Create().
		SetContinent(in.GetInfo().GetContinent()).
		SetCountry(in.GetInfo().GetCountry()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &npool.CreateTargetAreaResponse{
		Info: &npool.TargetAreaInfo{
			ID:        info.ID.String(),
			Continent: info.Continent,
			Country:   info.Country,
		},
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateTargetAreaRequest) (*npool.UpdateTargetAreaResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid target area id: %v", err)
	}

	info, err := db.Client().
		TargetArea.
		UpdateOneID(id).
		SetContinent(in.GetInfo().GetContinent()).
		SetCountry(in.GetInfo().GetCountry()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
