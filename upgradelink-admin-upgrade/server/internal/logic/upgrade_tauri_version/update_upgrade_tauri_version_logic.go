package upgrade_tauri_version

import (
	"context"
	"fmt"

	"upgradelink-admin-upgrade/server/ent"
	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradetauriversion"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeTauriVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewUpdateUpgradeTauriVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeTauriVersionLogic {
	return &UpdateUpgradeTauriVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeTauriVersionLogic) UpdateUpgradeTauriVersion(req *types.UpgradeTauriVersionInfo) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckUpdateUpgradeTauriVersion(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	// versionName转换为versionCode
	versionCode, err := utils.SemVerToInt(*req.VersionName)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errorx.NewCodeInvalidArgumentError("版本名称格式错误")
	}

	// 效验下请求数据是否属于当前操作者
	data, err := l.svcCtx.DB.UpgradeTauriVersion.Get(l.ctx, *req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	if data.CompanyID != companyID {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
	}

	// 判断安装包文件是否传了,如果没有传的话，把升级包信息赋值给安装包信息
	if req.InstallCloudFileId == nil {
		req.InstallCloudFileId = req.CloudFileId
	}

	reqModel := ent.UpgradeTauriVersion{
		ID:        *req.Id,
		CompanyID: companyID,
	}
	err = l.svcCtx.DB.UpgradeTauriVersion.UpdateOne(&reqModel).
		SetNotNilTauriID(req.TauriId).
		SetNotNilCloudFileID(req.CloudFileId).
		SetNotNilInstallCloudFileID(req.InstallCloudFileId).
		SetNotNilVersionName(req.VersionName).
		SetNotNilVersionCode(&versionCode).
		SetNotNilTarget(req.Target).
		SetNotNilArch(req.Arch).
		SetNotNilSignature(req.Signature).
		SetNotNilDescription(req.Description).
		SetNotNilIsDel(req.IsDel).
		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.UpdateSuccess}, nil
}

func (l *UpdateUpgradeTauriVersionLogic) CheckUpdateUpgradeTauriVersion(req *types.UpgradeTauriVersionInfo) error {
	// versionName转换为versionCode
	versionCode, err := utils.SemVerToInt(*req.VersionName)
	if err != nil {
		return errorx.NewCodeInvalidArgumentError("版本名称格式错误")
	}

	// 判断是否重复
	var predicates []predicate.UpgradeTauriVersion
	predicates = append(predicates, upgradetauriversion.TauriID(*req.TauriId))
	predicates = append(predicates, upgradetauriversion.IDNEQ(*req.Id))
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
	predicates1 = append(predicates1, upgradetauriversion.IDNEQ(*req.Id))
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
