package fractionwithdrawal

import (
	"context"

	fractionwithdrawal1 "github.com/NpoolPlatform/miningpool-gateway/pkg/fractionwithdrawal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionwithdrawal"
)

func (s *Server) GetFractionWithdrawal(ctx context.Context, in *npool.GetFractionWithdrawalRequest) (*npool.GetFractionWithdrawalResponse, error) {
	handler, err := fractionwithdrawal1.NewHandler(
		ctx,
		fractionwithdrawal1.WithEntID(&in.EntID, true),
		fractionwithdrawal1.WithAppID(&in.AppID, true),
		fractionwithdrawal1.WithUserID(&in.UserID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractionWithdrawal",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionWithdrawalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetFractionWithdrawal(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractionWithdrawal",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionWithdrawalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetFractionWithdrawalResponse{
		Info: info,
	}, nil
}

func (s *Server) GetUserFractionWithdrawals(ctx context.Context, in *npool.GetUserFractionWithdrawalsRequest) (*npool.GetUserFractionWithdrawalsResponse, error) {
	handler, err := fractionwithdrawal1.NewHandler(
		ctx,
		fractionwithdrawal1.WithAppID(&in.AppID, true),
		fractionwithdrawal1.WithUserID(&in.UserID, true),
		fractionwithdrawal1.WithOffset(in.Offset),
		fractionwithdrawal1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserFractionWithdrawals",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserFractionWithdrawalsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetUserFractionWithdrawals(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserFractionWithdrawals",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserFractionWithdrawalsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetUserFractionWithdrawalsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
