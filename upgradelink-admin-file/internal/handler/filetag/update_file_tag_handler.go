package filetag

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-file/internal/logic/filetag"
	"upgradelink-admin-file/internal/svc"
	"upgradelink-admin-file/internal/types"
)

// swagger:route post /file_tag/update filetag UpdateFileTag
//
// Update file tag information | 更新文件标签
//
// Update file tag information | 更新文件标签
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: FileTagInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateFileTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileTagInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := filetag.NewUpdateFileTagLogic(r.Context(), svcCtx)
		resp, err := l.UpdateFileTag(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
