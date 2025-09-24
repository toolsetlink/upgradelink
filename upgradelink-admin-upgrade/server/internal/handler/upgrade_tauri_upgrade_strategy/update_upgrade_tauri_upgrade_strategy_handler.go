package upgrade_tauri_upgrade_strategy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_tauri_upgrade_strategy"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_tauri_upgrade_strategy/update upgrade_tauri_upgrade_strategy UpdateUpgradeTauriUpgradeStrategy
//
// Update upgrade tauri upgrade strategy information | 更新UpgradeTauriUpgradeStrategy信息
//
// Update upgrade tauri upgrade strategy information | 更新UpgradeTauriUpgradeStrategy信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeTauriUpgradeStrategyInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateUpgradeTauriUpgradeStrategyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeTauriUpgradeStrategyInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_tauri_upgrade_strategy.NewUpdateUpgradeTauriUpgradeStrategyLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeTauriUpgradeStrategy(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
