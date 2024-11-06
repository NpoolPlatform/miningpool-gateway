package fractionwithdrawalrule

import (
	"context"

	fractionwithdrawalrule1 "github.com/NpoolPlatform/miningpool-gateway/pkg/fractionwithdrawalrule"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionwithdrawalrule"
)

func (s *Server) AdminGetFractionWithdrawalRules(ctx context.Context, in *npool.AdminGetFractionWithdrawalRulesRequest) (*npool.AdminGetFractionWithdrawalRulesResponse, error) {
	handler, err := fractionwithdrawalrule1.NewHandler(
		ctx,
		fractionwithdrawalrule1.WithOffset(in.Offset),
		fractionwithdrawalrule1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractionWithdrawalRules",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetFractionWithdrawalRulesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetFractionWithdrawalRules(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractionWithdrawalRules",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetFractionWithdrawalRulesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetFractionWithdrawalRulesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
