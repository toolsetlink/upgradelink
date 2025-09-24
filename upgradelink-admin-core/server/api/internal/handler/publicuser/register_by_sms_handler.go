package publicuser

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-core/server/api/internal/logic/publicuser"
	"upgradelink-admin-core/server/api/internal/svc"
	"upgradelink-admin-core/server/api/internal/types"
)

// swagger:route post /user/register_by_sms publicuser RegisterBySms
//
// Register by SMS | 短信注册
//
// Register by SMS | 短信注册
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: RegisterBySmsReq
//
// Responses:
//  200: BaseMsgResp

func RegisterBySmsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterBySmsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := publicuser.NewRegisterBySmsLogic(r.Context(), svcCtx)
		resp, err := l.RegisterBySms(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
