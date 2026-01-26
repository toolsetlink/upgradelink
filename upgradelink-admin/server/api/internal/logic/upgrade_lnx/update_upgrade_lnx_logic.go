package upgrade_lnx

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradelnx"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeLnxLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUpgradeLnxLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeLnxLogic {
	return &UpdateUpgradeLnxLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeLnxLogic) UpdateUpgradeLnx(req *types.UpgradeLnxInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckUpdateUpgradeLnx(req)
	if err != nil {
		return nil, err
	}

	reqModel := ent.UpgradeLnx{
		ID: *req.Id,
	}
	err = l.svcCtx.DB.UpgradeLnx.UpdateOne(&reqModel).
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

func (l *UpdateUpgradeLnxLogic) CheckUpdateUpgradeLnx(req *types.UpgradeLnxInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeLnx
	predicates = append(predicates, upgradelnx.IDNEQ(*req.Id))
	predicates = append(predicates, upgradelnx.Name(*req.Name))
	predicates = append(predicates, upgradelnx.IsDelEQ(0))
	predicates = append(predicates, upgradelnx.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeLnx.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError("应用名称重复")
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeLnx
	predicates1 = append(predicates1, upgradelnx.PackageName(*req.PackageName))
	predicates1 = append(predicates1, upgradelnx.IsDelEQ(0))
	predicates1 = append(predicates1, upgradelnx.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count1, err := l.svcCtx.DB.UpgradeLnx.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return http_error.NewCodeBadRequestError("应用包名名称重复")
	}

	return nil
}
