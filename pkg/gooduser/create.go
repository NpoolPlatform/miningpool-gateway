package gooduser

import (
	"context"

	goodusergwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/gooduser"
	goodusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	goodusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/gooduser"

	"github.com/google/uuid"
)

func (h *Handler) CreateGoodUser(ctx context.Context) (*goodusergwpb.GoodUser, error) {
	id := uuid.NewString()
	if h.EntID == nil {
		h.EntID = &id
	}

	info, err := goodusermwcli.CreateGoodUser(ctx, &goodusermwpb.GoodUserReq{
		EntID:       h.EntID,
		RootUserID:  h.RootUserID,
		CoinType:    h.CoinType,
		HashRate:    h.HashRate,
		RevenueType: h.RevenueType,
	})
	if err != nil {
		return nil, err
	}

	return mw2GW(info), nil
}
