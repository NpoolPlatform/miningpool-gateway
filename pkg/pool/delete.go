package pool

import (
	"context"

	poolgwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/pool"
	poolmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/pool"
)

func (h *Handler) DeletePool(ctx context.Context) (*poolgwpb.Pool, error) {
	err := h.checkPool(ctx)
	if err != nil {
		return nil, err
	}

	info, err := h.GetPool(ctx)
	if err != nil {
		return nil, err
	}

	err = poolmwcli.DeletePool(ctx, *h.ID, *h.EntID)
	if err != nil {
		return nil, err
	}
	return info, nil
}
