package upgrade_electron_version

import (
	"context"
	"fmt"
	"upgradelink-admin/server/api/internal/common"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradeelectronversion"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeElectronVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUpgradeElectronVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeElectronVersionLogic {
	return &UpdateUpgradeElectronVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeElectronVersionLogic) UpdateUpgradeElectronVersion(req *types.UpgradeElectronVersionInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckUpdateUpgradeElectronVersion(req)
	if err != nil {
		return nil, err
	}

	// versionName转换为versionCode
	versionCode, err := common.SemVerToInt(*req.VersionName)
	if err != nil {
		fmt.Println(err.Error())
		return nil, http_error.NewCodeBadRequestError("版本名称格式错误")
	}

	// 判断安装包文件是否传了,如果没有传的话，把升级包信息赋值给安装包信息
	if req.InstallCloudFileId == nil {
		req.InstallCloudFileId = req.CloudFileId
	}

	reqModel := ent.UpgradeElectronVersion{
		ID: *req.Id,
	}

	err = l.svcCtx.DB.UpgradeElectronVersion.UpdateOne(&reqModel).
		SetNotNilElectronID(req.ElectronId).
		SetNotNilCloudFileID(req.CloudFileId).
		SetNotNilInstallCloudFileID(req.InstallCloudFileId).
		SetNotNilVersionName(req.VersionName).
		SetNotNilVersionCode(&versionCode).
		SetNotNilArch(req.Arch).
		SetNotNilSha512(req.Sha512).
		SetNotNilInstallSha512(req.InstallSha512).
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

func (l *UpdateUpgradeElectronVersionLogic) CheckUpdateUpgradeElectronVersion(req *types.UpgradeElectronVersionInfo) error {
	// versionName转换为versionCode
	versionCode, err := common.SemVerToInt(*req.VersionName)
	if err != nil {
		return http_error.NewCodeBadRequestError("版本名称格式错误")
	}

	// 判断是否重复
	var predicates []predicate.UpgradeElectronVersion
	predicates = append(predicates, upgradeelectronversion.ElectronID(*req.ElectronId))
	predicates = append(predicates, upgradeelectronversion.IDNEQ(*req.Id))
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
		return http_error.NewCodeBadRequestError("当前应用版本名重复")
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeElectronVersion
	predicates1 = append(predicates1, upgradeelectronversion.ElectronID(*req.ElectronId))
	predicates1 = append(predicates1, upgradeelectronversion.IDNEQ(*req.Id))
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
		return http_error.NewCodeBadRequestError("当前应用版本号重复")
	}

	return nil
}
