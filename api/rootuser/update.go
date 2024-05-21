package rootuser

import (
	"context"

	rootuser1 "github.com/NpoolPlatform/miningpool-gateway/pkg/rootuser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/rootuser"
)

func (s *Server) AdminUpdateRootUser(ctx context.Context, in *npool.AdminUpdateRootUserRequest) (*npool.AdminUpdateRootUserResponse, error) {
	handler, err := rootuser1.NewHandler(
		ctx,
		rootuser1.WithID(&in.ID, true),
		rootuser1.WithEntID(&in.EntID, true),
		rootuser1.WithName(in.Name, false),
		rootuser1.WithEmail(in.Email, false),
		rootuser1.WithAuthToken(in.AuthToken, false),
		rootuser1.WithRemark(in.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateRootUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateRootUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateRootUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateRootUserResponse{
		Info: info,
	}, nil
}
