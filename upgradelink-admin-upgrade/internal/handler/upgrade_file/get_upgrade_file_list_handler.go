package upgrade_file

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_file"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_file/list upgrade_file GetUpgradeFileList
//
// Get upgrade file list | 获取UpgradeFile信息列表
//
// Get upgrade file list | 获取UpgradeFile信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeFileListReq
//
// Responses:
//  200: UpgradeFileListResp

func GetUpgradeFileListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeFileListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_file.NewGetUpgradeFileListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeFileList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
