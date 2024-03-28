package api

import (
	"context"

	v1 "github.com/NpoolPlatform/message/npool/miningpool/gw/v1"
	"github.com/NpoolPlatform/miningpool-gateway/api/pool"
	"github.com/NpoolPlatform/miningpool-gateway/api/rootuser"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	v1.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	v1.RegisterGatewayServer(server, &Server{})
	rootuser.Register(server)
	pool.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := v1.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := rootuser.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := pool.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
