package pool

import (
	"context"

	pool1 "github.com/NpoolPlatform/miningpool-gateway/pkg/pool"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/pool"
)

func (s *Server) AdminUpdatePool(ctx context.Context, in *npool.AdminUpdatePoolRequest) (*npool.AdminUpdatePoolResponse, error) {
	handler, err := pool1.NewHandler(
		ctx,
		pool1.WithID(&in.ID, true),
		pool1.WithEntID(&in.EntID, true),
		pool1.WithName(in.Name, false),
		pool1.WithSite(in.Site, false),
		pool1.WithLogo(in.Logo, false),
		pool1.WithDescription(in.Description, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePool",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdatePoolResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdatePool(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePool",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdatePoolResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdatePoolResponse{
		Info: info,
	}, nil
}
