package upgrade_configuration_version

import (
	"context"

	"upgradelink-admin-upgrade/ent"
	"upgradelink-admin-upgrade/ent/predicate"
	"upgradelink-admin-upgrade/ent/upgradeconfigurationversion"
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

type UpdateUpgradeConfigurationVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewUpdateUpgradeConfigurationVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeConfigurationVersionLogic {
	return &UpdateUpgradeConfigurationVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeConfigurationVersionLogic) UpdateUpgradeConfigurationVersion(req *types.UpgradeConfigurationVersionInfo) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckUpdateUpgradeConfigurationVersion(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	// 效验下请求数据是否属于当前操作者
	data, err := l.svcCtx.DB.UpgradeConfigurationVersion.Get(l.ctx, *req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	if data.CompanyID != companyID {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
	}

	reqModel := ent.UpgradeConfigurationVersion{
		ID:        *req.Id,
		CompanyID: companyID,
	}
	err = l.svcCtx.DB.UpgradeConfigurationVersion.UpdateOne(&reqModel).
		SetNotNilConfigurationID(req.ConfigurationId).
		SetNotNilContent(req.Content).
		SetNotNilVersionName(req.VersionName).
		SetNotNilVersionCode(req.VersionCode).
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

func (l *UpdateUpgradeConfigurationVersionLogic) CheckUpdateUpgradeConfigurationVersion(req *types.UpgradeConfigurationVersionInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeConfigurationVersion
	predicates = append(predicates, upgradeconfigurationversion.ConfigurationID(*req.ConfigurationId))
	predicates = append(predicates, upgradeconfigurationversion.IDNEQ(*req.Id))
	predicates = append(predicates, upgradeconfigurationversion.VersionName(*req.VersionName))
	predicates = append(predicates, upgradeconfigurationversion.IsDelEQ(0))
	predicates = append(predicates, upgradeconfigurationversion.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeConfigurationVersion.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("当前应用版本名重复")
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeConfigurationVersion
	predicates1 = append(predicates1, upgradeconfigurationversion.ConfigurationID(*req.ConfigurationId))
	predicates1 = append(predicates1, upgradeconfigurationversion.IDNEQ(*req.Id))
	predicates1 = append(predicates1, upgradeconfigurationversion.VersionCode(*req.VersionCode))
	predicates1 = append(predicates1, upgradeconfigurationversion.IsDelEQ(0))
	predicates1 = append(predicates1, upgradeconfigurationversion.CompanyIDEQ(l.companyID))
	count1, err := l.svcCtx.DB.UpgradeConfigurationVersion.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return errorx.NewCodeInvalidArgumentError("当前应用版本号重复")
	}

	return nil
}
