package pool

import (
	"context"

	pool1 "github.com/NpoolPlatform/miningpool-gateway/pkg/app/pool"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/app/pool"
)

//nolint:dupl
func (s *Server) AdminCreatePool(ctx context.Context, in *npool.AdminCreatePoolRequest) (*npool.AdminCreatePoolResponse, error) {
	handler, err := pool1.NewHandler(
		ctx,
		pool1.WithPoolID(&in.PoolID, true),
		pool1.WithTargetAppID(&in.TargetAppID, true),
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
