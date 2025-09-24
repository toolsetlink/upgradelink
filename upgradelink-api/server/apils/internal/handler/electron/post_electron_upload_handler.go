package electron

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"upgradelink-api/server/apils/internal/logic/electron"
	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"
)

func PostElectronUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostElectronUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := electron.NewPostElectronUploadLogic(r.Context(), svcCtx)
		resp, err := l.PostElectronUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
