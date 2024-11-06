package fractionwithdrawal

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	fractionwithdrawalgwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/fractionwithdrawal"
	"github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionwithdrawal"
	fractionwithdrawalmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/fractionwithdrawal"
)

func (h *Handler) GetFractionWithdrawal(ctx context.Context) (*fractionwithdrawalgwpb.FractionWithdrawal, error) {
	info, err := fractionwithdrawalmwcli.GetFractionWithdrawal(ctx, *h.EntID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid fractionwithdrawal")
	}
	if info.UserID != *h.UserID || info.AppID != *h.AppID {
		return nil, wlog.Errorf("permission denied")
	}
	return mw2GW(info), nil
}

func (h *Handler) GetUserFractionWithdrawals(ctx context.Context) ([]*fractionwithdrawalgwpb.FractionWithdrawal, uint32, error) {
	infos, total, err := fractionwithdrawalmwcli.GetFractionWithdrawals(ctx, &fractionwithdrawal.Conds{
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
