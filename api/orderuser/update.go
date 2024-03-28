package orderuser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/orderuser"
	orderuser1 "github.com/NpoolPlatform/miningpool-gateway/pkg/orderuser"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//nolint:dupl
func (s *Server) SetupRevenueAddress(ctx context.Context, in *npool.SetupRevenueAddressRequest) (*npool.SetupRevenueAddressResponse, error) {
	handler, err := orderuser1.NewHandler(
		ctx,
		orderuser1.WithEntID(&in.EntID, true),
		orderuser1.WithAppID(&in.AppID, true),
		orderuser1.WithUserID(&in.UserID, true),
		orderuser1.WithRevenueAddress(&in.RevenueAddress, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"SetupRevenueAddress",
			"In", in,
			"Error", err,
		)
		return &npool.SetupRevenueAddressResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.SetupRevenueAddress(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"SetupRevenueAddress",
			"In", in,
			"Error", err,
		)
		return &npool.SetupRevenueAddressResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.SetupRevenueAddressResponse{Info: info}, nil
}

//nolint:dupl
func (s *Server) SetupAutoPay(ctx context.Context, in *npool.SetupAutoPayRequest) (*npool.SetupAutoPayResponse, error) {
	handler, err := orderuser1.NewHandler(
		ctx,
		orderuser1.WithEntID(&in.EntID, true),
		orderuser1.WithAppID(&in.AppID, true),
		orderuser1.WithUserID(&in.UserID, true),
		orderuser1.WithAutoPay(&in.AutoPay, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"SetupAutoPay",
			"In", in,
			"Error", err,
		)
		return &npool.SetupAutoPayResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.SetupRevenueAddress(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"SetupAutoPay",
			"In", in,
			"Error", err,
		)
		return &npool.SetupAutoPayResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.SetupAutoPayResponse{
		Info: info,
	}, nil
}
