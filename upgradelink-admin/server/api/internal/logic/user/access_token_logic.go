package user

import (
	"context"

	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccessTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccessTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccessTokenLogic {
	return &AccessTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccessTokenLogic) AccessToken() (resp *types.RefreshTokenResp, err error) {
	// todo: add your logic here and delete this line

	return
}
