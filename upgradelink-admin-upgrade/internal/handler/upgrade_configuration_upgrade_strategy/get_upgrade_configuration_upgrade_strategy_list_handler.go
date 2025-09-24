package upgrade_configuration_upgrade_strategy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_configuration_upgrade_strategy"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_configuration_upgrade_strategy/list upgrade_configuration_upgrade_strategy GetUpgradeConfigurationUpgradeStrategyList
//
// Get upgrade configuration upgrade strategy list | 获取UpgradeConfigurationUpgradeStrategy信息列表
//
// Get upgrade configuration upgrade strategy list | 获取UpgradeConfigurationUpgradeStrategy信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeConfigurationUpgradeStrategyListReq
//
// Responses:
//  200: UpgradeConfigurationUpgradeStrategyListResp

func GetUpgradeConfigurationUpgradeStrategyListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeConfigurationUpgradeStrategyListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_configuration_upgrade_strategy.NewGetUpgradeConfigurationUpgradeStrategyListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeConfigurationUpgradeStrategyList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
