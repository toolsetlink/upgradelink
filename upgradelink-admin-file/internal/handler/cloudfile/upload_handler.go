package cloudfile

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-file/internal/logic/cloudfile"
	"upgradelink-admin-file/internal/svc"
)

// swagger:route post /cloud_file/upload cloudfile Upload
//
// Cloud file upload | 上传文件
//
// Cloud file upload | 上传文件
//
// Responses:
//  200: CloudFileInfoResp

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cloudfile.NewUploadLogic(r, svcCtx)
		resp, err := l.Upload()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
