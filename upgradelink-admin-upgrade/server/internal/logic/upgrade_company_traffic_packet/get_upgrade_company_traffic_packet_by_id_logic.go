package upgrade_company_traffic_packet

import (
	"context"
	"upgradelink-admin-upgrade/server/internal/utils"

	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeCompanyTrafficPacketByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeCompanyTrafficPacketByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeCompanyTrafficPacketByIdLogic {
	return &GetUpgradeCompanyTrafficPacketByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeCompanyTrafficPacketByIdLogic) GetUpgradeCompanyTrafficPacketById(req *types.IDReq) (*types.UpgradeCompanyTrafficPacketInfoResp, error) {
	data, err := l.svcCtx.DB.UpgradeCompanyTrafficPacket.Get(l.ctx, uint64(req.Id))
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	// 初始流量大小
	initialSize := utils.BytesToMBString(uint64(data.InitialSize))
	// 剩余流量大小
	remainingSize := utils.BytesToMBString(uint64(data.RemainingSize))

	return &types.UpgradeCompanyTrafficPacketInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  errormsg.Success,
		},

		Data: types.UpgradeCompanyTrafficPacketInfo{
			Id:            &data.ID,
			CompanyId:     &data.CompanyID,
			PacketId:      &data.PacketID,
			StartTime:     pointy.GetPointer(data.StartTime.UnixMilli()),
			EndTime:       pointy.GetPointer(data.EndTime.UnixMilli()),
			InitialSize:   &initialSize,
			RemainingSize: &remainingSize,
			Status:        &data.Status,
			ExchangeTime:  pointy.GetPointer(data.ExchangeTime.UnixMilli()),
			CreateAt:      pointy.GetUnixMilliPointer(data.CreateAt.UnixMilli()),
			UpdateAt:      pointy.GetUnixMilliPointer(data.UpdateAt.UnixMilli()),
		},
	}, nil

}
