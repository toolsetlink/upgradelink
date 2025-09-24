package upgrade_electron

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_electron"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_electron/update upgrade_electron UpdateUpgradeElectron
//
// Update upgrade electron information | 更新UpgradeElectron信息
//
// Update upgrade electron information | 更新UpgradeElectron信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeElectronInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateUpgradeElectronHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeElectronInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_electron.NewUpdateUpgradeElectronLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeElectron(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
