package upgrade_mac

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgrademac"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeMacLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUpgradeMacLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeMacLogic {
	return &UpdateUpgradeMacLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeMacLogic) UpdateUpgradeMac(req *types.UpgradeMacInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckUpdateUpgradeMac(req)
	if err != nil {
		return nil, err
	}

	reqModel := ent.UpgradeMac{
		ID: *req.Id,
	}
	err = l.svcCtx.DB.UpgradeMac.UpdateOne(&reqModel).
		SetNotNilKey(req.Key).
		SetNotNilName(req.Name).
		SetNotNilPackageName(req.PackageName).
		SetNotNilDescription(req.Description).
		SetNotNilIsDel(req.IsDel).
		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
		Exec(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}

func (l *UpdateUpgradeMacLogic) CheckUpdateUpgradeMac(req *types.UpgradeMacInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeMac
	predicates = append(predicates, upgrademac.IDNEQ(*req.Id))
	predicates = append(predicates, upgrademac.Name(*req.Name))
	predicates = append(predicates, upgrademac.IsDelEQ(0))
	predicates = append(predicates, upgrademac.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeMac.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError("应用名称重复")
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeMac
	predicates1 = append(predicates1, upgrademac.PackageName(*req.PackageName))
	predicates1 = append(predicates1, upgrademac.IsDelEQ(0))
	predicates1 = append(predicates1, upgrademac.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count1, err := l.svcCtx.DB.UpgradeMac.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return http_error.NewCodeBadRequestError("应用包名名称重复")
	}

	return nil
}
