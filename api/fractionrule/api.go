package fractionrule

import (
	"context"

	"github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionrule"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	fractionrule.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	fractionrule.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := fractionrule.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
