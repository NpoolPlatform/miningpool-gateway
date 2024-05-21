package coin

import (
	"context"

	coingwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/coin"
	coinmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/coin"
)

func (h *Handler) DeleteCoin(ctx context.Context) (*coingwpb.Coin, error) {
	err := h.checkCoin(ctx)
	if err != nil {
		return nil, err
	}

	info, err := h.GetCoin(ctx)
	if err != nil {
		return nil, err
	}

	err = coinmwcli.DeleteCoin(ctx, *h.ID, *h.EntID)
	if err != nil {
		return nil, err
	}
	return info, nil
}
