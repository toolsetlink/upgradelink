package upgrade_file_upgrade_strategy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_file_upgrade_strategy"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_file_upgrade_strategy/list upgrade_file_upgrade_strategy GetUpgradeFileUpgradeStrategyList
//
// Get upgrade file upgrade strategy list | 获取UpgradeFileUpgradeStrategy信息列表
//
// Get upgrade file upgrade strategy list | 获取UpgradeFileUpgradeStrategy信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeFileUpgradeStrategyListReq
//
// Responses:
//  200: UpgradeFileUpgradeStrategyListResp

func GetUpgradeFileUpgradeStrategyListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeFileUpgradeStrategyListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_file_upgrade_strategy.NewGetUpgradeFileUpgradeStrategyListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeFileUpgradeStrategyList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
