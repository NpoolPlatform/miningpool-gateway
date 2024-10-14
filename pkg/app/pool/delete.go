package pool

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	poolgwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/app/pool"
	apppoolmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/app/pool"
)

func (h *Handler) DeletePool(ctx context.Context) (*poolgwpb.Pool, error) {
	err := h.checkPool(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	info, err := fullPools(ctx, *h.EntID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	err = apppoolmwcli.DeletePool(ctx, *h.ID, *h.EntID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
