package upgrade_company_traffic_packet

import (
	"context"

	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeCompanyTrafficPacketLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeCompanyTrafficPacketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeCompanyTrafficPacketLogic {
	return &DeleteUpgradeCompanyTrafficPacketLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeCompanyTrafficPacketLogic) DeleteUpgradeCompanyTrafficPacket(req *types.IDsReq) (*types.BaseMsgResp, error) {
	return &types.BaseMsgResp{Msg: errormsg.DeleteSuccess}, nil
}
