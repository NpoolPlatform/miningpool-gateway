package fractionwithdrawalrule

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	fractionwithdrawalrulegwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionwithdrawalrule"
	fractionwithdrawalrulemwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/fractionwithdrawalrule"
)

func (h *Handler) DeleteFractionWithdrawalRule(ctx context.Context) (*fractionwithdrawalrulegwpb.FractionWithdrawalRule, error) {
	err := h.checkFractionWithdrawalRule(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	info, err := h.GetFractionWithdrawalRule(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	err = fractionwithdrawalrulemwcli.DeleteFractionWithdrawalRule(ctx, *h.ID, *h.EntID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
