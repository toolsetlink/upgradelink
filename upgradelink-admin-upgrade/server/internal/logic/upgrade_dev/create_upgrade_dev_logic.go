package upgrade_dev

import (
	"context"

	"upgradelink-admin-upgrade/server/ent"
	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradedev"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeDevLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	companyID int
}

func NewCreateUpgradeDevLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeDevLogic {
	return &CreateUpgradeDevLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeDevLogic) CreateUpgradeDev(req *types.UpgradeDevInfo) (*types.BaseMsgResp, error) {

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckCreateUpgradeDevReq(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	// 开启事务
	if err := base.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {

		isDel := int32(0)
		info, err := l.svcCtx.DB.UpgradeDev.Create().
			SetNotNilCompanyID(&companyID).
			SetNotNilKey(req.Key).
			SetNotNilIsDel(&isDel).
			SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
			SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
			Save(l.ctx)
		if err != nil {
			return err
		}

		// 判断是否有传递分组 id
		if req.DevGroupIds != nil && len(req.DevGroupIds) > 0 {
			// 插入分组表数据
			for _, groupId := range req.DevGroupIds {
				_, err = l.svcCtx.DB.UpgradeDevGroupRelation.Create().
					SetNotNilDevID(&info.ID).
					SetDevGroupID(int(groupId)).
					SetNotNilIsDel(&isDel).
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

	return &types.BaseMsgResp{Msg: errormsg.CreateSuccess}, nil
}

func (l *CreateUpgradeDevLogic) CheckCreateUpgradeDevReq(req *types.UpgradeDevInfo) error {
	// 判断设备唯一标识是否重复
	var predicates []predicate.UpgradeDev
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
