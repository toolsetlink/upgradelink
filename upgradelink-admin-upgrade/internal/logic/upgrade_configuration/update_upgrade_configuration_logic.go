package upgrade_configuration

import (
	"context"

	"upgradelink-admin-upgrade/ent"
	"upgradelink-admin-upgrade/ent/predicate"
	"upgradelink-admin-upgrade/ent/upgradeconfiguration"
	"upgradelink-admin-upgrade/internal/logic/base"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeConfigurationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewUpdateUpgradeConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeConfigurationLogic {
	return &UpdateUpgradeConfigurationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeConfigurationLogic) UpdateUpgradeConfiguration(req *types.UpgradeConfigurationInfo) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckUpdateUpgradeConfiguration(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	// 效验下请求数据是否属于当前操作者
	data, err := l.svcCtx.DB.UpgradeConfiguration.Get(l.ctx, *req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	if data.CompanyID != companyID {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
	}

	reqModel := ent.UpgradeConfiguration{
		ID:        *req.Id,
		CompanyID: companyID,
	}
	err = l.svcCtx.DB.UpgradeConfiguration.UpdateOne(&reqModel).
		SetNotNilKey(req.Key).
		SetNotNilName(req.Name).
		SetNotNilDescription(req.Description).
		SetNotNilIsDel(req.IsDel).
		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.UpdateSuccess}, nil
}

func (l *UpdateUpgradeConfigurationLogic) CheckUpdateUpgradeConfiguration(req *types.UpgradeConfigurationInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeConfiguration
	predicates = append(predicates, upgradeconfiguration.IDNEQ(*req.Id))
	predicates = append(predicates, upgradeconfiguration.Name(*req.Name))
	predicates = append(predicates, upgradeconfiguration.IsDelEQ(0))
	predicates = append(predicates, upgradeconfiguration.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeConfiguration.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("配置名称重复")
	}

	return nil
}
