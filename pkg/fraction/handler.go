package fraction

import (
	"context"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	orderusemwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/orderuser"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	fractiongw "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fraction"
	fractionmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fraction"
	constant "github.com/NpoolPlatform/miningpool-gateway/pkg/const"
)

type Handler struct {
	ID            *uint32
	EntID         *string
	AppID         *string
	UserID        *string
	OrderUserID   *string
	WithdrawState *basetypes.WithdrawState
	WithdrawAt    *uint32
	PromisePayAt  *uint32
	Msg           *string
	Offset        int32
	Limit         int32
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

func mw2GW(info *fractionmw.Fraction) *fractiongw.Fraction {
	if info == nil {
		return nil
	}
	return &fractiongw.Fraction{
		ID:            info.ID,
		EntID:         info.EntID,
		AppID:         info.AppID,
		UserID:        info.UserID,
		OrderUserID:   info.OrderUserID,
		WithdrawState: info.WithdrawState,
		WithdrawAt:    info.WithdrawAt,
		PromisePayAt:  info.PromisePayAt,
		Msg:           info.Msg,
		CreatedAt:     info.CreatedAt,
		UpdatedAt:     info.UpdatedAt,
	}
}

func mws2GWs(infos []*fractionmw.Fraction) []*fractiongw.Fraction {
	_infos := []*fractiongw.Fraction{}
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

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid appid")
			}
			return nil
		}
		exist, err := appmwcli.ExistApp(ctx, *id)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid appid")
		}
		h.AppID = id
		return nil
	}
}

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid userid")
			}
			return nil
		}

		if h.AppID == nil {
			return wlog.Errorf("invalid appid")
		}

		exist, err := usermwcli.ExistUser(ctx, *h.AppID, *id)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid user")
		}

		h.UserID = id
		return nil
	}
}

func WithOrderUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid orderuserid")
			}
			return nil
		}
		exist, err := orderusemwcli.ExistOrderUser(ctx, *id)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid orderuser")
		}
		h.OrderUserID = id
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
