package fraction

import (
	"context"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/google/uuid"

	fractiongwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fraction"
	"github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fraction"
	fractionmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/fraction"
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

func (h *Handler) CreateFraction(ctx context.Context) (*fractiongwpb.Fraction, error) {
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

	err = fractionmwcli.CreateFraction(ctx, &fraction.FractionReq{
		EntID:       h.EntID,
		AppID:       h.AppID,
		UserID:      h.UserID,
		OrderUserID: h.OrderUserID,
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetFraction(ctx)
}
