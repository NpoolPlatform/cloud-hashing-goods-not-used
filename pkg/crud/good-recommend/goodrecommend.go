package goodrecommend

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodrecommend"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func dbRowToGoodRecommend(info *ent.GoodRecommend) *npool.GoodRecommend {
	return &npool.GoodRecommend{
		ID:      info.ID.String(),
		UserID:  info.UserID.String(),
		GoodID:  info.GoodID.String(),
		Content: info.Content,
	}
}

func validateGoodRecommend(info *npool.GoodRecommend) error {
	_, err := uuid.Parse(info.GetUserID())
	if err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}

	_, err = uuid.Parse(info.GetGoodID())
	if err != nil {
		return xerrors.Errorf("invalid good id: %v", err)
	}

	return nil
}

func Create(ctx context.Context, in *npool.CreateGoodRecommendRequest) (*npool.CreateGoodRecommendResponse, error) {
	if err := validateGoodRecommend(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		GoodRecommend.
		Create().
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetContent(in.GetInfo().GetContent()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &npool.CreateGoodRecommendResponse{
		Info: dbRowToGoodRecommend(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateGoodRecommendRequest) (*npool.UpdateGoodRecommendResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid recommend id: %v", err)
	}

	info, err := db.Client().
		GoodRecommend.
		UpdateOneID(id).
		SetContent(in.GetInfo().GetContent()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to udpate content: %v", err)
	}

	return &npool.UpdateGoodRecommendResponse{
		Info: dbRowToGoodRecommend(info),
	}, nil
}

func GetByGood(ctx context.Context, in *npool.GetGoodRecommendsByGoodRequest) (*npool.GetGoodRecommendsByGoodResponse, error) {
	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	infos, err := db.Client().
		GoodRecommend.
		Query().
		Where(goodrecommend.GoodID(goodID)).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query good recommend: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty good recommend")
	}

	recommends := []*npool.GoodRecommend{}
	for _, info := range infos {
		recommends = append(recommends, dbRowToGoodRecommend(info))
	}

	return &npool.GetGoodRecommendsByGoodResponse{
		Infos: recommends,
	}, nil
}

func GetByUser(ctx context.Context, in *npool.GetGoodRecommendsByUserRequest) (*npool.GetGoodRecommendsByUserResponse, error) {
	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	infos, err := db.Client().
		GoodRecommend.
		Query().
		Where(goodrecommend.UserID(userID)).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query good recommend: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty good recommend")
	}

	recommends := []*npool.GoodRecommend{}
	for _, info := range infos {
		recommends = append(recommends, dbRowToGoodRecommend(info))
	}

	return &npool.GetGoodRecommendsByUserResponse{
		Infos: recommends,
	}, nil
}
