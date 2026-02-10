package tauri

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"upgradelink-api/server/api/internal/logic/tauri"
	"upgradelink-api/server/api/internal/svc"
	"upgradelink-api/server/api/internal/types"
)

func TauriActionUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TauriActionUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tauri.NewTauriActionUploadLogic(r.Context(), svcCtx)
		resp, err := l.TauriActionUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
