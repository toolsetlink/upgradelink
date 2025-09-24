package upgrade_tauri_version

import (
	"context"

	"upgradelink-admin-upgrade/ent/predicate"
	"upgradelink-admin-upgrade/ent/upgradetauriversion"
	"upgradelink-admin-upgrade/internal/logic/base"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeTauriVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewCreateUpgradeTauriVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeTauriVersionLogic {
	return &CreateUpgradeTauriVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeTauriVersionLogic) CreateUpgradeTauriVersion(req *types.UpgradeTauriVersionInfo) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckCreateUpgradeTauriVersion(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	// versionName转换为versionCode
	versionCode, err := utils.SemVerToInt(*req.VersionName)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError("版本名称格式错误")
	}

	// 判断安装包文件是否传了,如果没有传的话，把升级包信息赋值给安装包信息
	if req.InstallCloudFileId == nil {
		req.InstallCloudFileId = req.CloudFileId
	}

	isDel := int32(0)
	_, err = l.svcCtx.DB.UpgradeTauriVersion.Create().
		SetNotNilCompanyID(&companyID).
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
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.CreateSuccess}, nil
}

func (l *CreateUpgradeTauriVersionLogic) CheckCreateUpgradeTauriVersion(req *types.UpgradeTauriVersionInfo) error {

	// versionName转换为versionCode
	versionCode, err := utils.SemVerToInt(*req.VersionName)
	if err != nil {
		return errorx.NewCodeInvalidArgumentError("版本名称格式错误")
	}

	// 判断是否重复
	var predicates []predicate.UpgradeTauriVersion
	predicates = append(predicates, upgradetauriversion.TauriID(*req.TauriId))
	predicates = append(predicates, upgradetauriversion.VersionName(*req.VersionName))
	predicates = append(predicates, upgradetauriversion.Target(*req.Target))
	predicates = append(predicates, upgradetauriversion.Arch(*req.Arch))
	predicates = append(predicates, upgradetauriversion.IsDelEQ(0))
	predicates = append(predicates, upgradetauriversion.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeTauriVersion.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("当前应用版本名重复")
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeTauriVersion
	predicates1 = append(predicates1, upgradetauriversion.TauriID(*req.TauriId))
	predicates1 = append(predicates1, upgradetauriversion.VersionCode(versionCode))
	predicates1 = append(predicates1, upgradetauriversion.Target(*req.Target))
	predicates1 = append(predicates1, upgradetauriversion.Arch(*req.Arch))
	predicates1 = append(predicates1, upgradetauriversion.IsDelEQ(0))
	predicates1 = append(predicates1, upgradetauriversion.CompanyIDEQ(l.companyID))
	count1, err := l.svcCtx.DB.UpgradeTauriVersion.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return errorx.NewCodeInvalidArgumentError("当前应用版本号重复")
	}

	return nil
}
