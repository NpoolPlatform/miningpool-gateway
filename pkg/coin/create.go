package coin

import (
	"context"

	coingwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/coin"
	coinmwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	coinmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/coin"

	"github.com/google/uuid"
)

func (h *Handler) CreateCoin(ctx context.Context) (*coingwpb.Coin, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	err := coinmwcli.CreateCoin(ctx, &coinmwpb.CoinReq{
		EntID:                  h.EntID,
		PoolID:                 h.PoolID,
		CoinTypeID:             h.CoinTypeID,
		CoinType:               h.CoinType,
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
