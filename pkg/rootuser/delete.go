package rootuser

import (
	"context"
	"fmt"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	rootusergwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/rootuser"
	rootusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"
	rootusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/rootuser"
)

func (h *Handler) DeleteRootUser(ctx context.Context) (*rootusergwpb.RootUser, error) {
	info, err := rootusermwcli.GetRootUserOnly(ctx, &rootusermwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid rootuser")
	}

	_, err = rootusermwcli.DeleteRootUser(ctx, *h.ID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
