package fractionrule

import (
	"context"
	"fmt"

	fractionrulegwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionrule"
	fractionrulemwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"
	fractionrulemwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/fractionrule"
)

func (h *Handler) GetFractionRules(ctx context.Context) ([]*fractionrulegwpb.FractionRule, uint32, error) {
	infos, total, err := fractionrulemwcli.GetFractionRules(ctx, &fractionrulemwpb.Conds{}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	_infos := []*fractionrulegwpb.FractionRule{}
	for _, info := range infos {
		_info := mw2GW(info)
		_infos = append(_infos, _info)
	}
	return _infos, total, nil
}

func (h *Handler) GetFractionRule(ctx context.Context) (*fractionrulegwpb.FractionRule, error) {
	info, err := fractionrulemwcli.GetFractionRule(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid pool")
	}
	return mw2GW(info), nil
}
