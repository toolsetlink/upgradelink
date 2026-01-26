package upgrade_mac_upgrade_strategy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"upgradelink-admin/server/api/internal/logic/upgrade_mac_upgrade_strategy"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"
)

func UpdateUpgradeMacUpgradeStrategyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeMacUpgradeStrategyInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_mac_upgrade_strategy.NewUpdateUpgradeMacUpgradeStrategyLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeMacUpgradeStrategy(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
