package vendorlocation

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	_ "github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"                //nolint
	_ "github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/targetarea" //nolint

	_ "github.com/google/uuid" //nolint

	"golang.org/x/xerrors"
)

func Create(ctx context.Context, in *npool.CreateVendorLocationRequest) (*npool.CreateVendorLocationResponse, error) {
	return nil, xerrors.Errorf("NOT IMPLEMENTED")
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
