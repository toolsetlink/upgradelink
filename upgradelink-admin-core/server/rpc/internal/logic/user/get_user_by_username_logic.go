package user

import (
	"context"

	"upgradelink-admin-core/server/rpc/ent/user"
	"upgradelink-admin-core/server/rpc/internal/svc"
	"upgradelink-admin-core/server/rpc/internal/utils/dberrorhandler"
	"upgradelink-admin-core/server/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByUsernameLogic {
	return &GetUserByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByUsernameLogic) GetUserByUsername(in *core.UsernameReq) (*core.UserInfo, error) {
	result, err := l.svcCtx.DB.User.Query().Where(user.UsernameEQ(in.Username)).WithRoles().First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.UserInfo{
		Nickname:     &result.Nickname,
		Avatar:       &result.Avatar,
		Password:     &result.Password,
		RoleIds:      GetRoleIds(result.Edges.Roles),
		RoleCodes:    GetRoleCodes(result.Edges.Roles),
		Mobile:       &result.Mobile,
		Email:        &result.Email,
		Status:       pointy.GetPointer(uint32(result.Status)),
		Id:           pointy.GetPointer(result.ID.String()),
		Username:     &result.Username,
		HomePath:     &result.HomePath,
		Description:  &result.Description,
		DepartmentId: &result.DepartmentID,
		CompanyId:    &result.CompanyID,
		CreatedAt:    pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:    pointy.GetPointer(result.UpdatedAt.UnixMilli()),
	}, nil
}
