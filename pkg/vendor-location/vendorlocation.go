package vendorlocation

import (
	"context"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"                    //nolint
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/vendorlocation" //nolint

	"github.com/google/uuid" //nolint

	"golang.org/x/xerrors"
)

func Create(ctx context.Context, in *npool.CreateVendorLocationRequest) (*npool.CreateVendorLocationResponse, error) {
	info, err := db.Client().
		VendorLocation.
		Create().
		SetCountry(in.GetInfo().GetCountry()).
		SetProvince(in.GetInfo().GetProvince()).
		SetCity(in.GetInfo().GetCity()).
		SetAddress(in.GetInfo().GetAddress()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &npool.CreateVendorLocationResponse{
		Info: &npool.VendorLocationInfo{
			ID:       info.ID.String(),
			Country:  info.Country,
			Province: info.Province,
			City:     info.City,
			Address:  info.Address,
		},
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateVendorLocationRequest) (*npool.UpdateVendorLocationResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid vendor location id: %v", err)
	}

	info, err := db.Client().
		VendorLocation.
		UpdateOneID(id).
		SetCountry(in.GetInfo().GetCountry()).
		SetProvince(in.GetInfo().GetProvince()).
		SetCity(in.GetInfo().GetCity()).
		SetAddress(in.GetInfo().GetAddress()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &npool.UpdateVendorLocationResponse{
		Info: &npool.VendorLocationInfo{
			ID:       info.ID.String(),
			Country:  info.Country,
			Province: info.Province,
			City:     info.City,
			Address:  info.Address,
		},
	}, nil
}

func Get(ctx context.Context, in *npool.GetVendorLocationRequest) (*npool.GetVendorLocationResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid vendor location id: %v", err)
	}

	infos, err := db.Client().
		VendorLocation.
		Query().
		Where(
			vendorlocation.Or(
				vendorlocation.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query vendor location: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty reply of vendor location: %v", err)
	}

	return &npool.GetVendorLocationResponse{
		Info: &npool.VendorLocationInfo{
			ID:       id.String(),
			Country:  infos[0].Country,
			Province: infos[0].Province,
			City:     infos[0].City,
			Address:  infos[0].Address,
		},
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteVendorLocationRequest) (*npool.DeleteVendorLocationResponse, error) {
	info, err := db.Client().
		VendorLocation.
		UpdateOneID(uuid.MustParse(in.GetID())).
		SetDeleteAt(time.Now().UnixNano()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to delete target area: %v", err)
	}

	return &npool.DeleteVendorLocationResponse{
		Info: &npool.VendorLocationInfo{
			ID:       info.ID.String(),
			Country:  info.Country,
			Province: info.Province,
			City:     info.City,
			Address:  info.Address,
		},
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetVendorLocationsRequest) (*npool.GetVendorLocationsResponse, error) {
	return nil, xerrors.Errorf("NOT IMPLEMENTED")
}
