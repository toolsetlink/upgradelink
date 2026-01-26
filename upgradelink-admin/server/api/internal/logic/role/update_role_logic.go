package role

import (
	"context"
	"fmt"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/entx"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.RoleInfo) (resp *types.BaseMsgResp, err error) {
	err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		origin, err := tx.Role.Get(l.ctx, *req.Id)
		if err != nil {
			return err
		}

		err = tx.Role.UpdateOneID(*req.Id).
			SetNotNilStatus(pointy.GetStatusPointer(req.Status)).
			SetNotNilName(req.Name).
			SetNotNilCode(req.Code).
			SetNotNilDefaultRouter(req.DefaultRouter).
			SetNotNilRemark(req.Remark).
			SetNotNilSort(req.Sort).
			Exec(l.ctx)
		if err != nil {
			return err
		}

		if req.Code != nil && origin.Code != *req.Code {
			_, err = tx.QueryContext(l.ctx, fmt.Sprintf("update casbin_rules set v0='%s' WHERE v0='%s'", *req.Code, origin.Code))
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
