package fractionwithdrawalrule

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	fractionwithdrawalrulegwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionwithdrawalrule"
	fractionwithdrawalrulemwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionwithdrawalrule"
	fractionwithdrawalrulemwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/fractionwithdrawalrule"
)

func (h *Handler) GetFractionWithdrawalRules(ctx context.Context) ([]*fractionwithdrawalrulegwpb.FractionWithdrawalRule, uint32, error) {
	infos, total, err := fractionwithdrawalrulemwcli.GetFractionWithdrawalRules(ctx, &fractionwithdrawalrulemwpb.Conds{}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	_infos := []*fractionwithdrawalrulegwpb.FractionWithdrawalRule{}
	for _, info := range infos {
		_info := mw2GW(info)
		_infos = append(_infos, _info)
	}
	return _infos, total, nil
}

func (h *Handler) GetFractionWithdrawalRule(ctx context.Context) (*fractionwithdrawalrulegwpb.FractionWithdrawalRule, error) {
	info, err := fractionwithdrawalrulemwcli.GetFractionWithdrawalRule(ctx, *h.EntID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid fractionwithdrawalrule")
	}
	return mw2GW(info), nil
}
