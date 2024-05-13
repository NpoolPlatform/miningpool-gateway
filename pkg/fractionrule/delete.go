package fractionrule

import (
	"context"

	fractionrulegwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionrule"
	fractionrulemwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/fractionrule"
)

func (h *Handler) DeleteFractionRule(ctx context.Context) (*fractionrulegwpb.FractionRule, error) {
	err := h.checkFractionRule(ctx)
	if err != nil {
		return nil, err
	}

	info, err := h.GetFractionRule(ctx)
	if err != nil {
		return nil, err
	}

	err = fractionrulemwcli.DeleteFractionRule(ctx, *h.ID, *h.EntID)
	if err != nil {
		return nil, err
	}
	return info, nil
}
