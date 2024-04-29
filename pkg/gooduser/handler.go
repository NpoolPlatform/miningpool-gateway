package gooduser

import (
	"context"
	"fmt"

	coinmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/coin"
	rootusemwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/rootuser"

	goodusergw "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/gooduser"
	goodusermw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	constant "github.com/NpoolPlatform/miningpool-gateway/pkg/const"
)

type Handler struct {
	ID         *uint32
	EntID      *string
	CoinID     *string
	RootUserID *string
	HashRate   *float32
	Offset     int32
	Limit      int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
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
		MiningpoolType: info.MiningpoolType,
		CoinType:       info.CoinType,
		HashRate:       info.HashRate,
		ReadPageLink:   info.ReadPageLink,
		RevenueType:    info.RevenueType,
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
				return fmt.Errorf("invalid id")
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
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		h.EntID = id
		return nil
	}
}

func WithCoinID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid coinid")
			}
			return nil
		}
		exist, err := coinmwcli.ExistCoin(ctx, *id)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid coinid")
		}
		h.CoinID = id
		return nil
	}
}

func WithRootUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid rootuserid")
			}
			return nil
		}
		exist, err := rootusemwcli.ExistRootUser(ctx, *id)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid rootuserid")
		}
		h.RootUserID = id
		return nil
	}
}

func WithHashRate(hashrate *float32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if hashrate == nil {
			if must {
				return fmt.Errorf("invalid hashrate")
			}
			return nil
		}
		h.HashRate = hashrate
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
