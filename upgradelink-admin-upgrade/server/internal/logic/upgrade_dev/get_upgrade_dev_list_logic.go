package upgrade_dev

import (
	"context"
	"time"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradedev"
	"upgradelink-admin-upgrade/server/ent/upgradedevgrouprelation"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeDevListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	companyID int
}

func NewGetUpgradeDevListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeDevListLogic {
	return &GetUpgradeDevListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeDevListLogic) GetUpgradeDevList(req *types.UpgradeDevListReq) (*types.UpgradeDevListResp, error) {
	var predicates []predicate.UpgradeDev

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	predicates = append(predicates, upgradedev.CompanyIDEQ(companyID))

	// 删除状态
	predicates = append(predicates, upgradedev.IsDelEQ(0))

	if req.Key != nil {
		predicates = append(predicates, upgradedev.KeyContains(*req.Key))
	}
	if req.IsDel != nil {
		predicates = append(predicates, upgradedev.IsDelEQ(*req.IsDel))
	}
	if req.CreateAt != nil {
		predicates = append(predicates, upgradedev.CreateAtGTE(time.UnixMilli(*req.CreateAt)))
	}
	if req.UpdateAt != nil {
		predicates = append(predicates, upgradedev.UpdateAtGTE(time.UnixMilli(*req.UpdateAt)))
	}
	data, err := l.svcCtx.DB.UpgradeDev.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeDevListResp{}
	resp.Msg = errormsg.Success
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {

		// 查询出关联的设备组
		var predicates []predicate.UpgradeDevGroupRelation
		predicates = append(predicates, upgradedevgrouprelation.DevID(v.ID))
		predicates = append(predicates, upgradedevgrouprelation.IsDelEQ(0))
		devGroupList, err := l.svcCtx.DB.UpgradeDevGroupRelation.Query().Where(predicates...).All(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}
		var devGroupIds []uint64
		for _, group := range devGroupList {
			devGroupIds = append(devGroupIds, uint64(group.DevGroupID))
		}

		resp.Data.Data = append(resp.Data.Data,
			types.UpgradeDevInfo{
				Id:          &v.ID,
				Key:         &v.Key,
				IsDel:       &v.IsDel,
				DevGroupIds: devGroupIds,
				CreateAt:    pointy.GetUnixMilliPointer(v.CreateAt.UnixMilli()),
				UpdateAt:    pointy.GetUnixMilliPointer(v.UpdateAt.UnixMilli()),
			})
	}

	return resp, nil
}
