package upgrade_file_version

import (
	"context"

	"upgradelink-admin-upgrade/server/ent"
	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradefileversion"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeFileVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewUpdateUpgradeFileVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeFileVersionLogic {
	return &UpdateUpgradeFileVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeFileVersionLogic) UpdateUpgradeFileVersion(req *types.UpgradeFileVersionInfo) (*types.BaseMsgResp, error) {

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckUpdateUpgradeFileVersion(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	// 效验下请求数据是否属于当前操作者
	data, err := l.svcCtx.DB.UpgradeFileVersion.Get(l.ctx, *req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	if data.CompanyID != companyID {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
	}

	reqModel := ent.UpgradeFileVersion{
		ID:        *req.Id,
		CompanyID: companyID,
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
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.UpdateSuccess}, nil
}

func (l *UpdateUpgradeFileVersionLogic) CheckUpdateUpgradeFileVersion(req *types.UpgradeFileVersionInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeFileVersion
	predicates = append(predicates, upgradefileversion.IDNEQ(*req.Id))
	predicates = append(predicates, upgradefileversion.FileID(*req.FileId))
	predicates = append(predicates, upgradefileversion.VersionName(*req.VersionName))
	predicates = append(predicates, upgradefileversion.IsDelEQ(0))
	predicates = append(predicates, upgradefileversion.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeFileVersion.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("当前应用版本名重复")
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeFileVersion
	predicates1 = append(predicates1, upgradefileversion.IDNEQ(*req.Id))
	predicates1 = append(predicates1, upgradefileversion.FileID(*req.FileId))
	predicates1 = append(predicates1, upgradefileversion.VersionCode(*req.VersionCode))
	predicates1 = append(predicates1, upgradefileversion.IsDelEQ(0))
	predicates1 = append(predicates1, upgradefileversion.CompanyIDEQ(l.companyID))
	count1, err := l.svcCtx.DB.UpgradeFileVersion.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return errorx.NewCodeInvalidArgumentError("当前应用版本号重复")
	}

	return nil
}
