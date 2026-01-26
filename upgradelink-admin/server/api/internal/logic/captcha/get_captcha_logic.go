package captcha

import (
	"context"
	"net/http"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"

	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha() (resp *types.CaptchaResp, err error) {
	id, b64s, _, err := l.svcCtx.Captcha.Generate()
	if err != nil {
		logx.Errorw("fail to generate captcha", logx.Field("detail", err.Error()))
		return nil,
			http_error.NewApiError(http.StatusInternalServerError,
				l.svcCtx.Trans.Trans(l.ctx, "captcha.generateFailed"),
			)
	}

	resp = &types.CaptchaResp{
		BaseDataInfo: types.BaseDataInfo{
			Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.CaptchaInfo{
			CaptchaId: id,
			ImgPath:   b64s,
		},
	}
	return resp, nil
}
