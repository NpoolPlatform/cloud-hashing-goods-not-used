package vendorlocation

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"                      //nolint
	_ "github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/vendorlocation" //nolint

	_ "github.com/google/uuid" //nolint

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
	return nil, xerrors.Errorf("NOT IMPLEMENTED")
}

func Delete(ctx context.Context, in *npool.DeleteVendorLocationRequest) (*npool.DeleteVendorLocationResponse, error) {
	return nil, xerrors.Errorf("NOT IMPLEMENTED")
}

func GetAll(ctx context.Context, in *npool.GetVendorLocationsRequest) (*npool.GetVendorLocationsResponse, error) {
	return nil, xerrors.Errorf("NOT IMPLEMENTED")
}
