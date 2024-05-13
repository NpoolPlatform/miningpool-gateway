package fractionrule

import (
	"context"

	fractionrule1 "github.com/NpoolPlatform/miningpool-gateway/pkg/fractionrule"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionrule"
)

func (s *Server) AdminDeleteFractionRule(ctx context.Context, in *npool.AdminDeleteFractionRuleRequest) (*npool.AdminDeleteFractionRuleResponse, error) {
	handler, err := fractionrule1.NewHandler(
		ctx,
		fractionrule1.WithID(&in.ID, true),
		fractionrule1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteFractionRuleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteFractionRule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteFractionRuleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteFractionRuleResponse{
		Info: info,
	}, nil
}
