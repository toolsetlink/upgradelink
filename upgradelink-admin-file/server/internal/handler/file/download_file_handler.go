package file

import (
	"net/http"
	"os"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-file/server/internal/logic/file"
	"upgradelink-admin-file/server/internal/svc"
	"upgradelink-admin-file/server/internal/types"
)

// swagger:route get /file/download/{id} file DownloadFile
//
// Download file | 下载文件
//
// Download file | 下载文件
//
// Parameters:
//  + name: id
//    require: true
//    in: path
//
// Responses:
//  200: BaseMsgResp

func DownloadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UUIDPathReq
		if err := httpx.Parse(r, &req, false); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewDownloadFileLogic(r.Context(), svcCtx)
		filePath, err := l.DownloadFile(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			body, err := os.ReadFile(filePath)
			if err != nil {
				httpx.Error(w, errorx.NewApiError(http.StatusInternalServerError, err.Error()))
				return
			}
			w.Header().Set("Accept-Encoding", "identity;q=1, *;q=0")
			w.Write(body)
		}
	}
}
