package orderuser

import (
	"context"
	"fmt"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	orderusergwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/orderuser"
	orderusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	orderusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/orderuser"
)

func (h *Handler) checkAuth(ctx context.Context) (*orderusermwpb.OrderUser, error) {
	info, err := orderusermwcli.GetOrderUserOnly(ctx, &orderusermwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid orderuser")
	}
	if info.UserID != *h.UserID || info.AppID != *h.AppID {
		return nil, fmt.Errorf("permission denied")
	}
	return info, nil
}

func (h *Handler) UpdateOrderUser(ctx context.Context) (*orderusergwpb.OrderUser, error) {
	info, err := h.checkAuth(ctx)
	if err != nil {
		return nil, err
	}

	_info, err := orderusermwcli.UpdateOrderUser(ctx, &orderusermwpb.OrderUserReq{
		ID:             &info.ID,
		EntID:          &info.EntID,
		RevenueAddress: h.RevenueAddress,
		AutoPay:        h.AutoPay,
	})
	if err != nil {
		return nil, err
	}
	return mw2GW(_info), nil
}
