package coin

import (
	"context"
	"fmt"

	coingwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/coin"
	coinmwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	coinmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/coin"
)

func (h *Handler) GetCoins(ctx context.Context) ([]*coingwpb.Coin, uint32, error) {
	infos, total, err := coinmwcli.GetCoins(ctx, &coinmwpb.Conds{}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	_infos := []*coingwpb.Coin{}
	for _, info := range infos {
		_info := mw2GW(info)
		_infos = append(_infos, _info)
	}
	return _infos, total, nil
}

func (h *Handler) GetCoin(ctx context.Context) (*coingwpb.Coin, error) {
	info, err := coinmwcli.GetCoin(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid coin")
	}
	return mw2GW(info), nil
}
