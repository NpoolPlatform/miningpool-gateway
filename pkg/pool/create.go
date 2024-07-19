//nolint:dupl
package pool

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	poolgwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/pool"
	poolmwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"
	poolmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/pool"

	"github.com/google/uuid"
)

func (h *Handler) CreatePool(ctx context.Context) (*poolgwpb.Pool, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	err := poolmwcli.CreatePool(ctx, &poolmwpb.PoolReq{
		EntID:          h.EntID,
		MiningpoolType: h.MiningpoolType,
		Name:           h.Name,
		Site:           h.Site,
		Logo:           h.Logo,
		Description:    h.Description,
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetPool(ctx)
}
