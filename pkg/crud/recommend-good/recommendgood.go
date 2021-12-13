package recommendgood

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/recommendgood"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func dbRowToRecommendGood(info *ent.RecommendGood) *npool.RecommendGood {
	return &npool.RecommendGood{
		ID:            info.ID.String(),
		RecommenderID: info.RecommenderID.String(),
		GoodID:        info.GoodID.String(),
		Message:       info.Message,
	}
}

func validateRecommendGood(info *npool.RecommendGood) error {
	_, err := uuid.Parse(info.GetRecommenderID())
	if err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}

	_, err = uuid.Parse(info.GetGoodID())
	if err != nil {
		return xerrors.Errorf("invalid good id: %v", err)
	}

	return nil
}

func Create(ctx context.Context, in *npool.CreateRecommendGoodRequest) (*npool.CreateRecommendGoodResponse, error) {
	if err := validateRecommendGood(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		RecommendGood.
		Create().
		SetRecommenderID(uuid.MustParse(in.GetInfo().GetRecommenderID())).
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetMessage(in.GetInfo().GetMessage()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &npool.CreateRecommendGoodResponse{
		Info: dbRowToRecommendGood(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateRecommendGoodRequest) (*npool.UpdateRecommendGoodResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid recommend id: %v", err)
	}

	info, err := db.Client().
		RecommendGood.
		UpdateOneID(id).
		SetMessage(in.GetInfo().GetMessage()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to udpate Message: %v", err)
	}

	return &npool.UpdateRecommendGoodResponse{
		Info: dbRowToRecommendGood(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetRecommendGoodsRequest) (*npool.GetRecommendGoodsResponse, error) {
	offsetN, limitN := in.PageInfo.PageIndex, in.PageInfo.PageSize
	if limitN == 0 {
		limitN = 20
	}
	offsetN *= limitN
	infos, err := db.Client().
		RecommendGood.
		Query().
		Order(ent.Desc(recommendgood.FieldCreateAt)).
		Offset(int(offsetN)).
		Limit(int(limitN)).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query recommend good: %v", err)
	}
	countN, err := db.Client().
		RecommendGood.
		Query().Count(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query recommend length: %v", err)
	}

	recommends := []*npool.RecommendGood{}
	for _, info := range infos {
		recommends = append(recommends, dbRowToRecommendGood(info))
	}

	return &npool.GetRecommendGoodsResponse{
		Recommends: recommends,
		Total:      int32(countN),
	}, nil
}

func GetByRecommender(ctx context.Context, in *npool.GetRecommendGoodsByRecommenderRequest) (*npool.GetRecommendGoodsByRecommenderResponse, error) {
	RecommenderID, err := uuid.Parse(in.GetRecommenderID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	infos, err := db.Client().
		RecommendGood.
		Query().
		Where(recommendgood.RecommenderID(RecommenderID)).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query recommend good: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty recommend good")
	}

	recommends := []*npool.RecommendGood{}
	for _, info := range infos {
		recommends = append(recommends, dbRowToRecommendGood(info))
	}

	return &npool.GetRecommendGoodsByRecommenderResponse{
		Recommends: recommends,
	}, nil
}
