package cloudfile

import (
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-file/server/internal/logic/cloudfile"
	"upgradelink-admin-file/server/internal/svc"
	"upgradelink-admin-file/server/internal/types"
)

// swagger:route post /cloud_file/list cloudfile GetCloudFileList
//
// Get cloud file list | 获取CloudFile列表
//
// Get cloud file list | 获取CloudFile列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CloudFileListReq
//
// Responses:
//  200: CloudFileListResp

func GetCloudFileListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GetCloudFileListHandler")
		var req types.CloudFileListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cloudfile.NewGetCloudFileListLogic(r.Context(), svcCtx)
		resp, err := l.GetCloudFileList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
