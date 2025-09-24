package file

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-file/internal/logic/file"
	"upgradelink-admin-file/internal/svc"
)

// swagger:route post /upload file Upload
//
// Upload file | 上传文件
//
// Upload file | 上传文件
//
// Responses:
//  200: UploadResp

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := file.NewUploadLogic(r, svcCtx)
		resp, err := l.Upload()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
