package upgrade_apk

import (
	"context"

	"upgradelink-admin-upgrade/ent"
	"upgradelink-admin-upgrade/ent/predicate"
	"upgradelink-admin-upgrade/ent/upgradeapk"
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

type UpdateUpgradeApkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewUpdateUpgradeApkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeApkLogic {
	return &UpdateUpgradeApkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeApkLogic) UpdateUpgradeApk(req *types.UpgradeApkInfo) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckUpdateUpgradeApk(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	// 效验下请求数据是否属于当前操作者
	data, err := l.svcCtx.DB.UpgradeApk.Get(l.ctx, *req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	if data.CompanyID != companyID {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
	}

	reqModel := ent.UpgradeApk{
		ID:        *req.Id,
		CompanyID: companyID,
	}
	err = l.svcCtx.DB.UpgradeApk.UpdateOne(&reqModel).
		SetNotNilKey(req.Key).
		SetNotNilName(req.Name).
		SetNotNilPackageName(req.PackageName).
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

func (l *UpdateUpgradeApkLogic) CheckUpdateUpgradeApk(req *types.UpgradeApkInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeApk
	predicates = append(predicates, upgradeapk.IDNEQ(*req.Id))
	predicates = append(predicates, upgradeapk.Name(*req.Name))
	predicates = append(predicates, upgradeapk.IsDelEQ(0))
	predicates = append(predicates, upgradeapk.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeApk.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("应用名称重复")
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeApk
	predicates1 = append(predicates1, upgradeapk.PackageName(*req.PackageName))
	predicates1 = append(predicates1, upgradeapk.IsDelEQ(0))
	predicates1 = append(predicates1, upgradeapk.CompanyIDEQ(l.companyID))
	count1, err := l.svcCtx.DB.UpgradeApk.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return errorx.NewCodeInvalidArgumentError("应用包名名称重复")
	}

	return nil
}
