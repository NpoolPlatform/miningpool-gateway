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
func (s *Server) GetFraction(ctx context.Context, in *npool.GetFractionRequest) (*npool.GetFractionResponse, error) {
	handler, err := fraction1.NewHandler(
		ctx,
		fraction1.WithEntID(&in.EntID, true),
		fraction1.WithAppID(&in.AppID, true),
		fraction1.WithUserID(&in.UserID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFraction",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetFraction(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFraction",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetFractionResponse{
		Info: info,
	}, nil
}

func (s *Server) GetUserFractions(ctx context.Context, in *npool.GetUserFractionsRequest) (*npool.GetUserFractionsResponse, error) {
	handler, err := fraction1.NewHandler(
		ctx,
		fraction1.WithAppID(&in.AppID, true),
		fraction1.WithUserID(&in.UserID, true),
		fraction1.WithOffset(in.Offset),
		fraction1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserFractions",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserFractionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetUserFractions(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserFractions",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserFractionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetUserFractionsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
