package fractionwithdrawal

import (
	"context"

	"github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionwithdrawal"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	fractionwithdrawal.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	fractionwithdrawal.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := fractionwithdrawal.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
