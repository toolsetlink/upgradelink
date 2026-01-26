package cloudfile

import (
	"net/http"

	"upgradelink-admin/server/api/internal/logic/cloudfile"
	"upgradelink-admin/server/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cloudfile.NewUploadLogic(r.Context(), r, svcCtx)
		resp, err := l.Upload()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
