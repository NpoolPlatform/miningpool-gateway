package fractionrule

import (
	"context"

	fractionrule1 "github.com/NpoolPlatform/miningpool-gateway/pkg/fractionrule"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionrule"
)

func (s *Server) AdminUpdateFractionRule(ctx context.Context, in *npool.AdminUpdateFractionRuleRequest) (*npool.AdminUpdateFractionRuleResponse, error) {
	handler, err := fractionrule1.NewHandler(
		ctx,
		fractionrule1.WithID(&in.ID, true),
		fractionrule1.WithEntID(&in.EntID, true),
		fractionrule1.WithWithdrawInterval(in.WithdrawInterval, false),
		fractionrule1.WithPayoutThreshold(in.PayoutThreshold, false),
		fractionrule1.WithMinAmount(in.MinAmount, false),
		fractionrule1.WithWithdrawRate(in.WithdrawRate, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateFractionRuleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateFractionRule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateFractionRuleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateFractionRuleResponse{
		Info: info,
	}, nil
}
