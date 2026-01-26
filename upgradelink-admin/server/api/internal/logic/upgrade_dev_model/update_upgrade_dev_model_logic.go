package upgrade_dev_model

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradedevmodel"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeDevModelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUpgradeDevModelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeDevModelLogic {
	return &UpdateUpgradeDevModelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeDevModelLogic) UpdateUpgradeDevModel(req *types.UpgradeDevModelInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckUpdateUpgradeDevModel(req)
	if err != nil {
		return nil, http_error.NewCodeBadRequestError(l.svcCtx.Trans.Trans(l.ctx, err.Error()))
	}

	reqModel := ent.UpgradeDevModel{
		ID: *req.Id,
	}
	err = l.svcCtx.DB.UpgradeDevModel.UpdateOne(&reqModel).
		SetNotNilKey(req.Key).
		SetNotNilName(req.Name).
		SetNotNilIsDel(req.IsDel).
		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
		Exec(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}

func (l *UpdateUpgradeDevModelLogic) CheckUpdateUpgradeDevModel(req *types.UpgradeDevModelInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeDevModel
	predicates = append(predicates, upgradedevmodel.IDNEQ(*req.Id))
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
