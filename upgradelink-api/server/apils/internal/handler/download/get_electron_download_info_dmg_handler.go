package download

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"upgradelink-api/server/apils/internal/logic/download"
	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"
)

func GetElectronDownloadInfoDmgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetElectronDownloadInfoDmgReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := download.NewGetElectronDownloadInfoDmgLogic(r.Context(), svcCtx)
		resp, err := l.GetElectronDownloadInfoDmg(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
