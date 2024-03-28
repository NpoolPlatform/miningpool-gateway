package orderuser

import (
	"context"
	"fmt"

	orderusergwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/orderuser"
	orderusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/orderuser"
)

func (h *Handler) GetOrderUser(ctx context.Context) (*orderusergwpb.OrderUser, error) {
	info, err := orderusermwcli.GetOrderUser(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid orderuser")
	}
	if info.UserID != *h.UserID || info.AppID != *h.AppID {
		return nil, fmt.Errorf("permission denied")
	}
	return mw2GW(info), nil
}
