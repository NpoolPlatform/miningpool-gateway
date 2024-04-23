package rootuser

import (
	"context"

	rootuser1 "github.com/NpoolPlatform/miningpool-gateway/pkg/rootuser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/rootuser"
)

func (s *Server) AdminCreateRootUser(ctx context.Context, in *npool.AdminCreateRootUserRequest) (*npool.AdminCreateRootUserResponse, error) {
	handler, err := rootuser1.NewHandler(
		ctx,
		rootuser1.WithName(&in.Name, true),
		rootuser1.WithMiningpoolType(&in.MiningpoolType, true),
		rootuser1.WithEmail(&in.Email, true),
		rootuser1.WithAuthToken(&in.AuthToken, true),
		rootuser1.WithRemark(in.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateRootUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateRootUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateRootUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateRootUserResponse{
		Info: info,
	}, nil
}
