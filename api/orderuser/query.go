package orderuser

import (
	"context"

	orderuser1 "github.com/NpoolPlatform/miningpool-gateway/pkg/orderuser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/orderuser"
)

func (s *Server) GetOrderUser(ctx context.Context, in *npool.GetOrderUserRequest) (*npool.GetOrderUserResponse, error) {
	handler, err := orderuser1.NewHandler(
		ctx,
		orderuser1.WithEntID(&in.EntID, true),
		orderuser1.WithAppID(&in.AppID, true),
		orderuser1.WithUserID(&in.UserID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetOrderUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetOrderUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetOrderUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetOrderUserResponse{
		Info: info,
	}, nil
}

func (s *Server) AdminGetOrderUserProportion(ctx context.Context, in *npool.AdminGetOrderUserProportionRequest) (*npool.AdminGetOrderUserProportionResponse, error) {
	handler, err := orderuser1.NewHandler(
		ctx,
		orderuser1.WithEntID(&in.EntID, true),
		orderuser1.WithAppID(&in.TargetAppID, true),
		orderuser1.WithUserID(&in.TargetUserID, true),
		orderuser1.WithCoinTypeID(&in.CoinTypeID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetOrderUserProportion",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetOrderUserProportionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	proportion, err := handler.GetOrderUserProportion(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetOrderUserProportion",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetOrderUserProportionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetOrderUserProportionResponse{
		Proportion: proportion,
	}, nil
}
