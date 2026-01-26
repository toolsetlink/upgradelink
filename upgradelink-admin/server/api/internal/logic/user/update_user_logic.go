package user

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/encrypt"
	"upgradelink-admin/server/api/internal/common/utils/entx"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/common/utils/uuidx"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UserInfo) (resp *types.BaseMsgResp, err error) {
	err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		updateQuery := tx.User.UpdateOneID(uuidx.ParseUUIDString(*req.Id)).
			SetNotNilUsername(req.Username).
			SetNotNilNickname(req.Nickname).
			SetNotNilEmail(req.Email).
			SetNotNilMobile(req.Mobile).
			SetNotNilAvatar(req.Avatar).
			SetNotNilHomePath(req.HomePath).
			SetNotNilDescription(req.Description).
			SetNotNilStatus(pointy.GetStatusPointer(req.Status))

		if req.Password != nil {
			updateQuery = updateQuery.SetNotNilPassword(pointy.GetPointer(encrypt.BcryptEncrypt(*req.Password)))
		}

		if req.RoleIds != nil {
			err := tx.User.UpdateOneID(uuidx.ParseUUIDString(*req.Id)).ClearRoles().Exec(l.ctx)
			if err != nil {
				return err
			}

			updateQuery = updateQuery.AddRoleIDs(req.RoleIds...)
		}

		// todo 删除用户对应 token
		//if req.Password != nil || req.RoleIds != nil || req.PositionIds != nil || (req.Status != nil && uint8(*req.Status) != enum.StatusNormal) {
		//	_, err := token.NewBlockUserAllTokenLogic(l.ctx, l.svcCtx).BlockUserAllToken(&core.UUIDReq{Id: *in.Id})
		//	if err != nil {
		//		return err
		//	}
		//}

		return updateQuery.Exec(l.ctx)
	})
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)}, nil
}
