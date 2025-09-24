package upgrade_electron_upgrade_strategy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_electron_upgrade_strategy"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_electron_upgrade_strategy/update upgrade_electron_upgrade_strategy UpdateUpgradeElectronUpgradeStrategy
//
// Update upgrade electron upgrade strategy information | 更新UpgradeElectronUpgradeStrategy信息
//
// Update upgrade electron upgrade strategy information | 更新UpgradeElectronUpgradeStrategy信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeElectronUpgradeStrategyInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateUpgradeElectronUpgradeStrategyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeElectronUpgradeStrategyInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_electron_upgrade_strategy.NewUpdateUpgradeElectronUpgradeStrategyLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeElectronUpgradeStrategy(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
