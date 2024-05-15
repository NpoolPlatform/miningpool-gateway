package rootuser

import (
	"context"

	rootuser1 "github.com/NpoolPlatform/miningpool-gateway/pkg/rootuser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/rootuser"
)

func (s *Server) AdminGetRootUsers(ctx context.Context, in *npool.AdminGetRootUsersRequest) (*npool.AdminGetRootUsersResponse, error) {
	handler, err := rootuser1.NewHandler(
		ctx,
		rootuser1.WithOffset(in.Offset),
		rootuser1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRootUsers",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetRootUsersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetRootUsers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRootUsers",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetRootUsersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetRootUsersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
