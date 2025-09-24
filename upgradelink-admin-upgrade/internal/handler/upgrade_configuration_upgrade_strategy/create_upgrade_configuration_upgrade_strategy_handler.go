package upgrade_configuration_upgrade_strategy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_configuration_upgrade_strategy"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_configuration_upgrade_strategy/create upgrade_configuration_upgrade_strategy CreateUpgradeConfigurationUpgradeStrategy
//
// Create upgrade configuration upgrade strategy information | 创建UpgradeConfigurationUpgradeStrategy信息
//
// Create upgrade configuration upgrade strategy information | 创建UpgradeConfigurationUpgradeStrategy信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeConfigurationUpgradeStrategyInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeConfigurationUpgradeStrategyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeConfigurationUpgradeStrategyInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_configuration_upgrade_strategy.NewCreateUpgradeConfigurationUpgradeStrategyLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeConfigurationUpgradeStrategy(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
