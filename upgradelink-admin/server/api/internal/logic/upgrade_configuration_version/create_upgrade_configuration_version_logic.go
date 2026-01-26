package upgrade_configuration_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradeconfigurationversion"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeConfigurationVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUpgradeConfigurationVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeConfigurationVersionLogic {
	return &CreateUpgradeConfigurationVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeConfigurationVersionLogic) CreateUpgradeConfigurationVersion(req *types.UpgradeConfigurationVersionInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckCreateUpgradeConfigurationVersion(req)
	if err != nil {
		return nil, err
	}

	isDel := int32(0)
	_, err = l.svcCtx.DB.UpgradeConfigurationVersion.Create().
		SetNotNilCompanyID(companyctx.GetCompanyIDPointerFromCtx(l.ctx)).
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
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}

func (l *CreateUpgradeConfigurationVersionLogic) CheckCreateUpgradeConfigurationVersion(req *types.UpgradeConfigurationVersionInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeConfigurationVersion
	predicates = append(predicates, upgradeconfigurationversion.ConfigurationID(*req.ConfigurationId))
	predicates = append(predicates, upgradeconfigurationversion.VersionName(*req.VersionName))
	predicates = append(predicates, upgradeconfigurationversion.IsDelEQ(0))
	predicates = append(predicates, upgradeconfigurationversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeConfigurationVersion.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError("当前应用版本名重复")
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeConfigurationVersion
	predicates1 = append(predicates1, upgradeconfigurationversion.ConfigurationID(*req.ConfigurationId))
	predicates1 = append(predicates1, upgradeconfigurationversion.VersionCode(*req.VersionCode))
	predicates1 = append(predicates1, upgradeconfigurationversion.IsDelEQ(0))
	predicates1 = append(predicates1, upgradeconfigurationversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count1, err := l.svcCtx.DB.UpgradeConfigurationVersion.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return http_error.NewCodeBadRequestError("当前应用版本号重复")
	}

	return nil
}
