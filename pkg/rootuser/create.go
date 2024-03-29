package rootuser

import (
	"context"

	rootusergwpb "github.com/NpoolPlatform/message/npool/miningpool/gw/v1/rootuser"
	rootusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"
	rootusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/rootuser"

	"github.com/google/uuid"
)

func (h *Handler) CreateRootUser(ctx context.Context) (*rootusergwpb.RootUser, error) {
	id := uuid.NewString()
	if h.EntID == nil {
		h.EntID = &id
	}

	info, err := rootusermwcli.CreateRootUser(ctx, &rootusermwpb.RootUserReq{
		EntID:          h.EntID,
		Name:           h.Name,
		MiningpoolType: h.MiningpoolType,
		Email:          h.Email,
		AuthToken:      h.AuthToken,
		Remark:         h.Remark,
	})
	if err != nil {
		return nil, err
	}

	return mw2GW(info), nil
}
