package messagesender

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-core/server/api/internal/logic/messagesender"
	"upgradelink-admin-core/server/api/internal/svc"
	"upgradelink-admin-core/server/api/internal/types"
)

// swagger:route post /sms/send messagesender SendSms
//
// Send sms message | 发送短信
//
// Send sms message | 发送短信
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SendSmsReq
//
// Responses:
//  200: BaseMsgResp

func SendSmsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendSmsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := messagesender.NewSendSmsLogic(r.Context(), svcCtx)
		resp, err := l.SendSms(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
