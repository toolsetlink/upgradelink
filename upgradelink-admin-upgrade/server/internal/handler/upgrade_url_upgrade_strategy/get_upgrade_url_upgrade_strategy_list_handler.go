package upgrade_url_upgrade_strategy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_url_upgrade_strategy"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_url_upgrade_strategy/list upgrade_url_upgrade_strategy GetUpgradeUrlUpgradeStrategyList
//
// Get upgrade url upgrade strategy list | 获取UpgradeUrlUpgradeStrategy信息列表
//
// Get upgrade url upgrade strategy list | 获取UpgradeUrlUpgradeStrategy信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeUrlUpgradeStrategyListReq
//
// Responses:
//  200: UpgradeUrlUpgradeStrategyListResp

func GetUpgradeUrlUpgradeStrategyListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeUrlUpgradeStrategyListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_url_upgrade_strategy.NewGetUpgradeUrlUpgradeStrategyListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeUrlUpgradeStrategyList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
