package coin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	coingwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/coin"
	coinmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/coin"
)

func (h *Handler) DeleteCoin(ctx context.Context) (*coingwpb.Coin, error) {
	err := h.checkCoin(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	info, err := h.GetCoin(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	err = coinmwcli.DeleteCoin(ctx, *h.ID, *h.EntID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
