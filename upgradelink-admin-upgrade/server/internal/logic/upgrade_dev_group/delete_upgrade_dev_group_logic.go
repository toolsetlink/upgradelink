package upgrade_dev_group

import (
	"context"

	"upgradelink-admin-upgrade/server/ent"
	"upgradelink-admin-upgrade/server/ent/upgradedevgroup"
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

type DeleteUpgradeDevGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeDevGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeDevGroupLogic {
	return &DeleteUpgradeDevGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeDevGroupLogic) DeleteUpgradeDevGroup(req *types.IDsReq) (*types.BaseMsgResp, error) {

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}

	// 效验下请求数据是否属于当前操作者
	for _, id := range req.Ids {
		devData, err := l.svcCtx.DB.UpgradeDevGroup.Get(l.ctx, id)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}
		if devData.CompanyID != companyID {
			return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
		}
	}

	// 开启事务
	if err := base.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {

		intDel := int32(1)
		err = l.svcCtx.DB.UpgradeDevGroup.Update().
			Where(upgradedevgroup.IDIn(req.Ids...), upgradedevgroup.CompanyIDIn(companyID)).
			SetNotNilIsDel(&intDel).
			Exec(l.ctx)
		if err != nil {
			return err
		}

		// 删除关联关系
		err = l.svcCtx.DB.UpgradeDevGroupRelation.Update().
			Where(upgradedevgrouprelation.DevGroupIDIn(req.Ids...)).
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
