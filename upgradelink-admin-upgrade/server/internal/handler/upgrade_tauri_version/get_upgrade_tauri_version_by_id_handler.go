package upgrade_tauri_version

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_tauri_version"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_tauri_version upgrade_tauri_version GetUpgradeTauriVersionById
//
// Get upgrade tauri version by ID | 通过ID获取UpgradeTauriVersion信息
//
// Get upgrade tauri version by ID | 通过ID获取UpgradeTauriVersion信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: UpgradeTauriVersionInfoResp

func GetUpgradeTauriVersionByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_tauri_version.NewGetUpgradeTauriVersionByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeTauriVersionById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
