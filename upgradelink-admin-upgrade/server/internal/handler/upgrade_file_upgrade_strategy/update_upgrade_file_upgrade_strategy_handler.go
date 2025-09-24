package upgrade_file_upgrade_strategy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_file_upgrade_strategy"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_file_upgrade_strategy/update upgrade_file_upgrade_strategy UpdateUpgradeFileUpgradeStrategy
//
// Update upgrade file upgrade strategy information | 更新UpgradeFileUpgradeStrategy信息
//
// Update upgrade file upgrade strategy information | 更新UpgradeFileUpgradeStrategy信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeFileUpgradeStrategyInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateUpgradeFileUpgradeStrategyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeFileUpgradeStrategyInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_file_upgrade_strategy.NewUpdateUpgradeFileUpgradeStrategyLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeFileUpgradeStrategy(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
