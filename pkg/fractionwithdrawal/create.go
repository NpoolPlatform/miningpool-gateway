package fractionwithdrawal

import (
	"context"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/google/uuid"

	fractionwithdrawalgwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionwithdrawal"
	"github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionwithdrawal"
	fractionwithdrawalmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/fractionwithdrawal"
	orderusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/orderuser"
)

func (h *Handler) getUser(ctx context.Context) error {
	user, err := usermwcli.GetUser(ctx, *h.AppID, *h.UserID)
	if err != nil {
		return wlog.WrapError(err)
	}
	if user == nil {
		return wlog.Errorf("invalid user")
	}
	return nil
}

func (h *Handler) getApp(ctx context.Context) error {
	app, err := appmwcli.GetApp(ctx, *h.AppID)
	if err != nil {
		return wlog.WrapError(err)
	}
	if app == nil {
		return wlog.Errorf("invalid app")
	}
	return nil
}

func (h *Handler) CreateFractionWithdrawal(ctx context.Context) (*fractionwithdrawalgwpb.FractionWithdrawal, error) {
	if err := h.getApp(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := h.getUser(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	orderUser, err := orderusermwcli.GetOrderUser(ctx, *h.OrderUserID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if orderUser.AppID != *h.AppID || orderUser.UserID != *h.UserID {
		return nil, wlog.Errorf("permission denine")
	}

	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	err = fractionwithdrawalmwcli.CreateFractionWithdrawal(ctx, &fractionwithdrawal.FractionWithdrawalReq{
		EntID:       h.EntID,
		AppID:       h.AppID,
		UserID:      h.UserID,
		OrderUserID: h.OrderUserID,
		CoinTypeID:  h.CoinTypeID,
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetFractionWithdrawal(ctx)
}
