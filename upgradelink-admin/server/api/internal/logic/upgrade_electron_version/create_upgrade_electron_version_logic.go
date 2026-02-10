package upgrade_electron_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/enum"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradeelectronversion"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeElectronVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUpgradeElectronVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeElectronVersionLogic {
	return &CreateUpgradeElectronVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeElectronVersionLogic) CreateUpgradeElectronVersion(req *types.UpgradeElectronVersionInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckCreateUpgradeElectronVersion(req)
	if err != nil {
		return nil, err
	}

	// versionName转换为versionCode
	versionCode, err := common.SemVerToInt(*req.VersionName)
	if err != nil {
		return nil, http_error.NewCodeBadRequestError(l.svcCtx.Trans.Trans(l.ctx, i18n.VersionFormatError))
	}

	// 判断安装包文件是否传了,如果没有传的话，把升级包信息赋值给安装包信息
	if req.InstallCloudFileId == nil {
		req.InstallCloudFileId = req.CloudFileId
	}
	if req.InstallSha512 == nil {
		req.InstallSha512 = req.Sha512
	}

	intDelFalse := enum.IsDelFalse
	_, err = l.svcCtx.DB.UpgradeElectronVersion.Create().
		SetNotNilCompanyID(companyctx.GetCompanyIDPointerFromCtx(l.ctx)).
		SetNotNilElectronID(req.ElectronId).
		SetNotNilCloudFileID(req.CloudFileId).
		SetNotNilInstallCloudFileID(req.InstallCloudFileId).
		SetNotNilVersionName(req.VersionName).
		SetNotNilVersionCode(&versionCode).
		SetNotNilPlatform(req.Platform).
		SetNotNilArch(req.Arch).
		SetNotNilSha512(req.Sha512).
		SetNotNilInstallSha512(req.InstallSha512).
		SetNotNilDescription(req.Description).
		SetNotNilIsDel(&intDelFalse).
		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
		Save(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}

func (l *CreateUpgradeElectronVersionLogic) CheckCreateUpgradeElectronVersion(req *types.UpgradeElectronVersionInfo) error {

	// versionName转换为versionCode
	versionCode, err := common.SemVerToInt(*req.VersionName)
	if err != nil {
		return http_error.NewCodeBadRequestError(l.svcCtx.Trans.Trans(l.ctx, i18n.VersionFormatError))
	}

	// 判断是否重复
	var predicates []predicate.UpgradeElectronVersion
	predicates = append(predicates, upgradeelectronversion.ElectronID(*req.ElectronId))
	predicates = append(predicates, upgradeelectronversion.VersionName(*req.VersionName))
	predicates = append(predicates, upgradeelectronversion.Platform(*req.Platform))
	predicates = append(predicates, upgradeelectronversion.Arch(*req.Arch))
	predicates = append(predicates, upgradeelectronversion.IsDelEQ(0))
	predicates = append(predicates, upgradeelectronversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeElectronVersion.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError(l.svcCtx.Trans.Trans(l.ctx, i18n.AppVersionDuplicate))
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeElectronVersion
	predicates1 = append(predicates1, upgradeelectronversion.ElectronID(*req.ElectronId))
	predicates1 = append(predicates1, upgradeelectronversion.VersionCode(versionCode))
	predicates1 = append(predicates1, upgradeelectronversion.Platform(*req.Platform))
	predicates1 = append(predicates1, upgradeelectronversion.Arch(*req.Arch))
	predicates1 = append(predicates1, upgradeelectronversion.IsDelEQ(0))
	predicates1 = append(predicates1, upgradeelectronversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count1, err := l.svcCtx.DB.UpgradeElectronVersion.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return http_error.NewCodeBadRequestError(l.svcCtx.Trans.Trans(l.ctx, i18n.AppVersionCodeDuplicate))
	}

	return nil
}
