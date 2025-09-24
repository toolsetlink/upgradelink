package upgrade_tauri_upgrade_strategy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_tauri_upgrade_strategy"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_tauri_upgrade_strategy upgrade_tauri_upgrade_strategy GetUpgradeTauriUpgradeStrategyById
//
// Get upgrade tauri upgrade strategy by ID | 通过ID获取UpgradeTauriUpgradeStrategy信息
//
// Get upgrade tauri upgrade strategy by ID | 通过ID获取UpgradeTauriUpgradeStrategy信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: UpgradeTauriUpgradeStrategyInfoResp

func GetUpgradeTauriUpgradeStrategyByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_tauri_upgrade_strategy.NewGetUpgradeTauriUpgradeStrategyByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeTauriUpgradeStrategyById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
