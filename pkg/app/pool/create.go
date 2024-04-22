package pool

import (
	"context"

	poolgwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/app/pool"
	"github.com/NpoolPlatform/message/npool/miningpool/mw/v1/app/pool"
	apppoolmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/app/pool"
	"github.com/google/uuid"
)

func (h *Handler) CreatePool(ctx context.Context) (*poolgwpb.Pool, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	err := apppoolmwcli.CreatePool(ctx, &pool.PoolReq{
		EntID:  h.EntID,
		AppID:  h.TargetAppID,
		PoolID: h.PoolID,
	})
	if err != nil {
		return nil, err
	}
	return fullPools(ctx, *h.EntID)
}
