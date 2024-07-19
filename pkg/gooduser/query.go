package gooduser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodusergwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/gooduser"
	goodusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	goodusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/gooduser"
)

func (h *Handler) GetGoodUsers(ctx context.Context) ([]*goodusergwpb.GoodUser, uint32, error) {
	infos, total, err := goodusermwcli.GetGoodUsers(ctx, &goodusermwpb.Conds{}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	return mws2GWs(infos), total, nil
}

func (h *Handler) GetGoodUser(ctx context.Context) (*goodusergwpb.GoodUser, error) {
	info, err := goodusermwcli.GetGoodUser(ctx, *h.EntID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid gooduser")
	}
	return mw2GW(info), nil
}
