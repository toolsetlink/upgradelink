package cloudfile

import (
	"net/http"

	"upgradelink-admin/server/api/internal/logic/cloudfile"
	"upgradelink-admin/server/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PresignedUrlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cloudfile.NewPresignedUrlLogic(r.Context(), r, svcCtx)
		resp, err := l.GeneratePresignedUrl()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
