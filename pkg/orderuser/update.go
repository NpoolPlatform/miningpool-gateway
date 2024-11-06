package orderuser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	orderusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/orderuser"
)

func (h *Handler) UpdateOrderUser(ctx context.Context) error {
	err := h.checkOrderUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	proportion := h.Proportion.String()
	err = orderusermwcli.UpdateOrderUser(ctx, &orderuser.OrderUserReq{
		EntID:      h.EntID,
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		Proportion: &proportion,
	})
	return wlog.WrapError(err)
}
