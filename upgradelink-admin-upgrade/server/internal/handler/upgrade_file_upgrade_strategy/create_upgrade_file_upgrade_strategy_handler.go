package upgrade_file_upgrade_strategy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_file_upgrade_strategy"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_file_upgrade_strategy/create upgrade_file_upgrade_strategy CreateUpgradeFileUpgradeStrategy
//
// Create upgrade file upgrade strategy information | 创建UpgradeFileUpgradeStrategy信息
//
// Create upgrade file upgrade strategy information | 创建UpgradeFileUpgradeStrategy信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeFileUpgradeStrategyInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeFileUpgradeStrategyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeFileUpgradeStrategyInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_file_upgrade_strategy.NewCreateUpgradeFileUpgradeStrategyLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeFileUpgradeStrategy(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
