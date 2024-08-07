package pool

import (
	"context"

	pool1 "github.com/NpoolPlatform/miningpool-gateway/pkg/pool"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/pool"
)

func (s *Server) AdminGetPools(ctx context.Context, in *npool.AdminGetPoolsRequest) (*npool.AdminGetPoolsResponse, error) {
	handler, err := pool1.NewHandler(
		ctx,
		pool1.WithOffset(in.Offset),
		pool1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPools",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetPoolsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetPools(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPools",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetPoolsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetPoolsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
