package gooduser

import (
	"context"

	"github.com/NpoolPlatform/message/npool/miningpool/gw/v1/gooduser"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	gooduser.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	gooduser.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := gooduser.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
