package targetarea

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
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

	logger.Sugar().Infof("%v", in.GetInfo())
	logger.Sugar().Infof("%v", info)

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
