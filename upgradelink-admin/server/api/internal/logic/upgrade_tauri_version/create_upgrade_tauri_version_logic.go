package upgrade_tauri_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradetauriversion"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeTauriVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUpgradeTauriVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeTauriVersionLogic {
	return &CreateUpgradeTauriVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeTauriVersionLogic) CreateUpgradeTauriVersion(req *types.UpgradeTauriVersionInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckCreateUpgradeTauriVersion(req)
	if err != nil {
		return nil, err
	}

	// versionName转换为versionCode
	versionCode, err := common.SemVerToInt(*req.VersionName)
	if err != nil {
		return nil, http_error.NewCodeBadRequestError("版本名称格式错误")
	}

	// 判断安装包文件是否传了,如果没有传的话，把升级包信息赋值给安装包信息
	if req.InstallCloudFileId == nil {
		req.InstallCloudFileId = req.CloudFileId
	}

	isDel := int32(0)
	_, err = l.svcCtx.DB.UpgradeTauriVersion.Create().
		SetNotNilCompanyID(companyctx.GetCompanyIDPointerFromCtx(l.ctx)).
		SetNotNilTauriID(req.TauriId).
		SetNotNilCloudFileID(req.CloudFileId).
		SetNotNilInstallCloudFileID(req.InstallCloudFileId).
		SetNotNilVersionName(req.VersionName).
		SetNotNilVersionCode(&versionCode).
		SetNotNilTarget(req.Target).
		SetNotNilArch(req.Arch).
		SetNotNilSignature(req.Signature).
		SetNotNilDescription(req.Description).
		SetNotNilIsDel(&isDel).
		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
		Save(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}

func (l *CreateUpgradeTauriVersionLogic) CheckCreateUpgradeTauriVersion(req *types.UpgradeTauriVersionInfo) error {

	// versionName转换为versionCode
	versionCode, err := common.SemVerToInt(*req.VersionName)
	if err != nil {
		return http_error.NewCodeBadRequestError("版本名称格式错误")
	}

	// 判断是否重复
	var predicates []predicate.UpgradeTauriVersion
	predicates = append(predicates, upgradetauriversion.TauriID(*req.TauriId))
	predicates = append(predicates, upgradetauriversion.VersionName(*req.VersionName))
	predicates = append(predicates, upgradetauriversion.Target(*req.Target))
	predicates = append(predicates, upgradetauriversion.Arch(*req.Arch))
	predicates = append(predicates, upgradetauriversion.IsDelEQ(0))
	predicates = append(predicates, upgradetauriversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeTauriVersion.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError("当前应用版本名重复")
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeTauriVersion
	predicates1 = append(predicates1, upgradetauriversion.TauriID(*req.TauriId))
	predicates1 = append(predicates1, upgradetauriversion.VersionCode(versionCode))
	predicates1 = append(predicates1, upgradetauriversion.Target(*req.Target))
	predicates1 = append(predicates1, upgradetauriversion.Arch(*req.Arch))
	predicates1 = append(predicates1, upgradetauriversion.IsDelEQ(0))
	predicates1 = append(predicates1, upgradetauriversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count1, err := l.svcCtx.DB.UpgradeTauriVersion.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return http_error.NewCodeBadRequestError("当前应用版本号重复")
	}

	return nil
}
