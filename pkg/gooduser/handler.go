package gooduser

import (
	"context"

	coinmwcli "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	rootusemwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/rootuser"

	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	goodusergw "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/gooduser"
	goodusermw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	constant "github.com/NpoolPlatform/miningpool-gateway/pkg/const"
)

type Handler struct {
	ID          *uint32
	EntID       *string
	CoinTypeIDs []string
	RootUserID  *string
	Offset      int32
	Limit       int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	return handler, nil
}

func mw2GW(info *goodusermw.GoodUser) *goodusergw.GoodUser {
	if info == nil {
		return nil
	}
	return &goodusergw.GoodUser{
		ID:             info.ID,
		EntID:          info.EntID,
		Name:           info.Name,
		RootUserID:     info.RootUserID,
		PoolID:         info.PoolID,
		MiningpoolType: info.MiningpoolType,
		ReadPageLink:   info.ReadPageLink,
		CreatedAt:      info.CreatedAt,
		UpdatedAt:      info.UpdatedAt,
	}
}

func mws2GWs(infos []*goodusermw.GoodUser) []*goodusergw.GoodUser {
	_infos := []*goodusergw.GoodUser{}
	for _, v := range infos {
		_infos = append(_infos, mw2GW(v))
	}
	return _infos
}

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid entid")
			}
			return nil
		}
		h.EntID = id
		return nil
	}
}

func WithCoinTypeIDs(ids []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
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

		h.CoinTypeIDs = ids
		return nil
	}
}

func WithRootUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid rootuserid")
			}
			return nil
		}
		exist, err := rootusemwcli.ExistRootUser(ctx, *id)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid rootuserid")
		}
		h.RootUserID = id
		return nil
	}
}

func WithOffset(n int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = n
		return nil
	}
}

func WithLimit(n int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == 0 {
			n = constant.DefaultRowLimit
		}
		h.Limit = n
		return nil
	}
}
