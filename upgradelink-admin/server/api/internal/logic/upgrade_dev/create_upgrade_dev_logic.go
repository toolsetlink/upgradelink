package upgrade_dev

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
	"upgradelink-admin/server/api/internal/ent/upgradedev"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeDevLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUpgradeDevLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeDevLogic {
	return &CreateUpgradeDevLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeDevLogic) CreateUpgradeDev(req *types.UpgradeDevInfo) (*types.BaseMsgResp, error) {
	// 效验请求参数数据
	err := l.CheckCreateUpgradeDevReq(req)
	if err != nil {
		return nil, err
	}

	// 开启事务
	if err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {

		intDelFalse := enum.IsDelFalse
		info, err := l.svcCtx.DB.UpgradeDev.Create().
			SetNotNilCompanyID(companyctx.GetCompanyIDPointerFromCtx(l.ctx)).
			SetNotNilKey(req.Key).
			SetNotNilIsDel(&intDelFalse).
			SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
			SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
			Save(l.ctx)
		if err != nil {
			return err
		}

		// 判断是否有传递分组 id
		if req.DevGroupIds != nil && len(req.DevGroupIds) > 0 {
			// 插入分组表数据
			for _, groupId := range req.DevGroupIds {
				_, err = l.svcCtx.DB.UpgradeDevGroupRelation.Create().
					SetNotNilDevID(&info.ID).
					SetDevGroupID(int(groupId)).
					SetNotNilIsDel(&intDelFalse).
					Save(l.ctx)
				if err != nil {
					return err
				}
			}
		}

		return nil

	}); err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}

func (l *CreateUpgradeDevLogic) CheckCreateUpgradeDevReq(req *types.UpgradeDevInfo) error {
	// 判断设备唯一标识是否重复
	var predicates []predicate.UpgradeDev
	predicates = append(predicates, upgradedev.KeyContains(*req.Key))
	predicates = append(predicates, upgradedev.IsDelEQ(0))
	predicates = append(predicates, upgradedev.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeDev.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError(l.svcCtx.Trans.Trans(l.ctx, i18n.DeviceUniqueIdDuplicate))
	}

	return nil
}
