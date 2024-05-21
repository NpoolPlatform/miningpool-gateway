package fractionrule

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"

	fractionrulemwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"
	fractionrulemwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/fractionrule"
)

func (h *Handler) checkFractionRule(ctx context.Context) error {
	exist, err := fractionrulemwcli.ExistFractionRuleConds(ctx, &fractionrulemwpb.Conds{
		ID: &v1.Uint32Val{
			Op:    cruder.EQ,
			Value: *h.ID,
		},
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: *h.EntID,
		},
	})
	if err != nil {
		return err
	}

	if !exist {
		return fmt.Errorf("invalid fractionrule")
	}
	return nil
}
