package cloudfile

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"upgradelink-admin/server/api/internal/logic/cloudfile"
	"upgradelink-admin/server/api/internal/svc"
)

func InternalUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cloudfile.NewInternalUploadLogic(r.Context(), svcCtx)
		resp, err := l.InternalUpload()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
