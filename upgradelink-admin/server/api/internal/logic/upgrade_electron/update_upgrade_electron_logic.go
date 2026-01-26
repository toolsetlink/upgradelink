package upgrade_electron

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradeelectron"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeElectronLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUpgradeElectronLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeElectronLogic {
	return &UpdateUpgradeElectronLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeElectronLogic) UpdateUpgradeElectron(req *types.UpgradeElectronInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckUpdateUpgrade(req)
	if err != nil {
		return nil, err
	}

	reqModel := ent.UpgradeElectron{
		ID: *req.Id,
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
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}

func (l *UpdateUpgradeElectronLogic) CheckUpdateUpgrade(req *types.UpgradeElectronInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeElectron
	predicates = append(predicates, upgradeelectron.IDNEQ(*req.Id))
	predicates = append(predicates, upgradeelectron.Name(*req.Name))
	predicates = append(predicates, upgradeelectron.IsDelEQ(0))
	predicates = append(predicates, upgradeelectron.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeElectron.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError("应用名称重复")
	}

	return nil
}
