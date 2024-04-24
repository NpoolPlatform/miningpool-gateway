package pool

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	apppoolmwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/app/pool"
	apppoolmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/app/pool"
)

func (h *Handler) checkPool(ctx context.Context) error {
	exist, err := apppoolmwcli.ExistPoolConds(ctx, &apppoolmwpb.Conds{
		ID: &v1.Uint32Val{
			Op:    cruder.EQ,
			Value: *h.ID,
		},
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: *h.EntID,
		},
		AppID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: *h.AppID,
		},
	})
	if err != nil {
		return err
	}

	if !exist {
		return fmt.Errorf("invalid apppool")
	}
	return nil
}
