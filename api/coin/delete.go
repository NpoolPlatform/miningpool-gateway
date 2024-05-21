package coin

import (
	"context"

	coin1 "github.com/NpoolPlatform/miningpool-gateway/pkg/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/coin"
)

func (s *Server) AdminDeleteCoin(ctx context.Context, in *npool.AdminDeleteCoinRequest) (*npool.AdminDeleteCoinResponse, error) {
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithID(&in.ID, true),
		coin1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCoin",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCoin",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteCoinResponse{
		Info: info,
	}, nil
}
