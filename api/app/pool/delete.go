package pool

import (
	"context"

	pool1 "github.com/NpoolPlatform/miningpool-gateway/pkg/app/pool"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/app/pool"
)

func (s *Server) AdminDeletePool(ctx context.Context, in *npool.AdminDeletePoolRequest) (*npool.AdminDeletePoolResponse, error) {
	handler, err := pool1.NewHandler(
		ctx,
		pool1.WithID(&in.ID, true),
		pool1.WithEntID(&in.EntID, true),
		pool1.WithAppID(&in.TargetAppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeletePool",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeletePoolResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeletePool(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeletePool",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeletePoolResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeletePoolResponse{
		Info: info,
	}, nil
}
