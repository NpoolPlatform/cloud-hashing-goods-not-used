package targetarea

import (
	"context"
	"database/sql"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/targetarea"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
)

func Create(ctx context.Context, in *npool.CreateTargetAreaRequest) (*npool.CreateTargetAreaResponse, error) {
	info1, err := db.Client().
		TargetArea.
		Query().
		Where(
			targetarea.And(
				targetarea.Continent(in.GetInfo().GetContinent()),
				targetarea.Country(in.GetInfo().GetCountry()),
			),
		).
		Limit(1).
		All(ctx)

	logger.Sugar().Infof("%v: %v", info1, err)
	if err == nil {
		if 0 < len(info1) {
			return &npool.CreateTargetAreaResponse{
				Info: &npool.TargetAreaInfo{
					ID:        info1[0].ID.String(),
					Continent: info1[0].Continent,
					Country:   info1[0].Country,
				},
			}, nil
		}
	} else if err != sql.ErrNoRows {
		return nil, err
	}

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
	return nil, nil
}

func GetAll(ctx context.Context, in *npool.GetTargetAreasRequest) (*npool.GetTargetAreasResponse, error) {
	return nil, nil
}
