package rootuser

import (
	"context"

	rootusergwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/rootuser"
	rootusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"
	rootusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/rootuser"
)

func (h *Handler) UpdateRootUser(ctx context.Context) (*rootusergwpb.RootUser, error) {
	err := h.checkRootUser(ctx)
	if err != nil {
		return nil, err
	}

	err = rootusermwcli.UpdateRootUser(ctx, &rootusermwpb.RootUserReq{
		ID:        h.ID,
		EntID:     h.EntID,
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
