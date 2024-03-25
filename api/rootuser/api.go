package rootuser

import (
	"context"

	"github.com/NpoolPlatform/message/npool/miningpool/gw/v1/rootuser"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	rootuser.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	rootuser.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := rootuser.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
