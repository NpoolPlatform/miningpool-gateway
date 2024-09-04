package rootuser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	rootusergw "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/rootuser"
	rootusermw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"
	constant "github.com/NpoolPlatform/miningpool-gateway/pkg/const"
)

type Handler struct {
	ID        *uint32
	EntID     *string
	PoolID    *string
	Name      *string
	Email     *string
	AuthToken *string
	Authed    *bool
	Remark    *string
	Offset    int32
	Limit     int32
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

func mw2GW(info *rootusermw.RootUser) *rootusergw.RootUser {
	if info == nil {
		return nil
	}
	return &rootusergw.RootUser{
		ID:             info.ID,
		EntID:          info.EntID,
		PoolID:         info.PoolID,
		Name:           info.Name,
		Email:          info.Email,
		AuthToken:      info.AuthToken,
		Authed:         info.Authed,
		Remark:         info.Remark,
		MiningPoolType: info.MiningPoolType,
		CreatedAt:      info.CreatedAt,
		UpdatedAt:      info.UpdatedAt,
	}
}

func mws2GWs(infos []*rootusermw.RootUser) []*rootusergw.RootUser {
	_infos := []*rootusergw.RootUser{}
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

func WithPoolID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid poolid")
			}
			return nil
		}
		h.PoolID = id
		return nil
	}
}

func WithName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return wlog.Errorf("invalid name")
			}
			return nil
		}
		h.Name = name
		return nil
	}
}

func WithEmail(email *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if email == nil {
			if must {
				return wlog.Errorf("invalid email")
			}
			return nil
		}
		h.Email = email
		return nil
	}
}

func WithAuthToken(authtoken *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if authtoken == nil {
			if must {
				return wlog.Errorf("invalid authtoken")
			}
			return nil
		}
		h.AuthToken = authtoken
		return nil
	}
}

func WithRemark(remark *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if remark == nil {
			if must {
				return wlog.Errorf("invalid remark")
			}
			return nil
		}
		h.Remark = remark
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
