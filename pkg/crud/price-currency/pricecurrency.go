package pricecurrency

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/pricecurrency"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func Create(ctx context.Context, in *npool.CreatePriceCurrencyRequest) (*npool.CreatePriceCurrencyResponse, error) {
	info, err := db.Client().
		PriceCurrency.
		Create().
		SetName(in.GetInfo().GetName()).
		SetUnit(in.GetInfo().GetUnit()).
		SetSymbol(in.GetInfo().GetSymbol()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &npool.CreatePriceCurrencyResponse{
		Info: &npool.PriceCurrency{
			ID:     info.ID.String(),
			Name:   info.Name,
			Unit:   info.Unit,
			Symbol: info.Symbol,
		},
	}, nil
}

func Update(ctx context.Context, in *npool.UpdatePriceCurrencyRequest) (*npool.UpdatePriceCurrencyResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid price currency id: %v", err)
	}

	info, err := db.Client().
		PriceCurrency.
		UpdateOneID(id).
		SetName(in.GetInfo().GetName()).
		SetUnit(in.GetInfo().GetUnit()).
		SetSymbol(in.GetInfo().GetSymbol()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &npool.UpdatePriceCurrencyResponse{
		Info: &npool.PriceCurrency{
			ID:     info.ID.String(),
			Name:   info.Name,
			Unit:   info.Unit,
			Symbol: info.Symbol,
		},
	}, nil
}

func Get(ctx context.Context, in *npool.GetPriceCurrencyRequest) (*npool.GetPriceCurrencyResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid price currency id: %v", err)
	}

	infos, err := db.Client().
		PriceCurrency.
		Query().
		Where(
			pricecurrency.Or(
				pricecurrency.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query price currency: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty reply of price currency")
	}

	return &npool.GetPriceCurrencyResponse{
		Info: &npool.PriceCurrency{
			ID:     infos[0].ID.String(),
			Name:   infos[0].Name,
			Unit:   infos[0].Unit,
			Symbol: infos[0].Symbol,
		},
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetPriceCurrencysRequest) (*npool.GetPriceCurrencysResponse, error) {
	infos, err := db.Client().
		PriceCurrency.
		Query().
		Where(
			pricecurrency.And(
				pricecurrency.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println(infos)

	areas := []*npool.PriceCurrency{}
	for _, info := range infos {
		areas = append(areas, &npool.PriceCurrency{
			ID:     info.ID.String(),
			Name:   info.Name,
			Unit:   info.Unit,
			Symbol: info.Symbol,
		})
	}

	return &npool.GetPriceCurrencysResponse{
		Infos: areas,
	}, nil
}
