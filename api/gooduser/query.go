package gooduser

import (
	"context"

	gooduser1 "github.com/NpoolPlatform/miningpool-gateway/pkg/gooduser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/gooduser"
)

func (s *Server) AdminGetGoodUsers(ctx context.Context, in *npool.AdminGetGoodUsersRequest) (*npool.AdminGetGoodUsersResponse, error) {
	handler, err := gooduser1.NewHandler(
		ctx,
		gooduser1.WithOffset(in.Offset),
		gooduser1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodUsers",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetGoodUsersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetGoodUsers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodUsers",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetGoodUsersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetGoodUsersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
