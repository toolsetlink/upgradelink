package upgrade_configuration_version

import (
	"context"
	"fmt"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradeconfigurationversion"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeConfigurationVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewCreateUpgradeConfigurationVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeConfigurationVersionLogic {
	return &CreateUpgradeConfigurationVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeConfigurationVersionLogic) CreateUpgradeConfigurationVersion(req *types.UpgradeConfigurationVersionInfo) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckCreateUpgradeConfigurationVersion(req)
	if err != nil {
		fmt.Println(err)
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	isDel := int32(0)
	_, err = l.svcCtx.DB.UpgradeConfigurationVersion.Create().
		SetNotNilCompanyID(&companyID).
		SetNotNilConfigurationID(req.ConfigurationId).
		SetNotNilContent(req.Content).
		SetNotNilVersionName(req.VersionName).
		SetNotNilVersionCode(req.VersionCode).
		SetNotNilDescription(req.Description).
		SetNotNilIsDel(&isDel).
		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.CreateSuccess}, nil
}

func (l *CreateUpgradeConfigurationVersionLogic) CheckCreateUpgradeConfigurationVersion(req *types.UpgradeConfigurationVersionInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeConfigurationVersion
	predicates = append(predicates, upgradeconfigurationversion.ConfigurationID(*req.ConfigurationId))
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
