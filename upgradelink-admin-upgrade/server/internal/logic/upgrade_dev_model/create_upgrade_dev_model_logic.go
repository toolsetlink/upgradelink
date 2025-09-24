package upgrade_dev_model

import (
	"context"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradedevmodel"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeDevModelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewCreateUpgradeDevModelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeDevModelLogic {
	return &CreateUpgradeDevModelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeDevModelLogic) CreateUpgradeDevModel(req *types.UpgradeDevModelInfo) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckCreateUpgradeDevModel(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	isDel := int32(0)
	_, err = l.svcCtx.DB.UpgradeDevModel.Create().
		SetNotNilCompanyID(&companyID).
		SetNotNilKey(req.Key).
		SetNotNilName(req.Name).
		SetNotNilIsDel(&isDel).
		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.CreateSuccess}, nil
}

func (l *CreateUpgradeDevModelLogic) CheckCreateUpgradeDevModel(req *types.UpgradeDevModelInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeDevModel
	predicates = append(predicates, upgradedevmodel.Name(*req.Name))
	predicates = append(predicates, upgradedevmodel.IsDelEQ(0))
	predicates = append(predicates, upgradedevmodel.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeDevModel.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("设备机型名称重复")
	}

	return nil
}
