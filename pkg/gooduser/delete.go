package gooduser

import (
	"context"
	"fmt"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	goodusergwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/gooduser"
	goodusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	goodusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/gooduser"
)

func (h *Handler) DeleteGoodUser(ctx context.Context) (*goodusergwpb.GoodUser, error) {
	info, err := goodusermwcli.GetGoodUserOnly(ctx, &goodusermwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid gooduser")
	}

	_, err = goodusermwcli.DeleteGoodUser(ctx, info.ID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
