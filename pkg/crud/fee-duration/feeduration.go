package feeduration

import (
	"context"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/feeduration"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateFeeDuration(info *npool.FeeDuration) error {
	if _, err := uuid.Parse(info.GetFeeTypeID()); err != nil {
		return xerrors.Errorf("invalid fee duration id: %v", err)
	}
	if info.GetDuration() < 0 {
		return xerrors.New("duration(days) shall not be negative")
	}

	return nil
}

func Create(ctx context.Context, in *npool.CreateFeeDurationRequest) (*npool.CreateFeeDurationResponse, error) {
	if err := validateFeeDuration(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		FeeDuration.
		Create().
		SetFeeTypeID(uuid.MustParse(in.GetInfo().GetFeeTypeID())).
		SetDuration(in.GetInfo().Duration).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &npool.CreateFeeDurationResponse{
		Info: dbRowToFeeDuration(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateFeeDurationRequest) (*npool.UpdateFeeDurationResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid fee duration id: %v", err)
	}
	if in.GetInfo().GetDuration() < 0 {
		return nil, xerrors.New("duration(days) shall not be negative")
	}
	ftid, err := uuid.Parse(in.GetInfo().GetFeeTypeID())
	if err != nil {
		return nil, xerrors.New("invalid fee type id")
	}

	info, err := db.Client().
		FeeDuration.
		UpdateOneID(id).
		SetDuration(in.Info.Duration).
		SetFeeTypeID(ftid).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("failed to update fee duration id: %v", err)
	}

	return &npool.UpdateFeeDurationResponse{
		Info: dbRowToFeeDuration(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetFeeDurationRequest) (*npool.GetFeeDurationResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid fee duration id: %v", in.GetID())
	}

	infos, err := db.Client().
		FeeDuration.
		Query().
		Where(
			feeduration.Or(
				feeduration.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("failed to query fee duration: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty reply of fee duration")
	}

	return &npool.GetFeeDurationResponse{
		Info: &npool.FeeDuration{
			ID:        infos[0].ID.String(),
			FeeTypeID: infos[0].FeeTypeID.String(),
			Duration:  infos[0].Duration,
		},
	}, nil
}

func GetByFeeType(ctx context.Context, in *npool.GetFeeDurationsByFeeTypeRequest) (*npool.GetFeeDurationsByFeeTypeResponse, error) {
	id, err := uuid.Parse(in.GetFeeTypeID())
	if err != nil {
		return nil, xerrors.Errorf("invalid fee duration id: %v", err)
	}

	infos, err := db.Client().FeeDuration.Query().
		Where(feeduration.FeeTypeID(id)).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("failed to query fee durations: %v", err)
	}

	return &npool.GetFeeDurationsByFeeTypeResponse{
		Infos: dbRowsToFeeDurations(infos),
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteFeeDurationRequest) (*npool.DeleteFeeDurationResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid fee duration id: %v", err)
	}

	info, err := db.Client().
		FeeDuration.
		UpdateOneID(id).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to delete fee duration: %v", err)
	}

	return &npool.DeleteFeeDurationResponse{
		Info: dbRowToFeeDuration(info),
	}, nil
}

func dbRowToFeeDuration(info *ent.FeeDuration) *npool.FeeDuration {
	return &npool.FeeDuration{
		ID:        info.ID.String(),
		FeeTypeID: info.FeeTypeID.String(),
		Duration:  info.Duration,
	}
}

func dbRowsToFeeDurations(infos []*ent.FeeDuration) []*npool.FeeDuration {
	feeDurations := make([]*npool.FeeDuration, len(infos))
	for i := 0; i < len(infos); i++ {
		feeDurations[i] = dbRowToFeeDuration(infos[i])
	}
	return feeDurations
}
