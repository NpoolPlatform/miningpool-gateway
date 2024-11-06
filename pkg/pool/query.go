package pool

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	poolgwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/pool"
	"github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	"github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionwithdrawalrule"
	poolmwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"
	coinmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/coin"
	fractionwithdrawalrulemwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/fractionwithdrawalrule"
	poolmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/pool"
)

func (h *Handler) GetPools(ctx context.Context) ([]*poolgwpb.Pool, uint32, error) {
	infos, total, err := poolmwcli.GetPools(ctx, &poolmwpb.Conds{}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	_infos := []*poolgwpb.Pool{}
	for _, info := range infos {
		_info, err := h.fullPools(ctx, info)
		if err != nil {
			return nil, 0, wlog.WrapError(err)
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
		return nil, wlog.WrapError(err)
	}

	rules := []*fractionwithdrawalrule.FractionWithdrawalRule{}
	for _, info := range coins {
		_rules, _, err := fractionwithdrawalrulemwcli.GetFractionWithdrawalRules(ctx, &fractionwithdrawalrule.Conds{PoolCoinTypeID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: info.EntID,
		}}, 0, 0)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		rules = append(rules, _rules...)
	}

	return mw2GW(info, coins, rules), nil
}

func (h *Handler) GetPool(ctx context.Context) (*poolgwpb.Pool, error) {
	info, err := poolmwcli.GetPool(ctx, *h.EntID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid pool")
	}
	return h.fullPools(ctx, info)
}
