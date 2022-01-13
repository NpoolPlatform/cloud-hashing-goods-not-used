package deviceinfo

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"                //nolint
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/deviceinfo" //nolint

	"github.com/google/uuid" //nolint

	"golang.org/x/xerrors"
)

func validateDeviceInfo(in *npool.DeviceInfo) error {
	if in.GetType() == "" {
		return xerrors.Errorf("invalid device type")
	}
	if in.GetManufacturer() == "" {
		return xerrors.Errorf("invalid device manufacturer")
	}
	return nil
}

func Create(ctx context.Context, in *npool.CreateDeviceInfoRequest) (*npool.CreateDeviceInfoResponse, error) {
	if err := validateDeviceInfo(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		DeviceInfo.
		Create().
		SetType(in.GetInfo().GetType()).
		SetManufacturer(in.GetInfo().GetManufacturer()).
		SetPowerComsuption(in.GetInfo().GetPowerComsuption()).
		SetShipmentAt(in.GetInfo().GetShipmentAt()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &npool.CreateDeviceInfoResponse{
		Info: &npool.DeviceInfo{
			ID:              info.ID.String(),
			Type:            info.Type,
			Manufacturer:    info.Manufacturer,
			PowerComsuption: info.PowerComsuption,
			ShipmentAt:      info.ShipmentAt,
		},
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateDeviceInfoRequest) (*npool.UpdateDeviceInfoResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid device info id: %v", err)
	}

	if err := validateDeviceInfo(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		DeviceInfo.
		UpdateOneID(id).
		SetType(in.GetInfo().GetType()).
		SetManufacturer(in.GetInfo().GetManufacturer()).
		SetPowerComsuption(in.GetInfo().GetPowerComsuption()).
		SetShipmentAt(in.GetInfo().GetShipmentAt()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &npool.UpdateDeviceInfoResponse{
		Info: &npool.DeviceInfo{
			ID:              info.ID.String(),
			Type:            info.Type,
			Manufacturer:    info.Manufacturer,
			PowerComsuption: info.PowerComsuption,
			ShipmentAt:      info.ShipmentAt,
		},
	}, nil
}

func Get(ctx context.Context, in *npool.GetDeviceInfoRequest) (*npool.GetDeviceInfoResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid device info id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		DeviceInfo.
		Query().
		Where(
			deviceinfo.Or(
				deviceinfo.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query device info: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty reply of device info: %v", err)
	}

	return &npool.GetDeviceInfoResponse{
		Info: &npool.DeviceInfo{
			ID:              id.String(),
			Type:            infos[0].Type,
			Manufacturer:    infos[0].Manufacturer,
			PowerComsuption: infos[0].PowerComsuption,
			ShipmentAt:      infos[0].ShipmentAt,
		},
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteDeviceInfoRequest) (*npool.DeleteDeviceInfoResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		DeviceInfo.
		UpdateOneID(uuid.MustParse(in.GetID())).
		SetDeleteAt(time.Now().UnixNano()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to delete device info: %v", err)
	}

	return &npool.DeleteDeviceInfoResponse{
		Info: &npool.DeviceInfo{
			ID:              info.ID.String(),
			Type:            info.Type,
			Manufacturer:    info.Manufacturer,
			PowerComsuption: info.PowerComsuption,
			ShipmentAt:      info.ShipmentAt,
		},
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetDeviceInfosRequest) (*npool.GetDeviceInfosResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		DeviceInfo.
		Query().
		Where(
			deviceinfo.Or(
				deviceinfo.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	devices := []*npool.DeviceInfo{}
	for _, info := range infos {
		devices = append(devices, &npool.DeviceInfo{
			ID:              info.ID.String(),
			Type:            info.Type,
			Manufacturer:    info.Manufacturer,
			PowerComsuption: info.PowerComsuption,
			ShipmentAt:      info.ShipmentAt,
		})
	}

	return &npool.GetDeviceInfosResponse{
		Infos: devices,
	}, nil
}
