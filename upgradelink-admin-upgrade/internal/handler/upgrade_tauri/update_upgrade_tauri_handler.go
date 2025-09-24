package upgrade_tauri

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_tauri"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_tauri/update upgrade_tauri UpdateUpgradeTauri
//
// Update upgrade tauri information | 更新UpgradeTauri信息
//
// Update upgrade tauri information | 更新UpgradeTauri信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeTauriInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateUpgradeTauriHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeTauriInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_tauri.NewUpdateUpgradeTauriLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeTauri(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
