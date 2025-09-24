package upgrade_url_upgrade_strategy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_url_upgrade_strategy"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_url_upgrade_strategy/create upgrade_url_upgrade_strategy CreateUpgradeUrlUpgradeStrategy
//
// Create upgrade url upgrade strategy information | 创建UpgradeUrlUpgradeStrategy信息
//
// Create upgrade url upgrade strategy information | 创建UpgradeUrlUpgradeStrategy信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeUrlUpgradeStrategyInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeUrlUpgradeStrategyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeUrlUpgradeStrategyInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_url_upgrade_strategy.NewCreateUpgradeUrlUpgradeStrategyLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeUrlUpgradeStrategy(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
