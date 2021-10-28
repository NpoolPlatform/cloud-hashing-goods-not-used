package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-goods/api/echo"
	"github.com/NpoolPlatform/cloud-hashing-goods/message/npool"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterCloudHashingGoodsServer(server, &echo.Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterCloudHashingGoodsHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
