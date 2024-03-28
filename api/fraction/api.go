package fraction

import (
	"context"

	"github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fraction"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	fraction.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	fraction.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := fraction.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
