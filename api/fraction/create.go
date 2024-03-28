package fraction

import (
	"context"

	fraction1 "github.com/NpoolPlatform/miningpool-gateway/pkg/fraction"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fraction"
)

//nolint:dupl
func (s *Server) CreateFraction(ctx context.Context, in *npool.CreateFractionRequest) (*npool.CreateFractionResponse, error) {
	handler, err := fraction1.NewHandler(
		ctx,
		fraction1.WithAppID(&in.AppID, true),
		fraction1.WithUserID(&in.UserID, true),
		fraction1.WithOrderUserID(&in.OrderUserID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFraction",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateFraction(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFraction",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateFractionResponse{Info: info}, nil
}
