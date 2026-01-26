package upgrade_dev

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/entx"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradedev"
	"upgradelink-admin/server/api/internal/ent/upgradedevgrouprelation"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeDevLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUpgradeDevLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeDevLogic {
	return &UpdateUpgradeDevLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeDevLogic) UpdateUpgradeDev(req *types.UpgradeDevInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckUpdateUpgradeDevReq(req)
	if err != nil {
		return nil, err
	}

	// 开启事务
	if err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {

		// 更新设备表数据
		reqModel := ent.UpgradeDev{
			ID: *req.Id,
		}
		err = l.svcCtx.DB.UpgradeDev.UpdateOne(&reqModel).
			SetNotNilKey(req.Key).
			SetNotNilIsDel(req.IsDel).
			SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
			SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
			Exec(l.ctx)
		if err != nil {
			return err
		}

		// 删除历史分组关联关系， 新增新的关联关系
		intDel := int32(1)
		err = l.svcCtx.DB.UpgradeDevGroupRelation.Update().
			Where(upgradedevgrouprelation.DevIDEQ(*req.Id)).
			SetNotNilIsDel(&intDel).
			Exec(l.ctx)
		if err != nil {
			return err
		}

		isDel := int32(0)
		// 判断是否有传递分组 id
		if req.DevGroupIds != nil && len(req.DevGroupIds) > 0 {
			// 插入分组表数据
			for _, groupId := range req.DevGroupIds {
				_, err = l.svcCtx.DB.UpgradeDevGroupRelation.Create().
					SetNotNilIsDel(&isDel).
					SetNotNilDevID(req.Id).
					SetDevGroupID(int(groupId)).
					Save(l.ctx)
				if err != nil {
					return err
				}
			}
		}

		return nil

	}); err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}

func (l *UpdateUpgradeDevLogic) CheckUpdateUpgradeDevReq(req *types.UpgradeDevInfo) error {
	// 判断唯一标识是否重复
	var predicates []predicate.UpgradeDev
	predicates = append(predicates, upgradedev.IDNEQ(*req.Id))
	predicates = append(predicates, upgradedev.KeyContains(*req.Key))
	predicates = append(predicates, upgradedev.IsDelEQ(0))
	predicates = append(predicates, upgradedev.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeDev.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError("设备唯一标识重复")
	}

	return nil
}
