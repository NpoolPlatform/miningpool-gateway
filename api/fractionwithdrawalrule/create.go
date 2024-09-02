package fractionwithdrawalrule

import (
	"context"

	fractionwithdrawalrule1 "github.com/NpoolPlatform/miningpool-gateway/pkg/fractionwithdrawalrule"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionwithdrawalrule"
)

func (s *Server) AdminCreateFractionWithdrawalRule(ctx context.Context, in *npool.AdminCreateFractionWithdrawalRuleRequest) (*npool.AdminCreateFractionWithdrawalRuleResponse, error) {
	handler, err := fractionwithdrawalrule1.NewHandler(
		ctx,
		fractionwithdrawalrule1.WithPoolCoinTypeID(&in.PoolCoinTypeID, true),
		fractionwithdrawalrule1.WithWithdrawInterval(&in.WithdrawInterval, true),
		fractionwithdrawalrule1.WithPayoutThreshold(&in.PayoutThreshold, true),
		fractionwithdrawalrule1.WithLeastWithdrawalAmount(&in.LeastWithdrawalAmount, true),
		fractionwithdrawalrule1.WithWithdrawFee(&in.WithdrawFee, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFractionWithdrawalRule",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateFractionWithdrawalRuleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateFractionWithdrawalRule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFractionWithdrawalRule",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateFractionWithdrawalRuleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateFractionWithdrawalRuleResponse{
		Info: info,
	}, nil
}
