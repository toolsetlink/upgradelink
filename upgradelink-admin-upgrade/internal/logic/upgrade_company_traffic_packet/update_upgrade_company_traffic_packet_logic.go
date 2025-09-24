package upgrade_company_traffic_packet

import (
	"context"

	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeCompanyTrafficPacketLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUpgradeCompanyTrafficPacketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeCompanyTrafficPacketLogic {
	return &UpdateUpgradeCompanyTrafficPacketLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeCompanyTrafficPacketLogic) UpdateUpgradeCompanyTrafficPacket(req *types.UpgradeCompanyTrafficPacketInfo) (*types.BaseMsgResp, error) {
	return &types.BaseMsgResp{Msg: errormsg.UpdateSuccess}, nil
}
