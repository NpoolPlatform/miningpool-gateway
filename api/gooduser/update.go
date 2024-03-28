package gooduser

import (
	"context"

	gooduser1 "github.com/NpoolPlatform/miningpool-gateway/pkg/gooduser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/gooduser"
)

func (s *Server) UpdateGoodUser(ctx context.Context, in *npool.UpdateGoodUserRequest) (*npool.UpdateGoodUserResponse, error) {
	handler, err := gooduser1.NewHandler(
		ctx,
		gooduser1.WithID(&in.ID, true),
		gooduser1.WithRootUserID(in.RootUserID, true),
		gooduser1.WithHashRate(in.HashRate, false),
		gooduser1.WithRevenueType(in.RevenueType, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateGoodUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateGoodUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateGoodUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateGoodUserResponse{
		Info: info,
	}, nil
}
