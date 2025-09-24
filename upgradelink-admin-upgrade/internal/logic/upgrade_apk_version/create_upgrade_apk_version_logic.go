package upgrade_apk_version

import (
	"context"

	"upgradelink-admin-upgrade/ent/predicate"
	"upgradelink-admin-upgrade/ent/upgradeapkversion"
	"upgradelink-admin-upgrade/internal/logic/base"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeApkVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewCreateUpgradeApkVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeApkVersionLogic {
	return &CreateUpgradeApkVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeApkVersionLogic) CreateUpgradeApkVersion(req *types.UpgradeApkVersionInfo) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckCreateUpgradeApkVersion(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	isDel := int32(0)
	_, err = l.svcCtx.DB.UpgradeApkVersion.Create().
		SetNotNilCompanyID(&companyID).
		SetNotNilApkID(req.ApkId).
		SetNotNilCloudFileID(req.CloudFileId).
		SetNotNilVersionName(req.VersionName).
		SetNotNilVersionCode(req.VersionCode).
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

func (l *CreateUpgradeApkVersionLogic) CheckCreateUpgradeApkVersion(req *types.UpgradeApkVersionInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeApkVersion
	predicates = append(predicates, upgradeapkversion.ApkID(*req.ApkId))
	predicates = append(predicates, upgradeapkversion.VersionName(*req.VersionName))
	predicates = append(predicates, upgradeapkversion.IsDelEQ(0))
	predicates = append(predicates, upgradeapkversion.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeApkVersion.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("当前应用版本名重复")
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeApkVersion
	predicates1 = append(predicates1, upgradeapkversion.ApkID(*req.ApkId))
	predicates1 = append(predicates1, upgradeapkversion.VersionCode(*req.VersionCode))
	predicates1 = append(predicates1, upgradeapkversion.IsDelEQ(0))
	predicates1 = append(predicates1, upgradeapkversion.CompanyIDEQ(l.companyID))
	count1, err := l.svcCtx.DB.UpgradeApkVersion.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return errorx.NewCodeInvalidArgumentError("当前应用版本号重复")
	}

	return nil
}
