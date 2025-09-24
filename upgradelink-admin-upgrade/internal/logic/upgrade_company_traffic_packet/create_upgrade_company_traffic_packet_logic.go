package upgrade_company_traffic_packet

import (
	"context"
	"upgradelink-admin-upgrade/ent/predicate"
	"upgradelink-admin-upgrade/ent/upgradetrafficpacket"
	"upgradelink-admin-upgrade/internal/logic/base"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"
	"time"

	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeCompanyTrafficPacketLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewCreateUpgradeCompanyTrafficPacketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeCompanyTrafficPacketLogic {
	return &CreateUpgradeCompanyTrafficPacketLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeCompanyTrafficPacketLogic) CreateUpgradeCompanyTrafficPacket(req *types.UpgradeCompanyTrafficPacketCreateInfo) (*types.BaseMsgResp, error) {

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckCreate(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	// 通过 key查询出流量包信息
	// 查询出流量包 id
	packetData, err := l.svcCtx.DB.UpgradeTrafficPacket.Query().Where(upgradetrafficpacket.Key(*req.Key)).First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	// 效验流量包是否存在
	if packetData == nil {
		return nil, errorx.NewCodeInvalidArgumentError("流量包不存在")
	}
	packetDataId := int(packetData.ID)

	// 修改流量包状态为已兑换
	_, err = l.svcCtx.DB.UpgradeTrafficPacket.UpdateOneID(packetData.ID).SetStatus(2).Save(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	// 计算流量包有效期开始时间与结束时间
	startTime := time.Now().UnixMilli()
	endTime := time.Now().AddDate(0, 0, int(packetData.ValidDays)).UnixMilli()

	statusOne := int32(1)
	// 生成客户流量包数据
	_, err = l.svcCtx.DB.UpgradeCompanyTrafficPacket.Create().
		SetNotNilCompanyID(&companyID).
		SetNotNilPacketID(&packetDataId).
		SetNotNilStartTime(pointy.GetTimeMilliPointer(&startTime)).
		SetNotNilEndTime(pointy.GetTimeMilliPointer(&endTime)).
		SetNotNilInitialSize(&packetData.Size).
		SetNotNilRemainingSize(&packetData.Size).
		SetNotNilStatus(&statusOne).
		SetNotNilExchangeTime(pointy.GetTimeMilliPointer(&startTime)).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.CreateSuccess}, nil
}

func (l *CreateUpgradeCompanyTrafficPacketLogic) CheckCreate(req *types.UpgradeCompanyTrafficPacketCreateInfo) error {
	// 判断流量包是否已经被兑换
	var predicates []predicate.UpgradeTrafficPacket
	predicates = append(predicates, upgradetrafficpacket.Key(*req.Key))
	predicates = append(predicates, upgradetrafficpacket.StatusEQ(1))
	count, err := l.svcCtx.DB.UpgradeTrafficPacket.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count == 0 {
		return errorx.NewCodeInvalidArgumentError("流量包已经被兑换")
	}

	return nil
}
