package user

import (
	"context"
	"fmt"

	"github.com/suyuan32/simple-admin-common/i18n"

	"upgradelink-admin-core/server/api/internal/svc"
	"upgradelink-admin-core/server/api/internal/types"
	"upgradelink-admin-core/server/rpc/types/core"

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

	user, err := l.svcCtx.CoreRpc.GetUserById(l.ctx,
		&core.UUIDReq{Id: l.ctx.Value("userId").(string)})
	if err != nil {
		fmt.Println("Err: ", err)
		return nil, err
	}

	return &types.UserBaseIDInfoResp{
		BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)},
		Data: types.UserBaseIDInfo{
			UUID:        user.Id,
			Username:    user.Username,
			Nickname:    user.Nickname,
			Avatar:      user.Avatar,
			HomePath:    user.HomePath,
			Description: user.Description,
			//DepartmentName: l.svcCtx.Trans.Trans(l.ctx, *user.DepartmentName),
			DepartmentName: "",
			RoleName:       TransRoleName(l.svcCtx, l.ctx, user.RoleName),
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
