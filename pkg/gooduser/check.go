package gooduser

import (
	"context"

	coinmwcli "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	goodusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	goodusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/gooduser"
)

func (h *Handler) checkGoodUser(ctx context.Context) error {
	exist, err := goodusermwcli.ExistGoodUserConds(ctx, &goodusermwpb.Conds{
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
		return wlog.Errorf("invalid gooduser")
	}
	return nil
}

func (h *Handler) checkCoinTypeIDs(ctx context.Context, ids []string) error {
	coinInfos, _, err := coinmwcli.GetCoins(ctx, &coin.Conds{
		EntIDs: &v1.StringSliceVal{
			Op:    cruder.IN,
			Value: ids,
		},
	}, 0, int32(len(ids)))
	if err != nil {
		return wlog.WrapError(err)
	}

	for _, id := range ids {
		exist := false
		for _, coinInfo := range coinInfos {
			if coinInfo.EntID == id {
				exist = true
				break
			}
		}
		if !exist {
			return wlog.Errorf("invalid cointypeids")
		}
	}
	return nil
}
