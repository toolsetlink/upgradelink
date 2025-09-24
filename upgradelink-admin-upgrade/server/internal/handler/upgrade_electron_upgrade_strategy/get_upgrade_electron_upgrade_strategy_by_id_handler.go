package upgrade_electron_upgrade_strategy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_electron_upgrade_strategy"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_electron_upgrade_strategy upgrade_electron_upgrade_strategy GetUpgradeElectronUpgradeStrategyById
//
// Get upgrade electron upgrade strategy by ID | 通过ID获取UpgradeElectronUpgradeStrategy信息
//
// Get upgrade electron upgrade strategy by ID | 通过ID获取UpgradeElectronUpgradeStrategy信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: UpgradeElectronUpgradeStrategyInfoResp

func GetUpgradeElectronUpgradeStrategyByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_electron_upgrade_strategy.NewGetUpgradeElectronUpgradeStrategyByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeElectronUpgradeStrategyById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
