package rootuser

import (
	"context"

	rootusergwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/rootuser"
	rootusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/rootuser"
)

func (h *Handler) DeleteRootUser(ctx context.Context) (*rootusergwpb.RootUser, error) {
	info, err := rootusermwcli.GetRootUser(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	err = rootusermwcli.DeleteRootUser(ctx, *h.ID, *h.EntID)
	if err != nil {
		return nil, err
	}
	return mw2GW(info), nil
}
