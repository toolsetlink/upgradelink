package cloudfile

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-file/server/internal/logic/cloudfile"
	"upgradelink-admin-file/server/internal/svc"
	"upgradelink-admin-file/server/internal/types"
)

// swagger:route post /cloud_file/create cloudfile CreateCloudFile
//
// Create cloud file information | 创建云文件
//
// Create cloud file information | 创建云文件
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CloudFileInfo
//
// Responses:
//  200: BaseMsgResp

func CreateCloudFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CloudFileInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cloudfile.NewCreateCloudFileLogic(r.Context(), svcCtx)
		resp, err := l.CreateCloudFile(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
