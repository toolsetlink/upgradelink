package upgrade_configuration

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_configuration"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_configuration/create upgrade_configuration CreateUpgradeConfiguration
//
// Create upgrade configuration information | 创建UpgradeConfiguration信息
//
// Create upgrade configuration information | 创建UpgradeConfiguration信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeConfigurationInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeConfigurationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeConfigurationInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_configuration.NewCreateUpgradeConfigurationLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeConfiguration(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
