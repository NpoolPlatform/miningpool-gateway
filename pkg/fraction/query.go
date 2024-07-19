package fraction

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	fractiongwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fraction"
	"github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fraction"
	fractionmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/fraction"
)

func (h *Handler) GetFraction(ctx context.Context) (*fractiongwpb.Fraction, error) {
	info, err := fractionmwcli.GetFraction(ctx, *h.EntID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid fraction")
	}
	if info.UserID != *h.UserID || info.AppID != *h.AppID {
		return nil, wlog.Errorf("permission denied")
	}
	return mw2GW(info), nil
}

func (h *Handler) GetUserFractions(ctx context.Context) ([]*fractiongwpb.Fraction, uint32, error) {
	infos, total, err := fractionmwcli.GetFractions(ctx, &fraction.Conds{
		AppID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: *h.AppID,
		},
		UserID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: *h.UserID,
		},
	}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	return mws2GWs(infos), total, nil
}
