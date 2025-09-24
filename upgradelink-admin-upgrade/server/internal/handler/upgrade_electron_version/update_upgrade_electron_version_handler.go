package upgrade_electron_version

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_electron_version"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_electron_version/update upgrade_electron_version UpdateUpgradeElectronVersion
//
// Update upgrade electron version information | 更新UpgradeElectronVersion信息
//
// Update upgrade electron version information | 更新UpgradeElectronVersion信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeElectronVersionInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateUpgradeElectronVersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeElectronVersionInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_electron_version.NewUpdateUpgradeElectronVersionLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeElectronVersion(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
