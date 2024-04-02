package orderuser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/orderuser"
	orderuser1 "github.com/NpoolPlatform/miningpool-gateway/pkg/orderuser"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateOrderUser(ctx context.Context, in *npool.UpdateOrderUserRequest) (*npool.UpdateOrderUserResponse, error) {
	handler, err := orderuser1.NewHandler(
		ctx,
		orderuser1.WithEntID(&in.EntID, true),
		orderuser1.WithAppID(&in.AppID, true),
		orderuser1.WithUserID(&in.UserID, true),
		orderuser1.WithRevenueAddress(in.RevenueAddress, false),
		orderuser1.WithAutoPay(in.AutoPay, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOrderUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateOrderUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOrderUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateOrderUserResponse{Info: info}, nil
}
