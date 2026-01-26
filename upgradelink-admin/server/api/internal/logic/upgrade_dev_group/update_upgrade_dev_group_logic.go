package upgrade_dev_group

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradedevgroup"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeDevGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUpgradeDevGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeDevGroupLogic {
	return &UpdateUpgradeDevGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeDevGroupLogic) UpdateUpgradeDevGroup(req *types.UpgradeDevGroupInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckUpdateUpgradeDevGroup(req)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.DB.UpgradeDevGroup.UpdateOneID(*req.Id).
		SetNotNilName(req.Name).
		SetNotNilIsDel(req.IsDel).
		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
		Exec(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}

func (l *UpdateUpgradeDevGroupLogic) CheckUpdateUpgradeDevGroup(req *types.UpgradeDevGroupInfo) error {
	// 判断唯一标识是否重复
	var predicates []predicate.UpgradeDevGroup
	predicates = append(predicates, upgradedevgroup.IDNEQ(*req.Id))
	predicates = append(predicates, upgradedevgroup.Name(*req.Name))
	predicates = append(predicates, upgradedevgroup.IsDelEQ(0))
	predicates = append(predicates, upgradedevgroup.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeDevGroup.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError("设备唯一标识重复")
	}

	return nil
}
