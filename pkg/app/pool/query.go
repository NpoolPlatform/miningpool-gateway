package pool

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	poolgwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/app/pool"
	apppoolmwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/app/pool"
	"github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	"github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"
	apppoolmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/app/pool"
	coinmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/coin"
	fractionrulemwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/fractionrule"
	poolmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/pool"
)

func (h *Handler) GetPools(ctx context.Context) ([]*poolgwpb.Pool, uint32, error) {
	infos, total, err := apppoolmwcli.GetPools(ctx, &apppoolmwpb.Conds{
		AppID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: *h.AppID,
		},
	}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}

	_infos := []*poolgwpb.Pool{}
	for _, info := range infos {
		_info, err := fullPools(ctx, info.PoolID)
		if err != nil {
			return nil, 0, err
		}
		_infos = append(_infos, _info)
	}

	return _infos, total, nil
}

func fullPools(ctx context.Context, apppoolID string) (*poolgwpb.Pool, error) {
	appinfo, err := apppoolmwcli.GetPool(ctx, apppoolID)
	if err != nil {
		return nil, err
	}
	if appinfo == nil {
		return nil, fmt.Errorf("invalid apppool")
	}

	info, err := poolmwcli.GetPool(ctx, appinfo.PoolID)
	if err != nil {
		return nil, err
	}

	if info == nil {
		return nil, fmt.Errorf("invalid pool")
	}

	coins, _, err := coinmwcli.GetCoins(ctx, &coin.Conds{
		MiningpoolType: &v1.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(info.MiningpoolType),
		},
	}, 0, 0)
	if err != nil {
		return nil, err
	}

	rules, _, err := fractionrulemwcli.GetFractionRules(ctx, &fractionrule.Conds{MiningpoolType: &v1.Uint32Val{
		Op:    cruder.EQ,
		Value: uint32(info.MiningpoolType),
	}}, 0, 0)
	if err != nil {
		return nil, err
	}

	return mw2GW(appinfo, info, coins, rules), nil
}

func (h *Handler) GetPool(ctx context.Context) (*poolgwpb.Pool, error) {
	info, err := apppoolmwcli.GetPool(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid app pool")
	}
	if info.AppID != *h.AppID {
		return nil, fmt.Errorf("permission denied")
	}

	return fullPools(ctx, info.PoolID)
}
