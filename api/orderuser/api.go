package orderuser

import (
	"context"

	"github.com/NpoolPlatform/message/npool/miningpool/gw/v1/orderuser"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	orderuser.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	orderuser.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := orderuser.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
