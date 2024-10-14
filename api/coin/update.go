package coin

import (
	"context"

	coin1 "github.com/NpoolPlatform/miningpool-gateway/pkg/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/coin"
)

func (s *Server) AdminUpdateCoin(ctx context.Context, in *npool.AdminUpdateCoinRequest) (*npool.AdminUpdateCoinResponse, error) {
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithID(&in.ID, true),
		coin1.WithEntID(&in.EntID, true),
		coin1.WithFeeRatio(in.FeeRatio, false),
		coin1.WithFixedRevenueAble(in.FixedRevenueAble, false),
		coin1.WithLeastTransferAmount(in.LeastTransferAmount, false),
		coin1.WithBenefitIntervalSeconds(in.BenefitIntervalSeconds, false),
		coin1.WithRemark(in.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateCoinResponse{
		Info: info,
	}, nil
}
