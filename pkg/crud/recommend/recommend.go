package recommend

import (
	"context"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/recommend"

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
	if err := validateRecommend(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		Recommend.
		UpdateOneID(id).
		SetMessage(in.GetInfo().GetMessage()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update recommend: %v", err)
	}

	return &npool.UpdateRecommendResponse{
		Info: dbRowToRecommend(info),
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetRecommendsByAppRequest) (*npool.GetRecommendsByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		Recommend.
		Query().
		Where(
			recommend.And(
				recommend.AppID(appID),
				recommend.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query recommend by app id: %v", err)
	}

	recommends := []*npool.Recommend{}
	for _, info := range infos {
		recommends = append(recommends, dbRowToRecommend(info))
	}

	return &npool.GetRecommendsByAppResponse{
		Infos: recommends,
	}, nil
}

func GetByRecommender(ctx context.Context, in *npool.GetRecommendsByRecommenderRequest) (*npool.GetRecommendsByRecommenderResponse, error) {
	recommenderID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		return nil, xerrors.Errorf("invalid recommender id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		Recommend.
		Query().
		Where(
			recommend.RecommenderID(recommenderID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query recommend by app id: %v", err)
	}

	recommends := []*npool.Recommend{}
	for _, info := range infos {
		recommends = append(recommends, dbRowToRecommend(info))
	}

	return &npool.GetRecommendsByRecommenderResponse{
		Infos: recommends,
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteRecommendRequest) (*npool.DeleteRecommendResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		Recommend.
		UpdateOneID(id).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update recommend: %v", err)
	}

	return &npool.DeleteRecommendResponse{
		Info: dbRowToRecommend(info),
	}, nil
}
