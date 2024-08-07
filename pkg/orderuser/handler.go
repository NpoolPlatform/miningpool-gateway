package orderuser

import (
	"context"
	"fmt"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"

	orderusergw "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/orderuser"
	orderusermw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	constant "github.com/NpoolPlatform/miningpool-gateway/pkg/const"
)

type Handler struct {
	ID     *uint32
	EntID  *string
	AppID  *string
	UserID *string
	Offset int32
	Limit  int32
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
	if info == nil {
		return nil
	}
	return &orderusergw.OrderUser{
		ID:             info.ID,
		EntID:          info.EntID,
		AppID:          info.AppID,
		UserID:         info.UserID,
		RootUserID:     info.RootUserID,
		GoodUserID:     info.GoodUserID,
		Name:           info.Name,
		Proportion:     info.Proportion,
		RevenueAddress: info.RevenueAddress,
		ReadPageLink:   info.ReadPageLink,
		AutoPay:        info.AutoPay,
		MiningpoolType: info.MiningpoolType,
		CoinType:       info.CoinType,
		RevenueType:    info.RevenueType,
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
		exist, err := appmwcli.ExistApp(ctx, *id)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid appid")
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
