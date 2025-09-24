package upgrade_tauri_upgrade_strategy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_tauri_upgrade_strategy"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_tauri_upgrade_strategy/create upgrade_tauri_upgrade_strategy CreateUpgradeTauriUpgradeStrategy
//
// Create upgrade tauri upgrade strategy information | 创建UpgradeTauriUpgradeStrategy信息
//
// Create upgrade tauri upgrade strategy information | 创建UpgradeTauriUpgradeStrategy信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeTauriUpgradeStrategyInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeTauriUpgradeStrategyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeTauriUpgradeStrategyInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_tauri_upgrade_strategy.NewCreateUpgradeTauriUpgradeStrategyLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeTauriUpgradeStrategy(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
