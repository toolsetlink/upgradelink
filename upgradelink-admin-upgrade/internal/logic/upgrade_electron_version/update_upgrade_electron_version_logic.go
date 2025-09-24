package upgrade_electron_version

import (
	"context"
	"fmt"

	"upgradelink-admin-upgrade/ent"
	"upgradelink-admin-upgrade/ent/predicate"
	"upgradelink-admin-upgrade/ent/upgradeelectronversion"
	"upgradelink-admin-upgrade/internal/logic/base"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeElectronVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewUpdateUpgradeElectronVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeElectronVersionLogic {
	return &UpdateUpgradeElectronVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeElectronVersionLogic) UpdateUpgradeElectronVersion(req *types.UpgradeElectronVersionInfo) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckUpdateUpgradeElectronVersion(req)
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
	data, err := l.svcCtx.DB.UpgradeElectronVersion.Get(l.ctx, *req.Id)
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
	
	reqModel := ent.UpgradeElectronVersion{
		ID:        *req.Id,
		CompanyID: companyID,
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
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.UpdateSuccess}, nil
}

func (l *UpdateUpgradeElectronVersionLogic) CheckUpdateUpgradeElectronVersion(req *types.UpgradeElectronVersionInfo) error {
	// versionName转换为versionCode
	versionCode, err := utils.SemVerToInt(*req.VersionName)
	if err != nil {
		return errorx.NewCodeInvalidArgumentError("版本名称格式错误")
	}

	// 判断是否重复
	var predicates []predicate.UpgradeElectronVersion
	predicates = append(predicates, upgradeelectronversion.ElectronID(*req.ElectronId))
	predicates = append(predicates, upgradeelectronversion.IDNEQ(*req.Id))
	predicates = append(predicates, upgradeelectronversion.VersionName(*req.VersionName))
	predicates = append(predicates, upgradeelectronversion.Arch(*req.Arch))
	predicates = append(predicates, upgradeelectronversion.IsDelEQ(0))
	predicates = append(predicates, upgradeelectronversion.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeElectronVersion.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("当前应用版本名重复")
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeElectronVersion
	predicates1 = append(predicates1, upgradeelectronversion.ElectronID(*req.ElectronId))
	predicates1 = append(predicates1, upgradeelectronversion.IDNEQ(*req.Id))
	predicates1 = append(predicates1, upgradeelectronversion.VersionCode(versionCode))
	predicates1 = append(predicates1, upgradeelectronversion.Arch(*req.Arch))
	predicates1 = append(predicates1, upgradeelectronversion.IsDelEQ(0))
	predicates1 = append(predicates1, upgradeelectronversion.CompanyIDEQ(l.companyID))
	count1, err := l.svcCtx.DB.UpgradeElectronVersion.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return errorx.NewCodeInvalidArgumentError("当前应用版本号重复")
	}

	return nil
}
