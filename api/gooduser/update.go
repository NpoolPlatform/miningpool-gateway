package gooduser

import (
	"context"

	gooduser1 "github.com/NpoolPlatform/miningpool-gateway/pkg/gooduser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/gooduser"
)

func (s *Server) AdminUpdateGoodUser(ctx context.Context, in *npool.AdminUpdateGoodUserRequest) (*npool.AdminUpdateGoodUserResponse, error) {
	handler, err := gooduser1.NewHandler(
		ctx,
		gooduser1.WithID(&in.ID, true),
		gooduser1.WithEntID(&in.EntID, true),
		gooduser1.WithHashRate(in.HashRate, false),
		gooduser1.WithRevenueType(in.RevenueType, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateGoodUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateGoodUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateGoodUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateGoodUserResponse{
		Info: info,
	}, nil
}
