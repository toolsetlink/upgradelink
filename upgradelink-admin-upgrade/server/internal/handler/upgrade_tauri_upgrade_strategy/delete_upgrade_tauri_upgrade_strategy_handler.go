package upgrade_tauri_upgrade_strategy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_tauri_upgrade_strategy"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_tauri_upgrade_strategy/delete upgrade_tauri_upgrade_strategy DeleteUpgradeTauriUpgradeStrategy
//
// Delete upgrade tauri upgrade strategy information | 删除UpgradeTauriUpgradeStrategy信息
//
// Delete upgrade tauri upgrade strategy information | 删除UpgradeTauriUpgradeStrategy信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteUpgradeTauriUpgradeStrategyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_tauri_upgrade_strategy.NewDeleteUpgradeTauriUpgradeStrategyLogic(r.Context(), svcCtx)
		resp, err := l.DeleteUpgradeTauriUpgradeStrategy(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
