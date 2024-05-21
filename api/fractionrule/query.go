package fractionrule

import (
	"context"

	fractionrule1 "github.com/NpoolPlatform/miningpool-gateway/pkg/fractionrule"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionrule"
)

func (s *Server) AdminGetFractionRules(ctx context.Context, in *npool.AdminGetFractionRulesRequest) (*npool.AdminGetFractionRulesResponse, error) {
	handler, err := fractionrule1.NewHandler(
		ctx,
		fractionrule1.WithOffset(in.Offset),
		fractionrule1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractionRules",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetFractionRulesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetFractionRules(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractionRules",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetFractionRulesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetFractionRulesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
