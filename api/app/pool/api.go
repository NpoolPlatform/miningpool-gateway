package pool

import (
	"context"

	"github.com/NpoolPlatform/message/npool/miningpool/gw/v1/app/pool"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	pool.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	pool.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := pool.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
