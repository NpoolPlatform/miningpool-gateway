package fractionwithdrawalrule

import (
	"context"

	fractionwithdrawalrule1 "github.com/NpoolPlatform/miningpool-gateway/pkg/fractionwithdrawalrule"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionwithdrawalrule"
)

func (s *Server) AdminUpdateFractionWithdrawalRule(ctx context.Context, in *npool.AdminUpdateFractionWithdrawalRuleRequest) (*npool.AdminUpdateFractionWithdrawalRuleResponse, error) {
	handler, err := fractionwithdrawalrule1.NewHandler(
		ctx,
		fractionwithdrawalrule1.WithID(&in.ID, true),
		fractionwithdrawalrule1.WithEntID(&in.EntID, true),
		fractionwithdrawalrule1.WithWithdrawInterval(in.WithdrawInterval, false),
		fractionwithdrawalrule1.WithPayoutThreshold(in.PayoutThreshold, false),
		fractionwithdrawalrule1.WithLeastWithdrawalAmount(in.LeastWithdrawalAmount, false),
		fractionwithdrawalrule1.WithWithdrawFee(in.WithdrawFee, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFractionWithdrawalRule",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateFractionWithdrawalRuleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateFractionWithdrawalRule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFractionWithdrawalRule",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateFractionWithdrawalRuleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateFractionWithdrawalRuleResponse{
		Info: info,
	}, nil
}
