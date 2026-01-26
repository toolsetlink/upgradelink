package upgrade_apk_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradeapkversion"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeApkVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUpgradeApkVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeApkVersionLogic {
	return &UpdateUpgradeApkVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeApkVersionLogic) UpdateUpgradeApkVersion(req *types.UpgradeApkVersionInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckUpdateUpgradeApkVersion(req)
	if err != nil {
		return nil, err
	}

	reqModel := ent.UpgradeApkVersion{
		ID: *req.Id,
	}
	err = l.svcCtx.DB.UpgradeApkVersion.UpdateOne(&reqModel).
		SetNotNilApkID(req.ApkId).
		SetNotNilCloudFileID(req.CloudFileId).
		SetNotNilVersionName(req.VersionName).
		SetNotNilVersionCode(req.VersionCode).
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

func (l *UpdateUpgradeApkVersionLogic) CheckUpdateUpgradeApkVersion(req *types.UpgradeApkVersionInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeApkVersion
	predicates = append(predicates, upgradeapkversion.ApkID(*req.ApkId))
	predicates = append(predicates, upgradeapkversion.IDNEQ(*req.Id))
	predicates = append(predicates, upgradeapkversion.VersionName(*req.VersionName))
	predicates = append(predicates, upgradeapkversion.IsDelEQ(0))
	predicates = append(predicates, upgradeapkversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeApkVersion.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError("当前应用版本名重复")
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeApkVersion
	predicates1 = append(predicates1, upgradeapkversion.ApkID(*req.ApkId))
	predicates1 = append(predicates1, upgradeapkversion.IDNEQ(*req.Id))
	predicates1 = append(predicates1, upgradeapkversion.VersionCode(*req.VersionCode))
	predicates1 = append(predicates1, upgradeapkversion.IsDelEQ(0))
	predicates1 = append(predicates1, upgradeapkversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count1, err := l.svcCtx.DB.UpgradeApkVersion.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return http_error.NewCodeBadRequestError("当前应用版本号重复")
	}

	return nil
}
