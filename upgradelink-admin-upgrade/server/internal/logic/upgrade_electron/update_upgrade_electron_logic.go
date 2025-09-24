package upgrade_electron

import (
	"context"

	"upgradelink-admin-upgrade/server/ent"
	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradeelectron"
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

type UpdateUpgradeElectronLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewUpdateUpgradeElectronLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeElectronLogic {
	return &UpdateUpgradeElectronLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeElectronLogic) UpdateUpgradeElectron(req *types.UpgradeElectronInfo) (*types.BaseMsgResp, error) {

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckUpdateUpgrade(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	// 效验下请求数据是否属于当前操作者
	data, err := l.svcCtx.DB.UpgradeElectron.Get(l.ctx, *req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	if data.CompanyID != companyID {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
	}

	reqModel := ent.UpgradeElectron{
		ID:        *req.Id,
		CompanyID: companyID,
	}
	err = l.svcCtx.DB.UpgradeElectron.UpdateOne(&reqModel).
		SetNotNilKey(req.Key).
		SetNotNilName(req.Name).
		SetNotNilDescription(req.Description).
		SetNotNilGithubURL(req.GithubUrl).
		SetNotNilIsDel(req.IsDel).
		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.UpdateSuccess}, nil
}

func (l *UpdateUpgradeElectronLogic) CheckUpdateUpgrade(req *types.UpgradeElectronInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeElectron
	predicates = append(predicates, upgradeelectron.IDNEQ(*req.Id))
	predicates = append(predicates, upgradeelectron.Name(*req.Name))
	predicates = append(predicates, upgradeelectron.IsDelEQ(0))
	predicates = append(predicates, upgradeelectron.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeElectron.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("应用名称重复")
	}

	return nil
}
