package targetarea

import (
	"context"
	"database/sql"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/targetarea"
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

	if err == nil {
		if len(info1) > 0 {
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
	infos, err := db.Client().
		TargetArea.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}

	areas := []*npool.TargetAreaInfo{}
	for _, info := range infos {
		areas = append(areas, &npool.TargetAreaInfo{
			ID:        info.ID.String(),
			Continent: info.Continent,
			Country:   info.Country,
		})
	}

	return &npool.GetTargetAreasResponse{
		Infos: areas,
	}, nil
}
