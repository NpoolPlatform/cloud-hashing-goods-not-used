package recommend

import (
	"context"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

const (
	dbTimeout = 5 * time.Second
)

func validateRecommend(info *npool.Recommend) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		return xerrors.Errorf("invalid good id: %v", err)
	}
	if _, err := uuid.Parse(info.GetRecommenderID()); err != nil {
		return xerrors.Errorf("invalid recommender id: %v", err)
	}
	if info.GetMessage() == "" {
		return xerrors.Errorf("invalid recommend message")
	}
	return nil
}

func dbRowToRecommend(row *ent.Recommend) *npool.Recommend {
	return &npool.Recommend{
		ID:            row.ID.String(),
		AppID:         row.AppID.String(),
		GoodID:        row.GoodID.String(),
		RecommenderID: row.RecommenderID.String(),
		Message:       row.Message,
	}
}

func Create(ctx context.Context, in *npool.CreateRecommendRequest) (*npool.CreateRecommendResponse, error) {
	if err := validateRecommend(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		Recommend.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetRecommenderID(uuid.MustParse(in.GetInfo().GetRecommenderID())).
		SetMessage(in.GetInfo().GetMessage()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail careate recommend: %v", err)
	}

	return &npool.CreateRecommendResponse{
		Info: dbRowToRecommend(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateRecommendRequest) (*npool.UpdateRecommendResponse, error) {
	return nil, nil
}

func GetByApp(ctx context.Context, in *npool.GetRecommendsByAppRequest) (*npool.GetRecommendsByAppResponse, error) {
	return nil, nil
}

func Delete(ctx context.Context, in *npool.DeleteRecommendRequest) (*npool.DeleteRecommendResponse, error) {
	return nil, nil
}
