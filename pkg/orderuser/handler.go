package orderuser

import (
	"context"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/shopspring/decimal"

	orderusergw "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/orderuser"
	orderusermw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	"github.com/NpoolPlatform/miningpool-gateway/pkg/common"
	constant "github.com/NpoolPlatform/miningpool-gateway/pkg/const"
)

type Handler struct {
	ID         *uint32
	EntID      *string
	AppID      *string
	UserID     *string
	CoinTypeID *string
	Proportion *decimal.Decimal
	Offset     int32
	Limit      int32
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
		ReadPageLink:   info.ReadPageLink,
		MiningPoolType: info.MiningPoolType,
		CreatedAt:      info.CreatedAt,
		UpdatedAt:      info.UpdatedAt,
	}
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
		h.UserID = id
		return nil
	}
}

func WithCoinTypeID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid cointypeid")
			}
			return nil
		}
		ccHandler := common.CoinCheckHandler{CoinTypeID: id}
		err := ccHandler.CheckCoin(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.CoinTypeID = id
		return nil
	}
}

func WithProportion(proportion *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if proportion == nil {
			if must {
				return wlog.Errorf("invalid proportion")
			}
			return nil
		}
		_proportion, err := decimal.NewFromString(*proportion)
		if err != nil {
			return wlog.Errorf("invalid proportion,err: %v", err)
		}
		h.Proportion = &_proportion
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
