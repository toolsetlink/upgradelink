package upgrade_electron_upgrade_strategy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_electron_upgrade_strategy"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_electron_upgrade_strategy/delete upgrade_electron_upgrade_strategy DeleteUpgradeElectronUpgradeStrategy
//
// Delete upgrade electron upgrade strategy information | 删除UpgradeElectronUpgradeStrategy信息
//
// Delete upgrade electron upgrade strategy information | 删除UpgradeElectronUpgradeStrategy信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteUpgradeElectronUpgradeStrategyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_electron_upgrade_strategy.NewDeleteUpgradeElectronUpgradeStrategyLogic(r.Context(), svcCtx)
		resp, err := l.DeleteUpgradeElectronUpgradeStrategy(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
