package user

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/uuidx"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/user"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.UserBaseIDInfoResp, err error) {
	userId := ""
	userId = l.ctx.Value("userId").(string)
	result, err := l.svcCtx.DB.User.Query().Where(user.IDEQ(uuidx.ParseUUIDString(userId))).WithRoles().First(l.ctx)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, userId)
	}

	return &types.UserBaseIDInfoResp{
		BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)},
		Data: types.UserBaseIDInfo{
			UUID:        &userId,
			Username:    &result.Username,
			Nickname:    &result.Nickname,
			Avatar:      &result.Avatar,
			HomePath:    &result.HomePath,
			Description: &result.Description,
			RoleName:    GetRoleNames(result.Edges.Roles),
		},
	}, nil
}

// TransRoleName returns the i18n translation of role name slice.
func TransRoleName(svc *svc.ServiceContext, ctx context.Context, data []string) []string {
	var result []string
	for _, v := range data {
		result = append(result, svc.Trans.Trans(ctx, v))
	}
	return result
}

func GetRoleIds(data []*ent.Role) []uint64 {
	var ids []uint64
	for _, v := range data {
		ids = append(ids, v.ID)
	}
	return ids
}

func GetRoleNames(data []*ent.Role) []string {
	var codes []string
	for _, v := range data {
		codes = append(codes, v.Name)
	}
	return codes
}

func GetRoleCodes(data []*ent.Role) []string {
	var codes []string
	for _, v := range data {
		codes = append(codes, v.Code)
	}
	return codes
}
