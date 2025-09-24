package upgrade_apk_upgrade_strategy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_apk_upgrade_strategy"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_apk_upgrade_strategy/create upgrade_apk_upgrade_strategy CreateUpgradeApkUpgradeStrategy
//
// Create upgrade apk upgrade strategy information | 创建UpgradeApkUpgradeStrategy信息
//
// Create upgrade apk upgrade strategy information | 创建UpgradeApkUpgradeStrategy信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeApkUpgradeStrategyInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeApkUpgradeStrategyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeApkUpgradeStrategyInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_apk_upgrade_strategy.NewCreateUpgradeApkUpgradeStrategyLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeApkUpgradeStrategy(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
