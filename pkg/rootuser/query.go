package rootuser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	rootusergwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/rootuser"
	rootusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"
	rootusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/rootuser"
)

func (h *Handler) GetRootUsers(ctx context.Context) ([]*rootusergwpb.RootUser, uint32, error) {
	infos, total, err := rootusermwcli.GetRootUsers(ctx, &rootusermwpb.Conds{}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	return mws2GWs(infos), total, nil
}

func (h *Handler) GetRootUser(ctx context.Context) (*rootusergwpb.RootUser, error) {
	info, err := rootusermwcli.GetRootUser(ctx, *h.EntID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid rootuser")
	}
	return mw2GW(info), nil
}
