package upgrade_configuration

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradeconfiguration"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeConfigurationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUpgradeConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeConfigurationLogic {
	return &UpdateUpgradeConfigurationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeConfigurationLogic) UpdateUpgradeConfiguration(req *types.UpgradeConfigurationInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckUpdateUpgradeConfiguration(req)
	if err != nil {
		return nil, err
	}

	reqModel := ent.UpgradeConfiguration{
		ID: *req.Id,
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
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}

func (l *UpdateUpgradeConfigurationLogic) CheckUpdateUpgradeConfiguration(req *types.UpgradeConfigurationInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeConfiguration
	predicates = append(predicates, upgradeconfiguration.IDNEQ(*req.Id))
	predicates = append(predicates, upgradeconfiguration.Name(*req.Name))
	predicates = append(predicates, upgradeconfiguration.IsDelEQ(0))
	predicates = append(predicates, upgradeconfiguration.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeConfiguration.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError("配置名称重复")
	}

	return nil
}
