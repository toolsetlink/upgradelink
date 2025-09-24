package company_secret

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-core/server/api/internal/logic/company_secret"
	"upgradelink-admin-core/server/api/internal/svc"
	"upgradelink-admin-core/server/api/internal/types"
)

// swagger:route post /company_secret/list company_secret GetCompanySecretList
//
// getCompanySecretList
//
// getCompanySecretList
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CompanySecretListReq
//
// Responses:
//  200: CompanySecretListResp

func GetCompanySecretListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CompanySecretListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := company_secret.NewGetCompanySecretListLogic(r.Context(), svcCtx)
		resp, err := l.GetCompanySecretList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
