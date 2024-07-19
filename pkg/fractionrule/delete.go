package fractionrule

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	fractionrulegwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionrule"
	fractionrulemwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/fractionrule"
)

func (h *Handler) DeleteFractionRule(ctx context.Context) (*fractionrulegwpb.FractionRule, error) {
	err := h.checkFractionRule(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	info, err := h.GetFractionRule(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	err = fractionrulemwcli.DeleteFractionRule(ctx, *h.ID, *h.EntID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
