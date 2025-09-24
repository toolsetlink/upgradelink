package cloudfile

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-file/server/internal/logic/cloudfile"
	"upgradelink-admin-file/server/internal/svc"
	"upgradelink-admin-file/server/internal/types"
)

// swagger:route post /cloud_file cloudfile GetCloudFileById
//
// Get cloud file by ID | 通过ID获取云文件
//
// Get cloud file by ID | 通过ID获取云文件
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UUIDReq
//
// Responses:
//  200: CloudFileInfoResp

func GetCloudFileByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UUIDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cloudfile.NewGetCloudFileByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetCloudFileById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
