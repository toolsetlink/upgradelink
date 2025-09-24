package cloudfiletag

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-file/internal/logic/cloudfiletag"
	"upgradelink-admin-file/internal/svc"
	"upgradelink-admin-file/internal/types"
)

// swagger:route post /cloud_file_tag/update cloudfiletag UpdateCloudFileTag
//
// Update cloud file tag information | 更新云文件标签
//
// Update cloud file tag information | 更新云文件标签
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CloudFileTagInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateCloudFileTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CloudFileTagInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cloudfiletag.NewUpdateCloudFileTagLogic(r.Context(), svcCtx)
		resp, err := l.UpdateCloudFileTag(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
