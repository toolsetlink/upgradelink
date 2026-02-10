package file

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"upgradelink-api/server/api/internal/logic/file"
	"upgradelink-api/server/api/internal/svc"
	"upgradelink-api/server/api/internal/types"
)

func FileActionUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileActionUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewFileActionUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileActionUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
