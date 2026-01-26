package upgrade_apk

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradeapk"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeApkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUpgradeApkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeApkLogic {
	return &UpdateUpgradeApkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeApkLogic) UpdateUpgradeApk(req *types.UpgradeApkInfo) (*types.BaseMsgResp, error) {
	// 效验请求参数数据
	err := l.CheckUpdateUpgradeApk(req)
	if err != nil {
		return nil, err
	}

	reqModel := ent.UpgradeApk{
		ID: *req.Id,
	}
	err = l.svcCtx.DB.UpgradeApk.UpdateOne(&reqModel).
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

func (l *UpdateUpgradeApkLogic) CheckUpdateUpgradeApk(req *types.UpgradeApkInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeApk
	predicates = append(predicates, upgradeapk.IDNEQ(*req.Id))
	predicates = append(predicates, upgradeapk.Name(*req.Name))
	predicates = append(predicates, upgradeapk.IsDelEQ(0))
	predicates = append(predicates, upgradeapk.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeApk.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError("应用名称重复")
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeApk
	predicates1 = append(predicates1, upgradeapk.IDNEQ(*req.Id))
	predicates1 = append(predicates1, upgradeapk.PackageName(*req.PackageName))
	predicates1 = append(predicates1, upgradeapk.IsDelEQ(0))
	predicates1 = append(predicates1, upgradeapk.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count1, err := l.svcCtx.DB.UpgradeApk.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return http_error.NewCodeBadRequestError("应用包名名称重复")
	}

	return nil
}
