package goodreview

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodreview"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func dbRowToGoodReview(info *ent.GoodReview) *npool.GoodReviewInfo {
	return &npool.GoodReviewInfo{
		ID:         info.ID.String(),
		ReviewedID: info.ReviewedID.String(),
		ReviewerID: info.ReviewerID.String(),
		Type:       string(info.EntityType),
		State:      string(info.State),
		Message:    info.Message,
	}
}

func validateGoodReview(info *npool.GoodReviewInfo) error {
	_, err := uuid.Parse(info.GetReviewedID())
	if err != nil {
		return xerrors.Errorf("invalid reviewed id: %v", err)
	}

	return nil
}

func Create(ctx context.Context, in *npool.CreateGoodReviewRequest) (*npool.CreateGoodReviewResponse, error) {
	if err := validateGoodReview(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		GoodReview.
		Create().
		SetReviewedID(uuid.MustParse(in.GetInfo().GetReviewedID())).
		SetEntityType(goodreview.EntityType(in.GetInfo().GetType())).
		SetState(goodreview.State(in.GetInfo().GetState())).
		SetMessage(in.GetInfo().GetMessage()).
		SetReviewerID(uuid.MustParse(in.GetInfo().GetReviewerID())).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &npool.CreateGoodReviewResponse{
		Info: dbRowToGoodReview(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateGoodReviewRequest) (*npool.UpdateGoodReviewResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good extra info id: %v", err)
	}

	if err := validateGoodReview(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	_, err = db.Client().
		GoodReview.
		Query().
		Where(
			goodreview.And(
				goodreview.ID(id),
				goodreview.ReviewedID(uuid.MustParse(in.GetInfo().GetReviewedID())),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("mismatch good extra info: %v", err)
	}

	info, err := db.Client().
		GoodReview.
		UpdateOneID(id).
		SetReviewerID(uuid.MustParse(in.GetInfo().GetReviewerID())).
		SetState(goodreview.State(in.GetInfo().GetState())).
		SetMessage(in.GetInfo().GetMessage()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail set poster to good extra info: %v", err)
	}
	return &npool.UpdateGoodReviewResponse{
		Info: dbRowToGoodReview(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetGoodReviewRequest) (*npool.GetGoodReviewResponse, error) {
	if err := validateGoodReview(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	infos, err := db.Client().
		GoodReview.
		Query().
		Where(
			goodreview.Or(
				goodreview.ReviewedID(uuid.MustParse(in.GetInfo().GetReviewedID())),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query good extra info: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty reply of good extra info")
	}

	return &npool.GetGoodReviewResponse{
		Info: dbRowToGoodReview(infos[0]),
	}, nil
}
