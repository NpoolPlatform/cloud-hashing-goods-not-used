package targetarea

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/targetarea"

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
	return &npool.UpdateTargetAreaResponse{
		Info: &npool.TargetAreaInfo{
			ID:        info.ID.String(),
			Continent: info.Continent,
			Country:   info.Country,
		},
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteTargetAreaRequest) (*npool.DeleteTargetAreaResponse, error) {
	infos, err := db.Client().
		TargetArea.
		Query().
		Where(
			targetarea.Or(
				targetarea.ID(uuid.MustParse(in.GetID())),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query target area: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("invalid target area")
	}

	err = db.Client().
		TargetArea.
		DeleteOneID(uuid.MustParse(in.GetID())).
		Exec(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to delete target area: %v", err)
	}

	return &npool.DeleteTargetAreaResponse{
		Info: &npool.TargetAreaInfo{
			ID:        infos[0].ID.String(),
			Continent: infos[0].Continent,
			Country:   infos[0].Country,
		},
	}, nil
}

func DeleteByContinentCountry(ctx context.Context, in *npool.DeleteTargetAreaByContinentCountryRequest) (*npool.DeleteTargetAreaByContinentCountryResponse, error) {
	infos, err := db.Client().
		TargetArea.
		Query().
		Where(
			targetarea.Or(
				targetarea.Continent(in.GetContinent()),
				targetarea.Country(in.GetCountry()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query target area: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("invalid target area")
	}

	err = db.Client().
		TargetArea.
		DeleteOne(infos[0]).
		Exec(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to delete target area: %v", err)
	}

	return &npool.DeleteTargetAreaByContinentCountryResponse{
		Info: &npool.TargetAreaInfo{
			ID:        infos[0].ID.String(),
			Continent: infos[0].Continent,
			Country:   infos[0].Country,
		},
	}, nil
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
