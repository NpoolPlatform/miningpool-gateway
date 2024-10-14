package rootuser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	rootusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"
	rootusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/rootuser"
)

func (h *Handler) checkRootUser(ctx context.Context) error {
	exist, err := rootusermwcli.ExistRootUserConds(ctx, &rootusermwpb.Conds{
		ID: &v1.Uint32Val{
			Op:    cruder.EQ,
			Value: *h.ID,
		},
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: *h.EntID,
		},
	})
	if err != nil {
		return wlog.WrapError(err)
	}

	if !exist {
		return wlog.Errorf("invalid rootuser")
	}
	return nil
}
