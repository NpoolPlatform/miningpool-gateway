package gooduser

import (
	"context"

	goodusergwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/gooduser"
	goodusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	goodusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/gooduser"

	"github.com/google/uuid"
)

func (h *Handler) CreateGoodUser(ctx context.Context) (*goodusergwpb.GoodUser, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	err := goodusermwcli.CreateGoodUser(ctx, &goodusermwpb.GoodUserReq{
		EntID:      h.EntID,
		RootUserID: h.RootUserID,
		CoinID:     h.CoinID,
		HashRate:   h.HashRate,
	})
	if err != nil {
		return nil, err
	}

	return h.GetGoodUser(ctx)
}
