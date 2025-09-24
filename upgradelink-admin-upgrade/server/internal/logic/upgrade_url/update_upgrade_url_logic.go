package upgrade_url

import (
	"context"

	"upgradelink-admin-upgrade/server/ent"
	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradeurl"
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

type UpdateUpgradeUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewUpdateUpgradeUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeUrlLogic {
	return &UpdateUpgradeUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeUrlLogic) UpdateUpgradeUrl(req *types.UpgradeUrlInfo) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckUpdateUpgradeUrl(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	// 效验下请求数据是否属于当前操作者
	data, err := l.svcCtx.DB.UpgradeUrl.Get(l.ctx, *req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	if data.CompanyID != companyID {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
	}

	reqModel := ent.UpgradeUrl{
		ID:        *req.Id,
		CompanyID: companyID,
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
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.UpdateSuccess}, nil
}

func (l *UpdateUpgradeUrlLogic) CheckUpdateUpgradeUrl(req *types.UpgradeUrlInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeUrl
	predicates = append(predicates, upgradeurl.IDNEQ(*req.Id))
	predicates = append(predicates, upgradeurl.Name(*req.Name))
	predicates = append(predicates, upgradeurl.IsDelEQ(0))
	predicates = append(predicates, upgradeurl.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeUrl.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("应用名称重复")
	}

	return nil
}
