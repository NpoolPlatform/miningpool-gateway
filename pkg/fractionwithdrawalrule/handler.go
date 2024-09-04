package fractionwithdrawalrule

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	fractionwithdrawalrulegw "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionwithdrawalrule"
	fractionwithdrawalrulemw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionwithdrawalrule"
	constant "github.com/NpoolPlatform/miningpool-gateway/pkg/const"
	coinmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/coin"

	"github.com/shopspring/decimal"
)

type Handler struct {
	ID                    *uint32
	EntID                 *string
	PoolCoinTypeID        *string
	WithdrawInterval      *uint32
	PayoutThreshold       *string
	LeastWithdrawalAmount *string
	WithdrawFee           *string
	Offset                int32
	Limit                 int32
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

func mw2GW(info *fractionwithdrawalrulemw.FractionWithdrawalRule) *fractionwithdrawalrulegw.FractionWithdrawalRule {
	if info == nil {
		return nil
	}
	return &fractionwithdrawalrulegw.FractionWithdrawalRule{
		ID:                    info.ID,
		EntID:                 info.EntID,
		PoolID:                info.PoolID,
		PoolCoinTypeID:        info.PoolCoinTypeID,
		WithdrawInterval:      info.WithdrawInterval,
		PayoutThreshold:       info.PayoutThreshold,
		LeastWithdrawalAmount: info.LeastWithdrawalAmount,
		WithdrawFee:           info.WithdrawFee,
		MiningPoolType:        info.MiningPoolType,
		CoinType:              info.CoinType,
		CoinTypeID:            info.CoinTypeID,
		CreatedAt:             info.CreatedAt,
		UpdatedAt:             info.UpdatedAt,
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

func WithPoolCoinTypeID(poolcointypeid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if poolcointypeid == nil {
			if must {
				return wlog.Errorf("invalid poolcointypeid")
			}
			return nil
		}

		exist, err := coinmwcli.ExistCoin(ctx, *poolcointypeid)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid poolcointypeid")
		}
		h.PoolCoinTypeID = poolcointypeid
		return nil
	}
}

func WithWithdrawInterval(withdrawinterval *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if withdrawinterval == nil {
			if must {
				return wlog.Errorf("invalid withdrawinterval")
			}
			return nil
		}
		h.WithdrawInterval = withdrawinterval
		return nil
	}
}

func WithLeastWithdrawalAmount(minamount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if minamount == nil {
			if must {
				return wlog.Errorf("invalid minamount")
			}
			return nil
		}
		_, err := decimal.NewFromString(*minamount)
		if err != nil {
			return wlog.Errorf("invalid minamount,err: %v", err)
		}
		h.LeastWithdrawalAmount = minamount
		return nil
	}
}

func WithPayoutThreshold(payoutthreshold *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if payoutthreshold == nil {
			if must {
				return wlog.Errorf("invalid payoutthreshold")
			}
			return nil
		}
		_, err := decimal.NewFromString(*payoutthreshold)
		if err != nil {
			return wlog.Errorf("invalid payoutthreshold,err: %v", err)
		}
		h.PayoutThreshold = payoutthreshold
		return nil
	}
}

func WithWithdrawFee(withdrawrate *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if withdrawrate == nil {
			if must {
				return wlog.Errorf("invalid withdrawrate")
			}
			return nil
		}
		_, err := decimal.NewFromString(*withdrawrate)
		if err != nil {
			return wlog.Errorf("invalid withdrawrate,err: %v", err)
		}
		h.WithdrawFee = withdrawrate
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
