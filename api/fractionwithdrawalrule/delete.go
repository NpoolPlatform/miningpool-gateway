package fractionwithdrawalrule

import (
	"context"

	fractionwithdrawalrule1 "github.com/NpoolPlatform/miningpool-gateway/pkg/fractionwithdrawalrule"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionwithdrawalrule"
)

func (s *Server) AdminDeleteFractionWithdrawalRule(ctx context.Context,
	in *npool.AdminDeleteFractionWithdrawalRuleRequest) (*npool.AdminDeleteFractionWithdrawalRuleResponse, error) {
	handler, err := fractionwithdrawalrule1.NewHandler(
		ctx,
		fractionwithdrawalrule1.WithID(&in.ID, true),
		fractionwithdrawalrule1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteFractionWithdrawalRule",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteFractionWithdrawalRuleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteFractionWithdrawalRule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteFractionWithdrawalRule",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteFractionWithdrawalRuleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteFractionWithdrawalRuleResponse{
		Info: info,
	}, nil
}
