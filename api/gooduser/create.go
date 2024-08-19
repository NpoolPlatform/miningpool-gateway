package gooduser

import (
	"context"

	gooduser1 "github.com/NpoolPlatform/miningpool-gateway/pkg/gooduser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/gooduser"
)

func (s *Server) AdminCreateGoodUser(ctx context.Context, in *npool.AdminCreateGoodUserRequest) (*npool.AdminCreateGoodUserResponse, error) {
	handler, err := gooduser1.NewHandler(
		ctx,
		gooduser1.WithCoinTypeIDs(in.CoinTypeIDs, true),
		gooduser1.WithRootUserID(&in.RootUserID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateGoodUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateGoodUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateGoodUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateGoodUserResponse{
		Info: info,
	}, nil
}
