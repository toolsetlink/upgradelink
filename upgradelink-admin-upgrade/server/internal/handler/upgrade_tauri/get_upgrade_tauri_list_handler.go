package upgrade_tauri

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_tauri"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_tauri/list upgrade_tauri GetUpgradeTauriList
//
// Get upgrade tauri list | 获取UpgradeTauri信息列表
//
// Get upgrade tauri list | 获取UpgradeTauri信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeTauriListReq
//
// Responses:
//  200: UpgradeTauriListResp

func GetUpgradeTauriListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeTauriListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_tauri.NewGetUpgradeTauriListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeTauriList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
