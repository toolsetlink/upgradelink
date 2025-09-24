package upgrade_dev_group

import (
	"context"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradedevgroup"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeDevGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewUpdateUpgradeDevGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeDevGroupLogic {
	return &UpdateUpgradeDevGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeDevGroupLogic) UpdateUpgradeDevGroup(req *types.UpgradeDevGroupInfo) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckUpdateUpgradeDevGroup(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	// 效验下请求数据是否属于当前操作者
	data, err := l.svcCtx.DB.UpgradeDevGroup.Get(l.ctx, *req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	if data.CompanyID != companyID {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
	}

	err = l.svcCtx.DB.UpgradeDevGroup.UpdateOneID(*req.Id).
		SetNotNilName(req.Name).
		SetNotNilIsDel(req.IsDel).
		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.UpdateSuccess}, nil
}

func (l *UpdateUpgradeDevGroupLogic) CheckUpdateUpgradeDevGroup(req *types.UpgradeDevGroupInfo) error {
	// 判断唯一标识是否重复
	var predicates []predicate.UpgradeDevGroup
	predicates = append(predicates, upgradedevgroup.IDNEQ(*req.Id))
	predicates = append(predicates, upgradedevgroup.Name(*req.Name))
	predicates = append(predicates, upgradedevgroup.IsDelEQ(0))
	predicates = append(predicates, upgradedevgroup.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeDevGroup.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("设备唯一标识重复")
	}

	return nil
}
