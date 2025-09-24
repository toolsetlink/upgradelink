package upgrade_tauri_version

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_tauri_version"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_tauri_version/update upgrade_tauri_version UpdateUpgradeTauriVersion
//
// Update upgrade tauri version information | 更新UpgradeTauriVersion信息
//
// Update upgrade tauri version information | 更新UpgradeTauriVersion信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeTauriVersionInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateUpgradeTauriVersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeTauriVersionInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_tauri_version.NewUpdateUpgradeTauriVersionLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeTauriVersion(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
