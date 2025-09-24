package upgrade_dev

import (
	"context"

	"upgradelink-admin-upgrade/ent"
	"upgradelink-admin-upgrade/ent/predicate"
	"upgradelink-admin-upgrade/ent/upgradedev"
	"upgradelink-admin-upgrade/ent/upgradedevgrouprelation"
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

type UpdateUpgradeDevLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewUpdateUpgradeDevLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeDevLogic {
	return &UpdateUpgradeDevLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeDevLogic) UpdateUpgradeDev(req *types.UpgradeDevInfo) (*types.BaseMsgResp, error) {

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckUpdateUpgradeDevReq(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	// 效验下请求数据是否属于当前操作者
	data, err := l.svcCtx.DB.UpgradeDev.Get(l.ctx, *req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	if data.CompanyID != companyID {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
	}

	// 开启事务
	if err := base.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {

		// 更新设备表数据
		reqModel := ent.UpgradeDev{
			ID:        *req.Id,
			CompanyID: companyID,
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
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.UpdateSuccess}, nil
}

func (l *UpdateUpgradeDevLogic) CheckUpdateUpgradeDevReq(req *types.UpgradeDevInfo) error {
	// 判断唯一标识是否重复
	var predicates []predicate.UpgradeDev
	predicates = append(predicates, upgradedev.IDNEQ(*req.Id))
	predicates = append(predicates, upgradedev.KeyContains(*req.Key))
	predicates = append(predicates, upgradedev.IsDelEQ(0))
	predicates = append(predicates, upgradedev.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeDev.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("设备唯一标识重复")
	}

	return nil
}
