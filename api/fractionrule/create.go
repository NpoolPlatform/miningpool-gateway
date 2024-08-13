package fractionrule

import (
	"context"

	fractionrule1 "github.com/NpoolPlatform/miningpool-gateway/pkg/fractionrule"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionrule"
)

func (s *Server) AdminCreateFractionRule(ctx context.Context, in *npool.AdminCreateFractionRuleRequest) (*npool.AdminCreateFractionRuleResponse, error) {
	handler, err := fractionrule1.NewHandler(
		ctx,
		fractionrule1.WithPoolCoinTypeID(&in.PoolCoinTypeID, true),
		fractionrule1.WithWithdrawInterval(&in.WithdrawInterval, true),
		fractionrule1.WithPayoutThreshold(&in.PayoutThreshold, true),
		fractionrule1.WithMinAmount(&in.MinAmount, true),
		fractionrule1.WithWithdrawRate(&in.WithdrawRate, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateFractionRuleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateFractionRule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateFractionRuleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateFractionRuleResponse{
		Info: info,
	}, nil
}
