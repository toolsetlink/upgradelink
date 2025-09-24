package upgrade_electron_version

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_electron_version"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_electron_version/create upgrade_electron_version CreateUpgradeElectronVersion
//
// Create upgrade electron version information | 创建UpgradeElectronVersion信息
//
// Create upgrade electron version information | 创建UpgradeElectronVersion信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeElectronVersionInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeElectronVersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeElectronVersionInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_electron_version.NewCreateUpgradeElectronVersionLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeElectronVersion(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
