package pool

import (
	"context"

	poolgwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/app/pool"
	apppoolmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/app/pool"
)

func (h *Handler) DeletePool(ctx context.Context) (*poolgwpb.Pool, error) {
	err := h.checkPool(ctx)
	if err != nil {
		return nil, err
	}

	info, err := fullPools(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}

	err = apppoolmwcli.DeletePool(ctx, *h.ID, *h.EntID)
	if err != nil {
		return nil, err
	}
	return info, nil
}
