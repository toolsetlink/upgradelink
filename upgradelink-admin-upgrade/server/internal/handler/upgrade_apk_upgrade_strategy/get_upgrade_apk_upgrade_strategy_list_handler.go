package upgrade_apk_upgrade_strategy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_apk_upgrade_strategy"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_apk_upgrade_strategy/list upgrade_apk_upgrade_strategy GetUpgradeApkUpgradeStrategyList
//
// Get upgrade apk upgrade strategy list | 获取UpgradeApkUpgradeStrategy信息列表
//
// Get upgrade apk upgrade strategy list | 获取UpgradeApkUpgradeStrategy信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeApkUpgradeStrategyListReq
//
// Responses:
//  200: UpgradeApkUpgradeStrategyListResp

func GetUpgradeApkUpgradeStrategyListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeApkUpgradeStrategyListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_apk_upgrade_strategy.NewGetUpgradeApkUpgradeStrategyListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeApkUpgradeStrategyList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
