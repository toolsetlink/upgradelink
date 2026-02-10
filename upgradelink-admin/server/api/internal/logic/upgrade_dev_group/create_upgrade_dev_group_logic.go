package upgrade_dev_group

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/enum"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"
	"upgradelink-admin/server/api/internal/common/utils/entx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradedevgroup"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeDevGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUpgradeDevGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeDevGroupLogic {
	return &CreateUpgradeDevGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeDevGroupLogic) CreateUpgradeDevGroup(req *types.UpgradeDevGroupInfo) (*types.BaseMsgResp, error) {
	// 效验请求参数数据
	err := l.CheckCreateUpgradeDevGroup(req)
	if err != nil {
		return nil, err
	}

	// 开启事务
	if err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {

		intDelFalse := enum.IsDelFalse
		_, err = l.svcCtx.DB.UpgradeDevGroup.Create().
			SetNotNilCompanyID(companyctx.GetCompanyIDPointerFromCtx(l.ctx)).
			SetNotNilName(req.Name).
			SetNotNilIsDel(&intDelFalse).
			SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
			SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
			Save(l.ctx)

		if err != nil {
			return err
		}

		return nil

	}); err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}

func (l *CreateUpgradeDevGroupLogic) CheckCreateUpgradeDevGroup(req *types.UpgradeDevGroupInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeDevGroup
	predicates = append(predicates, upgradedevgroup.Name(*req.Name))
	predicates = append(predicates, upgradedevgroup.IsDelEQ(0))
	predicates = append(predicates, upgradedevgroup.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeDevGroup.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError(l.svcCtx.Trans.Trans(l.ctx, i18n.DeviceGroupNameDuplicate))
	}

	return nil
}
