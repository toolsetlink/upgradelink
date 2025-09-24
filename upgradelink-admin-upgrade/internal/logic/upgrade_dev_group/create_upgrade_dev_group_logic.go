package upgrade_dev_group

import (
	"context"

	"upgradelink-admin-upgrade/ent"
	"upgradelink-admin-upgrade/ent/predicate"
	"upgradelink-admin-upgrade/ent/upgradedevgroup"
	"upgradelink-admin-upgrade/internal/logic/base"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeDevGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewCreateUpgradeDevGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeDevGroupLogic {
	return &CreateUpgradeDevGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeDevGroupLogic) CreateUpgradeDevGroup(req *types.UpgradeDevGroupInfo) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckCreateUpgradeDevGroup(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	// 开启事务
	if err := base.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {

		isDel := int32(0)
		_, err = l.svcCtx.DB.UpgradeDevGroup.Create().
			SetNotNilCompanyID(&companyID).
			SetNotNilName(req.Name).
			SetNotNilIsDel(&isDel).
			SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
			SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
			Save(l.ctx)

		if err != nil {
			return err
		}

		return nil

	}); err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.CreateSuccess}, nil
}

func (l *CreateUpgradeDevGroupLogic) CheckCreateUpgradeDevGroup(req *types.UpgradeDevGroupInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeDevGroup
	predicates = append(predicates, upgradedevgroup.Name(*req.Name))
	predicates = append(predicates, upgradedevgroup.IsDelEQ(0))
	predicates = append(predicates, upgradedevgroup.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeDevGroup.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("设备分组名称重复")
	}

	return nil
}
