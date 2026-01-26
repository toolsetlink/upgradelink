package api

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"upgradelink-admin/server/api/internal/logic/api"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"
)

func DeleteApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := api.NewDeleteApiLogic(r.Context(), svcCtx)
		resp, err := l.DeleteApi(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
