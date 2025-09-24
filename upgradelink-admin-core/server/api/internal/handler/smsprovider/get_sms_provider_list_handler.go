package smsprovider

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-core/server/api/internal/logic/smsprovider"
	"upgradelink-admin-core/server/api/internal/svc"
	"upgradelink-admin-core/server/api/internal/types"
)

// swagger:route post /sms_provider/list smsprovider GetSmsProviderList
//
// Get sms provider list | 获取短信配置列表
//
// Get sms provider list | 获取短信配置列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SmsProviderListReq
//
// Responses:
//  200: SmsProviderListResp

func GetSmsProviderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SmsProviderListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := smsprovider.NewGetSmsProviderListLogic(r.Context(), svcCtx)
		resp, err := l.GetSmsProviderList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
