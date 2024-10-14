package orderuser

import (
	"context"

	orderuser1 "github.com/NpoolPlatform/miningpool-gateway/pkg/orderuser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/orderuser"
)

func (s *Server) AdminUpdateOrderUser(ctx context.Context, in *npool.AdminUpdateOrderUserRequest) (*npool.AdminUpdateOrderUserResponse, error) {
	handler, err := orderuser1.NewHandler(
		ctx,
		orderuser1.WithEntID(&in.EntID, true),
		orderuser1.WithAppID(&in.TargetAppID, true),
		orderuser1.WithUserID(&in.TargetUserID, true),
		orderuser1.WithCoinTypeID(&in.CoinTypeID, true),
		orderuser1.WithProportion(&in.Proportion, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateOrderUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	err = handler.UpdateOrderUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateOrderUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateOrderUserResponse{}, nil
}
