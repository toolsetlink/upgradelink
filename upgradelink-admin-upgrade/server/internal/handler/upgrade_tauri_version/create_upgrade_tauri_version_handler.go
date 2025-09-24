package upgrade_tauri_version

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_tauri_version"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_tauri_version/create upgrade_tauri_version CreateUpgradeTauriVersion
//
// Create upgrade tauri version information | 创建UpgradeTauriVersion信息
//
// Create upgrade tauri version information | 创建UpgradeTauriVersion信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeTauriVersionInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeTauriVersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeTauriVersionInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_tauri_version.NewCreateUpgradeTauriVersionLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeTauriVersion(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
