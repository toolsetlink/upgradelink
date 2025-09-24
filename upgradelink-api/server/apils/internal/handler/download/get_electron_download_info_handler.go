package download

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"upgradelink-api/server/apils/internal/logic/download"
	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"
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
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
