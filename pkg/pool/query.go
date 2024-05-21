package pool

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	poolgwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/pool"
	"github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	"github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"
	poolmwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"
	coinmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/coin"
	fractionrulemwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/fractionrule"
	poolmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/pool"
)

func (h *Handler) GetPools(ctx context.Context) ([]*poolgwpb.Pool, uint32, error) {
	infos, total, err := poolmwcli.GetPools(ctx, &poolmwpb.Conds{}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	_infos := []*poolgwpb.Pool{}
	for _, info := range infos {
		_info, err := h.fullPools(ctx, info)
		if err != nil {
			return nil, 0, err
		}
		_infos = append(_infos, _info)
	}

	return _infos, total, nil
}

func (h *Handler) fullPools(ctx context.Context, info *poolmwpb.Pool) (*poolgwpb.Pool, error) {
	coins, _, err := coinmwcli.GetCoins(ctx, &coin.Conds{
		PoolID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: info.EntID,
		},
	}, 0, 0)
	if err != nil {
		return nil, err
	}

	rules := []*fractionrule.FractionRule{}
	for _, info := range coins {
		_rules, _, err := fractionrulemwcli.GetFractionRules(ctx, &fractionrule.Conds{PoolCoinTypeID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: info.EntID,
		}}, 0, 0)
		if err != nil {
			return nil, err
		}
		rules = append(rules, _rules...)
	}

	return mw2GW(info, coins, rules), nil
}

func (h *Handler) GetPool(ctx context.Context) (*poolgwpb.Pool, error) {
	info, err := poolmwcli.GetPool(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid pool")
	}
	return h.fullPools(ctx, info)
}
