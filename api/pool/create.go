package pool

import (
	"context"

	pool1 "github.com/NpoolPlatform/miningpool-gateway/pkg/pool"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/pool"
)

func (s *Server) AdminCreatePool(ctx context.Context, in *npool.AdminCreatePoolRequest) (*npool.AdminCreatePoolResponse, error) {
	handler, err := pool1.NewHandler(
		ctx,
		pool1.WithMiningPoolType(&in.MiningPoolType, true),
		pool1.WithName(&in.Name, true),
		pool1.WithSite(&in.Site, true),
		pool1.WithLogo(&in.Logo, true),
		pool1.WithDescription(in.Description, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePool",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreatePoolResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreatePool(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePool",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreatePoolResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreatePoolResponse{
		Info: info,
	}, nil
}
