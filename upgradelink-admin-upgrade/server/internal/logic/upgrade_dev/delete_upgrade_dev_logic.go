package upgrade_dev

import (
	"context"

	"upgradelink-admin-upgrade/server/ent"
	"upgradelink-admin-upgrade/server/ent/upgradedev"
	"upgradelink-admin-upgrade/server/ent/upgradedevgrouprelation"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeDevLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeDevLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeDevLogic {
	return &DeleteUpgradeDevLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeDevLogic) DeleteUpgradeDev(req *types.IDsReq) (*types.BaseMsgResp, error) {

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}

	// 效验下请求数据是否属于当前操作者
	for _, id := range req.Ids {
		data, err := l.svcCtx.DB.UpgradeDev.Get(l.ctx, id)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}
		if data.CompanyID != companyID {
			return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
		}
	}

	// 开启事务
	if err := base.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {

		intDel := int32(1)
		err = l.svcCtx.DB.UpgradeDev.Update().
			Where(upgradedev.IDIn(req.Ids...), upgradedev.CompanyIDIn(companyID)).
			SetNotNilIsDel(&intDel).
			Exec(l.ctx)
		if err != nil {
			return err
		}

		// 删除关联关系
		err = l.svcCtx.DB.UpgradeDevGroupRelation.Update().
			Where(upgradedevgrouprelation.DevIDIn(req.Ids...)).
			SetNotNilIsDel(&intDel).
			Exec(l.ctx)
		if err != nil {
			return err
		}

		return nil

	}); err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.DeleteSuccess}, nil
}
