package rootuser

import (
	"context"

	rootuser1 "github.com/NpoolPlatform/miningpool-gateway/pkg/rootuser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/rootuser"
)

func (s *Server) GetRootUser(ctx context.Context, in *npool.GetRootUserRequest) (*npool.GetRootUserResponse, error) {
	handler, err := rootuser1.NewHandler(
		ctx,
		rootuser1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetRootUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetRootUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetRootUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetRootUserResponse{
		Info: info,
	}, nil
}

func (s *Server) GetRootUsers(ctx context.Context, in *npool.GetRootUsersRequest) (*npool.GetRootUsersResponse, error) {
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
		return &npool.GetRootUsersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetRootUsers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRootUsers",
			"In", in,
			"Error", err,
		)
		return &npool.GetRootUsersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetRootUsersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
