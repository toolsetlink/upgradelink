package upgrade_company_traffic_packet

import (
	"context"
	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradecompanytrafficpacket"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeCompanyTrafficPacketListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeCompanyTrafficPacketListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeCompanyTrafficPacketListLogic {
	return &GetUpgradeCompanyTrafficPacketListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeCompanyTrafficPacketListLogic) GetUpgradeCompanyTrafficPacketList(req *types.UpgradeCompanyTrafficPacketListReq) (*types.UpgradeCompanyTrafficPacketListResp, error) {
	var predicates []predicate.UpgradeCompanyTrafficPacket

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	predicates = append(predicates, upgradecompanytrafficpacket.CompanyIDEQ(companyID))

	if req.Status != nil {
		predicates = append(predicates, upgradecompanytrafficpacket.StatusEQ(*req.Status))
	}

	data, err := l.svcCtx.DB.UpgradeCompanyTrafficPacket.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeCompanyTrafficPacketListResp{}
	resp.Msg = errormsg.Success
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {

		// 查询出流量包信息
		packet, err := l.svcCtx.DB.UpgradeTrafficPacket.Get(l.ctx, uint64(v.PacketID))
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}

		// 初始流量大小
		initialSize := utils.BytesToMBString(uint64(v.InitialSize))
		// 剩余流量大小
		remainingSize := utils.BytesToMBString(uint64(v.RemainingSize))

		resp.Data.Data = append(resp.Data.Data,
			types.UpgradeCompanyTrafficPacketInfo{
				Id:            &v.ID,
				CompanyId:     &v.CompanyID,
				PacketId:      &v.PacketID,
				Name:          &packet.Name,
				Key:           &packet.Key,
				StartTime:     pointy.GetPointer(v.StartTime.UnixMilli()),
				EndTime:       pointy.GetPointer(v.EndTime.UnixMilli()),
				InitialSize:   &initialSize,
				RemainingSize: &remainingSize,
				Status:        &v.Status,
				ExchangeTime:  pointy.GetPointer(v.ExchangeTime.UnixMilli()),
				CreateAt:      pointy.GetUnixMilliPointer(v.CreateAt.UnixMilli()),
				UpdateAt:      pointy.GetUnixMilliPointer(v.UpdateAt.UnixMilli()),
			})
	}

	return resp, nil
}
