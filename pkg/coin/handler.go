package coin

import (
	"context"
	"fmt"

	mpbasetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coingw "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/coin"
	coinmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	"github.com/NpoolPlatform/miningpool-gateway/pkg/common"
	constant "github.com/NpoolPlatform/miningpool-gateway/pkg/const"
	poolmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/pool"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID                     *uint32
	EntID                  *string
	PoolID                 *string
	CoinTypeID             *string
	CoinType               *basetypes.CoinType
	RevenueType            *mpbasetypes.RevenueType
	FeeRatio               *string
	FixedRevenueAble       *bool
	LeastTransferAmount    *string
	BenefitIntervalSeconds *uint32
	Remark                 *string
	Offset                 int32
	Limit                  int32
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

func mw2GW(info *coinmw.Coin) *coingw.Coin {
	if info == nil {
		return nil
	}
	return &coingw.Coin{
		ID:                     info.ID,
		EntID:                  info.EntID,
		PoolID:                 info.PoolID,
		CoinTypeID:             info.CoinTypeID,
		CoinType:               info.CoinType,
		RevenueType:            info.RevenueType,
		MiningpoolType:         info.MiningpoolType,
		FeeRatio:               info.FeeRatio,
		FixedRevenueAble:       info.FixedRevenueAble,
		LeastTransferAmount:    info.LeastTransferAmount,
		BenefitIntervalSeconds: info.BenefitIntervalSeconds,
		Remark:                 info.Remark,
		CreatedAt:              info.CreatedAt,
		UpdatedAt:              info.UpdatedAt,
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

func WithCoinType(cointype *basetypes.CoinType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if cointype == nil {
			if must {
				return fmt.Errorf("invalid cointype")
			}
			return nil
		}

		if *cointype == basetypes.CoinType_DefaultCoinType {
			return fmt.Errorf("invalid cointype,not allow be default type")
		}
		h.CoinType = cointype
		return nil
	}
}

func WithRevenueType(revenuetype *mpbasetypes.RevenueType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if revenuetype == nil {
			if must {
				return fmt.Errorf("invalid revenuetype")
			}
			return nil
		}
		if *revenuetype == mpbasetypes.RevenueType_DefaultRevenueType {
			return fmt.Errorf("invalid revenuetype,not allow be default type")
		}
		h.RevenueType = revenuetype
		return nil
	}
}

func WithCoinTypeID(cointypeid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if cointypeid == nil {
			if must {
				return fmt.Errorf("invalid cointypeid")
			}
			return nil
		}
		_, err := uuid.Parse(*cointypeid)
		if err != nil {
			return err
		}

		ccHandler := common.CoinCheckHandler{CoinTypeID: cointypeid}
		err = ccHandler.CheckCoin(ctx)
		if err != nil {
			return err
		}

		h.CoinTypeID = cointypeid
		return nil
	}
}

func WithPoolID(poolid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if poolid == nil {
			if must {
				return fmt.Errorf("invalid poolid")
			}
			return nil
		}
		exist, err := poolmwcli.ExistPool(ctx, *poolid)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid poolid")
		}
		h.PoolID = poolid
		return nil
	}
}

func WithFeeRatio(feeratio *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if feeratio == nil {
			if must {
				return fmt.Errorf("invalid feeratio")
			}
			return nil
		}
		_feeratio, err := decimal.NewFromString(*feeratio)
		if err != nil {
			return err
		}

		if _feeratio.Sign() <= 0 {
			return fmt.Errorf("invalid feeratio")
		}

		h.FeeRatio = feeratio
		return nil
	}
}

func WithFixedRevenueAble(fixedrevenueable *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if fixedrevenueable == nil {
			if must {
				return fmt.Errorf("invalid fixedrevenueable")
			}
			return nil
		}
		h.FixedRevenueAble = fixedrevenueable
		return nil
	}
}

func WithLeastTransferAmount(leastTransferAmount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if leastTransferAmount == nil {
			if must {
				return fmt.Errorf("invalid leasttransferamount")
			}
			return nil
		}
		_leastTransferAmount, err := decimal.NewFromString(*leastTransferAmount)
		if err != nil {
			return err
		}
		if _leastTransferAmount.Sign() <= 0 {
			return fmt.Errorf("invalid leasttransferamount")
		}
		h.LeastTransferAmount = leastTransferAmount
		return nil
	}
}

func WithBenefitIntervalSeconds(benefitintervalseconds *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if benefitintervalseconds == nil {
			if must {
				return fmt.Errorf("invalid benefitintervalseconds")
			}
			return nil
		}
		h.BenefitIntervalSeconds = benefitintervalseconds
		return nil
	}
}

func WithRemark(remark *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if remark == nil {
			if must {
				return fmt.Errorf("invalid remark")
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
