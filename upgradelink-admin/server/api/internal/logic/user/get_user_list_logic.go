package user

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/role"
	"upgradelink-admin/server/api/internal/ent/user"

	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.UserListReq) (resp *types.UserListResp, err error) {
	var predicates []predicate.User

	if req.Mobile != nil {
		predicates = append(predicates, user.MobileEQ(*req.Mobile))
	}

	if req.Username != nil {
		predicates = append(predicates, user.UsernameContains(*req.Username))
	}

	if req.Email != nil {
		predicates = append(predicates, user.EmailEQ(*req.Email))
	}

	if req.Nickname != nil {
		predicates = append(predicates, user.NicknameContains(*req.Nickname))
	}

	if req.RoleIds != nil {
		predicates = append(predicates, user.HasRolesWith(role.IDIn(req.RoleIds...)))
	}

	if req.Description != nil {
		predicates = append(predicates, user.DescriptionContains(*req.Description))
	}

	users, err := l.svcCtx.DB.User.Query().Where(predicates...).WithRoles().Page(l.ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	resp = &types.UserListResp{}
	for _, v := range users.List {
		resp.Data.Data = append(resp.Data.Data, types.UserInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        pointy.GetPointer(v.ID.String()),
				CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			},
			Username:    &v.Username,
			Nickname:    &v.Nickname,
			Mobile:      &v.Mobile,
			RoleIds:     GetRoleIds(v.Edges.Roles),
			Email:       &v.Email,
			Avatar:      &v.Avatar,
			Status:      pointy.GetPointer(uint32(v.Status)),
			Description: &v.Description,
			HomePath:    &v.HomePath,
		})
	}
	resp.Data.Total = users.PageDetails.Total
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	return resp, nil
}
