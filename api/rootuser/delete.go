package rootuser

import (
	"context"

	rootuser1 "github.com/NpoolPlatform/miningpool-gateway/pkg/rootuser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/rootuser"
)

func (s *Server) DeleteRootUser(ctx context.Context, in *npool.DeleteRootUserRequest) (*npool.DeleteRootUserResponse, error) {
	handler, err := rootuser1.NewHandler(
		ctx,
		rootuser1.WithID(&in.ID, true),
		rootuser1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteRootUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteRootUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteRootUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteRootUserResponse{
		Info: info,
	}, nil
}
