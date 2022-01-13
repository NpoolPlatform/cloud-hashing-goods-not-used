package targetarea

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/targetarea"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func Create(ctx context.Context, in *npool.CreateTargetAreaRequest) (*npool.CreateTargetAreaResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
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

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
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

func Get(ctx context.Context, in *npool.GetTargetAreaRequest) (*npool.GetTargetAreaResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid target area id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		TargetArea.
		Query().
		Where(
			targetarea.Or(
				targetarea.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query target area: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty reply of target area")
	}

	return &npool.GetTargetAreaResponse{
		Info: &npool.TargetAreaInfo{
			ID:        infos[0].ID.String(),
			Continent: infos[0].Continent,
			Country:   infos[0].Country,
		},
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteTargetAreaRequest) (*npool.DeleteTargetAreaResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid target area id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		TargetArea.
		UpdateOneID(id).
		SetDeleteAt(time.Now().UnixNano()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to delete target area: %v", err)
	}

	return &npool.DeleteTargetAreaResponse{
		Info: &npool.TargetAreaInfo{
			ID:        info.ID.String(),
			Continent: info.Continent,
			Country:   info.Country,
		},
	}, nil
}

func DeleteByContinentCountry(ctx context.Context, in *npool.DeleteTargetAreaByContinentCountryRequest) (*npool.DeleteTargetAreaByContinentCountryResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		TargetArea.
		Query().
		Where(
			targetarea.And(
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

	info, err := Delete(ctx, &npool.DeleteTargetAreaRequest{
		ID: infos[0].ID.String(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail to delete target area: %v", err)
	}
	return &npool.DeleteTargetAreaByContinentCountryResponse{
		Info: info.Info,
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetTargetAreasRequest) (*npool.GetTargetAreasResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		TargetArea.
		Query().
		Where(
			targetarea.And(
				targetarea.DeleteAt(0),
			),
		).
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
