package orderuser

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	orderusergw "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/orderuser"
	orderusermw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	constant "github.com/NpoolPlatform/miningpool-gateway/pkg/const"
)

type Handler struct {
	ID             *uint32
	EntID          *string
	Name           *string
	RootUserID     *string
	GoodUserID     *string
	AppID          *string
	UserID         *string
	MiningpoolType *basetypes.MiningpoolType
	CoinType       *basetypes.CoinType
	Proportion     *float32
	RevenueAddress *string
	ReadPageLink   *string
	AutoPay        *bool
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

func mw2GW(info *orderusermw.OrderUser) *orderusergw.OrderUser {
	return &orderusergw.OrderUser{
		ID:             info.ID,
		EntID:          info.EntID,
		Name:           info.Name,
		RootUserID:     info.RootUserID,
		GoodUserID:     info.GoodUserID,
		AppID:          info.AppID,
		UserID:         info.UserID,
		MiningpoolType: info.MiningpoolType,
		CoinType:       info.CoinType,
		Proportion:     info.Proportion,
		ReadPageLink:   info.ReadPageLink,
		RevenueAddress: info.RevenueAddress,
		AutoPay:        info.AutoPay,
		CreatedAt:      info.CreatedAt,
		UpdatedAt:      info.UpdatedAt,
	}
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

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
			return nil
		}
		h.AppID = id
		return nil
	}
}

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid userid")
			}
			return nil
		}
		h.UserID = id
		return nil
	}
}

func WithRevenueAddress(addr *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if addr == nil {
			if must {
				return fmt.Errorf("invalid revenueaddress")
			}
			return nil
		}
		h.RevenueAddress = addr
		return nil
	}
}

func WithAutoPay(autopay *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if autopay == nil {
			if must {
				return fmt.Errorf("invalid autopay")
			}
			return nil
		}
		h.AutoPay = autopay
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
