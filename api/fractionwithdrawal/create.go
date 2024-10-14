package fractionwithdrawal

import (
	"context"

	fractionwithdrawal1 "github.com/NpoolPlatform/miningpool-gateway/pkg/fractionwithdrawal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionwithdrawal"
)

//nolint:dupl
func (s *Server) CreateFractionWithdrawal(ctx context.Context, in *npool.CreateFractionWithdrawalRequest) (*npool.CreateFractionWithdrawalResponse, error) {
	handler, err := fractionwithdrawal1.NewHandler(
		ctx,
		fractionwithdrawal1.WithAppID(&in.AppID, true),
		fractionwithdrawal1.WithUserID(&in.UserID, true),
		fractionwithdrawal1.WithOrderUserID(&in.OrderUserID, true),
		fractionwithdrawal1.WithCoinTypeID(&in.CoinTypeID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFractionWithdrawal",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionWithdrawalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateFractionWithdrawal(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFractionWithdrawal",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionWithdrawalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateFractionWithdrawalResponse{Info: info}, nil
}
