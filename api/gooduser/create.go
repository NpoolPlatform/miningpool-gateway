package gooduser

import (
	"context"

	gooduser1 "github.com/NpoolPlatform/miningpool-gateway/pkg/gooduser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/gooduser"
)

func (s *Server) CreateGoodUser(ctx context.Context, in *npool.CreateGoodUserRequest) (*npool.CreateGoodUserResponse, error) {
	handler, err := gooduser1.NewHandler(
		ctx,
		gooduser1.WithRootUserID(&in.RootUserID, true),
		gooduser1.WithCoinType(&in.CoinType, true),
		gooduser1.WithHashRate(&in.HashRate, true),
		gooduser1.WithRevenueType(&in.RevenueType, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateGoodUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateGoodUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateGoodUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateGoodUserResponse{
		Info: info,
	}, nil
}
