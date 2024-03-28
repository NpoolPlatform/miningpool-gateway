package pool

import (
	"context"
	"fmt"

	poolgw "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/app/pool"
	apppoolmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/app/pool"
	coinmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	fractionrulemw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"
	poolmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"
	constant "github.com/NpoolPlatform/miningpool-gateway/pkg/const"
)

type Handler struct {
	ID          *uint32
	EntID       *string
	PoolID      *string
	AppID       *string
	TargetAppID *string
	Offset      int32
	Limit       int32
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

func mw2GW(appinfo *apppoolmw.Pool, info *poolmw.Pool, coins []*coinmw.Coin, rules []*fractionrulemw.FractionRule) *poolgw.Pool {
	_coins := []*poolgw.Coin{}
	for _, v := range coins {
		_coins = append(_coins, &poolgw.Coin{
			CoinType:         v.CoinType,
			RevenueTypes:     v.RevenueTypes,
			FeeRate:          v.FeeRate,
			FixedRevenueAble: v.FixedRevenueAble,
			Threshold:        v.Threshold,
			Remark:           v.Remark,
		})
	}

	_rules := []*poolgw.FractionRule{}
	for _, v := range rules {
		_rules = append(_rules, &poolgw.FractionRule{
			CoinType:         v.CoinType,
			WithdrawInterval: v.WithdrawInterval,
			MinAmount:        v.MinAmount,
			WithdrawRate:     v.WithdrawRate,
		})
	}

	return &poolgw.Pool{
		ID:             appinfo.ID,
		EntID:          appinfo.EntID,
		AppID:          appinfo.AppID,
		PoolID:         info.EntID,
		Name:           info.Name,
		MiningpoolType: info.MiningpoolType,
		Site:           info.Site,
		Description:    info.Description,
		Coins:          _coins,
		FractionRules:  _rules,
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

func WithPoolID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid poolid")
			}
			return nil
		}
		h.PoolID = id
		return nil
	}
}

func WithTargetAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid targetappid")
			}
			return nil
		}
		h.TargetAppID = id
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
