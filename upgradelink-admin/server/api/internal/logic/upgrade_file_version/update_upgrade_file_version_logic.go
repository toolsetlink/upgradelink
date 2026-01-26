package upgrade_file_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradefileversion"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeFileVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUpgradeFileVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeFileVersionLogic {
	return &UpdateUpgradeFileVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeFileVersionLogic) UpdateUpgradeFileVersion(req *types.UpgradeFileVersionInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckUpdateUpgradeFileVersion(req)
	if err != nil {
		return nil, err
	}

	reqModel := ent.UpgradeFileVersion{
		ID: *req.Id,
	}
	err = l.svcCtx.DB.UpgradeFileVersion.UpdateOne(&reqModel).
		SetNotNilFileID(req.FileId).
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

func (l *UpdateUpgradeFileVersionLogic) CheckUpdateUpgradeFileVersion(req *types.UpgradeFileVersionInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeFileVersion
	predicates = append(predicates, upgradefileversion.IDNEQ(*req.Id))
	predicates = append(predicates, upgradefileversion.FileID(*req.FileId))
	predicates = append(predicates, upgradefileversion.VersionName(*req.VersionName))
	predicates = append(predicates, upgradefileversion.IsDelEQ(0))
	predicates = append(predicates, upgradefileversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeFileVersion.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError("当前应用版本名重复")
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeFileVersion
	predicates1 = append(predicates1, upgradefileversion.IDNEQ(*req.Id))
	predicates1 = append(predicates1, upgradefileversion.FileID(*req.FileId))
	predicates1 = append(predicates1, upgradefileversion.VersionCode(*req.VersionCode))
	predicates1 = append(predicates1, upgradefileversion.IsDelEQ(0))
	predicates1 = append(predicates1, upgradefileversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count1, err := l.svcCtx.DB.UpgradeFileVersion.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return http_error.NewCodeBadRequestError("当前应用版本号重复")
	}

	return nil
}
