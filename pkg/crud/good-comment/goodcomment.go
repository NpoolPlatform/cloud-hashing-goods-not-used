package goodcomment

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodcomment"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func dbRowToGoodComment(info *ent.GoodComment) *npool.GoodComment {
	return &npool.GoodComment{
		ID:        info.ID.String(),
		ReplyToID: info.ReplyToID.String(),
		UserID:    info.UserID.String(),
		AppID:     info.AppID.String(),
		GoodID:    info.GoodID.String(),
		Content:   info.Content,
		OrderID:   info.OrderID.String(),
	}
}

func validateGoodComment(info *npool.GoodComment) error {
	_, err := uuid.Parse(info.GetUserID())
	if err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}

	_, err = uuid.Parse(info.GetAppID())
	if err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}

	_, err = uuid.Parse(info.GetGoodID())
	if err != nil {
		return xerrors.Errorf("invalid good id: %v", err)
	}

	_, err = uuid.Parse(info.GetOrderID())
	if err != nil {
		return xerrors.Errorf("invalid order id: %v", err)
	}

	return nil
}

func validateReplyToID(ctx context.Context, info *npool.GoodComment) (*uuid.UUID, error) {
	replyToID, err := uuid.Parse(info.GetReplyToID())
	if err == nil {
		infos, err := db.Client().
			GoodComment.
			Query().
			Where(
				goodcomment.And(
					goodcomment.ID(replyToID),
				),
			).
			All(ctx)
		if err != nil {
			return nil, xerrors.Errorf("reply to not exist comment: %v", err)
		}
		if len(infos) == 0 {
			return nil, xerrors.Errorf("empty comment to reply")
		}
	}

	if err != nil {
		return nil, nil
	}

	return &replyToID, nil
}

func Create(ctx context.Context, in *npool.CreateGoodCommentRequest) (*npool.CreateGoodCommentResponse, error) {
	if err := validateGoodComment(in.GetComment()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	replyToID, err := validateReplyToID(ctx, in.GetComment())
	if err != nil {
		return nil, xerrors.Errorf("invalid reply to: %v", err)
	}

	info, err := db.Client().
		GoodComment.
		Create().
		SetUserID(uuid.MustParse(in.GetComment().GetUserID())).
		SetAppID(uuid.MustParse(in.GetComment().GetAppID())).
		SetGoodID(uuid.MustParse(in.GetComment().GetGoodID())).
		SetOrderID(uuid.MustParse(in.GetComment().GetOrderID())).
		SetContent(in.GetComment().GetContent()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if replyToID != nil {
		info, err = db.Client().
			GoodComment.
			UpdateOneID(info.ID).
			SetReplyToID(*replyToID).
			Save(ctx)
		if err != nil {
			return nil, xerrors.Errorf("cannot update reply to: %v", err)
		}
	}

	return &npool.CreateGoodCommentResponse{
		Comment: dbRowToGoodComment(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateGoodCommentRequest) (*npool.UpdateGoodCommentResponse, error) {
	id, err := uuid.Parse(in.GetComment().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid comment id: %v", err)
	}

	info, err := db.Client().
		GoodComment.
		UpdateOneID(id).
		SetContent(in.GetComment().GetContent()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to udpate content: %v", err)
	}

	return &npool.UpdateGoodCommentResponse{
		Comment: dbRowToGoodComment(info),
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetGoodCommentsRequest) (*npool.GetGoodCommentsResponse, error) {
	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	infos, err := db.Client().
		GoodComment.
		Query().
		Where(
			goodcomment.And(
				goodcomment.GoodID(goodID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query good comment: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty good comment")
	}

	comments := []*npool.GoodComment{}
	for _, info := range infos {
		comments = append(comments, dbRowToGoodComment(info))
	}

	return &npool.GetGoodCommentsResponse{
		Comments: comments,
	}, nil
}
