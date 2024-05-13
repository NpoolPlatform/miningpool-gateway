package coin

import (
	"context"

	coin1 "github.com/NpoolPlatform/miningpool-gateway/pkg/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/coin"
)

func (s *Server) AdminGetCoins(ctx context.Context, in *npool.AdminGetCoinsRequest) (*npool.AdminGetCoinsResponse, error) {
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithOffset(in.Offset),
		coin1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoins",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetCoinsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetCoins(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoins",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetCoinsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetCoinsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
