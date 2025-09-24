package upgrade_tauri_version

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_tauri_version"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_tauri_version/list upgrade_tauri_version GetUpgradeTauriVersionList
//
// Get upgrade tauri version list | 获取UpgradeTauriVersion信息列表
//
// Get upgrade tauri version list | 获取UpgradeTauriVersion信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeTauriVersionListReq
//
// Responses:
//  200: UpgradeTauriVersionListResp

func GetUpgradeTauriVersionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeTauriVersionListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_tauri_version.NewGetUpgradeTauriVersionListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeTauriVersionList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
