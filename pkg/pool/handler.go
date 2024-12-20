package pool

import (
	"context"
	"regexp"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	poolgw "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/pool"
	coinmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	fractionwithdrawalrulemw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionwithdrawalrule"
	poolmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"
	constant "github.com/NpoolPlatform/miningpool-gateway/pkg/const"
)

type Handler struct {
	ID             *uint32
	EntID          *string
	MiningPoolType *basetypes.MiningPoolType
	Name           *string
	Site           *string
	Logo           *string
	Description    *string
	Offset         int32
	Limit          int32
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

func mw2GW(info *poolmw.Pool, coins []*coinmw.Coin, rules []*fractionwithdrawalrulemw.FractionWithdrawalRule) *poolgw.Pool {
	if info == nil {
		return nil
	}
	_coins := []*poolgw.Coin{}
	for _, v := range coins {
		_coins = append(_coins, &poolgw.Coin{
			EntID:                  v.EntID,
			PoolID:                 v.PoolID,
			CoinTypeID:             v.CoinTypeID,
			CoinType:               v.CoinType,
			FeeRatio:               v.FeeRatio,
			FixedRevenueAble:       v.FixedRevenueAble,
			LeastTransferAmount:    v.LeastTransferAmount,
			BenefitIntervalSeconds: v.BenefitIntervalSeconds,
			Remark:                 v.Remark,
			CreatedAt:              v.CreatedAt,
			UpdatedAt:              v.UpdatedAt,
		})
	}

	_rules := []*poolgw.FractionWithdrawalRule{}
	for _, v := range rules {
		_rules = append(_rules, &poolgw.FractionWithdrawalRule{
			EntID:                 v.EntID,
			PoolCoinTypeID:        v.PoolCoinTypeID,
			WithdrawInterval:      v.WithdrawInterval,
			PayoutThreshold:       v.PayoutThreshold,
			LeastWithdrawalAmount: v.LeastWithdrawalAmount,
			WithdrawFee:           v.WithdrawFee,
			CreatedAt:             v.CreatedAt,
			UpdatedAt:             v.UpdatedAt,
		})
	}

	return &poolgw.Pool{
		ID:                      info.ID,
		EntID:                   info.EntID,
		MiningPoolType:          info.MiningPoolType,
		Name:                    info.Name,
		Site:                    info.Site,
		Logo:                    info.Logo,
		Description:             info.Description,
		Coins:                   _coins,
		FractionWithdrawalRules: _rules,
		CreatedAt:               info.CreatedAt,
		UpdatedAt:               info.UpdatedAt,
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

func WithMiningPoolType(miningpooltype *basetypes.MiningPoolType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if miningpooltype == nil {
			if must {
				return wlog.Errorf("invalid miningpooltype")
			}
			return nil
		}
		if *miningpooltype == basetypes.MiningPoolType_DefaultMiningPoolType {
			return wlog.Errorf("invalid miningpooltype,not allow be default type")
		}
		h.MiningPoolType = miningpooltype

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
		re := regexp.MustCompile("^[a-zA-Z0-9\u3040-\u31ff][[a-zA-Z0-9_\\-\\.\u3040-\u31ff]{3,32}$") //nolint
		if !re.MatchString(*name) {
			return wlog.Errorf("invalid name")
		}
		h.Name = name
		return nil
	}
}

func WithSite(site *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if site == nil {
			if must {
				return wlog.Errorf("invalid site")
			}
			return nil
		}
		h.Site = site
		return nil
	}
}

func WithLogo(logo *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if logo == nil {
			if must {
				return wlog.Errorf("invalid logo")
			}
			return nil
		}
		h.Logo = logo
		return nil
	}
}

func WithDescription(description *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if description == nil {
			if must {
				return wlog.Errorf("invalid description")
			}
			return nil
		}
		h.Description = description
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
