package gooduser

import (
	"context"

	goodusergwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/gooduser"
	goodusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/gooduser"
)

func (h *Handler) DeleteGoodUser(ctx context.Context) (*goodusergwpb.GoodUser, error) {
	_, err := goodusermwcli.DeleteGoodUser(ctx, *h.ID, *h.EntID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
