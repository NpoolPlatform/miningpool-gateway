package fractionwithdrawalrule

import (
	"context"

	"github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionwithdrawalrule"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	fractionwithdrawalrule.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	fractionwithdrawalrule.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := fractionwithdrawalrule.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
