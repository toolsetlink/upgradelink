package upgrade_tauri

import (
	"context"

	"upgradelink-admin-upgrade/server/ent/upgradetauri"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeTauriLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeTauriLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeTauriLogic {
	return &DeleteUpgradeTauriLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeTauriLogic) DeleteUpgradeTauri(req *types.IDsReq) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}

	// 效验下请求数据是否属于当前操作者
	for _, id := range req.Ids {
		data, err := l.svcCtx.DB.UpgradeTauri.Get(l.ctx, id)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}
		if data.CompanyID != companyID {
			return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
		}
	}

	intDel := int32(1)
	err = l.svcCtx.DB.UpgradeTauri.Update().
		Where(upgradetauri.IDIn(req.Ids...), upgradetauri.CompanyIDIn(companyID)).
		SetNotNilIsDel(&intDel).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.DeleteSuccess}, nil
}
