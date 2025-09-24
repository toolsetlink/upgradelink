package upgrade_configuration

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_configuration"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_configuration/update upgrade_configuration UpdateUpgradeConfiguration
//
// Update upgrade configuration information | 更新UpgradeConfiguration信息
//
// Update upgrade configuration information | 更新UpgradeConfiguration信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeConfigurationInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateUpgradeConfigurationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeConfigurationInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_configuration.NewUpdateUpgradeConfigurationLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeConfiguration(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
