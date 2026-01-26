package user

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/common/utils/uuidx"
	"upgradelink-admin/server/api/internal/ent/user"

	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(req *types.UUIDReq) (resp *types.UserInfoResp, err error) {
	result, err := l.svcCtx.DB.User.Query().Where(user.IDEQ(uuidx.ParseUUIDString(req.Id))).WithRoles().First(l.ctx)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.UserInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.UserInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        pointy.GetPointer(result.ID.String()),
				CreatedAt: pointy.GetPointer(result.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(result.UpdatedAt.UnixMilli()),
			},
			Status:      pointy.GetPointer(uint32(result.Status)),
			Username:    &result.Username,
			Nickname:    &result.Nickname,
			Description: &result.Description,
			HomePath:    &result.HomePath,
			RoleIds:     GetRoleIds(result.Edges.Roles),
			Mobile:      &result.Mobile,
			Email:       &result.Email,
			Avatar:      &result.Avatar,
		},
	}, nil
}
