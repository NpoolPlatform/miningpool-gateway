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
func (s *Server) GetPool(ctx context.Context, in *npool.GetPoolRequest) (*npool.GetPoolResponse, error) {
	handler, err := pool1.NewHandler(
		ctx,
		pool1.WithEntID(&in.EntID, true),
		pool1.WithAppID(&in.AppID, true),
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

//nolint:dupl
func (s *Server) GetPools(ctx context.Context, in *npool.GetPoolsRequest) (*npool.GetPoolsResponse, error) {
	handler, err := pool1.NewHandler(
		ctx,
		pool1.WithAppID(&in.AppID, true),
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

	return &npool.GetPoolsResponse{Infos: infos, Total: total}, nil
}

//nolint:dupl
func (s *Server) AdminGetPools(ctx context.Context, in *npool.AdminGetNPoolsRequest) (*npool.AdminGetNPoolsResponse, error) {
	handler, err := pool1.NewHandler(
		ctx,
		pool1.WithAppID(&in.TargetAppID, true),
		pool1.WithOffset(in.Offset),
		pool1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPools",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetNPoolsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetPools(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPools",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetNPoolsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetNPoolsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
