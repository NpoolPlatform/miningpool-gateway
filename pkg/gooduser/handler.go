package gooduser

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	goodusergw "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/gooduser"
	goodusermw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	constant "github.com/NpoolPlatform/miningpool-gateway/pkg/const"
)

type Handler struct {
	ID             *uint32
	EntID          *string
	GoodID         *string
	RootUserID     *string
	Name           *string
	MiningpoolType *basetypes.MiningpoolType
	CoinType       *basetypes.CoinType
	HashRate       *float32
	ReadPageLink   *string
	RevenueType    *basetypes.RevenueType
	Offset         int32
	Limit          int32
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

func WithRootUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid rootuserid")
			}
			return nil
		}
		h.RootUserID = id
		return nil
	}
}

func WithGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid goodid")
			}
			return nil
		}
		h.GoodID = id
		return nil
	}
}

func WithName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return fmt.Errorf("invalid name")
			}
			return nil
		}
		h.Name = name
		return nil
	}
}

func WithMiningpoolType(miningpooltype *basetypes.MiningpoolType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if miningpooltype == nil {
			if must {
				return fmt.Errorf("invalid miningpooltype")
			}
			return nil
		}
		if miningpooltype == basetypes.MiningpoolType_DefaultMiningpoolType.Enum() {
			return fmt.Errorf("invalid miningpooltype,not allow be default type")
		}
		h.MiningpoolType = miningpooltype
		return nil
	}
}

func WithCoinType(cointype *basetypes.CoinType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if cointype == nil {
			if must {
				return fmt.Errorf("invalid cointype")
			}
			return nil
		}
		if cointype == basetypes.CoinType_DefaultCoinType.Enum() {
			return fmt.Errorf("invalid cointype,not allow be default type")
		}
		h.CoinType = cointype
		return nil
	}
}

func WithRevenueType(revenuetype *basetypes.RevenueType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if revenuetype == nil {
			if must {
				return fmt.Errorf("invalid revenuetype")
			}
			return nil
		}
		h.RevenueType = revenuetype
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

func WithReadPageLink(readpagelink *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if readpagelink == nil {
			if must {
				return fmt.Errorf("invalid readpagelink")
			}
			return nil
		}
		h.ReadPageLink = readpagelink
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
