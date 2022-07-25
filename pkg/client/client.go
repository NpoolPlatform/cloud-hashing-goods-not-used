package client

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/cloud-hashing-goods/pkg/message/const"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.CloudHashingGoodsClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get good payment connection: %v", err)
	}
	defer conn.Close()

	cli := npool.NewCloudHashingGoodsClient(conn)

	return fn(_ctx, cli)
}

func GetGood(ctx context.Context, id string) (*npool.GoodInfo, error) {
	// conds: NOT USED NOW, will be used after refactor code
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingGoodsClient) (cruder.Any, error) {
		resp, err := cli.GetGood(ctx, &npool.GetGoodRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get good: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get good: %v", err)
	}
	return info.(*npool.GoodInfo), nil
}

func GetGoods(ctx context.Context) ([]*npool.GoodInfo, error) {
	// conds: NOT USED NOW, will be used after refactor code
	infos, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingGoodsClient) (cruder.Any, error) {
		resp, err := cli.GetGoods(ctx, &npool.GetGoodsRequest{})
		if err != nil {
			return nil, fmt.Errorf("fail get goods: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get goods: %v", err)
	}
	return infos.([]*npool.GoodInfo), nil
}
