package cloudfiletag

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-file/internal/logic/cloudfiletag"
	"upgradelink-admin-file/internal/svc"
	"upgradelink-admin-file/internal/types"
)

// swagger:route post /cloud_file_tag/list cloudfiletag GetCloudFileTagList
//
// Get cloud file tag list | 获取云文件标签列表
//
// Get cloud file tag list | 获取云文件标签列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CloudFileTagListReq
//
// Responses:
//  200: CloudFileTagListResp

func GetCloudFileTagListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CloudFileTagListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cloudfiletag.NewGetCloudFileTagListLogic(r.Context(), svcCtx)
		resp, err := l.GetCloudFileTagList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
