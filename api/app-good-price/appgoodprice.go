package appgoodprice

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	npool.UnimplementedCloudHashingGoodsServer
}

func (s *Server) SetAppGoodPrice(ctx context.Context, in *npool.SetAppGoodPriceRequest) (*npool.SetAppGoodPriceResponse, error) {
	return nil, nil
}
