package download

import (
	"net/http"

	"upgradelink-api/server/api/internal/logic/download"
	"upgradelink-api/server/api/internal/svc"
	"upgradelink-api/server/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetApkDownloadInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetApkDownloadInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := download.NewGetApkDownloadInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetApkDownloadInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			// 执行 HTTP 重定向
			http.Redirect(w, r, resp, http.StatusFound) // 302 临时重定向
			//httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
