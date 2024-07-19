package coin

import (
	"context"

	coingwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/coin"
	coinmwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	coinmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/coin"

	"github.com/google/uuid"
)

func (h *Handler) UpdateCoin(ctx context.Context) (*coingwpb.Coin, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	err := coinmwcli.UpdateCoin(ctx, &coinmwpb.CoinReq{
		ID:                     h.ID,
		EntID:                  h.EntID,
		FeeRatio:               h.FeeRatio,
		FixedRevenueAble:       h.FixedRevenueAble,
		LeastTransferAmount:    h.LeastTransferAmount,
		BenefitIntervalSeconds: h.BenefitIntervalSeconds,
		Remark:                 h.Remark,
	})
	if err != nil {
		return nil, err
	}

	return h.GetCoin(ctx)
}
