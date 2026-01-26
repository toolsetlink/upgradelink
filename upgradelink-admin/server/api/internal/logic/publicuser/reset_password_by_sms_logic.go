package publicuser

import (
	"context"

	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordBySmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetPasswordBySmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordBySmsLogic {
	return &ResetPasswordBySmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPasswordBySmsLogic) ResetPasswordBySms(req *types.ResetPasswordBySmsReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
