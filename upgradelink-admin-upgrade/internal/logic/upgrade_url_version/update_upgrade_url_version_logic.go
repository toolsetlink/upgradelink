package upgrade_url_version

import (
	"context"

	"upgradelink-admin-upgrade/ent"
	"upgradelink-admin-upgrade/ent/predicate"
	"upgradelink-admin-upgrade/ent/upgradeurlversion"
	"upgradelink-admin-upgrade/internal/logic/base"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeUrlVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewUpdateUpgradeUrlVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeUrlVersionLogic {
	return &UpdateUpgradeUrlVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeUrlVersionLogic) UpdateUpgradeUrlVersion(req *types.UpgradeUrlVersionInfo) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckUpdateUpgradeUrlVersion(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	// 效验下请求数据是否属于当前操作者
	data, err := l.svcCtx.DB.UpgradeUrlVersion.Get(l.ctx, *req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	if data.CompanyID != companyID {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
	}

	reqModel := ent.UpgradeUrlVersion{
		ID:        *req.Id,
		CompanyID: companyID,
	}
	err = l.svcCtx.DB.UpgradeUrlVersion.UpdateOne(&reqModel).
		SetNotNilURLID(req.UrlId).
		SetNotNilURLPath(req.UrlPath).
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

func (l *UpdateUpgradeUrlVersionLogic) CheckUpdateUpgradeUrlVersion(req *types.UpgradeUrlVersionInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeUrlVersion
	predicates = append(predicates, upgradeurlversion.URLID(*req.UrlId))
	predicates = append(predicates, upgradeurlversion.IDNEQ(*req.Id))
	predicates = append(predicates, upgradeurlversion.VersionName(*req.VersionName))
	predicates = append(predicates, upgradeurlversion.IsDelEQ(0))
	predicates = append(predicates, upgradeurlversion.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeUrlVersion.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("当前应用版本名重复")
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeUrlVersion
	predicates1 = append(predicates1, upgradeurlversion.URLID(*req.UrlId))
	predicates1 = append(predicates1, upgradeurlversion.IDNEQ(*req.Id))
	predicates1 = append(predicates1, upgradeurlversion.VersionCode(*req.VersionCode))
	predicates1 = append(predicates1, upgradeurlversion.IsDelEQ(0))
	predicates1 = append(predicates1, upgradeurlversion.CompanyIDEQ(l.companyID))
	count1, err := l.svcCtx.DB.UpgradeUrlVersion.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return errorx.NewCodeInvalidArgumentError("当前应用版本号重复")
	}

	return nil
}
