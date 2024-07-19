package gooduser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodusergwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/gooduser"
	goodusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/gooduser"
)

func (h *Handler) DeleteGoodUser(ctx context.Context) (*goodusergwpb.GoodUser, error) {
	err := h.checkGoodUser(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	info, err := goodusermwcli.GetGoodUser(ctx, *h.EntID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	err = goodusermwcli.DeleteGoodUser(ctx, *h.ID, *h.EntID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return mw2GW(info), nil
}
