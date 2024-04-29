package rootuser

import (
	"context"

	rootusergwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/rootuser"
	rootusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"
	rootusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/rootuser"

	"github.com/google/uuid"
)

func (h *Handler) CreateRootUser(ctx context.Context) (*rootusergwpb.RootUser, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	err := rootusermwcli.CreateRootUser(ctx, &rootusermwpb.RootUserReq{
		EntID:     h.EntID,
		PoolID:    h.PoolID,
		Name:      h.Name,
		Email:     h.Email,
		AuthToken: h.AuthToken,
		Remark:    h.Remark,
	})
	if err != nil {
		return nil, err
	}

	return h.GetRootUser(ctx)
}
