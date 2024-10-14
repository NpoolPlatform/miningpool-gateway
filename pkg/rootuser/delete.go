package rootuser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	rootusergwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/rootuser"
	rootusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/rootuser"
)

func (h *Handler) DeleteRootUser(ctx context.Context) (*rootusergwpb.RootUser, error) {
	err := h.checkRootUser(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	info, err := rootusermwcli.GetRootUser(ctx, *h.EntID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	err = rootusermwcli.DeleteRootUser(ctx, *h.ID, *h.EntID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return mw2GW(info), nil
}
