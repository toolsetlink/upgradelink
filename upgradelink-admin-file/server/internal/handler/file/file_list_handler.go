package file

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-file/server/internal/logic/file"
	"upgradelink-admin-file/server/internal/svc"
	"upgradelink-admin-file/server/internal/types"
)

// swagger:route post /file/list file FileList
//
// Get file list | 获取文件列表
//
// Get file list | 获取文件列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: FileListReq
//
// Responses:
//  200: FileListResp

func FileListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewFileListLogic(r.Context(), svcCtx)
		resp, err := l.FileList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
