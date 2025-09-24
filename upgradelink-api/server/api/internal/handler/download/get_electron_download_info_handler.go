package download

import (
	"net/http"

	"upgradelink-api/server/api/internal/logic/download"
	"upgradelink-api/server/api/internal/svc"
	"upgradelink-api/server/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetElectronDownloadInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetElectronDownloadInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := download.NewGetElectronDownloadInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetElectronDownloadInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			// 执行 HTTP 重定向
			http.Redirect(w, r, *resp, http.StatusFound) // 302 临时重定向
			//httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
