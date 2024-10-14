package api

import (
	"context"

	v1 "github.com/NpoolPlatform/message/npool/miningpool/gw/v1"
	apppool "github.com/NpoolPlatform/miningpool-gateway/api/app/pool"
	"github.com/NpoolPlatform/miningpool-gateway/api/coin"
	"github.com/NpoolPlatform/miningpool-gateway/api/fractionwithdrawal"
	"github.com/NpoolPlatform/miningpool-gateway/api/fractionwithdrawalrule"
	"github.com/NpoolPlatform/miningpool-gateway/api/gooduser"
	"github.com/NpoolPlatform/miningpool-gateway/api/orderuser"
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
	gooduser.Register(server)
	fractionwithdrawal.Register(server)
	orderuser.Register(server)
	apppool.Register(server)
	pool.Register(server)
	coin.Register(server)
	fractionwithdrawalrule.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := v1.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := rootuser.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := gooduser.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := fractionwithdrawal.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := orderuser.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := apppool.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := pool.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := coin.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := fractionwithdrawalrule.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
