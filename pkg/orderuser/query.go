package orderuser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	orderusergwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/orderuser"
	orderusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/orderuser"
)

func (h *Handler) GetOrderUser(ctx context.Context) (*orderusergwpb.OrderUser, error) {
	info, err := orderusermwcli.GetOrderUser(ctx, *h.EntID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid orderuser")
	}
	if info.UserID != *h.UserID || info.AppID != *h.AppID {
		return nil, wlog.Errorf("permission denied")
	}
	return mw2GW(info), nil
}

func (h *Handler) GetOrderUserProportion(ctx context.Context) (string, error) {
	err := h.checkOrderUser(ctx)
	if err != nil {
		return "0", wlog.WrapError(err)
	}

	proportion, err := orderusermwcli.GetOrderUserProportion(ctx, *h.EntID, *h.CoinTypeID)
	if err != nil {
		return "0", wlog.WrapError(err)
	}
	return proportion, nil
}
