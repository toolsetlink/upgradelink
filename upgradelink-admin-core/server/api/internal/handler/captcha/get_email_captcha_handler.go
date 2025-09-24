package captcha

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-core/server/api/internal/logic/captcha"
	"upgradelink-admin-core/server/api/internal/svc"
	"upgradelink-admin-core/server/api/internal/types"
)

// swagger:route post /captcha/email captcha GetEmailCaptcha
//
// Get Email Captcha | 获取邮箱验证码
//
// Get Email Captcha | 获取邮箱验证码
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: EmailCaptchaReq
//
// Responses:
//  200: BaseMsgResp

func GetEmailCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailCaptchaReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := captcha.NewGetEmailCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetEmailCaptcha(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
