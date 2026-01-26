package captcha

import (
	"context"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSmsCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsCaptchaLogic {
	return &GetSmsCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSmsCaptchaLogic) GetSmsCaptcha(req *types.SmsCaptchaReq) (resp *types.BaseMsgResp, err error) {
	return
}
