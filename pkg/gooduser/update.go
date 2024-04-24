package gooduser

import (
	"context"

	goodusergwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/gooduser"
	goodusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	goodusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/gooduser"
)

func (h *Handler) UpdateGoodUser(ctx context.Context) (*goodusergwpb.GoodUser, error) {
	err := h.checkGoodUser(ctx)
	if err != nil {
		return nil, err
	}

	err = goodusermwcli.UpdateGoodUser(ctx, &goodusermwpb.GoodUserReq{
		ID:          h.ID,
		EntID:       h.EntID,
		RootUserID:  h.RootUserID,
		HashRate:    h.HashRate,
		RevenueType: h.RevenueType,
	})
	if err != nil {
		return nil, err
	}
	return h.GetGoodUser(ctx)
}
