package filetag

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-file/internal/logic/filetag"
	"upgradelink-admin-file/internal/svc"
	"upgradelink-admin-file/internal/types"
)

// swagger:route post /file_tag/list filetag GetFileTagList
//
// Get file tag list | 获取文件标签列表
//
// Get file tag list | 获取文件标签列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: FileTagListReq
//
// Responses:
//  200: FileTagListResp

func GetFileTagListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileTagListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := filetag.NewGetFileTagListLogic(r.Context(), svcCtx)
		resp, err := l.GetFileTagList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
