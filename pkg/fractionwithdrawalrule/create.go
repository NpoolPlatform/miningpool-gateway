//nolint:dupl
package fractionwithdrawalrule

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	fractionwithdrawalrulegwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionwithdrawalrule"
	fractionwithdrawalrulemwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionwithdrawalrule"
	fractionwithdrawalrulemwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/fractionwithdrawalrule"

	"github.com/google/uuid"
)

func (h *Handler) CreateFractionWithdrawalRule(ctx context.Context) (*fractionwithdrawalrulegwpb.FractionWithdrawalRule, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	err := fractionwithdrawalrulemwcli.CreateFractionWithdrawalRule(ctx, &fractionwithdrawalrulemwpb.FractionWithdrawalRuleReq{
		EntID:                 h.EntID,
		PoolCoinTypeID:        h.PoolCoinTypeID,
		WithdrawInterval:      h.WithdrawInterval,
		PayoutThreshold:       h.PayoutThreshold,
		LeastWithdrawalAmount: h.LeastWithdrawalAmount,
		WithdrawFee:           h.WithdrawFee,
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetFractionWithdrawalRule(ctx)
}
