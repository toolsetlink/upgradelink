package upgrade_electron

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_electron"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_electron/create upgrade_electron CreateUpgradeElectron
//
// Create upgrade electron information | 创建UpgradeElectron信息
//
// Create upgrade electron information | 创建UpgradeElectron信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeElectronInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeElectronHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeElectronInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_electron.NewCreateUpgradeElectronLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeElectron(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
