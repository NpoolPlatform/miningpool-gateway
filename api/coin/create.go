package coin

import (
	"context"

	coin1 "github.com/NpoolPlatform/miningpool-gateway/pkg/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/coin"
)

func (s *Server) AdminCreateCoin(ctx context.Context, in *npool.AdminCreateCoinRequest) (*npool.AdminCreateCoinResponse, error) {
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithPoolID(&in.PoolID, true),
		coin1.WithCoinTypeID(&in.CoinTypeID, true),
		coin1.WithCoinType(&in.CoinType, true),
		coin1.WithRevenueType(&in.RevenueType, true),
		coin1.WithFeeRatio(&in.FeeRatio, true),
		coin1.WithFixedRevenueAble(&in.FixedRevenueAble, true),
		coin1.WithLeastTransferAmount(&in.LeastTransferAmount, true),
		coin1.WithBenefitIntervalSeconds(&in.BenefitIntervalSeconds, true),
		coin1.WithRemark(in.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateCoinResponse{
		Info: info,
	}, nil
}
