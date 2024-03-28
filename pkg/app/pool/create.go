package pool

import (
	"context"

	poolgwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/app/pool"
	"github.com/NpoolPlatform/message/npool/miningpool/mw/v1/app/pool"
	apppoolmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/app/pool"
)

func (h *Handler) CreatePool(ctx context.Context) (*poolgwpb.Pool, error) {
	info, err := apppoolmwcli.CreatePool(ctx, &pool.PoolReq{
		AppID:  h.TargetAppID,
		PoolID: h.PoolID,
	})
	if err != nil {
		return nil, err
	}
	return h.fullPools(ctx, info)
}
