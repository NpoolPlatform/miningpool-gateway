package fraction

import (
	"context"
	"fmt"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"

	fractiongwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fraction"
	"github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fraction"
	fractionmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/fraction"
	orderusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/orderuser"
)

func (h *Handler) getUser(ctx context.Context) error {
	user, err := usermwcli.GetUser(ctx, *h.AppID, *h.UserID)
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("invalid user")
	}
	return nil
}

func (h *Handler) getApp(ctx context.Context) error {
	app, err := appmwcli.GetApp(ctx, *h.AppID)
	if err != nil {
		return err
	}
	if app == nil {
		return fmt.Errorf("invalid app")
	}
	return nil
}

func (h *Handler) CreateFraction(ctx context.Context) (*fractiongwpb.Fraction, error) {
	if err := h.getApp(ctx); err != nil {
		return nil, err
	}

	if err := h.getUser(ctx); err != nil {
		return nil, err
	}

	orderUser, err := orderusermwcli.GetOrderUser(ctx, *h.OrderUserID)
	if err != nil {
		return nil, err
	}

	if orderUser.AppID != *h.AppID || orderUser.UserID != *h.UserID {
		return nil, fmt.Errorf("permission denine")
	}

	info, err := fractionmwcli.CreateFraction(ctx, &fraction.FractionReq{
		AppID:       h.AppID,
		UserID:      h.UserID,
		OrderUserID: h.OrderUserID,
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid fraction")
	}
	if info.UserID != *h.UserID || info.AppID != *h.AppID {
		return nil, fmt.Errorf("permission denied")
	}
	return mw2GW(info), nil
}
