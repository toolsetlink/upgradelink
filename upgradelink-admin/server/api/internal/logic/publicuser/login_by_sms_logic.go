package publicuser

import (
	"context"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginBySmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginBySmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginBySmsLogic {
	return &LoginBySmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginBySmsLogic) LoginBySms(req *types.LoginBySmsReq) (resp *types.LoginResp, err error) {
	return
}
