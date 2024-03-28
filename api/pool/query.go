package pool

import (
	"context"

	pool1 "github.com/NpoolPlatform/miningpool-gateway/pkg/pool"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/pool"
)

func (s *Server) GetPool(ctx context.Context, in *npool.GetPoolRequest) (*npool.GetPoolResponse, error) {
	handler, err := pool1.NewHandler(
		ctx,
		pool1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPool",
			"In", in,
			"Error", err,
		)
		return &npool.GetPoolResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetPool(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPool",
			"In", in,
			"Error", err,
		)
		return &npool.GetPoolResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetPoolResponse{
		Info: info,
	}, nil
}

func (s *Server) GetPools(ctx context.Context, in *npool.GetPoolsRequest) (*npool.GetPoolsResponse, error) {
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
		return &npool.GetPoolsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetPools(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPools",
			"In", in,
			"Error", err,
		)
		return &npool.GetPoolsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetPoolsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
