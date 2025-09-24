package cloudfile

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-file/internal/logic/cloudfile"
	"upgradelink-admin-file/internal/svc"
	"upgradelink-admin-file/internal/types"
)

// swagger:route post /cloud_file/delete cloudfile DeleteCloudFile
//
// Delete cloud file information | 删除云文件信息
//
// Delete cloud file information | 删除云文件信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UUIDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteCloudFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UUIDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cloudfile.NewDeleteCloudFileLogic(r.Context(), svcCtx)
		resp, err := l.DeleteCloudFile(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
