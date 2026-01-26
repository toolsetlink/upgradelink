package upgrade_url

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradeurl"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUpgradeUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeUrlLogic {
	return &UpdateUpgradeUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeUrlLogic) UpdateUpgradeUrl(req *types.UpgradeUrlInfo) (*types.BaseMsgResp, error) {
	// 效验请求参数数据
	err := l.CheckUpdateUpgradeUrl(req)
	if err != nil {
		return nil, err
	}

	reqModel := ent.UpgradeUrl{
		ID: *req.Id,
	}
	err = l.svcCtx.DB.UpgradeUrl.UpdateOne(&reqModel).
		SetNotNilKey(req.Key).
		SetNotNilName(req.Name).
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

func (l *UpdateUpgradeUrlLogic) CheckUpdateUpgradeUrl(req *types.UpgradeUrlInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeUrl
	predicates = append(predicates, upgradeurl.IDNEQ(*req.Id))
	predicates = append(predicates, upgradeurl.Name(*req.Name))
	predicates = append(predicates, upgradeurl.IsDelEQ(0))
	predicates = append(predicates, upgradeurl.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeUrl.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError("应用名称重复")
	}

	return nil
}
