package upgrade_configuration

import (
	"context"

	"upgradelink-admin-upgrade/ent/upgradeconfiguration"
	"upgradelink-admin-upgrade/internal/logic/base"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeConfigurationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeConfigurationLogic {
	return &DeleteUpgradeConfigurationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeConfigurationLogic) DeleteUpgradeConfiguration(req *types.IDsReq) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}

	// 效验下请求数据是否属于当前操作者
	for _, id := range req.Ids {
		data, err := l.svcCtx.DB.UpgradeConfiguration.Get(l.ctx, id)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}
		if data.CompanyID != companyID {
			return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
		}
	}

	intDel := int32(1)
	err = l.svcCtx.DB.UpgradeConfiguration.Update().
		Where(upgradeconfiguration.IDIn(req.Ids...), upgradeconfiguration.CompanyIDIn(companyID)).
		SetNotNilIsDel(&intDel).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.DeleteSuccess}, nil
}
