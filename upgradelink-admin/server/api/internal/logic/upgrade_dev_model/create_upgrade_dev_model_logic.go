package upgrade_dev_model

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradedevmodel"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeDevModelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUpgradeDevModelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeDevModelLogic {
	return &CreateUpgradeDevModelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeDevModelLogic) CreateUpgradeDevModel(req *types.UpgradeDevModelInfo) (*types.BaseMsgResp, error) {
	// 效验请求参数数据
	err := l.CheckCreateUpgradeDevModel(req)
	if err != nil {
		return nil, http_error.NewCodeBadRequestError(err.Error())
	}

	isDel := int32(0)
	_, err = l.svcCtx.DB.UpgradeDevModel.Create().
		SetNotNilCompanyID(companyctx.GetCompanyIDPointerFromCtx(l.ctx)).
		SetNotNilKey(req.Key).
		SetNotNilName(req.Name).
		SetNotNilIsDel(&isDel).
		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
		Save(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}

func (l *CreateUpgradeDevModelLogic) CheckCreateUpgradeDevModel(req *types.UpgradeDevModelInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeDevModel
	predicates = append(predicates, upgradedevmodel.Name(*req.Name))
	predicates = append(predicates, upgradedevmodel.IsDelEQ(0))
	predicates = append(predicates, upgradedevmodel.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeDevModel.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError("设备机型名称重复")
	}

	return nil
}
