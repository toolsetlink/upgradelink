package upgrade_configuration_version

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_configuration_version"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_configuration_version/create upgrade_configuration_version CreateUpgradeConfigurationVersion
//
// Create upgrade configuration version information | 创建UpgradeConfigurationVersion信息
//
// Create upgrade configuration version information | 创建UpgradeConfigurationVersion信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeConfigurationVersionInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeConfigurationVersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeConfigurationVersionInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_configuration_version.NewCreateUpgradeConfigurationVersionLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeConfigurationVersion(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
