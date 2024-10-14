package orderuser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	orderusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/orderuser"

	orderusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
)

func (h *Handler) checkOrderUser(ctx context.Context) error {
	exist, err := orderusermwcli.ExistOrderUserConds(ctx, &orderusermwpb.Conds{
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: *h.EntID,
		},
		AppID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: *h.AppID,
		},
		UserID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: *h.UserID,
		},
	})
	if err != nil {
		return wlog.WrapError(err)
	}

	if !exist {
		return wlog.Errorf("invalid orderuser")
	}
	return nil
}
