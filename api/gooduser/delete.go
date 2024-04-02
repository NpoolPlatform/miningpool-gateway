package gooduser

import (
	"context"

	gooduser1 "github.com/NpoolPlatform/miningpool-gateway/pkg/gooduser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/gooduser"
)

func (s *Server) DeleteGoodUser(ctx context.Context, in *npool.DeleteGoodUserRequest) (*npool.DeleteGoodUserResponse, error) {
	handler, err := gooduser1.NewHandler(
		ctx,
		gooduser1.WithID(&in.ID, true),
		gooduser1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteGoodUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteGoodUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteGoodUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteGoodUserResponse{
		Info: info,
	}, nil
}
