package cloudfiletag

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-file/internal/logic/cloudfiletag"
	"upgradelink-admin-file/internal/svc"
	"upgradelink-admin-file/internal/types"
)

// swagger:route post /cloud_file_tag/delete cloudfiletag DeleteCloudFileTag
//
// Delete cloud file tag information | 删除云文件标签信息
//
// Delete cloud file tag information | 删除云文件标签信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteCloudFileTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cloudfiletag.NewDeleteCloudFileTagLogic(r.Context(), svcCtx)
		resp, err := l.DeleteCloudFileTag(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
