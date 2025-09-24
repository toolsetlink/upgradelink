package upgrade_dev

import (
	"context"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradedevgrouprelation"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeDevByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeDevByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeDevByIdLogic {
	return &GetUpgradeDevByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeDevByIdLogic) GetUpgradeDevById(req *types.IDReq) (*types.UpgradeDevInfoResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}

	// 获取数据
	data, err := l.svcCtx.DB.UpgradeDev.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	// 效验下请求数据是否属于当前操作者
	if data.CompanyID != companyID {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
	}

	// 查询出关联的设备组
	var predicates []predicate.UpgradeDevGroupRelation
	predicates = append(predicates, upgradedevgrouprelation.DevID(req.Id))
	predicates = append(predicates, upgradedevgrouprelation.IsDelEQ(0))
	devGroupList, err := l.svcCtx.DB.UpgradeDevGroupRelation.Query().Where(predicates...).All(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	var devGroupIds []uint64
	for _, group := range devGroupList {
		devGroupIds = append(devGroupIds, uint64(group.DevGroupID))
	}

	return &types.UpgradeDevInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  errormsg.Success,
		},
		Data: types.UpgradeDevInfo{
			Id:          &data.ID,
			Key:         &data.Key,
			IsDel:       &data.IsDel,
			DevGroupIds: devGroupIds,
			CreateAt:    pointy.GetUnixMilliPointer(data.CreateAt.UnixMilli()),
			UpdateAt:    pointy.GetUnixMilliPointer(data.UpdateAt.UnixMilli()),
		},
	}, nil
}
