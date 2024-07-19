//nolint:dupl
package fractionrule

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	fractionrulegwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionrule"
	fractionrulemwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"
	fractionrulemwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/fractionrule"

	"github.com/google/uuid"
)

func (h *Handler) CreateFractionRule(ctx context.Context) (*fractionrulegwpb.FractionRule, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	err := fractionrulemwcli.CreateFractionRule(ctx, &fractionrulemwpb.FractionRuleReq{
		EntID:            h.EntID,
		PoolCoinTypeID:   h.PoolCoinTypeID,
		WithdrawInterval: h.WithdrawInterval,
		MinAmount:        h.MinAmount,
		WithdrawRate:     h.WithdrawRate,
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetFractionRule(ctx)
}
