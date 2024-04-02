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
func (s *Server) DeletePool(ctx context.Context, in *npool.DeletePoolRequest) (*npool.DeletePoolResponse, error) {
	handler, err := pool1.NewHandler(
		ctx,
		pool1.WithID(&in.ID, true),
		pool1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeletePool",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePoolResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeletePool(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeletePool",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePoolResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeletePoolResponse{
		Info: info,
	}, nil
}
