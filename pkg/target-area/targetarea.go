package targetarea

import (
	"context"
	"fmt"
	"time"

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
		SetUpdateAt(time.Now().UnixNano()).
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
	info, err := db.Client().
		TargetArea.
		UpdateOneID(uuid.MustParse(in.GetID())).
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
	infos, err := db.Client().
		TargetArea.
		Query().
		Where(
			targetarea.Or(
				targetarea.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println(infos)

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
